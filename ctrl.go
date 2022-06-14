package anet

import (
	"net"
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
