//package main
//
//import (
//	"database/sql"
//	_ "github.com/lib/pq"
//)
//
//const (
//	HOST = "localhost"
//	PORT = 5432
//)
//
//type Database struct {
//	Conn *sql.DB
//}
//
//
//
//func main() {
//	db, err := Initialize("program", "test", "persons")
//	if err != nil {
//		println(err.Error())
//		return
//	}
//	//db.comand_db(2)
//	//db.comand_db(3)
//	db.comand_db(0)
//	db.comand_db(1)
//	person1 := &Person{1, "ivan", 30, "piter", "proger"}
//	person2 := &Person{2, "ivan4", 190, "kazan", "tsar"}
//	err = db.AddItem(*person1)
//	err = db.AddItem(*person2)
//	if err != nil {
//		println(err.Error())
//	}
//
//	l, err := db.GetAllItems()
//	println(l.Persons[0].to_String())
//
//	db.DeleteItem(2)
//	l, err = db.GetAllItems()
//	println("\n")
//	println(l.Persons[0].to_String())
//
//	(person1).age = -3
//	db.UpdateItem(1, *person2)
//	l, err = db.GetAllItems()
//	println("\n")
//
//	println(l.Persons[0].to_String())
//	//res, err := db.comand_db(3)
//	//if err != nil {
//	//	println("err3")
//	//	println(err.Error())
//	//} else {
//	//	for res.Next(){
//	//		println("!")
//	//		var item Person
//	//		res.Scan(&item.id, &item.name,&item.age , &item.address, &item.work)
//	//		//println(strconv.Itoa(item.id), item.name, item.age , item.address, item.work)
//	//		println(item.to_String())
//	//	}
//	//}
//
//	db.Conn.Close()
//}
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

const (
	HOST = "localhost"
	PORT = 5432
)

var database = Database{}

func main() {
	println("Begin??????")
	database, err := Initialize("program", "test", "persons")
	if err != nil {
		println("cant connect")
		println(err.Error())
	}
	database.comand_db(0)
	database.comand_db(1)
	person1 := &Person{1, "ivan", 30, "piter", "proger"}
	//person3 := `{"id":1 , "name":"ivan","age": 30, "address":"piter", "work":"proger"}`
	//p1, _ := json.Marshal(&person1)
	//println(string(p1))
	person2 := &Person{2, "ivan4", 190, "kazan", "tsar"}
	err, n1 := database.AddItem(person1)
	err, n2 := database.AddItem(person2)
	println(n1, n2)
	//err = database.AddItem(*person3)

	defer database.Conn.Close()

	//l, err := database.GetAllItems()
	//if err != nil{
	//	println("gg")
	//	println(err.Error())
	//
	//}
	//println(l.Persons[0].to_String())
	httpHandler := NewHandler1(&database)
	//db, err := Initialize("program", "test", "persons")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}

	//db.comand_db(3)
	//db.comand_db(0)
	//db.comand_db(1)

	//db.Conn.Close()

	//person1 := &Person{1, "ivan", 30, "piter", "proger"}
	//person2 := &Person{2, "ivan4", 190, "kazan", "tsar"}
	//err = db.AddItem(*person1)
	//err = db.AddItem(*person2)
	//if err != nil {
	//	println(err.Error())
	//}
	//
	//l, err := db.GetAllItems()
	//println(l.Persons[0].to_String())

	//db.Conn.Close()

	addr := "/api/v1/:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}
	//dbUser, dbPassword, dbName :=
	//	os.Getenv("POSTGRES_USER"),
	//	os.Getenv("POSTGRES_PASSWORD"),
	//	os.Getenv("POSTGRES_DB")
	//database, err := Initialize("program", "test", "persons")
	//database.comand_db(3)
	//_, err = database.comand_db(0)
	//if err != nil{
	//	println(err.Error())
	//}
	//database.comand_db(1)
	//
	//db := db.Initialize(database)
	//all, _ := database.GetAllItems()
	//println(len(all.Items))
	//if err != nil {
	//	log.Fatalf("Could not set up database: %v", err)
	//}
	//defer database.Conn.Close()
	//

	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()

	test1()
	test2()
	test3()
	test3()
	test3()
	//test4()
	//test1()
	//test5()
	test1()

	defer Stop(server)
	log.Printf("Started server on %s", addr)
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
