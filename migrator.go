package migrations

import "gorm.io/gorm"

type Migrator interface {
	Id() string
	Up(db *gorm.DB)
	Down(db *gorm.DB)
}
