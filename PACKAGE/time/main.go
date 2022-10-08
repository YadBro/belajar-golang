package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("Now :", now)
	fmt.Println("Year :", now.Year())
	fmt.Println("Month :", now.Month())
	fmt.Println("Day :", now.Day())
	fmt.Println("Hour :", now.Hour())
	fmt.Println("Minute :", now.Minute())
	fmt.Println("Second :", now.Second())
	fmt.Println("Nanosecond :", now.Nanosecond())
	fmt.Println("Unix :", now.Unix())
	fmt.Println("UnixNano :", now.UnixNano())
	fmt.Println("UnixMilli :", now.UnixMilli())
	fmt.Println("UnixMicro :", now.UnixMicro())

	fmt.Println("\n\nUTC =======")
	now = now.UTC()
	fmt.Println("UTC Now :", now)
	fmt.Println("UTC Year :", now.Year())
	fmt.Println("UTC Month :", now.Month())
	fmt.Println("UTC Day :", now.Day())
	fmt.Println("UTC Hour :", now.Hour())
	fmt.Println("UTC Minute :", now.Minute())
	fmt.Println("UTC Second :", now.Second())
	fmt.Println("UTC Nanosecond :", now.Nanosecond())
	fmt.Println("UTC Unix :", now.Unix())
	fmt.Println("UTC UnixNano :", now.UnixNano())
	fmt.Println("UTC UnixMilli :", now.UnixMilli())

	fmt.Println("\n\nLOCAL =======")
	now = now.Local()
	fmt.Println("Local Now :", now)
	fmt.Println("Local Year :", now.Year())
	fmt.Println("Local Month :", now.Month())
	fmt.Println("Local Day :", now.Day())
	fmt.Println("Local Hour :", now.Hour())
	fmt.Println("Local Minute :", now.Minute())
	fmt.Println("Local Second :", now.Second())
	fmt.Println("Local Nanosecond :", now.Nanosecond())
	fmt.Println("Local Unix :", now.Unix())
	fmt.Println("Local UnixNano :", now.UnixNano())
	fmt.Println("Local UnixMilli :", now.UnixMilli())
	fmt.Println("Local UnixMicro :", now.UnixMicro())

	fmt.Println("\n\nDATE =======")
	utc := time.Date(2010, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, error := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
	fmt.Println(error)
}
