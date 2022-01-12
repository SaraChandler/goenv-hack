package pkg

import (
	"archive/tar"
	"compress/gzip"
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
	gzipStream, err := os.Open(source)
	if err != nil {
		return err
	}

	tarStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(tarStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		destPath := filepath.Join(destination, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			err = os.Mkdir(destPath, 0755)
			if err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(destPath)
			if err != nil {
				return err
			}
			_, err = io.Copy(outFile, tarReader)
			if err != nil {
				return err
			}
			outFile.Close()

		default:
			return err
		}

	}
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
