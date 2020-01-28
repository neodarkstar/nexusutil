package nexusutil

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const group = "com.ac"
const artifact = "acx-core_2.11"
const dseHost = "172.22.4.11"
const nexusHost = "172.24.1.30"

// Artifact contains the information of the cql
type Artifact struct {
	DownloadURL string `json:"downloadUrl"`
	Path        string `json:"path"`
	ID          string `json:"id"`
	Repository  string `json:"repository"`
	Format      string `json:"format"`
}

// NexusResults Response
type NexusResults struct {
	Items             []Artifact `json:"items"`
	ContinuationToken string
}

// GetDownloadURL Queries Nexus for the artifact download url
func GetDownloadURL(version string) string {
	url := fmt.Sprintf("http://%s/service/rest/v1/search/assets?group=%s&name=%s&maven.baseVersion=%s&maven.extension=cql", nexusHost, group, artifact, version)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	results := NexusResults{}

	err = json.NewDecoder(res.Body).Decode(&results)

	if err != nil {
		log.Fatal(err)
	}

	return results.Items[0].DownloadURL
}

// DownloadCQLFile Download the init.cql file from nexus
func DownloadCQLFile(version string) ([]byte, io.Reader) {
	url := GetDownloadURL(version)
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body, res.Body
}
