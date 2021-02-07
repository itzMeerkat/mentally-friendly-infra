package mysql_client

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func InitMysqlClient(masterDSN string, slavesDSN []string) (*gorm.DB, error) {
	master := mysql.Open(masterDSN)
	var slaves []gorm.Dialector
	for _,v := range slavesDSN {
		slaves = append(slaves, mysql.Open(v))
	}

	db, err := gorm.Open(master, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{master},
		Replicas: slaves,
		Policy: dbresolver.RandomPolicy{},
	}))
	if err != nil {
		return nil, err
	}
	return db, nil
}