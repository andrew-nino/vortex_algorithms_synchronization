package entity

import "time"

type Manager struct {
	Name        string `db:"name" json:"name" binding:"required" example:"Andrew"`
	Managername string `db:"managername" json:"managername" binding:"required" example:"Manager"`
	Password    string `db:"password_hash" json:"password" binding:"required" example:"qwerty"`
	Role        string `db:"role" json:"-"`
}

type Client struct {
	ID          int64     `db:"client_id" json:"client_id" binding:"required"`
	ClientName  string    `db:"client_name" json:"client_name"`
	Version     int       `db:"version" json:"version"`
	Image       string    `db:"image" json:"image"`
	CPU         string    `db:"cpu" json:"cpu"`
	Memory      string    `db:"memory" json:"memory"`
	Priority    float64   `db:"priority" json:"priority"`
	NeedRestart bool      `db:"needRestart" json:"needRestart"`
	SpawnedAt   string    `db:"spawned_at" json:"spawned_at"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
	UpdatedAt   time.Time `db:"update_at" json:"-"`
}
