package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/BTGCodes/TwitterBot/pkg/swearjar"
)

var (
	configPath string
	profanePath string
)

// Type TwitterTokens contains the API authenication for Twitter
type TwitterTokens struct  {
		ConsumerKey       string `json:"zhC07OL0ZcIGeMSBZgODBskMN"`
		ConsumerSecret    string `json:"KTT9BYpMZQQKEPTd4xZb8HLlvFpum57ondUI0BVKGXF02b8T3u"`
		AccessToken       string `json:"1195184321419714560-vs02UssUkUHOtrRfThPaErtPNrGit8"`
		AccessTokenSecret string `json:"p8l33RFbGOA5SfM86OZETFg3zO1rxYXlc3WTolKGqiU7g"`
}


type Config struct {
	Tokens TwitterTokens `json:"twitter"`
	IgnoreUsers []string `json:"ignoreUsers"`
	Track []string `json:"track"`
}



func Load() (settings Config, swears swearjar.Swears){

	// read in config file into memory
	configRaw, err := ioutil.ReadFile(configPath)
	if err != nil{
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}

	// fill in Settings struct
	err = json.Unmarshal(configRaw, &settings)
	if err != nil{
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}

	// fill in Swears dictionary
	swears, err = swearjar.Load(profanePath)
	if err != nil{
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}

	return
}

// before main execution, handle setup of env variables and read them into package level
// variable Setting for use by other packages
func init(){
	flag.StringVar(&configPath, "config", "./Config.JSON", "Path to Config.JSON settings file")
	flag.StringVar(&profanePath, "profane", "./Profane.JSON", "Path to BadWords.JSON profane file")
	flag.Parse()
}
