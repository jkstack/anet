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
	TypeHMStaticReq TypeName = iota + 30 // 30：获取主机信息请求
	TypeHMStaticRep                      // 31：主机信息返回内容
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
	case TypeLoggingConfig:
		return "logging_config"
	case TypeLoggingStatusReq:
		return "logging_status_req"
	case TypeLoggingStatusRep:
		return "logging_status_rep"
	case TypeLoggingReport:
		return "logging_report"
	default:
		return "unset"
	}
}
