package models

type User struct {
	ID          int    `json:"id" form:"id" db:"id"`
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	Email       string `json:"email" form:"email"`
	Nickname    string `json:"nickname" form:"nickname"`
	Phone       string `json:"phone" form:"phone"`
	Collections []Collection
	Histories   []History
}
