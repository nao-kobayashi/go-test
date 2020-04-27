package main

import (
	"errors"
	"fmt"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

type IDate interface {
	SetYear(year int) error
	SetMonth(month int) error
	SetDay(day int) error
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{
		Year:  2020,
		Month: 4,
		Day:   21,
	}
	fmt.Println(date)

	date2 := Date{}
	date2.SetYear(2020)
	date2.SetMonth(6)
	date2.SetDay(22)
	fmt.Println(date2)

	// date3 := Date{}
	// err := date3.SetYear(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err := date3.SetMonth(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err := date3.SetDay(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	id := GetIDate(&date)
	fmt.Println(id)

	var i interface{}
	t1, ok := i.(IDate)
	fmt.Println("type:", t1)
	fmt.Println("ok", ok)

}

func GetIDate(d *Date) IDate {
	return d
}
