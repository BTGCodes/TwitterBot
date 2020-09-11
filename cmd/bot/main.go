package main

import (
	"github.com/BTGCodes/TwitterBot/pkg/config"
	"github.com/davidk/anaconda"
	"github.com/sirupsen/logrus"
	"net/url"
)



func main() {

	settings, swears := config.Load()

	anaconda.SetConsumerKey(settings.Tokens.ConsumerKey)
	anaconda.SetConsumerSecret(settings.Tokens.ConsumerSecret)
	api := anaconda.NewTwitterApi(settings.Tokens.AccessToken, settings.Tokens.AccessTokenSecret)

	log := &logger{logrus.New()}
	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": settings.Track,
	})
	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			logrus.Warningf("Received unexpected value of type %T", v)
		}


		// check to ensure it's not a "retweet"
		if t.RetweetedStatus != nil {
			continue
		}

		// do not interact with users on ignore list
		for _, user := range settings.IgnoreUsers{
			if t.User.IdStr == user{
				continue
			}
		}

		// check against profane list
		if ok := swears.CheckTweet(t.Text); !ok{
			continue
		}

		//like tweet!
		_, err := api.Favorite(t.Id)
		if err != nil {
			logrus.Errorf("could not favorite twitter %d:", t.Id)
		}

		// retweet it!
		_, err = api.Retweet(t.Id, false)
		if err != nil {
			logrus.Errorf("could not retweet %d:", t.Id)
		}

	}
}



// Satisfy logging interface
type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
