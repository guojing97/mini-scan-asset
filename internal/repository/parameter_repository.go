package repository

import (
	"kevinhend97/mini-scan-asset/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ParameterRepository struct {
	Repository[entity.Parameter]
	log *logrus.Logger
}

func NewParameterRepository(log *logrus.Logger) *ParameterRepository {
	return &ParameterRepository{
		log: log,
	}
}

func (r *ParameterRepository) FindByDeviceId(db *gorm.DB, deviceId int32) ([]entity.Parameter, error) {
	var parameters []entity.Parameter
	if err := db.Where("device_id = ?", deviceId).Find(&parameters).Error; err != nil {
		return nil, err
	}

	return parameters, nil
}
