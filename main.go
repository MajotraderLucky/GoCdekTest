package main

import (
	"context"
	"fmt"

	"github.com/vseinstrumentiru/cdek"
)

func main() {
	fmt.Println("Hello world")
	client := cdek.NewClient("https://integration.edu.cdek.ru").SetAuth("EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI", "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG")

	type Address struct {
		Street  *string `xml:"Street,attr"`
		House   *string `xml:"House,attr"`
		Flat    *string `xml:"Flat,attr,omitempty"`
		Phone   *string `xml:"Phone,attr,omitempty"`
		PvzCode *string `xml:"PvzCode,attr,omitempty"`
	}

	ctx := context.Background()

	cities, err := client.GetCities(ctx, map[cdek.CityFilter]string{cdek.CityFilterPage: "1"})

	if err != nil {
		panic(err)
	}
	fmt.Println(cities)
}
