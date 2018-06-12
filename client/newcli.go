package client

import (
	"github.com/crowi/go-crowi"
    "github.com/pkg/errors"
)

// Client embeds crow.Client.
type Client struct {
	crowi.Client
}

// Init client.
func (c *Client) Init() (*crowi.Client, error) {
	config := crowi.Config{
		URL:                CrowiEnv.URI,
		Token:              CrowiEnv.AccessToken,
		InsecureSkipVerify: CrowiEnv.Insecure,
	}

	client, err := crowi.NewClient(config)
	if err != nil {
		return nil, errors.Wrap(err, "Error by crowi.NewClient") 
	}
	return client, nil
}
