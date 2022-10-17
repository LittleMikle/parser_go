package main

import (
	"fmt"
	mydb "github.com/LittleMikle/parser_go/db"
)

func main() {
	mydb.ConnectToDb()
	fmt.Println("connected to db")

	UrlParse()
}
