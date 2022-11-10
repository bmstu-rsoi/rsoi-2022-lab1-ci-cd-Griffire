
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Database struct {
	Conn *sql.DB
}



var database = Database{}

//const HOST_PORT = "5000"
//const HOST_ADDRESS = "localhost"

//const HOST_URL = "http://" + HOST_ADDRESS + ":" + HOST_PORT + "/api/v1"
var HOST = "localhost"
var PORT = "8080"
var HOST_URL = ""
var PORTDB = "5432"

func main() {
	herokuPort, exist := os.LookupEnv("PORT")
	if exist {
		println("HERRRRROOOOOKKKUUUUU_PORT", herokuPort)
		PORT = herokuPort
	}
	herokuHOST, exist := os.LookupEnv("HOST")
	if exist {
		println("HERRRRROOOOOKKKUUUUU_HOST", herokuHOST)
		HOST = herokuHOST
	}
	fmt.Println(os.Getenv("PORT"), os.Getenv("HOST"), os.Getenv("DATABASE_URL"))

	//HOST_URL = "http://" + HOST_ADDRESS + ":" + PORT + "/api/v1"
	println("Begin??????")
	database, err := Initialize("program", "test", "persons")
	if err != nil {
		println("cant connect")
		println(err.Error())
	}
	database.comand_db(0)
	database.comand_db(1)
	//person1 := &Person{1, "ivan", 30, "piter", "proger"}
	//person3 := `{"id":1 , "name":"ivan","age": 30, "address":"piter", "work":"proger"}`
	//p1, _ := json.Marshal(&person1)
	//println(string(p1))
	//person2 := &Person{2, "ivan4", 190, "kazan", "tsar"}
	//err, n1 := database.AddItem(person1)
	//err, n2 := database.AddItem(person2)
	//println(n1, n2)
	//err = database.AddItem(*person3)

	defer database.Conn.Close()

	httpHandler := NewHandler1(&database)

	HOST = "0.0.0.0"
	println("HOST: " + HOST + "\nPORT: " + PORT)
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()

	//test1()
	//test2()
	//test3()
	//test3()
	//test3()
	//test4()
	//test1()
	//test5()
	//test1()
	//test6()

	defer Stop(server)
	log.Printf("Started server on %s:%s", HOST, PORT)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")

}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
