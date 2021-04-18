package model

type Table struct {
	Id       string  `json:"id" xml:"id" form:"id" query:"id"`
	Name     string  `json:"name" xml:"name" form:"name" query:"name"`
	Surname    string  `json:"surname" xml:"surname" form:"surname" query:"surname"`
}
