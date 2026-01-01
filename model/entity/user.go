package entity

import "time"

type User struct {
	ID int `json:"id" gorm:"column:id;primaryKey"`
	Email string `json:"email" gorm:"column:email;uniqueIndex"`
	Password  string    `json:"-" gorm:"column:password"`
	RoleID int `json:"role_id" gorm:"column:role_id;index"`
	Role      *Role     `json:"role,omitempty" gorm:"foreignKey:RoleID;references:ID"` // Relasi Belongs To (User milik satu Role)
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (u User) TableName() string {
	return "users"
}




/* buat bikin migration: migrate create -ext sql -dir db/migrations -seq create_users_table

Buat up:
migrate -path db/migrations -database "mysql://root:root@tcp(localhost:8889)/early_setup" up

Buat down:
migrate -path db/migrations -database "mysql://root:root@tcp(localhost:8889)/early_setup" down

fore kalo ada error - 1 nyesuaiin di log
migrate -path db/migrations -database "mysql://root:root@tcp(localhost:8889)/early_setup" force 1

*/





