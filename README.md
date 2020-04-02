# crab
[![Travis CI](https://api.travis-ci.com/atomicptr/crab.svg?branch=master)](https://travis-ci.com/atomicptr/crab)
[![Go Report Card](https://goreportcard.com/badge/github.com/atomicptr/crab)](https://goreportcard.com/report/github.com/atomicptr/crab)
[![Coverage Status](https://coveralls.io/repos/github/atomicptr/crab/badge.svg?branch=master)](https://coveralls.io/github/atomicptr/crab?branch=master)

A versatile tool to crawl dozens of URLs from a given source, like a sitemap or an URL list.

Useful for:
* Warming site caches
* Checking response times
* Identifying dead or broken pages

## Install

### Binaries

[You can download the newest release from here for Linux, macOS and Windows.](https://github.com/atomicptr/crab/releases/)

### Snap

```bash
$ snap install crab
```

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/crab)

### Homebrew

```bash
$ brew install atomictr/tools/crab
```

### Scoop

```bash
$ scoop bucket add atomicptr https://github.com/atomicptr/scoop-bucket
$ scoop install crab
```

## Usage

Crawl singular URLs:

```bash
$ crab crawl https://domain.com https://domain.com/test
{"status": 200, "url": "https://domain.com", ...}
...
```

Crawl through a sitemap:

```bash
$ crab crawl:sitemap https://domain.com/sitemap.xml
```

Replace all URLs with a different one:

```bash
$ crab crawl:sitemap https://domain.com/sitemap.xml --prefix-url=https://staging.domain.com
```

Add some cookies/headers:

```bash
$ crab crawl:sitemap https://domain.com/sitemap.xml --cookie auth_token=12345 --header X-Bypass-Cache=1
```

### Filter by Status Code

You can filter the output by it's status code

```bash
# This will only return responses with a 200 OK
$ crab crawl:sitemap https://domain.com/sitemap.xml --filter-status=200
# This will only return responses that are not OK
$ crab crawl:sitemap https://domain.com/sitemap.xml --filter-status=!200
# This will only return responses between 500-599 (range)
$ crab crawl:sitemap https://domain.com/sitemap.xml --filter-status=500-599
# This will only return responses with 200 or 404 (multiple, be aware if one condition is true they all are)
$ crab crawl:sitemap https://domain.com/sitemap.xml --filter-status=200,404
# This will only return responses with a code greater than 500
$ crab crawl:sitemap https://domain.com/sitemap.xml --filter-status=>500
```

## License

MIT
