# RAPIP

Rapid.Application.Program.Interface.Prototyping

Turns static files into a HTTP responses.

-------

## Why?

Existing API prototyping and documentation tools are great!

**But...**

Almost always add a layer of abstraction. There is no standard for this extra layer. This just means more stuff to know and understand.

Are limited to one or two response types.

Lack version/source control.

Someone else made them and I really wanted to reinvent a wheel.

------

## How?

You have a file containing a valid **HTTP Response** which is publicly available on a supported source (currently: github repo | github gist)

You figure out the proxy url (hopefully there are docs explaining this below)

You make the correct HTTP request.

RAPIP does some routing and fetches your file and answers with the response found in your file.

------

## Supported Sources

| Source | Method | Source URL | Proxy URL |
|- | - | -| -|
| github: gists | GET | https://gist.githubusercontent.com/PATH-TO-RAW-GIST | http://gist-github.rapip.mysterious-mountain.stream/PATH-TO-RAW-GIST |
| github: repos | * | https://raw.githubusercontent.com/PATH-FOLDER-CONTAINING-RAW-FILES/HTTP-METHOD | http://github.rapip.mysterious-mountain.stream/PATH-FOLDER-CONTAINING-RAW-FILES |

### Github : gists

Gist sources are limited to GET requests. Create a gist (can be private) with any random name and place a valid HTTP Response in it. View the **raw** gist and copy the url **path**. Join `http://gist-github.rapip.mysterious-mountain.stream` and the raw gist path.

You now have a working GET endpoint!

### Github : repos

Repo sources must be public (you really shouldn't trust any third party proxy with your private stuff). Create a folder and place a file in it with the HTTP method as file name. (`/GET` for `GET` requests,...)

View a raw file on any branch (can be `master`), commit to leverage source control. Copy the path and join with `http://github.rapip.mysterious-mountain.stream`. Omit the filename as that is determined by the HTTP method of the request.

You now have a working endpoint for as many methods as you like!

------

### Maybe features :

- [ ] https support (should really add this)
- [ ] some dynamic aspect to responses (request query part can be used to pass keys/values)
- [ ] ....

----

### What is `mysterious-mountain.stream`

It's a cheap domain and has no meaning at all. It is my personal playground.

------

### Disclaimer

Your mileage may vary.
