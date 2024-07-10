package entity

type Manager struct {
	Name        string `db:"name" json:"name" binding:"required" example:"Andrew"`
	Managername string `db:"managername" json:"managername" binding:"required" example:"Manager"`
	Password    string `db:"password_hash" json:"password" binding:"required" example:"qwerty"`
	Role        string `db:"role" json:"-"`
}
