package anet

const (
	RPAPause = iota
	RPAStop
	RPAContinue
)

type RPARunArgs struct {
	URL     string `json:"url"`                // zip包下载地址
	IsDebug bool   `json:"is_debug,omitempty"` // 是否是调试模式运行
}

type RPALogData string

type RPACtrlReq struct {
	Status int `json:"status"` // 操作状态，0=pause, 1=stop, 2=continue
}

type RPACtrlRep struct {
	OK  bool   `json:"ok,omitempty"`  // 操作是否成功
	Msg string `json:"msg,omitempty"` // 失败原因
}

type RPAFinish struct {
	Code int    `json:"code,omitempty"` // 状态码
	Msg  string `json:"msg,omitempty"`  // 错误信息
}

type RPASelectorRep struct {
	Code int    `json:"code,omitempty"` // 状态码
	Msg  string `json:"msg,omitempty"`  // 错误信息
}

type RPASelectorResult struct {
	Code    int    `json:"code,omitempty"`    // 状态码
	Msg     string `json:"msg,omitempty"`     // 错误信息
	Content string `json:"content,omitempty"` // xml内容
	Image   string `json:"image,omitempty"`   // 图片
}

type RPASelectorValidateReq struct {
	Content string `json:"content"` // xml内容
}

type RPASelectorValidateRep struct {
	OK  bool   `json:"ok"`  // 是否成功
	Msg string `json:"msg"` // 错误信息
}
