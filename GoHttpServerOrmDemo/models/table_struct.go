package models

//SameMapper 支持结构体名称和对应的表名称相同
//以及结构体field名称与对应的表字段名称相同的命名；

type HotelDict struct {
	Id int `xorm:"not null pk autoincr INT(11)"`
	HotelName string `xorm:"not null VARCHAR(200)"`
	HotelAddress string `xorm:"not null VARCHAR(200)"`
}
