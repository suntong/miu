
# miu

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/miu?status.svg)](http://godoc.org/github.com/suntong/miu)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/miu)](https://goreportcard.com/report/github.com/suntong/miu)
[![travis Status](https://travis-ci.org/suntong/miu.svg?branch=master)](https://travis-ci.org/suntong/miu)

## TOC
- [miu - Monitor Internet Usage](#miu---monitor-internet-usage)
- [Usage](#usage)
  - [$ miu](#-miu)
  - [$ miu default](#-miu-default)
- [Author(s) & Contributor(s)](#author(s)-&-contributor(s))

## miu - Monitor Internet Usage

The `miu` makes it easy to monitor Internet usage from ISP. 

# Usage

### $ miu
```sh
Monitor Internet Usage
Version 0.1.0 built on 2017-10-10

Tool to monitor Internet usage from ISP

Options:

  -h, --help       display help information
  -i, --url       *usage URL from ISP
  -c, --css       *css selection for usage text from usage rul
  -t, --to         email address to send to
  -p, --template   mail sending command template string
  -b, --base       base href tag used in the wrapped up html
  -n, --no-html    do not wrap up the output with html tags
  -d, --days       days to shift from billing date
  -v, --verbose    Verbose mode (Multiple -v increase the verbosity)

Commands:

  default   Default ISP
```

### $ miu default
```sh
Default ISP

Usage:
  miu default -i http://url/ -c css.sel

Options:

  -h, --help       display help information
  -i, --url       *usage URL from ISP
  -c, --css       *css selection for usage text from usage rul
  -t, --to         email address to send to
  -p, --template   mail sending command template string
  -b, --base       base href tag used in the wrapped up html
  -n, --no-html    do not wrap up the output with html tags
  -d, --days       days to shift from billing date
  -v, --verbose    Verbose mode (Multiple -v increase the verbosity)
```

## Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

All patches welcome. 
