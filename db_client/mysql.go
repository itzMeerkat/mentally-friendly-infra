package db_client

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func InitDatabaseClient(masterDSN string, slaveSDNs []string, dbType string) *gorm.DB {
	if dbType == "mysql" {
		return InitMysqlClient(masterDSN, slaveSDNs)
	}
	panic("unknown database type")
}

func InitMysqlClient(masterDSN string, slavesDSN []string) *gorm.DB {
	master := mysql.Open(masterDSN)
	var slaves []gorm.Dialector
	for _,v := range slavesDSN {
		slaves = append(slaves, mysql.Open(v))
	}

	db, err := gorm.Open(master, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{master},
		Replicas: slaves,
		Policy: dbresolver.RandomPolicy{},
	}))
	if err != nil {
		panic(err)
	}
	return db
}