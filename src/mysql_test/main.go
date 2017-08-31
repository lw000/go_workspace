// mysql_test project main.go
package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/thinkboy/log4go"
)

var (
	db *sql.DB

//	mutex *sync.RWMutex = new(sync.RWMutex)
)

func InitMysql() {
	var err error
	db, err = sql.Open("mysql", "lw:qazxsw123@/app_project")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Hour)

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	db, err = sql.Open("mysql", "lw:qazxsw123@/app_project")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	//	db.SetConnMaxLifetime(time.Hour)

	if err = db.Ping(); err != nil {
		panic(err)
	}

	{
		rows, err := db.Query("SELECT * FROM quotation")
		if err == nil {
			log.Debug("(%v)", rows)

			var name string
			var sale_name string
			var quotation_number string
			var create_time string
			var i int = 0
			for rows.Next() {
				err := rows.Scan(&name, &sale_name, &quotation_number, &create_time)
				if err != nil {
					log.Error(err)
				}
				i = i + 1
				log.Debug("[%d] %v, %v, %v, %v", i, sale_name, name, quotation_number, create_time)
			}

			rows.Close()
		}
	}

	{
		rows, err := db.Query("SELECT * FROM department_dictionaries")
		if err == nil {
			log.Debug("(%v)", rows)

			var department string
			var position string
			var position_type int
			var department_type int
			var i int = 0
			for rows.Next() {
				err := rows.Scan(&department, &position, &position_type, &department_type)
				if err != nil {
					log.Error(err)
				}
				i = i + 1
				log.Debug("[%d] %v, %v, %v, %v", i, department, position, position_type, department_type)
			}

			rows.Close()
		}
	}

	//	interrupt := make(chan os.Signal, 1)
	//	signal.Notify(interrupt, os.Interrupt)
	//	for {
	//		select {
	//		case <-interrupt:
	//			log.Debug("interrupt")
	//			return
	//		}
	//	}
}
