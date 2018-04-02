package main

import (
	"flag"
	"os"
	"fmt"
	"github.com/bearchit/namer"
)

func main() {
	text := flag.String("text", "", "A sentence to make a name. (Required)")
	count := flag.Int("count", 10, "A count of names")

	flag.Parse()

	if *text == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	for i:=0; i<*count; i++ {
		fmt.Println(namer.Name(*text))
	}
}