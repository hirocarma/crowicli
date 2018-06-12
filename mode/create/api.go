package create

import (
	"context"
	"time"

	"github.com/hirocarma/crowicli/client"

	"github.com/pkg/errors"
)

// Create implements client.Mode.
type Create struct {
}

func init() {
	client.Register("create", Create{})
}

// API create markdown.
func (m Create) API() (string, error) {
	var n *client.Client
	cli, err := n.Init()
	if err != nil {
		return "", err
	}

	mdtxt, err := client.ReadMd(client.Opt.File)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := cli.Pages.Create(ctx, client.Opt.Path, mdtxt)

	if err != nil {
		return "", errors.Wrap(err, "error by Pages.Create")
	}
	if !res.OK {
		err = errors.New(res.Error)
		return "", errors.Wrap(err, "res.Error by Pages.Create")
	}
	return client.Opt.Path + " is created.", nil
}
