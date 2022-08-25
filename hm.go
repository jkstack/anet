package anet

import "time"

type HMCore struct {
	Processor int32   `json:"processor,omitempty"` // 第几个核心
	Model     string  `json:"model,omitempty"`     // Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz
	Core      int32   `json:"core,omitempty"`      // 所在物理核上的编号
	Cores     int32   `json:"cores,omitempty"`     // 某块CPU上的编号
	Physical  int32   `json:"physical,omitempty"`  // 物理CPU编号
	Mhz       float64 `json:"mhz,omitempty"`       // CPU频率
}

type HMDisk struct {
	Model      string   `json:"model"` // 品牌型号
	Total      uint64   `json:"total"` // 容量
	Type       string   `json:"type"`  // hdd,fdd,odd
	Partitions []string `json:"parts"` // 逻辑分区
}

type HMPartition struct {
	Name   string   `json:"name"`             // linux为挂载路径如/run，windows为盘符如C:
	FSType string   `json:"fstype,omitempty"` // NTFS
	Opts   []string `json:"opts"`             // rw,nosuid,nodev
	Total  uint64   `json:"total"`            // 总容量
	INodes uint64   `json:"inodes"`           // inode数量
}

type HMInterface struct {
	Index   int      `json:"index,omitempty"` // 网卡下标
	Name    string   `json:"name,omitempty"`  // 网卡名称
	Mtu     int      `json:"mtu,omitempty"`   // 网卡mtu
	Flags   []string `json:"flags"`           // 网卡附加参数
	Mac     string   `json:"mac,omitempty"`   // 网卡mac地址
	Address []string `json:"addrs"`           // 网卡上绑定的IP地址列表
}

type HMUser struct {
	Name string `json:"name"` // 用户名
	ID   string `json:"id"`   // 用户ID
	GID  string `json:"gid"`  // 用户组ID
}

type HMStaticPayload struct {
	Time time.Time `json:"time"` // 上报时间
	Host struct {
		Name   string        `json:"name,omitempty"`   // 主机名
		UpTime time.Duration `json:"uptime,omitempty"` // 主机启动时长
	} `json:"host,omitempty"`
	OS struct {
		Name            string    `json:"name,omitempty"`     // linux、windows
		PlatformName    string    `json:"pname,omitempty"`    // debian、centos
		PlatformVersion string    `json:"pversion,omitempty"` // 7.7.1908
		Install         time.Time `json:"install,omitempty"`  // 安装时间
	} `json:"os,omitempty"`
	Kernel struct {
		Version string `json:"version,omitempty"` // 3.10.0-1062.el7.x86_64
		Arch    string `json:"arch,omitempty"`    // amd64、i386
	} `json:"kernel"`
	CPU struct {
		Physical int      `json:"physical,omitempty"` // 物理核心数
		Logical  int      `json:"logical,omitempty"`  // 逻辑核心数
		Socket   int      `json:"socket,omitempty"`   // CPU数量
		Cores    []HMCore `json:"cores"`              // 每个核心的参数
	} `json:"cpu"`
	Memory struct {
		Physical uint64 `json:"physical,omitempty"` // 物理内存大小
		Swap     uint64 `json:"swap,omitempty"`     // swap内存大小
	} `json:"memory"`
	Disks      []HMDisk      `json:"disks"`             // 物理磁盘列表
	Partitions []HMPartition `json:"parts"`             // 逻辑分区列表
	GateWay    string        `json:"gateway,omitempty"` // 网关地址
	Interface  []HMInterface `json:"interface"`         // 网卡列表
	User       []HMUser      `json:"user"`              // 用户列表
}

type hmDynamicReqType int

const (
	HMReqCPU hmDynamicReqType = iota
	HMReqMemory
	HMReqDisk
	HMReqProcess
)

type HMDynamicReq struct {
	Req []hmDynamicReqType `json:"req"` // 请求内容
}

type HMDynamicRep struct {
	CPU     *HMDynamicCPU      `json:"cpu,omitempty"`     // CPU
	Memory  *HMDynamicMemory   `json:"mem,omitempty"`     // 内存
	Disk    []HMDynamicDisk    `json:"disk,omitempty"`    // 磁盘
	Process []HMDynamicProcess `json:"process,omitempty"` // 进程
}

type HMDynamicCPU struct {
	Usage float64 `json:"usage"` // CPU使用率
}

type HMDynamicMemory struct {
	Used      uint64  `json:"used"`      // 已使用字节数
	Free      uint64  `json:"free"`      // 剩余字节数
	Available uint64  `json:"available"` // 可用字节数
	Total     uint64  `json:"total"`     // 总字节数
	Usage     float64 `json:"usage"`     // 内存使用率
	SwapUsed  uint64  `json:"sused"`     // swap已使用字节数
	SwapFree  uint64  `json:"sfree"`     // swap剩余字节数
	SwapTotal uint64  `json:"stotal"`    // swap总字节数
}

type HMDynamicDisk struct {
	Name  string  `json:"name,omitempty"` // linux为挂载路径如/run，windows为盘符如C:
	Used  uint64  `json:"used"`           // 已使用字节数
	Free  uint64  `json:"free"`           // 剩余字节数
	Total uint64  `json:"total"`          // 总字节数
	Usage float64 `json:"usage"`          // 内存使用率
}

type HMDynamicProcess struct {
	ID            int32    `json:"id"`               // 进程ID
	ParentID      int32    `json:"pid"`              // 父进程ID
	User          string   `json:"user,omitempty"`   // 用户
	CpuUsage      float64  `json:"cpu"`              // CPU使用率
	RssMemory     uint64   `json:"rss"`              // 物理内存数
	VirtualMemory uint64   `json:"vms"`              // 虚拟内存数
	SwapMemory    uint64   `json:"swap"`             // swap内存数
	MemoryUsage   float64  `json:"mem"`              // 内存使用率
	Cmd           []string `json:"cmd,omitempty"`    // 命令行
	Listen        []uint32 `json:"listen,omitempty"` // 监听端口
	Connections   int      `json:"conns"`            // 连接数
}
