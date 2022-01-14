package pkg

import (
  "fmt"
  "path/filepath"
  "runtime"
  "os"
)

func UninstallVersion(version string) error {
  downloadPath, versionPath := getInstallLocation(version)
  err := os.RemoveAll(downloadPath)
  if err != nil {
    return err
  }

  err = os.RemoveAll(versionPath)
  if err != nil {
    return err
  }
  fmt.Printf("Version %v is deleted\n", version)
  return nil
}

func getInstallLocation(version string) (string, string){
  os := runtime.GOOS
  arch := runtime.GOARCH
  downloadFile := fmt.Sprintf("go%s.%s-%s.tar.gz", version, os, arch)

  downloadPath := filepath.Join(getPath(DownloadDir), downloadFile)

  versionPath := filepath.Join(getPath(InstallVersionsDir), version)

  return downloadPath, versionPath
}
