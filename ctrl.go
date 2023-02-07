package anet

import (
	"crypto/md5"
	"net"
	"time"
)

// ComePayload handshake request
type ComePayload struct {
	ID              string `json:"id"` // agent-id
	Name            string `json:"name"`
	Version         string `json:"version"`          // 版本号
	IP              net.IP `json:"ip"`               // ip地址
	MAC             string `json:"mac"`              // mac地址
	HostName        string `json:"host_name"`        // 主机名
	OS              string `json:"os"`               // 操作系统名称，windows或linux
	Platform        string `json:"platform"`         // 操作系统名称，debian、centos等
	PlatformVersion string `json:"platform_version"` // 操作系统版本号，如7.7.1908
	KernelVersion   string `json:"kernel_version"`   // 内核版本号，如3.10.0-1062.el7.x86_64
	Arch            string `json:"arch"`             // 操作系统位数，amd64、i386等
	CPU             string `json:"cpu"`              // CPU型号，如AMD Ryzen 7 4800U with Radeon Graphics
	CPUCore         uint64 `json:"cpu_core"`         // CPU核心数
	Memory          uint64 `json:"memory"`           // 内存大小
}

// HandshakePayload handshake response
type HandshakePayload struct {
	OK       bool     `json:"ok"`                 // 握手是否成功
	ID       string   `json:"id"`                 // 被分配的agent id
	Msg      string   `json:"msg"`                // 错误信息
	Redirect []string `json:"redirect,omitempty"` // 重定向地址
}

// LogFile log file info
type LogFile struct {
	Name    string    `json:"name"` // 文件名
	Size    uint64    `json:"size"` // 文件大小
	ModTime time.Time `json:"mod"`  // 修改时间
}

// LsLogPayload ls log response
type LsLogPayload struct {
	Files []LogFile `json:"files"` // 文件列表
}

// LogDownloadReq download log request
type LogDownloadReq struct {
	Files []string `json:"files"` // 请求文件列表
}

// LogDownloadInfo download file info
type LogDownloadInfo struct {
	OK        bool           `json:"ok"`                   // 是否成功
	ErrMsg    string         `json:"msg,omitempty"`        // 错误信息
	Size      uint64         `json:"size,omitempty"`       // 文件大小
	BlockSize uint64         `json:"block_size,omitempty"` // 分块大小
	MD5       [md5.Size]byte `json:"md5,omitempty"`        // 文件md5
}

// LogDownloadData download file data
type LogDownloadData struct {
	Offset uint64 `json:"offset"` // 块偏移量
	Data   string `json:"data"`   // 块内容
}
