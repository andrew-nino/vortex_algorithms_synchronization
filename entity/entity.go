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
	ClientName  string    `db:"client_name" json:"client_name" binding:"required"`
	Version     int       `db:"version" json:"version" binding:"required"`
	Image       string    `db:"image" json:"image" binding:"required"`
	CPU         string    `db:"cpu" json:"cpu" binding:"required"`
	Memory      string    `db:"memory" json:"memory" binding:"required"`
	Priority    float64   `db:"priority" json:"priority" binding:"required"`
	NeedRestart bool      `db:"need_restart" json:"need_restart"`
	SpawnedAt   string    `db:"spawned_at" json:"spawned_at" binding:"required"`
	CreatedAt   time.Time `db:"created_at" json:"-"`
	UpdatedAt   time.Time `db:"update_at" json:"-"`
}

type AlgorithmStatus struct {
	ClientID int64 `db:"client_id" json:"client_id" binding:"required"`
	VWAP     bool  `db:"vwap" json:"vwap"`
	TWAP     bool  `db:"twap" json:"twap"`
	HFT      bool  `db:"hft" json:"hft"`
}
