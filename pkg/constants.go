package pkg

const InstallBaseDir string = "/usr/local/goenv-hack/"
const InstallVersionsDir string = "/usr/local/goenv-hack/versions/"

// /usr/local/goenv-hack/<version>/bin
// link /usr/local/gonev-hack/bin -> /usr/local/goenv-hack/<version>/bin
// link /usr/local/gonev-hack/bin -> /usr/local/goenv-hack/1.16.8/bin
// goenv-hack use 1.15.5
// link /usr/local/gonev-hack/bin -> /usr/local/goenv-hack/1.15.5/bin
const InstallBin string = "/usr/local/gonev-hack/bin"
