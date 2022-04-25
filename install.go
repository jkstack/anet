package anet

const (
	InstallActionDone = iota
	InstallActionDownload
	InstallActionInstall
	InstallActionFile
)

type InstallConfig struct {
	URI string `json:"uri"`
	URL string `json:"url"`
	Dir string `json:"dir"`
}

type InstallArgs struct {
	URI     string          `json:"uri"`
	URL     string          `json:"url"`
	Dir     string          `json:"dir"`
	Timeout int             `json:"timeout"`
	Auth    string          `json:"auth,omitempty"`
	User    string          `json:"user,omitempty"`
	Pass    string          `json:"pass,omitempty"`
	Configs []InstallConfig `json:"conf,omitempty"`
}

type InstallRep struct {
	Time   int64  `json:"ts"`
	Action uint8  `json:"action"`
	Name   string `json:"name"`
	OK     bool   `json:"ok"`
	Msg    string `json:"msg"`
}
