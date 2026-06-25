package models

type User struct {
	Id    int    `json:"id" gorm:"column:user_id;primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
