package entity

import "time"

type Role struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey"`
    Name      string    `json:"name" gorm:"column:name;uniqueIndex"`
    Users     []User    `json:"users,omitempty" gorm:"foreignKey:RoleID"` // Relasi Has Many (Satu Role punya banyak User)
    CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (r Role) TableName() string {
    return "roles"
}