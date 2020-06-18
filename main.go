package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	video := "video.mp4"
	out := "out.jpg"

	err := ExtractAudioFromVideo(video, out)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fileExistError := FileExists(out)
	if fileExistError != nil {
		fmt.Printf("Error: %s\n", fileExistError.Error())
		return
	}
}

/**
	Check if file exist and its size
 */
func FileExists(f string) error {
	var err error
	if f, err := os.Stat(f); err == nil && f.Size() > 0 {
		return nil
	}
	return err
}

/**
	Extract audio from video and generate audio file
 */
func ExtractAudioFromVideo(inFile string, outFile string) error {
	// check if input file exists
	inFileError := FileExists(inFile)
	if inFileError != nil {
		fmt.Printf("Error: %s\n", inFileError.Error())
		return inFileError
	}

	// check if output file exists
	outFileError := FileExists(outFile)
	if outFileError != nil {
		fmt.Printf("Error: %s\n", outFileError.Error())
		return outFileError
	} else {
		// if exists then remove
		// removeFileError := os.Remove(outFile)
		// if removeFileError != nil {
		// 	fmt.Printf("Error: %s\n", removeFileError.Error())
		// 	return removeFileError
		// }
	}

	// extract audio from video using ffmpeg library
	cmd := exec.Command("ffmpeg", "-i", inFile, "-ss", "00:00:05.000", "-vframes", "1", outFile)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
