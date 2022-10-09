# gorm-migrations

### Simple library to migrate up or migrate down.

### Usage
```go
import "github.com/jqqjj/gorm-migrations"

type Demo struct{}
func (Demo) Id() string {
    return "demo"
}
func (Demo) Up(db *gorm.DB) {
}
func (Demo) Down(db *gorm.DB) {
}

type Demo2 struct{}
func (Demo2) Id() string {
    return "demo2"
}
func (Demo2) Up(db *gorm.DB) {
}
func (Demo2) Down(db *gorm.DB) {
}


migrations.AddMigrators(Demo{}, Demo2{})
if err := migrations.Migrate(gorm.DB); err != nil {
    log.Fatalln(err)
}
```

### API
```go
//add Migrators
AddMigrators(...Migrator) error
//migrate up all Migrators
Migrate(*gorm.DB) error
//migrate up one Migrator
MigrateUp(*gorm.DB) error
//migrate down one Migrator
MigrateDown(*gorm.DB) error
```
