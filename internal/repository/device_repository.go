package repository

import (
	"kevinhend97/mini-scan-asset/internal/entity"

	"github.com/sirupsen/logrus"
)

type DeviceRepository struct {
	Repository[entity.Device]
	log *logrus.Logger
}

func NewDeviceRepository(log *logrus.Logger) *DeviceRepository {
	return &DeviceRepository{
		log: log,
	}
}
