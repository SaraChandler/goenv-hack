# goenv-hack
Goenv like tool for Hackathon 2022

## Features

### Install
Install a given version of Go

Usage: `goenv-hack install <version>`
Version can be in the format of `major.minor` or `major.minor.patch`. The OS and architecture of
the installed version will match the host machine.


### Use
Change currently linked Go version to the specified installed version.

Usage: `goenv-hack use <version>`
If the given version is not installed, it will be installed automatically.


### List
List currently installed go versions

Usage: `goenv-hack list`


### Uninstall/Teardown TBD
Uninstall Versions

Usage: `goenv-hack uninstall <version>`
Uninstalls specified version if it is not currently active.
