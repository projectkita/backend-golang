package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	. "github.com/projectkita/project-harapan-backend-golang/src/config"
)

type DBsql struct {
	ConfigMaster DBConfig
	ConfigSlave  DBConfig
	Master       *sql.DB
	Slave        *sql.DB
}

type DBConfig struct {
	Host             string
	Port             string
	Username         string
	Password         string
	DbName           string
	ConnectionString string
}

var DB *DBsql

func (db *DBsql) New(master, slave DBConfig) (*DBsql, error) {
	dbMaster, err := db.connect(master.Host, master.Port, master.Username, master.Password, master.DbName, master.ConnectionString)
	if err != nil {
		return nil, err
	}

	dbSlave, err := db.connect(slave.Host, slave.Port, slave.Username, slave.Password, slave.DbName, slave.ConnectionString)
	if err != nil {
		dbMaster.Close()
		return nil, err
	}

	db.ConfigMaster = master
	db.ConfigSlave = slave
	db.Master = dbMaster
	db.Slave = dbSlave

	return db, nil
}

func (db *DBsql) connect(host string, port string, username string, password string, dbName string, connectionString string) (*sql.DB, error) {

	dbTemp, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return dbTemp, nil
}

func Init(config AppConfig) error {

	DB = &DBsql{}

	pgMaster := DBConfig{
		Host:             config.AppDB.Master.Host,
		Port:             config.AppDB.Master.Port,
		Username:         config.AppDB.Master.Username,
		Password:         config.AppDB.Master.Password,
		DbName:           config.AppDB.Master.DbName,
		ConnectionString: config.AppDB.Master.ConnectionString,
	}

	pgSlave := DBConfig{
		Host:             config.AppDB.Slave.Host,
		Port:             config.AppDB.Slave.Port,
		Username:         config.AppDB.Slave.Username,
		Password:         config.AppDB.Slave.Password,
		DbName:           config.AppDB.Slave.DbName,
		ConnectionString: config.AppDB.Slave.ConnectionString,
	}

	_, err := DB.New(pgMaster, pgSlave)
	return err
}
