package interfacexplain

import "fmt"

type Database interface{
	Connection()
}

type Postgres struct{}

func (p Postgres) Connection(){
	fmt.Println("Connecting to the postgres db")
}

type MySql struct{}

func (m MySql) Connection(){
	fmt.Println("connecting to the Mysql")
}

func dbConnector(i Database){
   i.Connection()
}

func DbConnection(){
	p := Postgres{}
	dbConnector(p)
}