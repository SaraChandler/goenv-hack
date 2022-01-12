package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hashicorp/errwrap"
)

// Simple download function to download a given url to a path
func Download(url string, path string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errwrap.Wrapf("error in Download: {{err}}", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected 200, got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// Create path and file
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	filename := filepath.Base(url)
	fullFilename := filepath.Join(path, filename)
	file, err := os.Create(fullFilename)
	if err != nil {
		return err
	}

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Printf("Downloaded %s, size: %d\n", fullFilename, size)

	return nil
}

// Extract a given compressed file to a destination directory
func Extract(source string, destination string) error {
	return nil
}

// Check that goenv-hack is initialized
// check that InstallBaseDir exists
// check that InstallBin is in path
func CheckInit() error {
	return nil
}

// create InstallBaseDir and add InstallBin to path
func Init() error {
	return nil
}
