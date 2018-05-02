package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GitHub struct{}

func (GitHub) IsActive() bool {
	return true
}

func (GitHub) GetLatestVersion() (string, error) {
	response, err := http.Get("https://api.github.com/repos/alfio-event/alf.io/releases")
	if err != nil {
		log.Printf("could not fetch version from Github: %v", err)
		return "", err
	}
	release, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("could not read release page content from Github: %v", err)
		return "", err
	}
	defer response.Body.Close()

	if r := findLatestRelease(release); r != nil {
		return r["tag_name"].(string), nil
	}
	return "", errors.New("could not get the latest release version from Github")
}

func findLatestRelease(rel []byte) map[string]interface{} {
	// it would be great if there was a Github client that could expose a struct
	// to map response to, so instead of unmarshalling to interface{} and casting
	// there would be a concrete type to work with
	var dat []map[string]interface{}
	if err := json.Unmarshal(rel, &dat); err != nil {
		return nil
	}
	for _, r := range dat {
		if (!r["prerelease"].(bool)) && (!r["draft"].(bool)) {
			return r
		}
	}
	return nil
}

func (GitHub) GetArtifactUrl(version string) string {
	return fmt.Sprintf("https://github.com/alfio-event/alf.io/releases/download/%s/alfio-%s-boot.war", version, version)
}

func (GitHub) Name() string {
	return "GitHub Alf.io Repository"
}
