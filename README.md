# TwitterBot
simple twitter bot to retweet and like tweets based on some topics/hashtags to follow. Works in real time, using the Twitter streaming API. 

## Features 

- Follow & Like topics/hashtags
- Ignore certain users by ID, remove bad actors who may highjack tweet hashtag or topic
- Ignore certain tweets containing profane topics, filter list included


## Usage

This section describes how to download, setup and run the twitter bot.

### build from source
**Required**: You need to have Go tooling installed and git.

git clone repository:

`git clone https://github.com/BTGCodes/TwitterBot`

cd into `TwitterBot` directory and run build:

`go build ./cmd/bot`


### Pre-built binaries

contained in project is `build` directory, containing the bot already compiled for major operating systems.
## To Run

The binary expects the following two files to be located in the same directory, `Config.JSON` and `Profane.JSON` If you wish to place these two files elsewhere the following command line flags will be helpful:

- **config** - the path to the `Config.JSON` file
- **profane** - the path to the `Profane.JSON` file

to run, with `Config.JSON` and `Profane.JSON` in same directory as binary.

`./TwitterBot`

to run, with ith `Config.JSON` and `Profane.JSON` in different directories as binary.

`./TwitterBot -config path_To_Config/Config.JSON -profane path_To_Profane/Profane.JSON`


## Important Configuration Files

**Config.JSON** - contains the configuartion settings for bot.

Twitter API authenication tokens, you'll have to register for a dev account on twitter to get these values:
  - **consumerKey**: `"TWITTER_API_CONSUMER_KEY"`
  - **consumerSecret**: `"TWITTER_API_CONSUMER_SECRET"`
  - **accessToken**: `"TWITTER_API_ACCESSTOKEN"`
  - **accessSecret**: `"TWITTER_API_ACCESS_SECRET"`
  
The ignoreUsers filed, expects user IDs of accounts to ignore, if you need a tool to convert twitter @handle to id, use https://tweeterid.com/
  - **ignoreUsers** :`["12345", "292929", "393939"]`
 
The Track field, contians all the topics, hashtags you want to follow and interact with.
  - **track** : ```["#SomalisInTech", "#BlackTechTwitter", "Somalis in Technology"]```
  
  
  **Profane.JSON** - a dictionary of bad words and why they're profane. Add onto the list any words you wish to filter out in tweets you interact with.
  
  
  
## Issues & Feature Development

- Needs rate limiting, currently will hit rate limits for popular topics/hashtags 
- Add in flag to turn off profane filtering
- Introduce event hooks for meetup.com, to automatically promote events 
- Introduce planned tweets with time windows
 