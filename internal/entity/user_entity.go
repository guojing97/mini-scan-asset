package entity

type User struct {
	ID       int32  `gorm:"column:id;primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Token    string `gorm:"column:token"`
}
