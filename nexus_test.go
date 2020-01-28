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
	reader := DownloadCQLFile("1.1.6-RC1")

	if reader == nil {
		t.Error("Reader Error")
	}
}

// func TestParseInitCQL(t *testing.T) {
// 	reader := DownloadCQLFile("1.1.6-RC1")

// 	ParseInitCQL(reader)

// 	if reader == nil {
// 		t.Error("Reader Error")
// 	}
// }
