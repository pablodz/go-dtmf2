package test

import (
	"testing"

	"github.com/pablodz/go-dtmf2/dtmf"
)

func TestDTMFDecoding(t *testing.T) {
	fileName := "123456654321.raw"
	decodedValue, err := dtmf.DecodeDTMFFromFile(fileName, 8000.0, 12)
	if decodedValue != "123456654321" {
		t.Errorf("Decoded value is incorrect, got: %s, want: %s", decodedValue, "123456654321")
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	fileName = "147258369.raw"
	decodedValue, err = dtmf.DecodeDTMFFromFile(fileName, 8000.0, 12)
	if decodedValue != "147258369" {
		t.Errorf("Decoded value is incorrect, got: %s, want: %s", decodedValue, "147258369")
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
