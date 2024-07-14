package model

import "gorm.io/gorm"

// IDC 数据中心层
type IDC struct {
	gorm.Model
	Name   string `gorm:"column:name;type:varchar(50);not null" json:"name" validate:"required"`
	Code   string `gorm:"column:code;type:varchar(50);not null;unique" json:"code" validate:"required"`
	Remark string `gorm:"column:remark;type:text"`
	Locate string `gorm:"column:locate;type:varchar(500);not null" json:"locate" validate:"required"`
	// 供应商信息
	Vendor    string `gorm:"column:vendor;type:varchar(50)" json:"vendor"`
	Telephone string `gorm:"column:telephone;type:varchar(50)" json:"telephone"`
	Contact   string `gorm:"column:contact;type:varchar(50)" json:"contact"`
	Email     string `gorm:"column:email;type:varchar(50)" json:"email"`
	Website   string `gorm:"column:website;type:varchar(500)" json:"website"`
	OrderSite string `gorm:"column:order_site;type:varchar(500)" json:"order_site"`

	//	以下为关联信息
	Rooms []Room `gorm:"foreignKey:IDCID;references:ID" json:"-"`
}

// Room 机房层
type Room struct {
	gorm.Model
	IDCID  uint   `gorm:"column:idc_id;type:int(11);not null;UniqueIndex:code_idc_project" json:"idc_id" validate:"required"`
	IDC    IDC    `gorm:"foreignKey:IDCID;references:ID" json:"idc"`
	Name   string `gorm:"column:name;type:varchar(50);not null" json:"name" validate:"required"`
	Code   string `gorm:"column:code;type:varchar(50);not null;UniqueIndex:code_idc_project" json:"code" validate:"required"`
	Remark string `gorm:"column:remark;type:text" json:"remark"`
	//	以下为关联信息
	Cabinet []Cabinet `gorm:"foreignKey:RoomID;references:ID" json:"cabinet"`
}

// Cabinet 机柜层
type Cabinet struct {
	gorm.Model
	RoomID uint   `gorm:"column:room_id;type:int(11);not null;UniqueIndex:code_room_project" json:"room_id" validate:"required"`
	Room   Room   `gorm:"foreignKey:RoomID;references:ID" json:"room"`
	Name   string `gorm:"column:name;type:varchar(50);not null" json:"name" validate:"required"`
	Code   string `gorm:"column:code;type:varchar(50);not null;UniqueIndex:code_room_project" json:"code" validate:"required"`
	Remark string `gorm:"column:remark;type:text" json:"remark"`
	Rack   []Rack `gorm:"foreignKey:CabinetID;references:ID" json:"rack"`
}

// Rack 机架层
type Rack struct {
	gorm.Model
	CabinetID uint    `gorm:"column:cabinet_id;type:int(11);not null;UniqueIndex:code_cabinet_project" json:"cabinet_id" validate:"required"`
	Cabinet   Cabinet `gorm:"foreignKey:CabinetID;references:ID" json:"cabinet"`
	Name      string  `gorm:"column:name;type:varchar(50);not null" json:"name" validate:"required"`
	Code      string  `gorm:"column:code;type:varchar(50);not null;UniqueIndex:code_cabinet_project" json:"code" validate:"required"`
	Remark    string  `gorm:"column:remark;type:text" json:"remark"`
	// 初始化网络资源
	InternalIP string `gorm:"column:internal_ip;type:varchar(50);not null" json:"internal_ip"`
	ManagerIP  string `gorm:"column:manager_ip;type:varchar(50);not null" json:"manager_ip"`
}
