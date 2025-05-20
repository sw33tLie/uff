# uff - unleashed ffuf fork

Custom [ffuf](https://github.com/sw33tLie/uff) fork that relies on modified net/http and net/url libraries to avoid strict header and URL parsing, with some additional customizations too.

## Installation

To install `uff`, run the following command:

```bash
go install github.com/sw33tLie/uff@latest
```

# Use cases

### Absolute URI FUZZING:

`echo hi | uff -c -u http://example.com -w - -opaque "http://127.0.0.1/FUZZ"`
 
```
GET http://127.0.0.1/hi HTTP/1.1
Host: example.com


```

Absolute URI fuzzing is often valuable, but not supported in regular ffuf.

### Invalid header:

`echo "hi" | uff -c -u http://example.com/FUZZ -w - -H '   I AM AN INVALID: HEADER'`

```http
GET /hi HTTP/1.1
Host: example.com
   I AM AN INVALID: HEADER


```

This allows all sorts of malformed headers.
You can even have a line without a colon!

### No header canonization

`echo hi | uff -c -u http://example.com/FUZZ -w - -H 'lowercase-header: yes'`

```http
GET /hi HTTP/1.1
Host: example.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36
lowercase-header: yes
Accept-Encoding: gzip, deflate, br
Connection: keep-alive


```

Note how `lowercase-header` starts with a lowercase `l`.

### Abusing the HTTP method for total control of the first request line

```
echo "hi" | uff -u "http://1y0urh3t90pydzt4ws733i1gv71zppde.oastify.com/" -w - -X "GET /FUZZ HTTP/1.0
x: x"
```

```http
GET /hi HTTP/1.0
x: x / HTTP/1.1
Host: 1y0urh3t90pydzt4ws733i1gv71zppde.oastify.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36
Accept-Encoding: gzip


```

Note how the first line `GET /hi HTTP/1.0` is actually part of the HTTP method, but we can fuzz it!!!
In this example we're sending an `HTTP/1.0` request, something you can't do with regular ffuf.

Lots of things suddently become possible :)

## Other customizations

- Legit user agent instead of ffuf's default `Fuzz Faster U Fool` one.
- New flag: `-no-content-length` to send a body in a request even without a `Content-Length` header.

## Version

Currently based on ffuf 2.10-dev

