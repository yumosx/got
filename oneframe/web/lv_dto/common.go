package lv_dto

type BusinessType int

const (
	Business_Other BusinessType = 0
	Business_Add   BusinessType = 1
	Business_Edit  BusinessType = 2
	Business_Del   BusinessType = 3
)

const (
	Success      = 200
	Error        = 500
	Unauthorized = 403
	Fail         = -1
)

const (
	ErrorPage  = "error/error.html"
	UnauthPage = "error/unauth.html"
)

type CommonResponse struct {
	Code         int          `json:"code"`
	Msg          string       `json:"msg"`
	Data         any          `json:"data"`
	BusinessType BusinessType `json:"otype"`
}

type CaptchaResponse struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	Image          any    `json:"img"`
	UUID           string `json:"uuid"`
	CaptchaEnabled bool   `json:"captchaEnabled"`
}

type TableResponse struct {
	Total any    `json:"total"`
	Rows  any    `json:"rows"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
}

type ZTree struct {
	Id       int64  `json:"id"`
	Pid      int64  `json:"pId"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	Checked  bool   `json:"checked"`
	Open     bool   `json:"open"`
	NoCheck  bool   `json:"nocheck"`
	NodeType string `json:"nodeType"`
}

type DeleteReq struct {
	Ids string `form:"ids"  binding:"required"`
}

type DetailReq struct {
	Id int64 `json:"id"`
}

type EditReq struct {
	Id int64 `json:"id"`
}
