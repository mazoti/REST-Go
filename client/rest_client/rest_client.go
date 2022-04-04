/*
	Client object of the project.
	Contains one object for the client and one for the errors (ClientError)
*/

package rest_client

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"image"
	_ "image/jpeg"
	"net/http"
	"os"
	"runtime"
)

type Client struct {
	image string
}

func (client Client) String() string {
	return client.image
}

/*
	Reads the JPG file, converts to BASE64 string, encodes in JSON and POSTs in the URL.
	Unzip the response, decodes the BASE64 JSON and prints de string in the stdout.
*/
func Run(URL string, image_path string) (*Client, *ClientError) {

	// 1 core is enough
	runtime.GOMAXPROCS(1)

	// Read the image file from command line argument
	inputImage, err := os.Open(image_path)
	if err != nil {
		return nil, &ClientError{FILE_NOT_FOUND_MSG, FILE_NOT_FOUND}
	}

	defer inputImage.Close()

	// Does not send images with size > MAX_SIZE
	imageSize, err := inputImage.Stat()
	if err != nil {
		return nil, &ClientError{STAT_IMAGE_SIZE_MSG, STAT_IMAGE_SIZE}
	}

	if imageSize.Size() > MAX_SIZE {
		return nil, &ClientError{IMAGE_SIZE_BIG_MSG, IMAGE_SIZE_BIG}
	}

	// Check for a valid JPG file
	if _, _, err = image.Decode(inputImage); err != nil {
		return nil, &ClientError{INVALID_IMAGE_MSG, INVALID_IMAGE}
	}

	inputImage.Seek(0, 0)

	// Copy all image data to RAM
	imageData := make([]byte, imageSize.Size())
	if _, err = inputImage.Read(imageData); err != nil {
		return nil, &ClientError{READ_IMAGE_MSG, READ_IMAGE}
	}

	// Serialize JPG bytes in JSON
	jsonData, err := json.Marshal(imageData)
	if err != nil {
		return nil, &ClientError{JSON_ENCODE_MSG, JSON_ENCODE}
	}

	// POST JSON data to server with no compression (JPG is already compressed)
	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, &ClientError{MALFORMED_URL_MSG, MALFORMED_URL}
	}

	// Set headers, timeout and POST image data BASE64 encoded in JSON to server
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Accept-Encoding", "gzip")

	client := &http.Client{Timeout: TIMEOUT}
	response, err := client.Do(request)
	if err != nil {
		return nil, &ClientError{POST_ERROR_MSG, POST_ERROR}
	}

	defer response.Body.Close()

	gzReader, err := gzip.NewReader(response.Body)
	if err != nil {
		return nil, &ClientError{GZIP_ERROR_MSG, GZIP_ERROR}
	}

	defer gzReader.Close()

	// Receive ASCII image in a JSON gzip compressed response
	var imageAscii []byte

	if err := json.NewDecoder(gzReader).Decode(&imageAscii); err != nil {
		return nil, &ClientError{DECODE_ERROR_MSG, DECODE_ERROR}
	}

	// Server error, description is on imageAscii variable
	if response.StatusCode != http.StatusOK {
		return nil, &ClientError{string(imageAscii), SERVER_ERROR}
	}

	return &Client{string(imageAscii)}, nil
}
