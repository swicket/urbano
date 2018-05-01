package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GitHub struct{}

func (GitHub) IsActive() bool {
	return true
}

func (GitHub) GetLatestVersion() string {
	res, err := http.Get("https://api.github.com/repos/alfio-event/alf.io/releases")
	if err != nil {
		log.Fatal(err)
		return Error
	}
	rel, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return Error
	}
	r := findLatestRelease(rel)
	if r != nil {
		return r["tag_name"].(string)
	}
	return Error
}

func findLatestRelease(rel []byte) map[string]interface{} {
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
