package utils

import (
	"os/exec"
)

func Ffmpeg(videoPath, outputPath string) error {
	cmd := "ffmpeg -i " + videoPath + " -f image2 -t 0.001 " + outputPath
	err := exec.Command("cmd", "/c", cmd).Run()
	return err
}
