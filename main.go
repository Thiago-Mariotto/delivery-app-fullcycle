package main

import (
	"fmt"

	"github.com/thiago-mariotto/go-simulator/application/route"
)

func main() {
	route := route.Route{
		Id:       "1",
		ClientId: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPosition()

	fmt.Println(stringJson[i])

}
