package user

import (
	"net/http"

	"redler/internal/pkg/lib"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB      *gorm.DB
	Service *UserService
}

// Register godoc
// @Summary      create a new user
// @Description  create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user	body      	RegisterDto  true  "Add user"
// @Success      201	{object}	lib.HttpResponseStruct
// @Failure      400 	{object} 	lib.HttpResponseStruct
// @Failure      500 	{object} 	lib.HttpResponseStruct
// @Router       /user	[post]
func (ctrl *UserController) Register(c *gin.Context) {
	dto := RegisterDto{}
	if c.ShouldBindJSON(&dto) != nil {
		HttpErrors[ParseBodyErr].Send(c)
		return
	}

	result, err := ctrl.Service.Register(&dto)
	if err != nil {
		HttpErrors[err.Code].Send(c)
		return
	}

	lib.HttpResponse(http.StatusCreated).Data(result).Send(c)
}

// GetUserByEmail godoc
// @Summary      get user by email
// @Description  get user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        email	path      	string  true  "User Email"
// @Success      200	{object}	lib.HttpResponseStruct
// @Failure      404 	{object} 	lib.HttpResponseStruct
// @Failure      500 	{object} 	lib.HttpResponseStruct
// @Router       /user/{email}		[get]
func (ctrl *UserController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	result, err := ctrl.Service.GetUserByEmail(email)
	if err != nil {
		HttpErrors[err.Code].Send(c)
		return
	}

	lib.HttpResponse(http.StatusOK).Data(result).Send(c)
}
