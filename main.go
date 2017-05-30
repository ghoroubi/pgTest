package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"github.com/spf13/viper"
)

var schema = `
	CREATE TABLE IF NOT EXISTS users(
		ID  SERIAL PRIMARY KEY,
		UserName text,
		Password text,
		DisplayName text,
		Email text UNIQUE NOT NULL
		);`

var DB *sqlx.DB
func main() {
	config:=viper.New()
	config.SetConfigName("conf")
	config.AddConfigPath(".")
	config.SetConfigType("yml")
	var uc UserController
	var err error
	confErr:=config.ReadInConfig()
	if confErr!=nil {
		panic(confErr)
	}
	dbEng:=config.GetString("DataBase.Engin")
	router := gin.Default()
	DB, err = sqlx.Connect("postgres", "user=postgres password=123456 dbname=pgTest sslmode=disable")
	//dataSrc := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DBUser, DBPassword, DBName)
	//db, err := sqlx.Connect(Engin, dataSrc)
	if err != nil {
		log.Panic(err)
	}
	DB.MustExec(schema)
	router.POST("/CreateUser", uc.CreateUser)
	//router.GET("/Users")
	router.Run(":8585")

}
