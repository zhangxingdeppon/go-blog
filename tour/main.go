package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	// flag.StringVar(&name, "name", "GO language", "help information")
	// flag.StringVar(&name, "n", "go language", "help information")

	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "go lanaguage", "help information")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "PHP lanaguage", "help information")

	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)
}
