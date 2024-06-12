# uff - unleashed ffuf fork

Custom [ffuf](https://github.com/ffuf/ffuf) fork that relies on modified net/http and net/url libraries to avoid strict header and URL parsing.
This effectively makes it possible to send various malformed requests, such as:

### Invalid url encoded character:

```
GET /%9f HTTP/1.1
Host: example.com


```

### Invalid header:

```http
GET /%9f HTTP/1.1
Host: example.com
   I AM AN INVALID: HEADER


```

### No header canonization

```http
GET /hi HTTP/1.1
Host: t93lmkg4rn3bdmoq9ujd9acr8ie94xvlk.oastify.com
User-Agent: Fuzz Faster U Fool v2.1.0 sw33tlie unleashed-dev
lowercase-header: wexxxh
Accept-Encoding: gzip, deflate, br
Connection: keep-alive


```

Note how `lowercase-header` starts with a lowercase `l`.

## Version

Currently based on ffuf 2.10-dev

## Installation

To install `uff`, run the following command:

```bash
go install github.com/sw33tLie/uff@latest
```