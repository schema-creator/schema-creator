// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProjectHasUser = "project_has_users"

// ProjectHasUser mapped from table <project_has_users>
type ProjectHasUser struct {
	ProjectID string `gorm:"column:project_id;not null" json:"project_id"`
	UserID    string `gorm:"column:user_id;not null" json:"user_id"`
	Role      string `gorm:"column:role;not null" json:"role"`
}

// TableName ProjectHasUser's table name
func (*ProjectHasUser) TableName() string {
	return TableNameProjectHasUser
}
