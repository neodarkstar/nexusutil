package nexusutil

import (
	"encoding/json"
	"fmt"
	"io"
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
func DownloadCQLFile(version string) io.ReadCloser {
	url := GetDownloadURL(version)
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	// defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return res.Body
}

// // ParseInitCQL Applies the variables to the init cql file
// func ParseInitCQL(init io.Reader) {
// 	keyspace := "unit_tests"
// 	class := "SimpleStrategy"
// 	replicationFactor := 3
// 	datacenter := "replication_factor"

// 	scanner := bufio.NewScanner(init)

// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		regexpKeyspace := regexp.MustCompile(`<KEYSPACE>`)
// 		regexpClass := regexp.MustCompile(`<STRATEGY>`)
// 		regexpDatacenter := regexp.MustCompile(`<DATACENTER>`)
// 		regexpReplFactor := regexp.MustCompile(`<REPLICATION_FACTOR>`)
// 		regexpSemicolon := regexp.MustCompile(`;`)

// 		state1 := regexpKeyspace.ReplaceAllLiteralString(line, keyspace)
// 		state2 := regexpClass.ReplaceAllLiteralString(state1, class)
// 		state3 := regexpReplFactor.ReplaceAllLiteralString(state2, strconv.Itoa(replicationFactor))
// 		state4 := regexpDatacenter.ReplaceAllLiteralString(state3, datacenter)
// 		state5 := regexpSemicolon.ReplaceAllLiteralString(state4, "")

// 		fmt.Printf("Executing: %s\n", state5)
// 	}
// }
