package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; comment: 'Unique identifier for a role';"`
	RoleName string `gorm:"column:role_name; unique; comment: 'Role name'"`
	RoleNote string `gorm:"column:role_note; type: text; comment: 'Role note'"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
