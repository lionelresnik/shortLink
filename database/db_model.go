package database

import "time"

type Keys struct {
	ID           uint      `gorm:"id,primaryKey,autoIncrement"`
	UpdatedAt    time.Time `gorm:"updated_at"`
	UserKey      string    `gorm:"user_key"`
	VisitCounter int       `gorm:"visit_counter"`
}

type Domains struct {
	ID     uint   `gorm:"id,primaryKey,autoIncrement"`
	Domain string `gorm:"domain"`
}
