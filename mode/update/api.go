package update

import (
	"context"
	"time"

	"github.com/hirocarma/crowicli/client"

	"github.com/pkg/errors"
)

// Update implements client.Mode.
type Update struct {
}

func init() {
	client.Register("update", Update{})
}

// API update markdown.
func (m Update) API() (string, error) {
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

	mdtxt, err := client.ReadMd(client.Opt.File)
	if err != nil {
		return "", err
	}

	res2, err := cli.Pages.Update(ctx, res.Page.ID, mdtxt)
	if err != nil {
		return "", errors.Wrap(err, "error by Pages.Update")
	}
	if !res2.OK {
		err = errors.New(res2.Error)
		return "", errors.Wrap(err, "res.Error by Pages.Update")
	}
	return client.Opt.Path + " is updated.", nil
}
