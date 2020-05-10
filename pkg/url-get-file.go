package pkg

import (
	"errors"
	"io"
	"net/http"
	"os"
)

type File struct {
	url            *string
	targetFilePath *string
}

func (h *File) HttpGet() error {
	// Get the data
	resp, err := http.Get(*h.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(*h.targetFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func NewFile(config map[string]*string) (*File, error) {

	if *config["Url"] == "" {
		return nil,errors.New("Unable to find property Url")
	}

	if *config["TargetFilePath"] == "" {
		return nil,errors.New("Unable to find property TargetFilePath")
	}

	return &File{
		url:            config["Url"],
		targetFilePath: config["TargetFilePath"],
	}, nil
}
