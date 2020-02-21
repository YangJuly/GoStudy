package httpserver

import (
	"GoStudy/GoHttpServerOrmDemo/models"
	"encoding/json"
	"fmt"
	"net/http"
)

//获取所有的hotel数据
func getAllHotelData()  ([]byte,error){

	v, err := models.GetAllHotel()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return json.MarshalIndent(v,"","")
}

func HandleIndex(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println("PATH: ", r.URL.Path)
	fmt.Println("SCHEME: ", r.URL.Scheme)
	fmt.Println("METHOD: ", r.Method)
	fmt.Println()

	resp,err := getAllHotelData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(resp))
}

