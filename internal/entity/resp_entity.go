package entity

//http status字段
const (
	HttpSuccessful = 1
	HttpFailure    = 0
)

type AddContentResp struct {
	ErrCode   int64  `json:"err_code"`
	ContentID string `json:"content_id"`
	ErrMsg    string `json:"err_msg"`
}

type OperateContentResp struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

type GetContentListResp struct {
	ContentList []BaseContent `json:"content_list"`
}

type GetContentDetailResp Content

type HttpResponseObject struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	ErrCode int         `json:"err_code"`
	Data    interface{} `json:"data"`
}
