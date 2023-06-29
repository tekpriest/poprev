package database

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/utils"
)

type DatabaseConnection interface {
	GetDB() *gorm.DB
	GetRBD() *redis.Client
}

type databaseConnection struct {
	DB  *gorm.DB
	rDB *redis.Client
}

// GetRBD implements DatabaseConnection.
func (d *databaseConnection) GetRBD() *redis.Client {
	return d.rDB
}

// GetDB implements DatabaseConnection.
func (d *databaseConnection) GetDB() *gorm.DB {
	return d.DB
}

func NewConnection(config *config.Config) DatabaseConnection {
	db, err := gorm.Open(mysql.Open(config.DBURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	utils.PanicOnError(err, "there was an error connecting to the database")

	rDB := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Username: config.RedisUser,
		Password: config.RedisPass,
		DB:       config.RedisDB,
	})

	return &databaseConnection{
		DB:  db,
		rDB: rDB,
	}
}
