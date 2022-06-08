package main

import (
	"fmt"
	"go-sqlfile/models"

	"io/ioutil"
	"log"
	"net/http"
	"os"

	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Gorm postgres dialect interface
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	r := mux.NewRouter()

	if e != nil {
		// Be sure to remove any comments in the .env file
		// If needed use command "set .env"
		log.Fatalln("::::: Error loading .env file (main) :::::")
	}

	localhost := os.Getenv("DB_Host")
	port := os.Getenv("PORT")
	
	insertIntoDB()
	
	r.HandleFunc("/", greet).Methods("GET")

	log.Printf("..... Server running on %v:%v .....", localhost, port)
	log.Fatalln(http.ListenAndServe("localhost:"+port, r))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Greetings World!")
}

func connectDB() (*gorm.DB, error) {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatalln("::::: Error loading .env file (db) :::::")
	// }

	dbuser := os.Getenv("DB_User")
	dbpwd := os.Getenv("DB_Password")
	dbname := os.Getenv("DB_Name")
	dbhost := os.Getenv("DB_Host")

	// DB connection string
	dburi := fmt.Sprintf("host=%v user=%v dbname=%v sslmode=disable password=%v", dbhost, dbuser, dbname, dbpwd)

	// Connect
	db, err := gorm.Open("postgres", dburi)

	if err != nil {
		fmt.Println("::::: Error\n", err)
		panic(err)
	}

	log.Println("::::: Connected  to  Database :::::")
	return db, err
}

func insertIntoDB(){
	db, _ := connectDB()

	db.AutoMigrate(models.User{})

	// db.Exec("INSERT INTO users (id, username) VALUES (1, 'test_user5')")

	file, err := ioutil.ReadFile("bdd/testdata.sql")

	if err != nil {
		fmt.Println("::::: Error\n", err)
	}

	queries := strings.Split(string(file), ";")

	for _, q := range queries {
		// db.Raw(q)
		if q != ""{
			db.Exec(q)
			// fmt.Println("::::::", q)
		}
	}

}