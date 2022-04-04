/*
	Server object of the project.
	Contains one object for the server and one for the errors (ServerError)
*/

package rest_server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

// All colors converted to char
var Dictionary [256][256][256]byte

// Initialize the dictionary first for faster access
func init() {
	for red := 0; red < 256; red++ {
		for green := 0; green < 256; green++ {
			for blue := 0; blue < 256; blue++ {
				gray := 0.21*float64(red) + 0.72*float64(green) + 0.07*float64(blue)
				index := math.Ceil(float64(len(DICTIONARY)-1) * gray / 255)
				Dictionary[red][green][blue] = DICTIONARY[int(index)]
			}
		}
	}
}

type Server struct{}

/*
	Starts the server in the IP:PORT and stays listening.
	Stops on any error or when user enters CTRL+C.
	Everything is logged on stdout.
*/
func Start(URL_Port string) *ServerError {

	// Use all cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	handler := Server{}

	srv := &http.Server{
		Addr:           URL_Port,
		Handler:        handler,
		ReadTimeout:    READ_TIMEOUT,
		WriteTimeout:   WRITE_TIMEOUT,
		MaxHeaderBytes: MAX_HEADERS_SIZE}

	// Get CRTL+C signal to quit in the right way
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	var bind_err *ServerError

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			bind_err = &ServerError{err.Error(), BIND_ERROR}
			sig <- os.Kill
		}
	}()

	// Rendezvous on error or shutdown
	status, more_values := <-sig
	if more_values == false {
		close(sig)
	}

	if status == os.Kill {
		return bind_err
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		return &ServerError{err.Error(), SHUTDOWN_ERROR}
	}

	return nil
}

/*
	Private method to convert any image to ASCII
	Got from the javascript version at https://www.jonathan-petitcolas.com/2017/12/28/converting-image-to-ascii-art.html
*/
func (server Server) toASCII(img image.Image) string {

	var result strings.Builder

	// Reserve 1 char for each pixel + 1 EOL (\n) for each line
	result.Grow(img.Bounds().Max.Y*(img.Bounds().Max.X + 1))

	// Calculate dictionary index with RGB color
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			// Ignore alpha channel and convert to 8 bits
			red, green, blue, _ := img.At(x, y).RGBA()
			result.WriteByte(Dictionary[red>>8][green>>8][blue>>8])
		}
		result.WriteByte('\n')
	}

	return result.String()
}

func (server Server) sendJSON(httpCode int, asciiData string, writer *http.ResponseWriter, request *http.Request, messages ...error) {

	(*writer).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*writer).Header().Set("Content-Encoding", "gzip")

	(*writer).WriteHeader(httpCode)

	// Compress string, convert to BASE64 and sends to client in JSON format
	gzWriter := gzip.NewWriter(*writer)
	defer gzWriter.Close()

	if err := json.NewEncoder(gzWriter).Encode([]byte(asciiData)); err != nil {
		log.Println(err)
		return
	}

	log.Printf("%s %s %s => %d %s", request.RemoteAddr, request.Method, request.RequestURI, httpCode, http.StatusText(httpCode))

	for _, err := range messages {
		log.Println(err)
	}
}

func (server Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		if request.RequestURI == URL {
			marshData, err := ioutil.ReadAll(request.Body)
			if err != nil {
				server.sendJSON(http.StatusInternalServerError, READ_DATA, &writer, request, err)
				return
			}

			if len(marshData) > MAX_SIZE {
				server.sendJSON(http.StatusRequestEntityTooLarge, IMAGE_SIZE, &writer, request)
				return
			}

			// Convert BASE64 inside JSON to binary data
			var data []byte
			if err = json.Unmarshal(marshData, &data); err != nil {
				server.sendJSON(http.StatusInternalServerError, DECODE_ERROR, &writer, request, err)
				return
			}

			// Check for a valid JPG file
			img, _, err := image.Decode(bytes.NewReader(data))
			if err != nil {
				server.sendJSON(http.StatusUnsupportedMediaType, INVALID_JPEG, &writer, request, err)
				return
			}

			// No errors, send ASCII image with 200 OK return code
			server.sendJSON(http.StatusOK, server.toASCII(img), &writer, request)
			return
		}

		server.sendJSON(http.StatusNotFound, NOT_FOUND, &writer, request)
		return
	}

	server.sendJSON(http.StatusMethodNotAllowed, METHOD_NOT_ALLOWED, &writer, request)
}
