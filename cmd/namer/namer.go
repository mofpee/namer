package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bbtoeach/namer"
)

func main() {
	text := flag.String("text", "", "A sentence to make a name. (Required)")
	count := flag.Int("count", 10, "A count of names")
	limit := flag.Int("limit", 6, "A limit length of name")
	algorithm := flag.String("alg", "default", "An algorithm to generate names")

	flag.Parse()

	if *text == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	n := namer.New(*algorithm)
	for i := 0; i < *count; i++ {
		fmt.Println(n.NameMixedShuffle(*text, *limit))
	}
}
