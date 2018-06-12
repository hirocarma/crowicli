// Package get return crowi markdown body.
package get

import (
	"context"
	"time"

	"github.com/hirocarma/crowicli/client"

	"github.com/pkg/errors"
)

// Get implements client.Mode.
type Get struct {
}

func init() {
	client.Register("get", Get{})
}

// API returns markdown body.
func (m Get) API() (string, error) {
	var n *client.Client
	cli, err := n.Init()
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := cli.Pages.Get(ctx, client.Opt.Path)
	if err != nil {
		return "", errors.Wrap(err, "error by Pages.Get")
	}
	if !res.OK {
		err = errors.New(res.Error)
		return "", errors.Wrap(err, "res.Error by Pages.Get")
	}
	return res.Page.Revision.Body, nil
}
