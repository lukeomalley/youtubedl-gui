package main

import (
	"os/exec"
)

type Downloader struct{}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (d *Downloader) DownloadFromYoutube(url string) int {
	// youtube-dl -x -f bestaudio 'https://www.youtube.com/watch?v=Y4rd8caU4ow'
	// -o "C:/Users/user/Downloads/youtube-dl/%(title)s.%(ext)s"
	cmd := exec.Command("youtube-dl", "-x", "-f", "bestaudio", "-o", "/home/luke/yt/%(title)s.%(ext)s", url)

	err := cmd.Run()
	if err != nil {
		return 1
	}

	return 0
}
