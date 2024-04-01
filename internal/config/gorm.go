package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	username := viper.GetString("DATABASE_USERNAME")
	password := viper.GetString("DATABASE_PASSWORD")
	host := viper.GetString("DATABASE_HOST")
	port := viper.GetString("DATABASE_PORT")
	database := viper.GetString("DATABASE_NAME")
	idleConnection := viper.GetInt("DATABASE_POOL_IDLE")
	maxConnection := viper.GetInt("DATABASE_POOL_MAX")
	maxLifetimeConnection := viper.GetInt("DATABASE_POOL_LIFETIME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := db.DB()

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifetimeConnection))

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (log *logrusWriter) Printf(message string, args ...interface{}) {
	log.Logger.Tracef(message, args...)
}
