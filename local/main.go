package main

import (
	"database/sql"
	"net/http"

	"github.com/importfmt/config"
	"github.com/importfmt/logger"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	frogConfig   *config.Config
	db           *sql.DB
	rdb          redis.Conn
	initMySQLErr error
	initRedisErr error
)

func main() {
	frogConfig = config.NewConfig()
	frogConfig.Init("/config/config.json")

	logger.Init(frogConfig.LogPath)
	logger.Info.Println("logger init")

	initTemplate(frogConfig.TemplatePath)

	db, initMySQLErr = sql.Open("mysql", frogConfig.MySQLUsername+":"+frogConfig.MySQLPassword+"@/"+frogConfig.MySQLDatabase)
	checkErr(initMySQLErr, "connectMySQLErr")
	defer db.Close()

	rdb, initRedisErr = redis.Dial("tcp", "localhost:6379")
	checkErr(initRedisErr, "connectRedisErr")
	defer rdb.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", initDynamic)
	mux.HandleFunc("/login", initLoginPage)
	mux.HandleFunc("/forgot", initForgotPage)
	mux.HandleFunc("/register", initRegisterPage)
	mux.HandleFunc("/console", initConsolePage)
	mux.HandleFunc("/upload", initUploadPage)
	mux.HandleFunc("/category", initCategoryPage)

	mux.HandleFunc("/test", initTestPage)

	mux.HandleFunc("/requestregister", register)
	mux.HandleFunc("/requestforgot", forgot)
	mux.HandleFunc("/requestlogin", login)
	mux.HandleFunc("/logout", logout)

	mux.HandleFunc("/upload-record", uploadRecord)
	mux.HandleFunc("/search-records-by-organism-name", searchRecordsByOrganismName)
	mux.HandleFunc("/search-records-by-category", searchRecordsByCategory)
	mux.HandleFunc("/search-record-by-record-id", searchRecordByRecordID)	

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(frogConfig.PublicPath))))
	mux.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir(frogConfig.ResourcePath))))

	err := http.ListenAndServe(":80", mux)
	checkErr(err, "ListenAndServe err")
}
