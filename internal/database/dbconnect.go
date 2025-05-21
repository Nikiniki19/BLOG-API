package database

import (
	"blog-api/internal/model"
	stdlog "log"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "host=ep-twilight-leaf-a4sw3hym.us-east-1.aws.neon.tech user=neondb_owner password=npg_FYHLDX4Qf3ns dbname=neondb port=5432 sslmode=require TimeZone=Asia/Shanghai"

	gormLogger := logger.New(
		stdlog.New(os.Stdout, "\r\n", stdlog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to the database")
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get sql instance")
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Error().Err(err).Msg("Connection to the db is closed")
		return nil, err
	}
	err = db.AutoMigrate(&model.BlogPost{})
	if err != nil {
		log.Error().Err(err).Msg("Unable to auto migrate the table")
		return nil, err
	}
	return db, nil

}
