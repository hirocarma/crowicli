# crowicli

Command line client for [crowi-plus](https://github.com/weseek/crowi-plus).

## Overview

Use [go-crowi](https://github.com/crowi/go-crowi.git)(The official Go client for Crowi).

## Features

You can add and update Crowi by specifying the markdown file from the
command line.Also Provides a list of crowi documents and documents.

## Installation

``` console
$ go get github.com/hirocarma/crowicli
```

Add .env file in the current directory. Or add .env.crowicli in your $HOME.
If both files are present, .env takes precedence.

Example:
``` :.envsample
CROWI_ACCESSTOKEN='xxxxxxxxxxxxxxxxxxxxxxxxxxx'
CROWI_URI='http://localhost:3000'
CROWI_INSECURE=false
CROWI_USER=taro
```
* `CROWI_URI`: *Required.* The URI to use for the requests.

* `CROWI_ACCESSTOKEN`: *Required.* It can be obtained from the
[Settings] -> [API Settings] tab.

* ` CROWI_INSECURE`: *Optional.* Please set to true when using a site using a
self-signed certificate in a test environment etc, (default `false`).

* `CROWI_USER`: *Optional.* Default is $USER(OS environment).

## Usage

``` shell
  -f string
        markdown file (default "./file.md")
  -m string
        mode:create,update,get,list (default "list")
  -p string
        wiki path
```

Example:

* `crowicli -m list` or `crowicli`  
  get list of crowi documents.
* `crowicli -m get -p /user/taro/memo/2018/02/14/task`  
  get crowi document. (Output to stdout.)
* `crowicli -m update -p /user/taro/memo/2018/02/14/ -f ~/task.md`  
  update crowi document. '/user/taro/memo/2018/02/14/task' will be updated.
* `crowicli -m create -p /user/taro/memo/2018/02/15/ -f ~/report.md`  
  create crowi document. '/user/taro/memo/2018/02/15/report' will be created.

## License
MIT

