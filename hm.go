package anet

import (
	"time"

	"github.com/jkstack/jkframe/utils"
)

type HMCore struct {
	Processor int32           `json:"processor,omitempty"` // 第几个核心
	Model     string          `json:"model,omitempty"`     // Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz
	Core      int32           `json:"core,omitempty"`      // 所在物理核上的编号
	Cores     int32           `json:"cores,omitempty"`     // 某块CPU上的编号
	Physical  int32           `json:"physical,omitempty"`  // 物理CPU编号
	Mhz       utils.Float64P2 `json:"mhz,omitempty"`       // CPU频率
}

type HMDisk struct {
	Model      string   `json:"model,omitempty"` // 品牌型号
	Total      uint64   `json:"total,omitempty"` // 容量
	Type       string   `json:"type,omitempty"`  // hdd,fdd,odd
	Partitions []string `json:"parts,omitempty"` // 逻辑分区
}

type HMPartition struct {
	Name   string   `json:"name,omitempty"`   // linux为挂载路径如/run，windows为盘符如C:
	FSType string   `json:"fstype,omitempty"` // NTFS
	Opts   []string `json:"opts,omitempty"`   // rw,nosuid,nodev
	Total  uint64   `json:"total,omitempty"`  // 总容量
	INodes uint64   `json:"inodes,omitempty"` // inode数量
}

type HMInterface struct {
	Index   int      `json:"index,omitempty"` // 网卡下标
	Name    string   `json:"name,omitempty"`  // 网卡名称
	Mtu     int      `json:"mtu,omitempty"`   // 网卡mtu
	Flags   []string `json:"flags,omitempty"` // 网卡附加参数
	Mac     string   `json:"mac,omitempty"`   // 网卡mac地址
	Address []string `json:"addrs,omitempty"` // 网卡上绑定的IP地址列表
}

type HMUser struct {
	Name string `json:"name,omitempty"` // 用户名
	ID   string `json:"id,omitempty"`   // 用户ID
	GID  string `json:"gid,omitempty"`  // 用户组ID
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
		Startup         time.Time `json:"startup,omitempty"`  // 启动时间
	} `json:"os,omitempty"`
	Kernel struct {
		Version string `json:"version,omitempty"` // 3.10.0-1062.el7.x86_64
		Arch    string `json:"arch,omitempty"`    // amd64、i386
	} `json:"kernel,omitempty"`
	CPU struct {
		Physical int      `json:"physical,omitempty"` // 物理核心数
		Logical  int      `json:"logical,omitempty"`  // 逻辑核心数
		Cores    []HMCore `json:"cores,omitempty"`    // 每个核心的参数
	} `json:"cpu,omitempty"`
	Memory struct {
		Physical uint64 `json:"physical,omitempty"` // 物理内存大小
		Swap     uint64 `json:"swap,omitempty"`     // swap内存大小
	} `json:"memory,omitempty"`
	Disks      []HMDisk      `json:"disks,omitempty"`     // 物理磁盘列表
	Partitions []HMPartition `json:"parts,omitempty"`     // 逻辑分区列表
	GateWay    string        `json:"gateway,omitempty"`   // 网关地址
	Interface  []HMInterface `json:"interface,omitempty"` // 网卡列表
	User       []HMUser      `json:"user,omitempty"`      // 用户列表
}

type HMDynamicReqType int

const (
	HMReqUsage HMDynamicReqType = iota
	HMReqProcess
	HMReqConnections
)

type HMDynamicReq struct {
	Req []HMDynamicReqType `json:"req"` // 请求内容
}

type HMDynamicRep struct {
	Usage       *HMDynamicUsage       `json:"usage,omitempty"`   // CPU
	Process     []HMDynamicProcess    `json:"process,omitempty"` // 进程
	Connections []HMDynamicConnection `json:"conns,omitempty"`   // 连接
}

type HMDynamicUsage struct {
	Cpu struct {
		Usage utils.Float64P2 `json:"usage,omitempty"` // CPU使用率
	} `json:"cpu"`
	Memory struct {
		Used      uint64          `json:"used,omitempty"`      // 已使用字节数
		Free      uint64          `json:"free,omitempty"`      // 剩余字节数
		Available uint64          `json:"available,omitempty"` // 可用字节数
		Total     uint64          `json:"total,omitempty"`     // 总字节数
		Usage     utils.Float64P2 `json:"usage,omitempty"`     // 内存使用率
		SwapUsed  uint64          `json:"sused,omitempty"`     // swap已使用字节数
		SwapFree  uint64          `json:"sfree,omitempty"`     // swap剩余字节数
		SwapTotal uint64          `json:"stotal,omitempty"`    // swap总字节数
	} `json:"mem"`
	Partitions []HMDynamicPartition `json:"parts,omitempty"`      // 分区
	Interface  []HMDynamicInterface `json:"intferface,omitempty"` // 网卡
}

type HMDynamicPartition struct {
	Name       string          `json:"name,omitempty"`   // linux为挂载路径如/run，windows为盘符如C:
	Used       uint64          `json:"used,omitempty"`   // 已使用字节数
	Free       uint64          `json:"free,omitempty"`   // 剩余字节数
	Usage      utils.Float64P2 `json:"usage,omitempty"`  // 磁盘使用率
	InodeUsed  uint64          `json:"iused,omitempty"`  // inode使用数量
	InodeFree  uint64          `json:"ifree,omitempty"`  // inode剩余数量
	InodeUsage utils.Float64P2 `json:"iusage,omitempty"` // inode使用率
}

type HMDynamicInterface struct {
	Name        string `json:"name,omitempty"`  // 网卡名称
	BytesSent   uint64 `json:"bsent,omitempty"` // 发送字节数
	BytesRecv   uint64 `json:"brecv,omitempty"` // 接收字节数
	PacketsSent uint64 `json:"psent,omitempty"` // 发送数据包数量
	PacketsRecv uint64 `json:"precv,omitempty"` // 接收数据包数量
}

type HMDynamicConnection struct {
	Fd     uint32 `json:"fd,omitempty"`     // 句柄号
	Type   string `json:"type,omitempty"`   // 连接类型
	Local  string `json:"local,omitempty"`  // 本地地址
	Remote string `json:"remote,omitempty"` // 远程地址
	Status string `json:"status,omitempty"` // 连接状态
}

type HMDynamicProcess struct {
	ID            int32           `json:"id,omitempty"`     // 进程ID
	ParentID      int32           `json:"pid,omitempty"`    // 父进程ID
	User          string          `json:"user,omitempty"`   // 用户
	CpuUsage      utils.Float64P2 `json:"cpu,omitempty"`    // CPU使用率
	RssMemory     uint64          `json:"rss,omitempty"`    // 物理内存数
	VirtualMemory uint64          `json:"vms,omitempty"`    // 虚拟内存数
	SwapMemory    uint64          `json:"swap,omitempty"`   // swap内存数
	MemoryUsage   utils.Float64P2 `json:"mem,omitempty"`    // 内存使用率
	Cmd           []string        `json:"cmd,omitempty"`    // 命令行
	Listen        []uint32        `json:"listen,omitempty"` // 监听端口
	Connections   int             `json:"conns,omitempty"`  // 连接数
}
