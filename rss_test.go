package myepisodes

import(
  "testing"
  "fmt"
	"io/ioutil"
	"os"
)

func TestParseRss(t *testing.T){
	file, e := ioutil.ReadFile("test-files/today.xml")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
  rss := parseRss(file)
  if len(rss.Items.ItemList) != 2{
    t.Errorf("Wrong number of items")
  }
}
func TestParseFeed(t * testing.T){
	file, e := ioutil.ReadFile("test-files/today.xml")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
  episodes := parseFeed(file)
  if len(episodes) != 2{
    t.Errorf("Wrong number of items")
  }
  if episodes[0].Title != "Revenge"{
    t.Errorf("Wrong title: 'Revenge' vs %v", episodes[0].Title)
  }
}
func TestExtractEpisode(t * testing.T){
  title := "[ NCIS ][ 10x22 ][ Revenge ][ 01-May-2013 ]"
  testEpisode := Episode{
    Series: "NCIS",
    EpisodeNumber: 22,
    SeasonNumber: 10,
    Title: "Revenge",
    Date: "01-May-2013",
  }
  episode := extractEpisode(title)
  if episode != testEpisode{
    t.Errorf("Episode wrong: %v", episode)
  }
}

func TestMd5Pwd(t * testing.T){
  s := "The fog is getting thicker!"
  smd5 := md5Pwd(s)
  if smd5 != "bd009e4d93affc7c69101d2e0ec4bfde" {
    t.Errorf("md5 wrong: %v", smd5)
  }
}

