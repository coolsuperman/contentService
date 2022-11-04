package content

import (
	"contentService/internal/entity"
	"contentService/pkg/proto/content"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

//对rpc链接方式进行注入，现在用的是go-zero的zrpc
type ApiContent struct {
	ConnClient *grpc.ClientConn
}

func NewApiContent(conn *grpc.ClientConn) *ApiContent {
	return &ApiContent{
		ConnClient: conn,
	}
}

//添加文章
func (a *ApiContent) AddContent(c *gin.Context) {
	var tmpReq entity.AddContentReq
	err := c.ShouldBindJSON(&tmpReq)
	if err != nil {
		HttpFailResp(c, http.StatusForbidden, "params err", 1000, nil)
		return
	}
	manager := content.NewContentManagerClient(a.ConnClient)
	req := content.AddContentReq{
		Content: &content.ContentDetail{
			Title:       tmpReq.Title,
			HeadPhoto:   tmpReq.HeadPhoto,
			AuthorID:    tmpReq.AuthorID,
			Description: tmpReq.Description,
			Name:        tmpReq.Name,
			Tag:         tmpReq.Tag,
		},
	}
	RpcResp, err := manager.AddContent(context.Background(), &req)
	if err != nil {
		HttpFailResp(c, http.StatusInternalServerError, "call add AddContent rpc server failed", 1001, nil)
		return
	}
	resp := entity.AddContentResp{
		ErrCode:   RpcResp.ErrCode,
		ContentID: RpcResp.ContentID,
		ErrMsg:    RpcResp.ErrMsg,
	}
	HttpSuccessResp(c, http.StatusOK, "OK", resp)
}

//删除文章
func (a *ApiContent) DeleteContent(c *gin.Context) {
	comment, contentID := c.Query("comment"), c.Query("contentID")
	if contentID == "" || comment == "" {
		HttpFailResp(c, http.StatusForbidden, "params err", 1000, nil)
		return
	}
	req := &content.OperateContentReq{
		Action:    "del",
		Comment:   comment,
		ContentID: contentID,
	}
	manager := content.NewContentManagerClient(a.ConnClient)
	RpcResp, err := manager.OperateContent(context.Background(), req)
	if err != nil {
		HttpFailResp(c, http.StatusInternalServerError, "call add DeleteContent rpc server failed", 1001, nil)
		return
	}
	resp := entity.OperateContentResp{
		ErrCode: RpcResp.ErrCode,
		ErrMsg:  RpcResp.ErrMsg,
	}
	HttpSuccessResp(c, http.StatusOK, "OK", resp)
}

//获取文章列表
func (a *ApiContent) GetContentList(c *gin.Context) {
	var tmpReq entity.GetContentListReq
	err := c.ShouldBindQuery(&tmpReq)
	if err != nil {
		HttpFailResp(c, http.StatusForbidden, "params err", 1000, nil)
		return
	}
	req := &content.GetContentListReq{
		Tag:    tmpReq.Tag,
		Start:  tmpReq.Start,
		Stop:   tmpReq.Stop,
		Status: tmpReq.Status,
	}
	manager := content.NewContentManagerClient(a.ConnClient)
	RpcResp, err := manager.GetContentLit(context.Background(), req)
	if err != nil {
		HttpFailResp(c, http.StatusInternalServerError, "call add GetContentList rpc server failed", 1001, nil)
		return
	}
	var resp entity.GetContentListResp
	for _, v := range RpcResp.List {
		resp.ContentList = append(resp.ContentList, entity.BaseContent{
			Title:     v.Title,
			HeadPhoto: v.HeadPhoto,
			ContentID: v.ContentID,
			Name:      v.Name,
			AuthorID:  v.AuthorID,
		})
	}
	HttpSuccessResp(c, http.StatusOK, "OK", resp)
}

//获取文章详情
func (a *ApiContent) GetContentDetail(c *gin.Context) {
	contentID := c.Query("content_id")
	if contentID == "" {
		HttpFailResp(c, http.StatusForbidden, "params err", 1000, nil)
		return
	}
	req := &content.GetContentDetailReq{
		ContentID: contentID,
	}
	manager := content.NewContentManagerClient(a.ConnClient)
	RpcResp, err := manager.GetContentDetail(context.Background(), req)
	if err != nil {
		HttpFailResp(c, http.StatusInternalServerError, "call add GetContentDetail rpc server failed", 1001, nil)
		return
	}
	resp := entity.GetContentDetailResp{
		Title:       RpcResp.Content.Title,
		HeadPhoto:   RpcResp.Content.HeadPhoto,
		ContentID:   RpcResp.Content.ContentID,
		Name:        RpcResp.Content.Name,
		AuthorID:    RpcResp.Content.AuthorID,
		Status:      RpcResp.Content.Status,
		CreateTime:  RpcResp.Content.AddTime,
		UpdateTime:  RpcResp.Content.UpdateTime,
		Description: RpcResp.Content.Description,
	}
	HttpSuccessResp(c, http.StatusOK, "OK", resp)
}

func HttpSuccessResp(ctx *gin.Context, httpCode int, msg string, data interface{}) {
	resp := entity.HttpResponseObject{
		Status:  entity.HttpSuccessful,
		Message: msg,
		ErrCode: 200,
		Data:    data,
	}

	ctx.JSON(httpCode, resp)
}

func HttpFailResp(ctx *gin.Context, httpCode int, msg string, errCode int, data interface{}) {
	resp := entity.HttpResponseObject{
		Status:  entity.HttpFailure,
		Message: msg,
		ErrCode: errCode,
		Data:    data,
	}
	ctx.JSON(httpCode, resp)
}
