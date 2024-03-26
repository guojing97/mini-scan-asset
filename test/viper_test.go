package test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	var config *viper.Viper = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("../")

	// read config
	err := config.ReadInConfig()

	assert.Nil(t, err)

	assert.Equal(t, "GO Scan", config.GetString("app.name"))
	assert.Equal(t, "1.0.0", config.GetString("app.version"))
	assert.Equal(t, "Kevin Hendrawan Hartono", config.GetString("app.author"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestENVFile(t *testing.T) {
	var config *viper.Viper = viper.New()
	config.SetConfigName("local")
	config.SetConfigType("env")
	config.AddConfigPath("../")

	// read config
	err := config.ReadInConfig()

	assert.Nil(t, err)

	assert.Equal(t, "go-scan", config.GetString("APP_NAME"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
	assert.Equal(t, "Kevin Hendrawan Hartono", config.GetString("APP_AUTHOR"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))
}

func TestENV(t *testing.T) {
	var config *viper.Viper = viper.New()
	config.SetConfigName("local")
	config.SetConfigType("env")
	config.AddConfigPath("../")
	config.AutomaticEnv()

	// read config
	err := config.ReadInConfig()

	assert.Nil(t, err)

	assert.Equal(t, "go-scan", config.GetString("APP_NAME"))
	assert.Equal(t, "1.0.0", config.GetString("APP_VERSION"))
	assert.Equal(t, "Kevin Hendrawan Hartono", config.GetString("APP_AUTHOR"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, config.GetInt("DATABASE_PORT"))

	// assert.Equal(t, "Hello", config.GetString("FROM_ENV"))
}
