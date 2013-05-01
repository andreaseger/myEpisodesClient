package myepisodes

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestGetRSS(t *testing.T) {
	uri := buildURI("today", "hubert.h", "foooo")
	_,e := getRss(uri)
	if m,_:=regexp.MatchString(`ContentLength`, e.Error());!m {
		t.Errorf("getRss failed %v", e)
	}
}

func TestBuildURI(t *testing.T) {
	uri := buildURI("testfeed", "testuid", "testpwd")
	if m,_ := regexp.MatchString(`myepisodes\.com`, uri); !m {
		t.Errorf("wrong base url: %v", uri)
	}
	if m,_ := regexp.MatchString(`feed=testfeed&uid=testuid&pwdmd5=`, uri); !m {
		t.Errorf("wrong feed and user: %v", uri)
	}
	pwdmd5 := md5Pwd("testpwd")
	if m,_ := regexp.MatchString(pwdmd5, uri); !m {
		t.Errorf("pwdmd5 missing: %v", uri)
	}
}

func TestParseRss(t *testing.T) {
	file, e := ioutil.ReadFile("test-files/today.xml")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	rss := parseRss(file)
	if len(rss.Items.ItemList) != 2 {
		t.Errorf("Wrong number of items")
	}
}
func TestParseFeed(t *testing.T) {
	file, e := ioutil.ReadFile("test-files/today.xml")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	episodes := parseFeed(file)
	if len(episodes) != 2 {
		t.Errorf("Wrong number of items")
	}
	if episodes[0].Title != "Revenge" {
		t.Errorf("Wrong title: 'Revenge' vs %v", episodes[0].Title)
	}
}
func TestExtractEpisode(t *testing.T) {
	title := "[ NCIS ][ 10x22 ][ Revenge ][ 01-May-2013 ]"
	testEpisode := Episode{
		Series:        "NCIS",
		EpisodeNumber: 22,
		SeasonNumber:  10,
		Title:         "Revenge",
		Date:          "01-May-2013",
	}
	episode := extractEpisode(title)
	if episode != testEpisode {
		t.Errorf("Episode wrong: %v", episode)
	}
}

func TestMd5Pwd(t *testing.T) {
	s := "The fog is getting thicker!"
	smd5 := md5Pwd(s)
	if smd5 != "bd009e4d93affc7c69101d2e0ec4bfde" {
		t.Errorf("md5 wrong: %v", smd5)
	}
}
