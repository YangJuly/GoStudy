package models

import "fmt"

type HotelDictInfo struct {
	Id   int
	HotelName string
	HotelAddress string
}


func GetAllHotel() ([]HotelDict, error) {
	var err error
	//var hotelInfo []HotelDict
	hotelInfo := make([]HotelDict, 0)

	err = Engine.Table("hotel_dict").
		Find(&hotelInfo)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return hotelInfo, nil
}