package user

import (
	"redler/internal/pkg/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(prefix string, rg *gin.RouterGroup, config *config.Config, db *gorm.DB) {
	router := rg.Group(prefix)

	userCtrl := UserController{
		DB: db,
		Service: &UserService{
			Config: config,
			Db:     db,
		},
	}

	router.POST("", userCtrl.Register)
	router.GET("/:email", userCtrl.GetUserByEmail)
}
