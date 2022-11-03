package content

import (
	"contentService/internal/entity"
	"contentService/internal/rpc/datamanager"
	"contentService/pkg/proto/content"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
	"strings"
	"time"
)

type rpcContent struct {
	content.UnimplementedContentManagerServer
	rpcPort     int
	rpcListenIP string
}

func (rpc *rpcContent) GetContentListRK() string {
	return "content:list:kv"
}

func (rpc *rpcContent) GetContentInfoRK(contentID string) string {
	return fmt.Sprintf("content:info:kv:%s", contentID)
}

func (rpc *rpcContent) GetContentLit(_ context.Context, req *content.GetContentListReq) (resp *content.GetContentListResp, err error) {
	//缓存第一页，因为文章信息更改对文章列表信息的实时要求不高且大量用户基本停留在第一页，所以可以把第一页缓存个30秒
	resp = &content.GetContentListResp{}
	list, err := datamanager.GetRedisInstance().Get(rpc.GetContentListRK())
	if list != "" && err == nil {
		var ret []entity.BaseContent
		err = json.Unmarshal([]byte(list), &ret)
		if err != nil {
			return nil, err
		}
		for _, v := range ret {
			resp.List = append(resp.List, &content.Content{
				Title:     v.Title,
				HeadPhoto: v.HeadPhoto,
				ContentID: v.ContentID,
				Name:      v.Name,
				AuthorID:  v.AuthorID,
				Tag:       v.Tag,
			})
		}
		return
	}
	ret, err := datamanager.GetMysqlInstance().GetContentListByTag(int(req.Stop-req.Start)+1, int(req.Start), int(req.Tag), int(req.Status))
	if err != nil {
		return nil, err
	}
	for _, v := range ret {
		resp.List = append(resp.List, &content.Content{
			Title:     v.Title,
			HeadPhoto: v.HeadPhoto,
			ContentID: v.ContentID,
			Name:      v.Name,
			AuthorID:  v.AuthorID,
			Tag:       v.Tag,
		})
	}
	jsonStr, err := json.Marshal(ret)
	if err == nil {
		datamanager.GetRedisInstance().SetNX(rpc.GetContentListRK(), string(jsonStr), 30*time.Second)
	}
	return
}

func (rpc *rpcContent) GetContentDetail(_ context.Context, req *content.GetContentDetailReq) (resp *content.GetContentDetailResp, err error) {
	//给详情页加个redis 缓存，每次更新的时候删除缓存，然后懒加载
	ret, err := datamanager.GetRedisInstance().Get(rpc.GetContentInfoRK(req.ContentID))
	var data *entity.Content
	if err == nil && ret != "" {
		data = &entity.Content{}
		err = json.Unmarshal([]byte(ret), data)
	} else {
		//从mysql取
		err = nil
		data, err = datamanager.GetMysqlInstance().GetContentDetail(req.GetContentID())
		if err == nil {
			str, _ := json.Marshal(data)
			datamanager.GetRedisInstance().SetNX(rpc.GetContentInfoRK(req.ContentID), string(str), 12*time.Hour)
		}
	}
	if data != nil {
		resp = &content.GetContentDetailResp{
			Content: &content.ContentDetail{
				Title:       data.Title,
				HeadPhoto:   data.HeadPhoto,
				Status:      data.Status,
				AddTime:     data.CreateTime,
				UpdateTime:  data.UpdateTime,
				ContentID:   data.ContentID,
				AuthorID:    data.AuthorID,
				Description: data.Description,
				Name:        data.Name,
				Tag:         data.Tag,
			},
		}
	}
	return
}
func (rpc *rpcContent) AddContent(_ context.Context, req *content.AddContentReq) (*content.AddContentResp, error) {
	uuID := strings.ReplaceAll(uuid.New().String(), "-", "")
	err := datamanager.GetMysqlInstance().InsertContent(entity.Content{
		Title:       req.Content.Title,
		Tag:         req.Content.Tag,
		HeadPhoto:   req.Content.HeadPhoto,
		ContentID:   uuID,
		Name:        req.Content.Name,
		AuthorID:    req.Content.AuthorID,
		Status:      0,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime:  time.Now().Format("2006-01-02 15:04:05"),
		Description: req.Content.Description,
	})
	if err != nil {
		return &content.AddContentResp{
			ErrCode:   -1,
			ContentID: "",
			ErrMsg:    err.Error(),
		}, err
	}
	return &content.AddContentResp{
		ErrCode:   0,
		ContentID: uuID,
		ErrMsg:    "",
	}, nil
}

func (rpc *rpcContent) OperateContent(_ context.Context, req *content.OperateContentReq) (resp *content.OperateContentResp, err error) {
	switch req.Action {
	case "del":
		err = datamanager.GetMysqlInstance().UpdateContent(req.ContentID, entity.Content{
			Status: 5,
		})
	}
	//清除缓存
	datamanager.GetRedisInstance().Del(rpc.GetContentInfoRK(req.ContentID))
	if err == nil {
		resp = &content.OperateContentResp{
			ErrCode: 0,
			ErrMsg:  "",
		}
	}
	return
}

func NewRpcContentServer(port int, addr string) *rpcContent {
	return &rpcContent{
		rpcPort:     port,
		rpcListenIP: addr,
	}
}

func (rpc *rpcContent) Run() {
	address := rpc.rpcListenIP + ":" + strconv.Itoa(rpc.rpcPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("listening err:" + err.Error())
	}
	//grpc server
	srv := grpc.NewServer()
	defer srv.GracefulStop()
	content.RegisterContentManagerServer(srv, rpc)
	reflection.Register(srv)
	if err = srv.Serve(listener); err != nil {
		fmt.Println("Serve failed ", err.Error())
		return
	}
	fmt.Println("rpc content ok")
}
