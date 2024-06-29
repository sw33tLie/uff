# uff - unleashed ffuf fork

Custom [ffuf](https://github.com/ffuf/ffuf) fork that relies on modified net/http and net/url libraries to avoid strict header and URL parsing.

## Installation

To install `uff`, run the following command:

```bash
go install github.com/sw33tLie/uff@latest
```

# Use cases

This effectively makes it possible to send various malformed or unsupported requests, such as:

### Absolute URI FUZZING:

`uff -c -u http://example.com -w vhosts.txt -opaque "http://FUZZ/"`
 
```
GET http://anything-here/ HTTP/1.1
Host: example.com


```

### Arbitrary HTTP method:

`uff -c -u http://example.com/FUZZ -w wordlist.txt -X ASDASD`
 
```
ASDASD /hi HTTP/1.1
Host: example.com


```

This is not possible in the normal ffuf because the net/http library only allows RFC-compliant HTTP methods.


### Invalid url encoded character:

`echo "%9f" | uff -c -u http://example.com/FUZZ -w -`

```
GET /%9f HTTP/1.1
Host: example.com


```

### Invalid header:

`echo "%9f" | uff -c -u http://example.com/FUZZ -w - -H '   I AM AN INVALID: HEADER'`

```http
GET /%9f HTTP/1.1
Host: example.com
   I AM AN INVALID: HEADER


```

### No header canonization

`uff -c -u http://example.com/FUZZ -w wordlist.txt -H 'lowercase-header: weh'`

```http
GET /hi HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36
lowercase-header: weh
Accept-Encoding: gzip, deflate, br
Connection: keep-alive


```

Note how `lowercase-header` starts with a lowercase `l`.

## Other customizations

- Legit user agent instead of ffuf's default `Fuzz Faster U Fool` one.

# Caveats

As of now, since the automatic content-length header calculation code was removed from net/http, you need to set your own Content-Length header if you're using the -d flag for POST data.

## Version

Currently based on ffuf 2.10-dev

