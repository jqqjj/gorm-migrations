package migrations

type Model struct {
	Id        int    `gorm:"primaryKey;autoIncrement"`
	Migration string `gorm:"not null"`
	Batch     int    `gorm:"not null"`
}

func (Model) TableName() string {
	return "migrations"
}
