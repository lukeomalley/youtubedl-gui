package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"os/user"
)

type DownloadFormat int

const (
	AUDIO_ONLY = iota
	VIDEO_AND_AUDIO
)

type Downloader struct{}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (d *Downloader) DownloadFromYoutube(url string, format DownloadFormat, dir string) int {
	directory := fmt.Sprintf("%s/%s/%%(title)s.%%(ext)s", d.userHomeDir(), dir)

	var cmd *exec.Cmd

	switch format {
	case AUDIO_ONLY:
		cmd = exec.Command("youtube-dl", "-x", "-f", "bestaudio", "-o", directory, url)
	case VIDEO_AND_AUDIO:
		cmd = exec.Command("youtube-dl", "-f", "bestvideo+bestaudio", "-o", directory, url)
	default:
		cmd = exec.Command("youtube-dl", "-x", "-f", "bestaudio", "-o", directory, url)
	}

	err := cmd.Run()
	if err != nil {
		return 1
	}

	return 0
}

func (d *Downloader) ListHomeDirectories() []string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	files, err := ioutil.ReadDir(user.HomeDir)
	if err != nil {
		log.Fatal(err)
	}

	var dirs []string
	for _, f := range files {
		if string(f.Name()[0]) == "." || !f.IsDir() {
			continue
		}

		dirs = append(dirs, string(f.Name()))
	}

	return dirs
}

func (d *Downloader) userHomeDir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return user.HomeDir
}
