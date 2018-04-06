# gomorse [![GoDoc](https://godoc.org/github.com/NBR41/gomorse?status.svg)](https://godoc.org/github.com/NBR41/gomorse) [![Build Status](https://travis-ci.org/NBR41/gomorse.svg?branch=master)](https://travis-ci.org/NBR41/gomorse) [![Coverage Status](http://codecov.io/gh/NBR41/gomorse/branch/master/graph/badge.svg)](http://codecov.io/gh/NBR41/gomorse)
Simple command to code phrases into beep morse code.

## Purpose
Insprired by [github.com/dbatbold/beep](https://github.com/dbatbold/beep), this library get a phrase as input, check it validity, and play on speakers beep morse code corresponding to the input phrase.

## Install


To compile the package require libasound2-dev, to install:
```bash
sudo apt-get install libasound2-dev  # for Debian and Ubuntu
```

After you get the package, and install dependancies
```bash
go get github.com/NBR41/gomorse
cd $GOPATH/src/github.com/NBR41/gomorse/cmd/gomorse
dep ensure
go install # or go build
```

## Usage
```bash
./gomorse -h
gomorse is a tool to transcript phrases to morse code.
3 parameters:
- volume
- beep duration
- beep frequence

Usage:
  gomorse [flags]

Flags:
  -d, --duration int      beep duration in ms (default 250)
  -f, --frequence float   frequency in Hertz (1-22050) (default 523.25)
  -h, --help              help for gomorse
  -v, --volume int        beep volume (0 to 100) (default 100)
```
