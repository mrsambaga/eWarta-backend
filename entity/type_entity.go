package entity

type Type struct {
	Id   uint64 `gorm:"PrimaryKey"`
	Type string
}
