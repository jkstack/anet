package anet

import "time"

type HMCore struct {
	Processor int32  `json:"processor,omitempty"` // 第几个核心
	Model     string `json:"model,omitempty"`     // Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz
	Core      int32  `json:"core,omitempty"`      // 所在物理核上的编号
	Cores     int32  `json:"cores,omitempty"`     // 某块CPU上的编号
	Physical  int32  `json:"physical,omitempty"`  // 物理CPU编号
}

type HMDisk struct {
	Name   string   `json:"name,omitempty"`   // linux为挂载路径如/run，windows为盘符如C:
	FSType string   `json:"fstype,omitempty"` // NTFS
	Opts   []string `json:"opts"`             // rw,nosuid,nodev
	Total  uint64   `json:"total,omitempty"`  // 总容量
	INodes uint64   `json:"inodes,omitempty"` // inode数量
}

type HMInterface struct {
	Index   int      `json:"index,omitempty"` // 网卡下标
	Name    string   `json:"name,omitempty"`  // 网卡名称
	Mtu     int      `json:"mtu,omitempty"`   // 网卡mtu
	Flags   []string `json:"flags"`           // 网卡附加参数
	Mac     string   `json:"mac,omitempty"`   // 网卡mac地址
	Address []string `json:"addrs"`           // 网卡上绑定的IP地址列表
}

type HMStaticPayload struct {
	Time time.Time `json:"time"` // 上报时间
	Host struct {
		Name   string        `json:"name,omitempty"`   // 主机名
		UpTime time.Duration `json:"uptime,omitempty"` // 主机启动时长
	} `json:"host,omitempty"`
	OS struct {
		Name            string `json:"name,omitempty"`     // linux、windows
		PlatformName    string `json:"pname,omitempty"`    // debian、centos
		PlatformVersion string `json:"pversion,omitempty"` // 7.7.1908
	} `json:"os,omitempty"`
	Kernel struct {
		Version string `json:"version,omitempty"` // 3.10.0-1062.el7.x86_64
		Arch    string `json:"arch,omitempty"`    // amd64、i386
	} `json:"kernel"`
	CPU struct {
		Physical int      `json:"physical,omitempty"` // 物理核心数
		Logical  int      `json:"logical,omitempty"`  // 逻辑核心数
		Cores    []HMCore `json:"cores"`              // 每个核心的参数
	} `json:"cpu"`
	Memory struct {
		Physical uint64 `json:"physical,omitempty"` // 物理内存大小
		Swap     uint64 `json:"swap,omitempty"`     // swap内存大小
	} `json:"memory"`
	Disks     []HMDisk      `json:"disks"`     // 磁盘列表
	Interface []HMInterface `json:"interface"` // 网卡列表
}
