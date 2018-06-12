// Package list return crowi markdown list.
package list

import (
	"context"
	"strings"
	"time"

	"github.com/hirocarma/crowicli/client"

	"github.com/crowi/go-crowi"
	"github.com/pkg/errors"
)

// List implements client.Mode.
type List struct {
}

func init() {
	client.Register("list", List{})
}

// API returns markdown list.
func (m List) API() (string, error) {
	var n *client.Client
	cli, err := n.Init()
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := cli.Pages.List(
		ctx,
		"",
		client.CrowiEnv.User,
		&crowi.PagesListOptions{
			ListOptions: crowi.ListOptions{Pagenation: true},
		})
	if err != nil {
		return "", errors.Wrap(err, "error by Pages.List")
	}
	if !res.OK {
		err = errors.New(res.Error)
		return "", errors.Wrap(err, "res.Error by Pages.List")
	}

	pathlist := []string{}
	for _, p := range res.Pages {
		pathlist = append(pathlist, p.Revision.Path)
	}

	return strings.Join(pathlist, "\n"), nil

}
