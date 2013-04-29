package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var configfile = flag.String("config", "my_episodes.json", "config file for my episodes")

func main() {
	flag.Parse()
	fmt.Println("configfile: ", *configfile)
	file, e := ioutil.ReadFile(*configfile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))
}
