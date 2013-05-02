package myepisodes

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"net/http"
	"io/ioutil"
	"errors"
)

//func get_feed(feed string)

type Episode struct {
	ShowID				int
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
	GUID				string `xml:"guid"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

const (
	feedBaseUri = "https://www.myepisodes.com/rss.php"
	loginBaseUri = "https://www.myepisodes.com/login.php"
	updateBaseUri = "https://www.myepisodes.com/myshows.php"
)
var titleRegex = regexp.MustCompile(`\[ (?P<series>.+) \]\[ (?P<season>\d+)x(?P<episode>\d+) \]\[ (?P<title>.+) \]\[ (?P<date>.+) \]`)
var guidRegex = regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)

func md5Pwd(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func buildLoginURI(uid, pwd string)(string) {
	return loginBaseUri + "?action=Login&username=" + uid + "&password=" + pwd
}
func buildFeedURI(feedname, uid, pwd string)(uri string) {
	pwdmd5 := md5Pwd(pwd)
	uri = feedBaseUri + "?feed=" + feedname +
		"&uid=" + uid +
		"&pwdmd5=" + pwdmd5 +
		"&showignored=0&onlyunacquired=1&sort=ASC"
	return
}
func buildUpdateURI(episode Episode)(string) {
	return updateBaseUri + "?action=Update&showid=" + string(episode.ShowID) +
			"&season=" + string(episode.SeasonNumber) +
			"&episode=" + string(episode.EpisodeNumber) + "&seen=0"
}

func GetCookie(uid,pwd string)([]string) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", buildLoginURI(uid, pwd), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)

	fmt.Printf("cookie: %v\n", resp.Header["set-cookie"])
	return resp.Header["set-cookie"]
}

func GetFeed(feedname, uid, pwd string) (episodes []Episode) {
	uri := buildFeedURI(feedname, uid, pwd)
	body,_ := getRss(uri)
	return parseFeed(body)
}

func (e Episode) MarkAquired(cookie []string)(bool) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", buildUpdateURI(e), nil)
	req.Header.Add("Cookie", "cookie")
	resp, _ := client.Do(req)
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

func getRss(uri string) ([]byte,error) {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Printf("Error fetching feed: %v\n", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, errors.New("StatusCode: " + string(resp.StatusCode))
	}
	if resp.ContentLength == 0{
		return nil, errors.New("ContentLength = 0")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)
	}
	return body, nil
}

func parseRss(feed []byte) (rss RSS) {
	err := xml.Unmarshal(feed, &rss)
	if err != nil {
		fmt.Printf("Error parsing xml: %v\n", err)
	}
	return
}
func parseFeed(feed []byte) (episodes []Episode) {
	rss := parseRss(feed)
	for _, item := range rss.Items.ItemList {
		episode := extractEpisode(item.Title)
		episode.ShowID,_ = strconv.Atoi(guidRegex.FindStringSubmatch(item.GUID)[1])
		episodes = append(episodes, episode)
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
