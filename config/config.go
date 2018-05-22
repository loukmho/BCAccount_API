package config

import (
	"encoding/json"
	"os"
	"fmt"
)

const MoneyPrecision  = 2

//TODO: เมื่อรันจริงต้องเปลี่ยนเป็น Docker Network Bridge IP เช่น 172.17.0.3 เป็นต้น
type Config struct {
	DBHost string `json:"db_host"`
	DBName string `json:"db_name"`
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
	DBPort string `json:"port"`
}

func LoadDSN(fileName string,dbType int) string {

	// dbType 0 = mySql , 1 = MsSql
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Err Open file %v: Error is:", file, err)
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	c := new(Config)
	err = decoder.Decode(&c)
	if err != nil {
		fmt.Println("error Decode Json:", err)
	}
	dsn := ""
	if dbType == 0 {
		dsn = c.DBUser + ":" + c.DBPass + "@" + c.DBHost + "/" + c.DBName + "?parseTime=true"
	}
	if dbType == 1 {
		dsn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", c.DBHost, c.DBUser, c.DBPass, c.DBPort, c.DBName)
		fmt.Println("DSN =", dsn)
	}

	return dsn
}