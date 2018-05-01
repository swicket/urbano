package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/swicket/urbano/provider"
)

func main() {
	providers := getActiveProviders()

	if len(providers) == 0 {
		log.Fatal("no active providers have been found. Aborting.")
	}
	outDir := "/home/alfio"
	if len(os.Args) > 1 {
		outDir = os.Args[1]
	}
	p := providers[0]
	v := p.GetLatestVersion()
	if v == provider.Error {
		log.Fatal("Got Error from provider [", p.Name(), "]")
	}
	url := p.GetArtifactUrl(v)
	outFile := filepath.Join(outDir, fmt.Sprintf("/alfio-%s-boot.war", v))
	log.Print("Provider [", p.Name(), "] says: ", v)
	if _, err := os.Stat(outFile); os.IsNotExist(err) {
		log.Print("About to download ", url, " => ", outFile)
		downloadNewRelease(url, outFile)
		performDeployment(outFile)
	} else {
		log.Print("File ", outFile, " exists, nothing to be done here.")
	}
}

//source: https://github.com/cavaliercoder/grab
func downloadNewRelease(url string, to string) {
	client := grab.NewClient()
	req, _ := grab.NewRequest(to, url)

	// start download
	log.Print("Downloading ", req.URL(), "...")
	resp := client.Do(req)
	log.Print("  ", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()
Loop:
	for {
		select {
		case <-t.C:
			log.Printf("  transferred %v / %v bytes (%.2f%%)",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		log.Fatal("Download failed:", err)
	}
	log.Printf("Download saved to %v", resp.Filename)
}

//steps:
//1 - stop the service
//	here we expect that the user running urbano has the following privilege:
//  {username} ALL=NOPASSWD: /usr/bin/systemctl stop alfio
//  {username} ALL=NOPASSWD: /usr/bin/systemctl start alfio
//
//2 - replace the symlink [dir]/alfio-boot.war
//
//3 - start the service

func performDeployment(artifactPath string) error {

	log.Print("Performing deployment of ", artifactPath)

	if err := callSystemCtl("stop"); err != nil {
		log.Fatal(err)
	}

	log.Print("delete symlink ")

	symlinkPath := filepath.Join(filepath.Dir(artifactPath), "alfio-boot.war")
	if _, err := os.Lstat(symlinkPath); err == nil {
		log.Print("removing symlink ", symlinkPath)
		os.Remove(symlinkPath)
	}

	log.Print("recreate symlink ")

	if err := os.Symlink(artifactPath, symlinkPath); err != nil {
		log.Fatal(err)
	}

	log.Print("restart service")

	return callSystemCtl("start")
}

func callSystemCtl(command string) error {
	cmd := exec.Command("sudo", "/usr/bin/systemctl", command, "alfio")
	return cmd.Run()
}

func getActiveProviders() (ret []provider.VersionProvider) {
	for _, p := range allProviders() {
		if p.IsActive() {
			ret = append(ret, p)
		}
	}
	return ret
}

func allProviders() []provider.VersionProvider {
	return []provider.VersionProvider{provider.Swicket{}, provider.GitHub{}}
}
