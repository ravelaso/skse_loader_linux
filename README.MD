# SKSE Loader (LINUX)

This aims to wrap the skse64_loader and avoid "Couldn't load arguments" / "Couldn't read arguments" if you intend to load the game from Steam in Linux.


## Prerequisites

1. Install the proper skse64 version.
2. Install mods manually which requires skse.

## Installation

1. Rename the 'SkyrimSELauncher.exe' to 'OldSkyrimSELauncher.exe'
2. Paste the new SkyrimSELauncher.exe
3. Launch the game from Steam.

## Compile

run: 
```sh
GOOS=windows GOARCH=amd64 go build -o SkyrimSELauncher.exe main.go
```