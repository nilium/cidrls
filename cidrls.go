package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/apparentlymart/go-cidr/cidr"
)

func main() {
	for _, in := range os.Args[1:] {
		_, ipnet, err := net.ParseCIDR(in)
		if err != nil {
			log.Fatal("cannot parse ", in, ": ", err)
		}
		f, l := cidr.AddressRange(ipnet)
		fmt.Println(f, "-", l)
	}
}
