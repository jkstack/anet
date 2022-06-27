package anet

import "crypto/md5"

// PluginInfo plugin info
type PluginInfo struct {
	Name    string         `json:"name"`    // 插件名称
	Version string         `json:"version"` // 插件版本号
	MD5     [md5.Size]byte `json:"md5"`     // 插件文件的md5
	URI     string         `json:"uri"`     // 插件下载的uri
}
