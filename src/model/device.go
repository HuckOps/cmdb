package model

type DeviceType uint

const (
	// Server 物理机
	Server DeviceType = iota
	// VirtualServer 虚拟机
	VirtualServer
	// Router 路由器
	Router
	// TOR TOR交换机
	TOR
	// Firewall 防火墙
	Firewall
	// OtherDevice 其他网络设备（如F5等）
	OtherDevice
)

// 设备表，存机器和交换机等设备，虚拟机元数据也存在这

type Device struct {
	SN string `gorm:"column:sn;type:varchar(50);" json:"sn" validate:"required"`
	Asset string `gorm:"column:asset_id;type:varchar(50);" json:"asset_id" validate:"required"`
	Hostname string `gorm:"column:hostname;type:varchar(50);unique" json:"hostname" validate:"required"`

	// 仅对物理机和TOR交换机有效
	RackID uint `gorm:"column:rack_id;type:int(11);not null" json:"rack_id" validate:"required"`
	Rack Rack `gorm:"foreignKey:RackID;references:ID" json:"rack"`
	IPS []IP `gorm:"foreignKey:DeviceID;references:ID" json:"ips"`
}
