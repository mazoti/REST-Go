package rest_server

import "time"

const (

	// Usage return codes to OS
	USAGE_ERROR    = 2
	BIND_ERROR     = 3
	SHUTDOWN_ERROR = 4

	// POST URL address to receive the image file
	URL = "/toascii"

	// Max http headers size in bytes
	MAX_HEADERS_SIZE = 1024

	// Max image size in bytes to process
	MAX_SIZE = 1048576

	// Server timeouts in seconds
	READ_TIMEOUT  = 30 * time.Second
	WRITE_TIMEOUT = 60 * time.Second

	// Darker characters on the left
	DICTIONARY = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,\"^`'. "

	// If users need translations
	USAGE_MSG = "REST-Go server to convert JPG image in ASCII\n\nUsage:\n\tserver [IP:port]\n\tserver 127.0.0.1:8080"
	START_MSG = "Server started and running, CTRL+C to quit"
	END_MSG   = "CRTL+C, server finished"

	READ_DATA          = "cannot read data"
	IMAGE_SIZE         = "image received is bigger than MAX_SIZE"
	DECODE_ERROR       = "cannot decode data"
	INVALID_JPEG       = "invalid JPG image"
	NOT_FOUND          = "not found"
	METHOD_NOT_ALLOWED = "method not allowed"
)
