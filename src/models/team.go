package models

type Team struct {
	Team_name string `json: "team_name" binding:"required"`
	Admin_email string `json: "admin_email" binding:"required"`
	Admin_password string `json: "admin_password" binding:"required"`
}
