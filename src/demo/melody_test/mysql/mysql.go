package main

import (
	"time"
	//	_ "github.com/go-sql-driver/mysql"
	log "github.com/thinkboy/log4go"
)

var (
//	Db *sql.DB

//	mutex *sync.RWMutex = new(sync.RWMutex)
)

//func InitMysql() {
//	var err error

//	//sqlauth := config.Conf.MySqlAccount + ":" + config.Conf.MySqlPassWord + "@tcp(" + config.Conf.MysqlAddress + ":" + config.Conf.MySqlPort + ")/" + config.Conf.MysqlRedbagDataBase + "?charset=utf8mb4"

//	Db, err = sql.Open("mysql", "lw:qazxsw123@/app_project")
//	if err != nil {
//		panic(err)
//	}
//	Db.SetMaxOpenConns(10)
//	Db.SetMaxIdleConns(2)
//	Db.SetConnMaxLifetime(time.Hour)

//	if err = Db.Ping(); err != nil {
//		panic(err)
//	}

//	log.Debug("1")
//}

func main() {
	defer time.Sleep(time.Microsecond * 10)

	log.Debug("mysql test")
}
