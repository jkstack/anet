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

type LoggingReportK8sData struct {
	QPS            float64 `json:"qps,omitempty"`
	AvgCost        float64 `json:"avg,omitempty"`
	P0             int64   `json:"p0,omitempty"`
	P50            int64   `json:"p50,omitempty"`
	P90            int64   `json:"p90,omitempty"`
	P99            int64   `json:"p99,omitempty"`
	P100           int64   `json:"p100,omitempty"`
	CountService   int     `json:"ct_svr,omitempty"`
	CountPod       int     `json:"ct_pds,omitempty"`
	CountContainer int     `json:"ct_cts,omitempty"`
}

type LoggingReportDockerData struct {
	Count int `json:"ct,omitempty"`
}

type LoggingReportFileData struct {
	Count int `json:"ct,omitempty"`
}

type LoggingReportInfo struct {
	QPS     float64 `json:"qps,omitempty"`
	AvgCost float64 `json:"avg,omitempty"`
	P0      int64   `json:"p0,omitempty"`
	P50     int64   `json:"p50,omitempty"`
	P90     int64   `json:"p90,omitempty"`
	P99     int64   `json:"p99,omitempty"`
	P100    int64   `json:"p100,omitempty"`
	Count   uint64  `json:"ct,omitempty"`
	Bytes   uint64  `json:"b,omitempty"`
	HttpErr uint64  `json:"he,omitempty"`
	SrvErr  uint64  `json:"se,omitempty"`
}

type LoggingReportAgentInfo struct {
	GoVersion  string             `json:"gv"`
	Threads    int                `json:"ts"`
	Routines   int                `json:"rs"`
	Startup    int64              `json:"st"`
	HeapInuse  uint64             `json:"hi"`
	GC         map[string]float64 `json:"gc"`
	InPackets  uint64             `json:"ip"`
	InBytes    uint64             `json:"ib"`
	OutPackets uint64             `json:"op"`
	OutBytes   uint64             `json:"ob"`
}

type LoggingReport struct {
	CountK8s    uint64                  `json:"ct_k8s,omitempty"`
	CountDocker uint64                  `json:"ct_dr,omitempty"`
	CountFile   uint64                  `json:"ct_f,omitempty"`
	K8s         LoggingReportK8sData    `json:"k8s"`
	Docker      LoggingReportDockerData `json:"dr"`
	File        LoggingReportFileData   `json:"f"`
	Info        LoggingReportInfo       `json:"rp"`
	AgentInfo   LoggingReportAgentInfo  `json:"ag"`
}
