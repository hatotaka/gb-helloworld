package main

import (
	"helloworldlib"

	"github.com/k0kubun/pp"
)

func main() {
	pp.Println("start")
	helloworldlib.Hello("wrold")
	pp.Println("end")
}
