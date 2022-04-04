package rest_client

import "time"

const (

	// Max image size to send to server in bytes (1MB)
	MAX_SIZE = 1048576

	// POST timeout in seconds
	TIMEOUT = 30 * time.Second

	// All error codes
	USAGE_ERROR     = 1
	FILE_NOT_FOUND  = 2
	INVALID_IMAGE   = 3
	STAT_IMAGE_SIZE = 4
	IMAGE_SIZE_BIG  = 5
	READ_IMAGE      = 6
	JSON_ENCODE     = 7
	MALFORMED_URL   = 8
	POST_ERROR      = 9
	GZIP_ERROR      = 10
	DECODE_ERROR    = 11
	SERVER_ERROR    = 12

	// If users need translations
	USAGE_MSG           = "\nREST-Go client to convert image in ASCII\n\nUsage:\n\tclient [URL] [JPG_image]\n\tclient http://127.0.0.1:8080/toascii image.jpg"
	FILE_NOT_FOUND_MSG  = "image file not found"
	STAT_IMAGE_SIZE_MSG = "cannot stat image file size"
	IMAGE_SIZE_BIG_MSG  = "image is too big"
	INVALID_IMAGE_MSG   = "invalid image file"
	READ_IMAGE_MSG      = "cannot read image data"
	JSON_ENCODE_MSG     = "cannot encode data in json"
	MALFORMED_URL_MSG   = "malformed URL"
	POST_ERROR_MSG      = "cannot POST data to server"
	GZIP_ERROR_MSG      = "cannot decompress image data"
	DECODE_ERROR_MSG    = "cannot decode image data"
)
