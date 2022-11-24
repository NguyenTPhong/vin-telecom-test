package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"
	"vinbigdata/config"
	"vinbigdata/database/migration"
	db2 "vinbigdata/package/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func GetTestDB() (*gorm.DB, func()) {
	err := godotenv.Load(basepath + "/../../.env")
	fmt.Println(err)
	DBPort := config.GetInt64("DB_PORT", 3000)
	url := fmt.Sprintf("host=127.0.0.1 port=%v user=dbUser password=dbPassword sslmode=disable", DBPort)
	fmt.Println(url)
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("can not connect to database", err.Error())
	}

	databaseName := fmt.Sprintf("telecom_test_%v", time.Now().Nanosecond())
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", databaseName))
	if err != nil {
		log.Fatal("can not create test database")
	}

	dbTest, connectErr := db2.NewDatabase(fmt.Sprintf("host=127.0.0.1 port=%v user=dbUser password=dbPassword dbname=%v sslmode=disable", DBPort, databaseName), 2, 1)
	if connectErr != nil {
		log.Fatal(errors.New("no db connection"))
	}

	migration.CreateTable(dbTest)

	return dbTest, func() {
		migration.DropTable(dbTest)
		db2.CloseDatabase(dbTest)
		_, e := db.Exec(fmt.Sprintf("DROP DATABASE %s;", databaseName))
		if e != nil {
			fmt.Println("fail to delete database", e)
		}
		db.Close()
	}
}
