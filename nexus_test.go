package nexusutil

import (
	"testing"
)

func TestGetDownloadUrl(t *testing.T) {
	url := GetDownloadURL("1.1.6-RC1")

	if url == "" {
		t.Error("Error")
	}
}

func TestDownloadCQLFile(t *testing.T) {
	body, reader := DownloadCQLFile("1.1.6-RC1")

	if len(body) == 0 {
		t.Error("Error")
	}

	if reader == nil {
		t.Error("Reader Error")
	}
}
