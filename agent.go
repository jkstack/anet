package anet

type AgentInfo struct {
	Version          string             `json:"v"`   // 版本号
	GoVersion        string             `json:"gv"`  // go版本号
	CpuUsage         float32            `json:"cu"`  // cpu使用率（百分比）
	MemoryUsage      float32            `json:"mu"`  // 内存使用率（百分比）
	Threads          int                `json:"ts"`  // 线程数
	Routines         int                `json:"rs"`  // 协程数
	Startup          int64              `json:"st"`  // 启动时间戳
	HeapInuse        uint64             `json:"hi"`  // 堆内存用量
	GC               map[string]float64 `json:"gc"`  // gc信息
	InPackets        uint64             `json:"ip"`  // 接收数据包数量
	InBytes          uint64             `json:"ib"`  // 接收数据字节数
	OutPackets       uint64             `json:"op"`  // 发送数据包数量
	OutBytes         uint64             `json:"ob"`  // 发送数据字节数
	PluginExecd      uint64             `json:"pe"`  // 已运行插件数
	PluginRunning    uint64             `json:"pr"`  // 正在运行插件数量
	ReconnectCount   int                `json:"rcc"` // 重连次数
	PluginUseCount   map[string]uint64  `json:"puc"` // 每个插件使用次数
	PluginOutPackets map[string]uint64  `json:"pop"` // 每个插件发送数据包数量
	PluginOutBytes   map[string]uint64  `json:"pob"` // 每个插件发送字节数
	ReadChanSize     int                `json:"rcs"` // 接收队列长度
	WriteChanSize    int                `json:"wcs"` // 发送队列长度
}
