package anet

type IPMICommonRequest struct {
	Interface string `json:"interface"` // 连接方式：lan=v1.5, lanplus=v2.0, auto=lanplus first
	Addr      string `json:"addr"`      // IPMI地址
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
}

type IPMIDeviceInfo struct {
	OEM             string `json:"oem"`              // 生产厂商
	FirmwareVersion string `json:"firmware_version"` // 固件版本
	IPMIVersion     string `json:"ipmi_version"`     // IPMI版本
}

type sensorUnitValue struct {
	Percent bool   `json:"percent"`       // 是否是百分比
	Base    string `json:"base"`          // 基本单位
	Op      string `json:"op,omitempty"`  // 乘或除
	Mod     string `json:"mod,omitempty"` // 第二项单位
}

type sensorCritical struct {
	NonCritical    *float64 `json:"non_critical,omitempty"`    // 恢复数值
	Critical       *float64 `json:"critical,omitempty"`        // 告警数值
	NonRecoverable *float64 `json:"non_recoverable,omitempty"` // 严重告警数值
}

type sensorValue struct {
	Unit    sensorUnitValue `json:"unit"`            // 当前传感器的单位信息
	Current float64         `json:"current"`         // 当前数值
	Lower   *sensorCritical `json:"lower,omitempty"` // 最低告警数值
	Upper   *sensorCritical `json:"upper,omitempty"` // 最高告警数值
}

type IPMISensorInfo struct {
	ID       uint16       `json:"id"`               // 传感器序号
	SensorID uint8        `json:"sensor_id"`        // 传感器ID
	EntityID string       `json:"entity_id"`        // 实体ID
	Name     string       `json:"name"`             // 传感器名称
	Type     string       `json:"type"`             // 传感器类型
	Discrete bool         `json:"discrete"`         // 是否是离散传感器
	Values   *sensorValue `json:"values,omitempty"` // 传感器数值
}
