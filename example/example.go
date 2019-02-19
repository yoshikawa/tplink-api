package main

import (
	"fmt"
	"log"

	tplink "github.com/Pluslab/tplink-api"
)

func main() {
	api, err := tplink.Connect("yoshikawataiki@gmail.com", "hogehage")
	if err != nil {
		log.Fatal(err)
	}
	hs105, err := api.GetHS105("HS105_NAME")
	if err != nil {
		log.Fatal(err)
	}
	// err = hs105.TurnOn()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print("success!")

	info, err := hs105.GetInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(info)
}
