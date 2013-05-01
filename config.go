package myepisodes

import(
  "fmt"
	"io/ioutil"
  "encoding/json"
	"os"
)

type Config struct{
  UserID string
  Password string
}
func ReadConfig(filename string) (c Config){
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
  err := json.Unmarshal(file, &c)
  if err != nil {
    fmt.Println("JSON parse error: %v", err)
  }
  return
}
