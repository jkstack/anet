package anet

type TypeName int

// global
const (
	TypeError     TypeName = iota + 0 // 0：错误消息，可能是一个全局错误
	TypeCome                          // 1：握手消息，agent发起
	TypeHandshake                     // 2：握手消息，服务端返回
)

const TypeAgentInfo TypeName = iota + 7 // 7: agent状态上报

const (
	TypeFoo TypeName = iota + 8 // 8: foo消息
	TypeBar                     // 9: bar消息
)

// exec
const (
	TypeExec     TypeName = iota + 10 // 10：执行命令
	TypeExecd                         // 11：脚本或命令启动结果
	TypeExecData                      // 12：脚本返回内容
	TypeExecDone                      // 13：脚本执行结果
	TypeExecKill                      // 14：kill命令
)

// file
const (
	TypeLsReq         TypeName = iota + 20 // 20：ls命令请求参数
	TypeLsRep                              // 21：ls命令返回结果
	TypeUpload                             // 22：文件上传时的文件信息
	TypeUploadRep                          // 23：文件写入时的结果
	TypeDownloadReq                        // 24：文件下载请求参数
	TypeDownloadRep                        // 25：文件下载请求时的文件信息
	TypeDownloadData                       // 26：文件下载时的内容
	TypeDownloadError                      // 27：文件下载时的错误信息
)

// host.monitor
const (
	TypeHMStaticReq           TypeName = iota + 30 // 30: 获取主机信息请求
	TypeHMStaticRep                                // 31: 主机信息返回内容
	TypeHMDynamicReq                               // 32: 获取动态数据
	TypeHMDynamicRep                               // 33: 返回动态数据
	TypeHMQueryCollect                             // 34: 获取自动采集状态(req)
	TypeHMCollectStatus                            // 35: 返回自动采集状态(rep)
	TypeHMChangeCollectStatus                      // 36: 切换自动采集状态
	TypeHMReportAgentStatus                        // 37: agent状态上报数据
)

// install
const (
	TypeInstallArgs TypeName = iota + 40
	TypeInstallRep
)

// logging
const (
	TypeLoggingConfig TypeName = iota + 50
	TypeLoggingStatusReq
	TypeLoggingStatusRep
	TypeLoggingReport
)

// rpa
const (
	TypeRPARun TypeName = iota + 60
	TypeRPALog
	TypeRPAControlReq
	TypeRPAControlRep
	TypeRPAFinish
)

// agent log
const (
	TypeLogLsReq        TypeName = iota + 10000 // 10000：查询log文件列表
	TypeLogLsRep                                // 10001: 返回log文件列表
	TypeLogDownloadReq                          // 10002：请求log文件
	TypeLogDownloadInfo                         // 10003：返回log文件详情
	TypeLogDownloadData                         // 10004: 返回log文件内容
)

func (name TypeName) String() string {
	switch name {
	case TypeError:
		return "error"
	case TypeCome:
		return "come"
	case TypeAgentInfo:
		return "agent_info"
	case TypeFoo:
		return "foo"
	case TypeBar:
		return "bar"
	case TypeExec:
		return "exec"
	case TypeExecd:
		return "execd"
	case TypeExecData:
		return "exec_data"
	case TypeExecDone:
		return "exec_done"
	case TypeExecKill:
		return "exec_kill"
	case TypeLsReq:
		return "ls_req"
	case TypeLsRep:
		return "ls_rep"
	case TypeUpload:
		return "upload"
	case TypeUploadRep:
		return "upload_rep"
	case TypeDownloadReq:
		return "download_req"
	case TypeDownloadRep:
		return "download_rep"
	case TypeDownloadData:
		return "download_data"
	case TypeDownloadError:
		return "download_error"
	case TypeHMStaticReq:
		return "hm_static_req"
	case TypeHMStaticRep:
		return "hm_static_rep"
	case TypeHMDynamicReq:
		return "hm_dynamic_req"
	case TypeHMDynamicRep:
		return "hm_dynamic_rep"
	case TypeHMQueryCollect:
		return "hm_query_collect"
	case TypeHMCollectStatus:
		return "hm_collect_status"
	case TypeHMChangeCollectStatus:
		return "hm_change_collect_status"
	case TypeHMReportAgentStatus:
		return "hm_report_agent_status"
	case TypeLoggingConfig:
		return "logging_config"
	case TypeLoggingStatusReq:
		return "logging_status_req"
	case TypeLoggingStatusRep:
		return "logging_status_rep"
	case TypeLoggingReport:
		return "logging_report"
	case TypeLogLsReq:
		return "request_log_files"
	case TypeLogLsRep:
		return "response_log_files"
	case TypeLogDownloadReq:
		return "download_log_files"
	case TypeLogDownloadInfo:
		return "download_log_info"
	case TypeLogDownloadData:
		return "download_log_data"
	case TypeRPARun:
		return "rpa_run"
	case TypeRPALog:
		return "rpa_log"
	case TypeRPAControlReq:
		return "rpa_ctrl_req"
	case TypeRPAControlRep:
		return "rpa_ctrl_rep"
	case TypeRPAFinish:
		return "rpa_finish"
	default:
		return "unset"
	}
}
