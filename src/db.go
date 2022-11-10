package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strconv"
)

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	println(HOST, PORTDB, username, password, database)
	//dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//HOST, PORT, username, password, database)
	dsn := "postgres://vnkldgkbuxtsdk:edc93be519c3326ee21f9197426f8e0f29c17483ed84f9a261927987a85cb075@ec2-34-248-169-69.eu-west-1.compute.amazonaws.com:5432/dd9rcurfahnu5r"
	username = "vnkldgkbuxtsdk"
	password = "edc93be519c3326ee21f9197426f8e0f29c17483ed84f9a261927987a85cb075"
	database = "dd9rcurfahnu5r"
	HOST = "ec2-34-248-169-69.eu-west-1.compute.amazonaws.com"
	PORTDB = "5432"
	println(HOST, PORTDB, username, password, database)
	//dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
	//	username, password, HOST, PORTDB, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}

func (db Database) comand_db(n int) (*sql.Rows, error) {
	b, err := ioutil.ReadFile("./db_sql/db_c" + strconv.Itoa(n) + ".sql")
	//print("db_sql/db_c" + strconv.Itoa(n) + ".sql")
	if err != nil {
		println(err.Error())
		return nil, err
	}
	str := string(b)
	//println(str)
	res, err := db.Conn.Query(str)
	if err != nil {
		println("err1")
		println(err.Error())
		return nil, err
	}
	return res, nil
	//rr, _ := res.Columns()
	//fmt.Println(rr)
}

func (db Database) GetAllItems() (*personList, error) {
	list := &personList{}
	rows, err := db.Conn.Query("SELECT * FROM persons ORDER BY Id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item Person
		err := rows.Scan(&item.Id, &item.Name, &item.Age, &item.Address, &item.Work)
		if err != nil {
			return list, err
		}
		list.Persons = append(list.Persons, item)
	}
	return list, nil
}

func (db Database) AddItem(item *Person) (error, int) {
	//var Id int
	var name string
	var age int
	var address string
	var work string
	query := `INSERT INTO persons(name, age, address, work) VALUES ( $1, $2, $3, $4) RETURNING Name, Age, Address, Work`
	err := db.Conn.QueryRow(
		query, item.Name, item.Age, item.Address, item.Work).Scan(
		&name, &age, &address, &work)
	if err != nil {
		println("001")
		println(err.Error())
		return err, -1
	}
	idA, _ := db.GetAllItems()
	//println(idA.Persons[len(idA.Persons-1)])
	item.Id = idA.Persons[0].Id
	item.Name = name
	item.Age = age
	item.Address = address
	item.Work = work
	return nil, item.Id
}

func (db Database) GetItemById(itemId int) (Person, error) {
	item := Person{}
	query := `SELECT * FROM persons WHERE id = $1;`
	row := db.Conn.QueryRow(query, itemId)
	switch err := row.Scan(&item.Id, &item.Name, &item.Age, &item.Address, &item.Work); err {
	case sql.ErrNoRows:
		return item, err
	default:
		return item, err
	}
}

func (db Database) DeleteItem(itemId int) error {
	query := `DELETE FROM persons WHERE id = $1;`
	_, err := db.Conn.Exec(query, itemId)
	switch err {
	case sql.ErrNoRows:
		return err
	default:
		return err
	}
}

func (p1 *Person) personUpdate(p2 Person) {
	if p1.Id == 0 {
		p1.Id = p2.Id
	}
	if p1.Age == 0 {
		p1.Age = p2.Age
	}
	if p1.Work == "" {
		p1.Work = p2.Work
	}
	if p1.Address == "" {
		p1.Address = p2.Address
	}
	if p1.Name == "" {
		p1.Name = p2.Name
	}

}

func (db Database) UpdateItem(itemId int, itemData Person) (Person, error) {
	item := Person{}
	item0, _ := db.GetItemById(itemId)
	itemData.personUpdate(item0)

	query := `UPDATE persons SET name=$1, age=$2, address= $3, work=$4 WHERE id=$5 RETURNING id, name, age, address, work;`
	err := db.Conn.QueryRow(query, itemData.Name, itemData.Age, itemData.Address, itemData.Work, itemId).Scan(&item.Id, &item.Name, &item.Age, &item.Address, &item.Work)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, err
		}
		return item, err
	}
	println("hhheeerree")
	println(item.Name, item.Work)
	return item, nil
}
