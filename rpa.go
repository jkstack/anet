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
