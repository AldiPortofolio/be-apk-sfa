package postgresrose

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Postgres some standard stuff
	"github.com/kelseyhightower/envconfig"
	ottologer "ottodigital.id/library/logger"
)

type DbEnv struct {
	DbUser    string `envconfig:"DB_ROSE_POSTGRES_USER" default:"ottoagcfg"`
	DbPass    string `envconfig:"DB_ROSE_POSTGRES_PASS" default:"dTj*&56$es"`
	DbName    string `envconfig:"DB_ROSE_POSTGRES_NAME" default:"rosedb"`
	DbAddres  string `envconfig:"DB_ROSE_POSTGRES_ADDRESS" default:"13.228.23.160"`
	DbPort    string `envconfig:"DB_ROSE_POSTGRES_PORT" default:"8432"`
	DbDebug   bool   `envconfig:"DB_ROSE_POSTGRES_DEBUG" default:"true"`
	DbType    string `envconfig:"DB_ROSE_POSTGRES_TYPE" default:"postgres"`
	SslMode   string `envconfig:"DB_ROSE_POSTGRES_SSL_MODE" default:"disable"`
	DbTimeout string `envconfig:"DB_ROSE_POSTGRES_TIMEOUT" default:"30"`
}

var (
	// Dbcon ..
	Dbcon *gorm.DB

	// Errdb ..
	Errdb error
	dbEnv DbEnv
)

// init ..
func init() {
	fmt.Println("DB POSTGRES ROSE")

	err := envconfig.Process("Database_ROSE", &dbEnv)
	if err != nil {
		fmt.Println("Failed to get Database ROSE env:", err)
	}

	if DbOpen() != nil {
		fmt.Println("Can't Open ", dbEnv.DbName, " DB", DbOpen())
	}
	Dbcon = GetDbCon()
	Dbcon = Dbcon.LogMode(true)
}

// DbOpen ..
func DbOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", dbEnv.DbAddres, dbEnv.DbPort, dbEnv.DbUser, dbEnv.DbPass, dbEnv.DbName, dbEnv.SslMode, dbEnv.DbTimeout)
	Dbcon, Errdb = gorm.Open("postgres", args)
	sugarLogger := ottologer.GetLogger()
	if Errdb != nil {
		sugarLogger.Error(fmt.Sprintf("open db Err ", Errdb))
		return Errdb
	}

	if errping := Dbcon.DB().Ping(); errping != nil {
		sugarLogger.Error(fmt.Sprintf("Db Not Connect test Ping :", errping))
		fmt.Println("Can't Open db Postgres")
		return errping
	}
	sugarLogger.Info("Connect DB success")
	return nil
}

// GetDbCon ..
func GetDbCon() *gorm.DB {
	//TODO looping try connection until timeout
	// using channel timeout
	sugarLogger := ottologer.GetLogger()
	if errping := Dbcon.DB().Ping(); errping != nil {
		sugarLogger.Error(fmt.Sprintf("Db Not Connect test Ping :", errping))
		//errping = nil
		if errping = DbOpen(); errping != nil {
			sugarLogger.Error(fmt.Sprintf("try to connect again but error :", errping))
		}
	}
	Dbcon.LogMode(true)
	return Dbcon
}
