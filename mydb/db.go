package mydb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func CreateTableSQL() {
	// 連接到 MySQL 資料庫
	DB, err := sql.Open("mysql", "ads_user:ads_password@tcp(localhost:3306)/ads_database")
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()

	// 創建 ad 資料表的 SQL 語句
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS ad (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		startAt DATETIME,
		endAt DATETIME,
		ageStart INT,
		ageEnd INT,
		gender SET('M', 'F'),
		country VARCHAR(255),
		platform SET('android', 'ios', 'web')
	);
	`

	// 執行 SQL 命令以創建資料表
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("ad 資料表創建成功！")
}

// 初始化資料庫連接
func InitializeDB() error {
	var err error
	DB, err = sql.Open("mysql", "ads_user:ads_password@tcp(localhost:3306)/ads_database")
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Database connection established")
	return nil
}
