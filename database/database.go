package database

import (
	"github.com/librespeed/speedtest/config"
	"github.com/librespeed/speedtest/database/bolt"
	"github.com/librespeed/speedtest/database/memory"
	"github.com/librespeed/speedtest/database/mysql"
	"github.com/librespeed/speedtest/database/none"
	"github.com/librespeed/speedtest/database/postgresql"
	"github.com/librespeed/speedtest/database/schema"

	log "github.com/sirupsen/logrus"
)

var (
	DB DataAccess
)

type DataAccess interface {
	Insert(*schema.TelemetryData) error
	FetchByUUID(string) (*schema.TelemetryData, error)
	FetchLast100() ([]schema.TelemetryData, error)
}

func SetDBInfo(conf *config.Config) {
    switch conf.DatabaseType {
    case "postgresql":
	log.Infof("Connecting to PostgreSQL database at %s as user %s to database '%s'", 
	    conf.DatabaseHostname,
	    conf.DatabaseUsername,
	    conf.DatabaseName)
	DB = postgresql.Open(conf.DatabaseHostname, conf.DatabaseUsername, conf.DatabasePassword, conf.DatabaseName)

    case "mysql":
	log.Infof("Connecting to MySQL database at %s as user %s to database '%s'", 
	    conf.DatabaseHostname,
	    conf.DatabaseUsername,
	    conf.DatabaseName)
	DB = mysql.Open(conf.DatabaseHostname, conf.DatabaseUsername, conf.DatabasePassword, conf.DatabaseName)

    case "bolt":
	log.Infof("Opening BoltDB file at %s", conf.DatabaseFile)
	DB = bolt.Open(conf.DatabaseFile)

    case "memory":
	log.Info("Using in-memory database")
	DB = memory.Open("")

    case "none":
	log.Info("Database functionality disabled (none)")
	DB = none.Open("")

    default:
	log.Fatalf("Unsupported database type: %s", conf.DatabaseType)
    }

    if DB != nil {
	log.Infof("Successfully connected to %s database '%s'", 
	    conf.DatabaseType,
	    conf.DatabaseName)
    } else {
	log.Errorf("Failed to connect to %s database '%s'", 
	    conf.DatabaseType,
	    conf.DatabaseName)
    }
}
