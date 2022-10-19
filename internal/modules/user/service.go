package user

import (
	"redler/internal/pkg/config"
	"redler/internal/pkg/lib"
	"redler/models"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var module = "user_service"

type UserService struct {
	Config *config.Config
	Db     *gorm.DB
}

func (service *UserService) Register(registerDto *RegisterDto) (result lib.H, e *lib.ServiceError) {
	user := models.UserModel{
		Email:         registerDto.Email,
		Name:          registerDto.Name,
		ServicesOwned: pq.StringArray(registerDto.ServicesOwned),
	}

	if created := service.Db.Create(&user); created.Error != nil {
		e = lib.Error(EmailExistsErr)

		logrus.WithFields(logrus.Fields{"module": module}).
			Error(created.Error)
	}

	result = user.Transform()
	return
}

func (service *UserService) GetUserByEmail(email string) (result lib.H, e *lib.ServiceError) {
	user := models.UserModel{}

	err := service.Db.Where("email=?", email).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		e = lib.Error(UserNotFoundErr)

		logrus.WithFields(logrus.Fields{"module": module}).
			Error(err)
	}

	if err != nil {
		e = lib.Error(SomethingWentWrongErr)

		logrus.WithFields(logrus.Fields{"module": module}).
			Error(err)
	}

	result = user.Transform()
	return
}
