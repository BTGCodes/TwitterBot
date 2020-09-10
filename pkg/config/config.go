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
		ConsumerKey       string `json:"consumerKey"`
		ConsumerSecret    string `json:"consumerSecret"`
		AccessToken       string `json:"accessToken"`
		AccessTokenSecret string `json:"accessSecret"`
}


type Config struct {
	Tokens TwitterTokens `json:"twitter"`
	IgnoreUsers []string `json:"ignoreUsers"`
	track []string `json:"track"`
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
