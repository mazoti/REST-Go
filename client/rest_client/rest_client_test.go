package rest_client

import (
	"os"
	"testing"
)

func TestClientImageNotFound(t *testing.T) {
	args := []string{"http://127.0.0.1:8080/toascii", "none"}

	_, err := Run(args[0], args[1])
	if err.Code() != FILE_NOT_FOUND {
		t.Fatalf("FAILED")
	}
}

func TestClientImageTooBig(t *testing.T) {
	args := []string{"http://127.0.0.1:8080/toascii", "a.out"}

	var data [MAX_SIZE + 1]byte
	slice := data[:]
	os.WriteFile(args[1], slice, 0644)
	defer os.Remove(args[1])

	_, err := Run(args[0], args[1])
	if err.Code() != IMAGE_SIZE_BIG {
		t.Fatalf("FAILED")
	}
}

func TestClientImageJPG(t *testing.T) {
	args := []string{"http://127.0.0.1:8080/toascii", "a.out"}
	var data [2]byte
	slice := data[:]
	os.WriteFile(args[1], slice, 0644)
	defer os.Remove(args[1])

	_, err := Run(args[0], args[1])
	if err.Code() != INVALID_IMAGE {
		t.Fatalf("FAILED")
	}
}
