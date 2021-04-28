package main

import (
	"example.com/m/v2/httpapi"
	"log"
)

func main() {
	if err := httpapi.InstanceEPM().Execute(httpapi.InstanceRT().GetDefault()); err != nil {
		log.Println("main", err)
	}
}
