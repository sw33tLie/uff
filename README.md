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

Currently based on ffuf 2.10