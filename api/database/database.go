package api

import (
	"gomboc/api/models"
	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OpaqueDatabase struct {
	Database *gorm.DB `default:"null"`
}

var mutex = &sync.Mutex{}    // mutex lock to prevent collision
var instance *OpaqueDatabase // instance of this simple struct

func StartSQLiteDatabase(filepath string) {
	if instance == nil {
		panic("database must to be initialized")
	}

	if instance.Database == nil {
		log.Trace().Msgf("Start SQLite database %s", filepath)

		db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})

		if err != nil {
			log.Panic().Msgf("Unable to connect to database '%s': '%s'", filepath, err.Error())
			panic("unable to connect to database")
		}

		instance.Database = db
	}
}

func AutoMigrate() {
	instance.Database.AutoMigrate(&models.NodeModel{})
	instance.Database.AutoMigrate(&models.UserModel{})
}

func New() OpaqueDatabase {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()

		if instance == nil {
			instance = &OpaqueDatabase{Database: nil}
		}
	}

	return *instance
}
