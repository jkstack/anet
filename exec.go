package anet

import "time"

// ExecPayload exec payload
type ExecPayload struct {
	Cmd     string   `json:"cmd"`               // 命令路径
	Args    []string `json:"args,omitempty"`    // 命令参数
	Timeout int      `json:"timeout"`           // 超时时间，单位秒
	Auth    string   `json:"auth,omitempty"`    // 是否提权，su、sudo或空
	User    string   `json:"user,omitempty"`    // 提权用户
	Pass    string   `json:"pass,omitempty"`    // 提权密码
	WorkDir string   `json:"workdir,omitempty"` // 工作目录
	Env     []string `json:"env,omitempty"`     // 运行时的环境变量
	DeferRM string   `json:"drm,omitempty"`
}

// ExecdPayload execd payload
type ExecdPayload struct {
	OK   bool      `json:"ok"`            // 是否启动成功
	Msg  string    `json:"msg,omitempty"` // 启动失败时的详情
	Pid  int       `json:"pid"`           // 进程id
	Time time.Time `json:"t"`             // 启动时间
}

// ExecData exec data
type ExecData struct {
	Data string `json:"data"`
}

// ExecDone exec done
type ExecDone struct {
	Code int       `json:"code"` // 退出时的返回码
	Time time.Time `json:"t"`    // 结束时间
}

// ExecKill exec kill
type ExecKill struct {
	Pid int `json:"pid"` // 进程ID
}
