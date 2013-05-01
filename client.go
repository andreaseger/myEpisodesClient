package myepisodes

import (
	"flag"
	"fmt"
)

var configfile = flag.String("config", "my_episodes.json", "config file for my episodes")

func main() {
	flag.Parse()
	config := ReadConfig(*configfile)
	fmt.Println("%s", config.UserID)
}
