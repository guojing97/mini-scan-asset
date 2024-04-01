package entity

type Parameter struct {
	ID            int32  `gorm:"column:id;primaryKey"`
	DeviceId      int32  `gorm:"column:device_id"`
	ParameterName string `gorm:"column:parameter_name"`
	Good          string `gorm:"column:good"`
	Bad           string `gorm:"column:bad"`
	Notes         string `gorm:"column:notes"`
}
