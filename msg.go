package anet

// Msg message
type Msg struct {
	Type      TypeName    `json:"type"`             // 消息类型，详见消息类型章节
	Important bool        `json:"-"`                // 是否是关键消息
	TaskID    string      `json:"tid,omitempty"`    // 任务ID
	Plugin    *PluginInfo `json:"plugin,omitempty"` // 所需调用的插件信息
	ErrorMsg  string      `json:"errmsg,omitempty"` // 错误消息的详情
	// ctrl
	AgentInfo       *AgentInfo        `json:"ai,omitempty"`        // agent状态上报
	Come            *ComePayload      `json:"come,omitempty"`      // 握手请求
	Handshake       *HandshakePayload `json:"handshake,omitempty"` // 握手返回结果
	LsLog           *LsLogPayload     `json:"ls_log,omitempty"`    // 查询log文件结果
	LogDownload     *LogDownloadReq   `json:"log_dreq,omitempty"`  // 下载log文件
	LogDownloadInfo *LogDownloadInfo  `json:"log_dinfo,omitempty"` // 下载log文件信息
	LogDownloadData *LogDownloadData  `json:"log_ddata,omitempty"` // 下载log文件内容
	// exec plugin
	Exec     *ExecPayload  `json:"exec,omitempty"`      // 执行命令消息参数
	Execd    *ExecdPayload `json:"execd,omitempty"`     // 执行命令脚本或命令的启动结果，包括pid等
	ExecData *ExecData     `json:"exec_data,omitempty"` // 执行命令消息返回内容
	ExecDone *ExecDone     `json:"exec_done,omitempty"` // 执行命令返回结果，exit_code等
	ExecKill *ExecKill     `json:"exec_kill,omitempty"` // kill命令参数
	// file plugin
	LSReq         *LsReq         `json:"ls_req,omitempty"`         // ls请求参数
	LSRep         *LsRep         `json:"ls_rep,omitempty"`         // ls请求返回结果
	Upload        *Upload        `json:"upload,omitempty"`         // 文件上传时的文件基本信息，md5等
	UploadRep     *UploadRep     `json:"upload_rep,omitempty"`     // 文件创建结果
	DownloadReq   *DownloadReq   `json:"download_req,omitempty"`   // 文件下载请求，文件路径等
	DownloadRep   *DownloadRep   `json:"download_rep,omitempty"`   // 文件下载的基本信息，md5等
	DownloadData  *DownloadData  `json:"download_data,omitempty"`  // 文件内容
	DownloadError *DownloadError `json:"download_error,omitempty"` // 文件读取时的报错信息
	// host.monitor
	HMStatic        *HMStaticPayload       `json:"hm_static,omitempty"`        // 主机信息返回结果
	HMDynamicReq    *HMDynamicReq          `json:"hm_dynamic_req,omitempty"`   // 动态数据请求
	HMDynamicRep    *HMDynamicRep          `json:"hm_dynamic_rep,omitempty"`   // 动态数据响应
	HMCollectStatus *HMReportStatusPayload `json:"hm_status_rep,omitempty"`    // 自动采集状态返回
	HMChangeStatus  *HMChangeReportStatus  `json:"hm_change_status,omitempty"` // 修改自动采集状态
	HMAgentStatus   *HMAgentStatus         `json:"hm_agent_status,omitempty"`  // metrics-agent状态上报
	// install
	InstallArgs *InstallArgs `json:"inst_args,omitempty"` // 安装请求参数
	InstallRep  *InstallRep  `json:"inst_rep,omitempty"`  // 安装结果
	// logging
	LoggingConfig    *LoggingConfig    `json:"logging_config,omitempty"` // 创建任务
	LoggingStatusReq *LoggingStatusReq `json:"logging_req,omitempty"`    // 切换状态请求
	LoggingStatusRep *LoggingStatusRep `json:"logging_rep,omitempty"`    // 切换状态结果
	LoggingReport    *LoggingReport    `json:"logging_report,omitempty"` // 状态上报
	// rpa
	RPARun     *RPARunArgs `json:"rpa_run,omitempty"`      // RPA运行参数
	RPALog     *RPALogData `json:"rpa_log,omitempty"`      // RPA日志上报
	RPACtrlReq *RPACtrlReq `json:"rpa_ctrl_req,omitempty"` // RPA控制请求
	RPACtrlRep *RPACtrlRep `json:"rpa_ctrl_rep,omitempty"` // RPA控制结果
	RPAFinish  *RPAFinish  `json:"rpa_finish,omitempty"`   // RPA运行结束
}
