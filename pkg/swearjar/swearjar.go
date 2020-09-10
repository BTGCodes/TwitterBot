package swearjar

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

// Swears type alias to map[string][]string
type Swears map[string][]string

// Load a default set of swears and returns a Swears instance
// or can load a JSON file which will unmarshal to Swears (map[string[]string)
func Load(config ...string) (swears Swears, err error) {

	var swearsJson []byte

	if config != nil && config[0] != "" {
		absPath, err := filepath.Abs(config[0])
		if err != nil {
			return nil, err
		}

		swearsJson, err = ioutil.ReadFile(absPath)
		if err != nil {
			return nil, err
		}
	}

	err = json.Unmarshal(swearsJson, &swears)
	if err != nil {
		return nil, err
	}

	return
}

// Profane takes an input word against the swears list and returns bool at an occurrence.
// This calls Scorecard but discards the reason
func (swears Swears) Profane(input string) (bool, error) {
	profane, _, err := swears.Scorecard(input)
	return profane, err
}

// Scorecard checks input against swear list and returns bool, slice of reasons + error.
// This will return at first occurrence
func (swears Swears) Scorecard(input string) (profane bool, reasons []string, err error) {
	for word, reason := range swears {
		wordPattern := `\b` + word + `\b`
		match, err := regexp.MatchString(wordPattern, input)

		if err != nil {
			return false, nil, err
		}

		if match {
			return true, reason, nil
		}
	}

	return false, nil, nil
}

// CheckTweet checks a tweet against profane word list and returns bool.
func (swears Swears) CheckTweet(tweet string) bool{
	words := strings.Fields(tweet)
	for _, word := range words {
		ok, _, _ := swears.Scorecard(word)
		if ok{
			return false
		}
	}
	return true
}