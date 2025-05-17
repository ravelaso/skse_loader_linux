# Bethesda Script Extender Loader (LINUX)

This aims to wrap the SKSE, FOSE, NVSE and avoid "Couldn't load arguments" / "Couldn't read arguments" 
if you intend to load the game from Steam in Linux.


## Prerequisites

1. Install the proper SE version for your game (FOSE, SKSE, NVES).
2. Install mods manually which requires SE, normally in (Data/YourSEVersion/Plugins/).

## Installation

1. Rename the original game launcher, example: 'SkyrimSELauncher.exe' to 'OldSkyrimSELauncher.exe'
2. Paste the new launcher and rename it to your game, like: SkyrimSELauncher.exe
3. Launch the game from Steam.
4. Script Extender for your game should hookup and work as intended.
5. Log file will be created in the game dir as: linux_loader.log

## Compile

run: 
```sh
GOOS=windows GOARCH=amd64 go build -o Launcher.exe main.go
```