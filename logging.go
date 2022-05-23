package anet

const (
	LoggingTypeK8s = iota
	LoggingTypeDocker
	LoggingTypeFile
)

const (
	LoggingAlreadyRunningMsg = "project is already running"
	LoggingNotRunningMsg     = "project is not running"
)

type LoggingConfigK8s struct {
	Namespace string   `json:"ns"`
	Names     []string `json:"nms"`
	Dir       string   `json:"d"`
	Api       string   `json:"a"`
	Token     string   `json:"tk"`
}

type LoggingConfigDocker struct {
	ContainerName string `json:"cn"`
	ContainerTag  string `json:"ct"`
	Dir           string `json:"d"`
}

type LoggingConfigFile struct {
	Dir string `json:"d"`
}

type LoggingConfig struct {
	Pid      int64                `json:"pid"`
	T        int                  `json:"t"`
	Exclude  string               `json:"ext,omitempty"`
	Batch    int                  `json:"b"`
	Buffer   int                  `json:"bf"`
	Interval int                  `json:"itv"`
	Report   string               `json:"rp"`
	K8s      *LoggingConfigK8s    `json:"k8s,omitempty"`
	Docker   *LoggingConfigDocker `json:"d,omitempty"`
	File     *LoggingConfigFile   `json:"f,omitempty"`
}

type LoggingStatusReq struct {
	ID      int64 `json:"id"`
	Running bool  `json:"run"`
}

type LoggingStatusRep struct {
	OK  bool   `json:"ok"`
	Msg string `json:"msg"`
}

type LoggingReportTask struct {
	ID      int64 `json:"id"`
	Pods    int   `json:"pods"`
	Running bool  `json:"run"`
}

type LoggingReport struct {
	Tasks []LoggingReportTask `json:"tasks"`
}
