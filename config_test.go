package myepisodes

import "testing"

func TestReadConfig(t *testing.T){
  config := ReadConfig("test-files/config.json")
  if config.UserID != "test-userid"{
    t.Errorf("Wrong userid, %v", config.UserID)
  }
  if config.Password != "test-password"{
    t.Errorf("Wrong password, %v", config.Password)
  }
}

