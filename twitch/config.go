package twitch

import (
	"fmt"

	gotwitch "github.com/catsby/go-twitch/twitch"
)

type Config struct {
	ApiKey string
}

type twitchClient struct {
	conn *gotwitch.Client
}

func (c *Config) Client() (interface{}, error) {
	var client twitchClient

	if c.ApiKey == "" {
		return nil, fmt.Errorf("[Err] No API key for twitch")
	}

	fconn, err := gotwitch.NewClient("", c.ApiKey)
	if err != nil {
		return nil, err
	}

	client.conn = fconn
	return &client, nil
}
