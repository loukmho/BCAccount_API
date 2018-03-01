package ctrl


import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/loukmho/bcaccount_api/config"
	"github.com/jmoiron/sqlx"
)

var HeaderKeys = make(map[string]interface{})

func setHeader() {

	HeaderKeys = map[string]interface{}{
		"Server":                       "smart_pump_invoice",
		"Host":                         "nopadol.net:6000",
		"Content_Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE",
		"Access-Control-Allow-Headers": "Origin, Content-Type, X-Auth-Token",
	}
}

var Dbc *sqlx.DB

func init(){
	Dbc = ConnectSQL()
}

func ConnectSQL()(msdb *sqlx.DB){
	//db_host := "192.168.0.7"
	//db_name := "bcnp"
	//db_host := "it11n"
	//db_name := "pc2017"
	//db_user := "sa"
	//db_pass := "[ibdkifu"
	//port := "1433"

	dsn := config.LoadDSN("config.json",1)//fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	msdb = sqlx.MustConnect("mssql", dsn)
	if msdb.Ping() != nil {
		fmt.Println("Error ")
	}

	return msdb
}



