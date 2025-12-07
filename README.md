# 🔗 URL Shortener CLI

A simple and extensible command-line URL shortener written in Go, using TinyURL and is.gd APIs.
Built with Cobra CLI, featuring persistent JSON history storage and clean modular architecture.

## Features

- Shorten URLs from the command line
- Supports TinyURL and is.gd
- Automatic URL normalization (https://)
- Optional URL availability check
- Persistent JSON history stored in user's Documents folder
- Cross-platform (Windows / macOS / Linux)
- Modular Go project structure
- Built with Cobra CLI

## Requirements

- Windows (tested)  
- Go 1.22+
- Internet connection

## Installation

**1. Clone the repository**
```console
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
```
**Run without building**
```console
go run .
```
**Build executable**
```console
go build -o url-shortener
```
## Available Commands

### Shorten a URL

```console
url-shortener shorten [url] [flags]
```
#### Examples
**Shorten using TinyURL (default)**
```console
url-shortener shorten https://google.com
```
**Shorten using is.gd**
```console
url-shortener shorten https://google.com --service isgd
```
**Check URL availability before shortening**
```console
url-shortener shorten github.com/golang/go --check
```
**Flags for *shorten***

| Flag  | Description |
| ------------- |:-------------:|
|--service   | URL shortening service (tinyurl, isgd)     |
| --check      | Check if the URL is reachable before shortening   |

### View history
**Displays all previously shortened URLs (persistent)**
```console
url-shortener history
```
### Clear history
```console
url-shortener clear-history
```
### Show version
```console
url-shortener version
```
### Help
```console
url-shortener --help
url-shortener shorten --help
```

## History Storage
**Shortened URLs are stored in a JSON file located at:**
```console
~/Documents/url-shortener/history.json
```
#### Example:
```json
[
  {
    "original": "https://google.com",
    "short": "https://tinyurl.com/abc123",
    "service": "TinyURL"
  }
]
```







