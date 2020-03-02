package main

import (
	"GoStudy/GoHttpServerOrmDemo/httpserver"
	"github.com/go-redis/redis"

	//"GoStudy/GoHttpServerOrmDemo/httpserver"
	//_ "GoStudy/GoHttpServerOrmDemo/models"
	"fmt"
	"net/http"

	//_ "github.com/go-sql-driver/mysql"
	//"net/http"
)

func main(){
	fmt.Println("Start")

	//query redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	//httpServer
	http.HandleFunc("/", httpserver.HandleIndex)
	err = http.ListenAndServe(":8000", nil)
	if err != nil{
		fmt.Println(err)
	}
}
