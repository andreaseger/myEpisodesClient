package myepisodes

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
)

//func get_feed(feed string)

type Episode struct {
	Series        string
	SeasonNumber  int
	EpisodeNumber int
	Title         string
	Date          string
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Items   items    `xml:"channel"`
}
type items struct {
	XMLName  xml.Name `xml:"channel"`
	ItemList []item   `xml:"item"`
}
type item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

var titleRegex = regexp.MustCompile(`\[ (?P<series>.+) \]\[ (?P<season>\d+)x(?P<episode>\d+) \]\[ (?P<title>.+) \]\[ (?P<date>.+) \]`)

func md5Pwd(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	return fmt.Sprintf("%x", h.Sum(nil))
}
func buildURI(feedname, uid, pwd string)(uri string) {
	pwdmd5 := md5Pwd(pwd)
	uri = "https://www.myepisodes.com/rss.php?feed=" + feedname +
		"&uid=" + uid +
		"&pwdmd5=" + pwdmd5 +
		"&showignored=0&onlyunacquired=1&sort=ASC"
	return
}

func getFeed(feedname, uid, pwd string) (episodes []Episode) {
	uri := buildURI(feedname, uid, pwd)
	body := getRss(uri)
	return parseFeed(body)
}
func getRss(uri string) (body []byte) {
	return
}

func parseRss(feed []byte) (rss RSS) {
	err := xml.Unmarshal(feed, &rss)
	if err != nil {
		fmt.Println("Error parsing xml: %v", err)
	}
	return
}
func parseFeed(feed []byte) (episodes []Episode) {
	rss := parseRss(feed)
	for _, item := range rss.Items.ItemList {
		episodes = append(episodes, extractEpisode(item.Title))
	}
	return
}

func extractEpisode(title string) (e Episode) {
	match := titleRegex.FindStringSubmatch(title)
	e.Series = match[1]
	e.SeasonNumber, _ = strconv.Atoi(match[2])
	e.EpisodeNumber, _ = strconv.Atoi(match[3])
	e.Title = match[4]
	e.Date = match[5]
	return
}
