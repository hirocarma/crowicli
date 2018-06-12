package client

import (
    "fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

// ReadMd reads markdown file.
func ReadMd(filename string) (string, error) {
	fdata, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrap(err, "error markdown file read")

	}
	return fmt.Sprintf(string(fdata)),nil
}
