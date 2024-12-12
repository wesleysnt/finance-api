package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/wesleysnt/go-base/app/facades"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbCon struct {
	username, password, db, host, port string
}

var (
	dbInstance *gorm.DB
	err        error
)

func InitDb(env Database) (db *gorm.DB, err error) {
	dbCon := dbCon{
		username: env.Username,
		password: env.Password,
		db:       env.Database,
		host:     env.Host,
		port:     string(env.Port),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbCon.host, dbCon.username, dbCon.password, dbCon.db, dbCon.port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return
}

func ConnectDB(env Database) {
	if dbInstance == nil {
		dbInstance, err = InitDb(env)
	}

	if err != nil {
		panic("Failed to connect to database")
	}
	sqlDB, err := dbInstance.DB()

	if err != nil {
		panic("Failed to connect to database")
	}

	setPool(sqlDB)

	facades.MakeOrm(dbInstance)
}

func setPool(sqlDB *sql.DB) {
	var idleCons int64 = 2
	var openCons int64 = 10

	if os.Getenv("DB_MAX_OPEN_CONNS") != "" {
		openCons, _ = strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONNS"), 10, 32)
	}
	if os.Getenv("DB_MAX_IDLE_CONNS") != "" {
		idleCons, _ = strconv.ParseInt(os.Getenv("DB_MAX_IDLE_CONNS"), 10, 32)
	}

	sqlDB.SetMaxIdleConns(int(idleCons))
	sqlDB.SetMaxOpenConns(int(openCons))
	sqlDB.SetConnMaxLifetime(time.Hour)

}
