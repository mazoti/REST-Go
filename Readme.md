# REST-Go
## _Client-Server REST project written in Go_

This program converts JPG images to ASCII characters and display it on screen.
First the client reads the image, encodes in JSON and sends to server. The server decodes the JSON, converts the bitmap image to ASCII, compress the string, encodes in JSON and sends back to the client. The client decodes the JSON, decompress and displays on the screen.

#### Requirement
- A modern go compiler

#### Building the client
Make sure the requirement is in your PATH:
```sh
cd REST-Go
cd client
go build
```
If you are using the GNU Go compiler:
```sh
cd REST-Go
cd client
go build -compiler=gccgo
```

#### Building the server

```sh
cd REST-Go
cd server
go build
```
If you are using the GNU Go compiler:
```sh
cd REST-Go
cd server
go build -compiler=gccgo
```

### Make a simple test
Find a small JPG file and run each command in a different terminal **(run the server first):**
```sh
server 127.0.0.1:8080
client http://127.0.0.1:8080/toascii small_image.jpg 
```
The client prints the ASCII image on the screen.

#### Unit tests

```sh
cd REST-Go
cd client
cd rest_client
go test
```

```sh
cd REST-Go
cd server
cd rest_server
go test
```

For code coverage:

```sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

It is also good to run
```sh
go vet
```
on each module directory to look for errors.

#### Benchmarks (server only)

```sh
cd REST-Go
cd server
cd rest_server
go test -bench=.
```

Bugs, optimizations, support, send by email to my last name at gmail

## License

3-clause BSD

**Thanks for your time and have fun!**