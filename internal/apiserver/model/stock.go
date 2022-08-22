package model

// Stock defines the db table structure for report
type Stock struct {
	//ObjectMeta `json:"metadata:omitempty"`
	//ObjectMeta `json:"metadata"`
	ObjectMeta

	Code        string  `gorm:"column:code" json:"code"`
	Price       float32 `gorm:"column:price" json:"price"`
	Pe          float32 `gorm:"column:pe" json:"pe"`
	Description string  `gorm:"column:description" json:"description"`
}

// TableName maps to mysql table name.
func (r *Stock) TableName() string {
	return "stock"
}
