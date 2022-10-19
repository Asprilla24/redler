package user

type RegisterDto struct {
	Email         string   `json:"email" form:"email" binding:"required,email,max=100"`
	Name          string   `json:"name" form:"name" binding:"required,max=100"`
	ServicesOwned []string `json:"services_owned" form:"services_owned" binding:"required"`
}
