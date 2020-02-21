package main

import (
	"GoStudy/GoHttpServerOrmDemo/httpserver"
	_ "GoStudy/GoHttpServerOrmDemo/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main(){
	fmt.Println("Start")

	//httpServer
	http.HandleFunc("/", httpserver.HandleIndex)
	err := http.ListenAndServe(":8000", nil)
	if err != nil{
		fmt.Println(err)
	}
}
