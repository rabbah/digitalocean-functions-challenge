package main

import (
	"flag"
)


func init() {
	flag.StringVar(&sammyname, "name", "", "The name to give to your new Sammy.")
	flag.StringVar(&sammytype, "type", "", "The type to give to your new Sammy.")
	flag.Parse()
}

func main() {
	doit()
}
