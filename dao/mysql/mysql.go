package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/youta?charset=utf8&parseTime=true&loc=Local"
	db, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info), //打印sql
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			//DryRun: true,
			DisableAutomaticPing:                     true,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		log.Fatalf("database error: %v", err)
	}
}

func NewDB() *gorm.DB {
	return db
}
