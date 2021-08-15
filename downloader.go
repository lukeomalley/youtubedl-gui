package main

import (
	"os/exec"
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

func (d *Downloader) DownloadFromYoutube(url string, format DownloadFormat) int {

	var cmd *exec.Cmd
	switch format {
	case AUDIO_ONLY:
		cmd = exec.Command("youtube-dl", "-x", "-f", "bestaudio", "-o", "/home/luke/yt/%(title)s.%(ext)s", url)
	case VIDEO_AND_AUDIO:
		cmd = exec.Command("youtube-dl", "-f", "bestvideo+bestaudio", "-o", "/home/luke/yt/%(title)s.%(ext)s", url)
	default:
		cmd = exec.Command("youtube-dl", "-x", "-f", "bestaudio", "-o", "/home/luke/yt/%(title)s.%(ext)s", url)
	}

	err := cmd.Run()
	if err != nil {
		return 1
	}

	return 0
}
