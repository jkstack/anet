package anet

type SNMPReq struct {
	Host      string `json:"host"`          // 服务地址
	Community string `json:"community"`     // snmp community
	OID       string `json:"oid,omitempty"` // snmp oid
}

type SNMPItem struct {
	OID   string `json:"oid"`   // snmp oid
	Name  string `json:"name"`  // snmp name
	Value string `json:"value"` // snmp value
}

type SNMPRep struct {
	OK    bool       `json:"ok"`              // 是否成功
	Msg   string     `json:"msg,omitempty"`   // 错误信息
	Items []SNMPItem `json:"items,omitempty"` // snmp items
}
