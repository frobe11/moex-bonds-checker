package main

import (
	"fmt"
	"log"

	moex "github.com/frobe11/moex-bonds-checker/iternal/services"
)

func main() {
	bonds, err := moex.Getbonds()
	fmt.Print(bonds)
	
	if err != nil {
		log.Fatal(err)
		return
	}
}
