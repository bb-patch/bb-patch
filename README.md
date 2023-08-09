# Beasts and Bumpkins Fan Patch

This project is fan patch for the 1997 game Beasts and Bumpkins developed by Worldweaver Ltd, published under Electronic Arts.

Beasts and Bumpkins currently only exists on a wishlist for GOG.com and abandonware websites and old dusty shelves of people in their thirties. The game is not easily playable in the current age of windows 10 for various reasons, and playing it requires a lot of work. But not any more!

This fan patch makes it possible to just pop in the CD / fetch it from some abandonware website, install it, patch in some extra features and run it!

## Usage

The project has only been tested with BEASTS.EXE having md5sum 2d2852fcb22e6f65a0cf3415abaf4411. This is generally the one distributed on abandonware websites's ISO files or with the CD that was distributed in the EU. The RIP file versions on abandonware websites usually have a different md5sum and the patches has not been tested with this.

To install the patch:
1. Mount CD file to virtual drive, or insert your own CD.
2. Download the bb-patch-x.xx.exe file from the [release page](https://github.com/bb-patch/bb-patch/releases).
3. Place the bb-patch.exe in the folder where you want the game files.
4. Execute the bb-patch.exe.
5. Follow the UI to install and patch the game.

Alternatively, you can extract all the game files to a folder and place the bb-patch-x.xx.exe there and start it to patch it. The installation process of the patch just copies the files from specified folder to the folder the bb-patch.exe file resides in.

## Features

The patch contains the following configurations:

1. Fixes the windows registry needed to play the game.
2. Remove the need for a CD.
3. Fixes the gameplay lag that exists on new computers.
4. Enabling developer console on F11 (cheats 😀).
5. Change the resolution of the game. Though this currently contain some graphical glitches and some icons not being where they should be, the game is very much playable. The original game was in 640x480, so bumping it up to 1280x720 is a nice quality of life improvement.

## Build
The project uses Go with the UI toolkit `fyne`. Follow the installation guideline over at https://github.com/fyne-io/fyne to install fyne.

Otherwise it is just `go build` or `./build.sh` to build a release.

## Acknowledgments
One of the small patches uses the DxWrapper from 
https://github.com/elishacloud/dxwrapper. The files needed are included in the repository and embedded into one of the .go files.

This [reddit thread](https://www.reddit.com/r/widescreengamingforum/comments/14ktbw8/beasts_and_bumpkins/) contained some nice insperation and help with the resolution patch.

This [blog](https://www.dynamic-mess.com/jeux/retro-engineering-beasts-and-bumpkins/) shows how cheats are enabled.

## Contributing
If you'd like to contribute to this project, please follow the guidelines in CONTRIBUTING.md.

## License
This project is licensed under the GNU General Public License v3.0.