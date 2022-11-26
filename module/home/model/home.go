package hmmdl

import "bds/common"

const EntityName = "Home"

type Home struct {
	common.SQLModel `json:",inline"`
	HomeType        string         `json:"home_type" gorm:"column:home_type;"`
	BusinessType    string         `json:"business_type" gorm:"column:business_type;"`
	Area            float32        `json:"area" gorm:"column:area;"`
	Address         string         `json:"address" gorm:"column:address;"`
	Direction       string         `json:"direction" gorm:"column:direction;"`
	Facade          float32        `json:"facade" gorm:"column:facade;"`
	Price           float32        `json:"price" gorm:"column:price;"`
	TaxPaid         bool           `json:"tax_paid" gorm:"column:tax_paid;"`
	PK              string         `json:"pk" gorm:"column:pk;"`
	PL              string         `json:"pl" gorm:"column:pl;"`
	LegalStatus     string         `json:"legal_status" gorm:"column:legal_status;"`
	Contact         string         `json:"contact" gorm:"column:contact;"`
	Describe        string         `json:"describe" gorm:"column:describe;"`
	Thumbnail       *common.Image  `json:"logo" gorm:"column:thumbnail;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Home) TableName() string {
	return "home"
}

type HomeCreate struct {
	common.SQLModel `json:",inline"`
	HomeType        string         `json:"home_type" gorm:"column:home_type;"`
	BusinessType    string         `json:"business_type" gorm:"column:business_type;"`
	Area            float32        `json:"area" gorm:"column:area;"`
	Address         string         `json:"address" gorm:"column:address;"`
	Direction       string         `json:"direction" gorm:"column:direction;"`
	Facade          float32        `json:"facade" gorm:"column:facade;"`
	Price           float32        `json:"price" gorm:"column:price;"`
	TaxPaid         bool           `json:"tax_paid" gorm:"column:tax_paid;"`
	PK              string         `json:"pk" gorm:"column:pk;"`
	PL              string         `json:"pl" gorm:"column:pl;"`
	LegalStatus     string         `json:"legal_status" gorm:"column:legal_status;"`
	Contact         string         `json:"contact" gorm:"column:contact;"`
	Describe        string         `json:"describe" gorm:"column:describe;"`
	Thumbnail       *common.Image  `json:"logo" gorm:"column:thumbnail;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (data *HomeCreate) Validate() error {
	return nil
}

func (HomeCreate) TableName() string {
	return Home{}.TableName()
}

type HomeUpdate struct {
	common.SQLModel `json:",inline"`
	HomeType        string         `json:"home_type" gorm:"column:home_type;"`
	BusinessType    string         `json:"business_type" gorm:"column:business_type;"`
	Area            float32        `json:"area" gorm:"column:area;"`
	Address         string         `json:"address" gorm:"column:address;"`
	Direction       string         `json:"direction" gorm:"column:direction;"`
	Facade          float32        `json:"facade" gorm:"column:facade;"`
	Price           float32        `json:"price" gorm:"column:price;"`
	TaxPaid         bool           `json:"tax_paid" gorm:"column:tax_paid;"`
	PK              string         `json:"pk" gorm:"column:pk;"`
	PL              string         `json:"pl" gorm:"column:pl;"`
	LegalStatus     string         `json:"legal_status" gorm:"column:legal_status;"`
	Contact         string         `json:"contact" gorm:"column:contact;"`
	Describe        string         `json:"describe" gorm:"column:describe;"`
	Thumbnail       *common.Image  `json:"logo" gorm:"column:thumbnail;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (HomeUpdate) TableName() string {
	return Home{}.TableName()
}
