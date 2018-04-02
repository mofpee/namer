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
	limit := flag.Int("limit", 6, "A limit length of name")

	flag.Parse()

	if *text == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	n := namer.New()
	for i:=0; i<*count; i++ {
		fmt.Println(n.Name(*text, *limit))
	}
}