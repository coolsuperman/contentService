//单元测试
package content

import (
	"contentService/pkg/proto/content"
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"testing"
	"time"
)

func TestRpcContent_AddContent(t *testing.T) {
	handler := rpcContent{}
	fmt.Println(strings.ReplaceAll(uuid.New().String(), "-", ""))
	ret, err := handler.AddContent(context.Background(), &content.AddContentReq{
		Content: &content.ContentDetail{
			Title:       "测试2",
			HeadPhoto:   "aaaaaaa",
			AddTime:     time.Now().Format("2006-01-02 15:04:05"),
			UpdateTime:  time.Now().Format("2006-01-02 15:04:05"),
			AuthorID:    "222333",
			Description: "jjjjjjjjjjjjjjjjjjjj",
			Name:        "ssssssss",
			Tag:         1,
		},
	})
	fmt.Println(ret, err)
}

func TestRpcContent_GetContentDetail(t *testing.T) {
	handler := rpcContent{}
	ret, err := handler.GetContentDetail(context.Background(), &content.GetContentDetailReq{
		ContentID: "845d2b1168964df8860a6ccb5a511184",
	})
	fmt.Println(ret, err)
}

func TestRpcContent_GetContentLit(t *testing.T) {
	handler := rpcContent{}
	ret, err := handler.GetContentLit(context.Background(), &content.GetContentListReq{
		Tag:    1,
		Start:  0,
		Stop:   5,
		Status: 0,
	})
	fmt.Println(ret, err)
}

func TestRpcContent_OperateContent(t *testing.T) {
	handler := rpcContent{}
	ret, err := handler.OperateContent(context.Background(), &content.OperateContentReq{
		Action:    "del",
		Comment:   "删除",
		ContentID: "845d2b1168964df8860a6ccb5a511184",
	})
	fmt.Println(ret, err)
}
