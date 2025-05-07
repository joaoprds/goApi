package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConectionDataBase = ""
	ApiPortRun              = 0
)

func LoadEnv() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiPortRun, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiPortRun = 9000
	}

	StringConectionDataBase = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	//fmt.Println("String de conexão com o banco de dados: ", StringConectionDataBase)
}
