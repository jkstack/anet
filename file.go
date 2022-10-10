package anet

import (
	"crypto/md5"
	"os"
	"time"
)

// LsReq ls request
type LsReq struct {
	Dir string `json:"dir"` // ls目录地址
}

// FileInfo file info for ls
type FileInfo struct {
	Name    string      `json:"name"`     // 文件名称
	Mod     os.FileMode `json:"mod"`      // 文件权限
	User    string      `json:"user"`     // 所属用户
	Group   string      `json:"group"`    // 所属组
	Size    uint64      `json:"size"`     // 文件大小
	ModTime time.Time   `json:"mod_time"` // 文件修改时间
	IsLink  bool        `json:"is_link"`  // 是否是软链接
	LinkDir string      `json:"link_dir"` // 连接路径
}

// LsRep ls response
type LsRep struct {
	Dir    string     `json:"dir"`             // 路径
	OK     bool       `json:"ok"`              // 是否成功
	ErrMsg string     `json:"msg,omitempty"`   // 出错时的信息
	Files  []FileInfo `json:"files,omitempty"` // 文件列表
}

// Upload upload
type Upload struct {
	Dir      string         `json:"dir"`                 // 路径
	Name     string         `json:"name"`                // 文件名
	Auth     string         `json:"auth,omitempty"`      // 是否提权，su、sudo或空
	User     string         `json:"user,omitempty"`      // 提权用户
	Pass     string         `json:"pass,omitempty"`      // 提权密码
	Mod      os.FileMode    `json:"mod"`                 // 文件权限
	OwnUser  string         `json:"own_user,omitempty"`  // 上传文件所属用户
	OwnGroup string         `json:"own_group,omitempty"` // 上传文件所属分组
	Size     uint64         `json:"size,omitempty"`      // 上传文件大小
	MD5      [md5.Size]byte `json:"md5"`                 // 文件md5
	MD5Check bool           `json:"md5_check"`           // 是否校验文件md5
	Data     string         `json:"data,omitempty"`      // 如果文件内容小于1M则直接传输
	Token    string         `json:"token,omitempty"`     // 大于1M时下载的token信息
	URI      string         `json:"uri,omitempty"`       // 从server下载时的uri
}

// UploadRep upload response
type UploadRep struct {
	Dir    string `json:"dir"`           // 路径
	OK     bool   `json:"ok"`            // 是否成功
	ErrMsg string `json:"msg,omitempty"` // 错误时的消息内容
}

// DownloadReq download request
type DownloadReq struct {
	Dir string `json:"dir"` // 文件路径
}

// DownloadRep download response
type DownloadRep struct {
	Dir       string         `json:"dir"`           // 文件路径
	OK        bool           `json:"ok"`            // 是否读取成功
	ErrMsg    string         `json:"msg,omitempty"` // 错误信息
	Size      uint64         `json:"size"`          // 文件大小
	BlockSize uint64         `json:"block_size"`    // 分块大小
	MD5       [md5.Size]byte `json:"md5"`           // 文件md5
}

// DownloadData download data
type DownloadData struct {
	Offset uint64 `json:"offset"` // 块偏移量
	Data   string `json:"data"`   // 块内容
}

type DownloadError struct {
	Msg string `json:"msg"`
}
