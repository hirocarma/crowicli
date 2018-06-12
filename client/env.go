package client

import (
    "os"

    "github.com/joho/godotenv"
    "github.com/kelseyhightower/envconfig"
    "github.com/pkg/errors"
)

func envload() error {
    err := godotenv.Load()
    if err != nil {
        err := godotenv.Load(os.Getenv("HOME") + "/.env.crowicli")
        if err != nil {
            return errors.Wrap(err, "error loading .env file")
        }
    }
    return nil
}

type envapp struct {
    URI         string `required:"true"`
    AccessToken string `required:"true"`
    Insecure    bool   `default:"false"`
    User        string
}

func (p *envapp) envconf() (*envapp, error) {
    var e envapp
    err := envconfig.Process("crowi", &e)
    if err != nil {
        return nil, errors.Wrap(err, "error by envconfig.Process")
    }
    return &e, nil
}

// CrowiEnv has OS environments variables for crowi.
var CrowiEnv *envapp

// Env read .env and stores environment variables in Envapp.
func Env() error {
    err := envload()
    if err != nil {
        return err
    }
    var e envapp
    CrowiEnv, err = e.envconf()
    if err != nil {
        return err
    }
    if CrowiEnv.User == "" {
        CrowiEnv.User = os.Getenv("USER")
    }
    return nil
}
