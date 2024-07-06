package model

import (
	"gorm.io/gorm"
)

type SegmentType uint

// 网段类型
const (
	// SegmentTypeRack 机架默认网段， 必须和机柜绑定，多个机柜可以关联同一网段
	SegmentTypeRack SegmentType =iota
	// SegmentTypeAdmin 带外网段，必须和机柜绑定，多个机柜可以关联同一网段
	SegmentTypeAdmin

	// SegmentTypeInner 业务机内网段
	SegmentTypeInner
	// SegmentTypeOuter 业务机外网段
	SegmentTypeOuter
	// SegmentTypePurchased 已采购未启用
	SegmentTypePurchased
)

type NetworkSegment struct {
	gorm.Model
	Remark  string `gorm:"column:remark;type:text" json:"remark"`
	Segment string `gorm:"column:segment;type:varchar(50);not null;" json:"segment"`
	Netmask string `gorm:"column:netmask;type:varchar(50);not null;" json:"netmask"`
	Gateway string `gorm:"column:gateway;type:varchar(50);not null;" json:"gateway"`
	Type SegmentType `gorm:"column:type;type:int(11);not null" json:"type" validate:"min=0,max=4,required"`

	// 网段配置的TOR设备
	DeviceID uint `gorm:"column:device_id;type:int(11);not null" json:"device_id"`
	Device Device `gorm:"foreignKey:DeviceID;references:ID" json:"device"`

	// 网段属性
	ISP string `gorm:"column:isp;type:varchar(50);not null" json:"isp"`
	VlanID uint `gorm:"column:vlan_id;type:int(11)" json:"vlan_id"`


	//	网段必须绑定数据中心
	IDCID uint `gorm:"column:idc_id;type:int(11);not null" json:"idc_id" validate:"required"`
	IDC IDC `gorm:"foreignKey:IDCID;references:ID" json:"idc"`
	//	强制分配的ip绑定到多个机柜
	CabinetID uint `gorm:"column:cabinet_id;type:int(11);not null" json:"cabinet_id" validate:"required"`
	Cabinet Cabinet `gorm:"foreignKey:CabinetID;references:ID" json:"cabinet"`
}


type IPType uint

const(
	// Main 主IP
	Main IPType =iota
	// Backup 备用IP
	Backup
	// Inner 内网IP
	Inner
	// Manage 管理网IP
	Manage
)

type IP struct {
	gorm.Model
	IP string `gorm:"column:ip;type:varchar(50);not null;uniqueIndex:ip" json:"ip"`
	NetworkSegmentID uint `gorm:"column:network_segment_id;type:int(11);not null" json:"network_segment_id"`
	NetworkSegment NetworkSegment `gorm:"foreignKey:NetworkSegmentID;references:ID" json:"network_segment"`

	IPType IPType `gorm:"column:ip_type;type:int(11);not null" json:"ip_type" validate:"min=0,max=3,required"`

	// 关联到设备上
	DeviceID uint `gorm:"column:device_id;type:int(11);not null" json:"device_id"`
	Device Device `gorm:"foreignKey:DeviceID;references:ID" json:"device"`
	Project string `gorm:"column:project;type:varchar(50);not null;uniqueIndex:ip_project" json:"project"`
}