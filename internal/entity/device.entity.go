package entity

type Device struct {
	ID          int32  `gorm:"column:id;primaryKey"`
	NameDevice  string `gorm:"column:name_device"`
	Description string `gorm:"column:description"`
}
