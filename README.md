# WSO - ScummAtlas

This is Warriorstar's fork of the [ScummAtlas][] project by Miguel Gonzalez (ktzar).

[ScummAtlas]: https://github.com/ktzar/scummatlas

## Installation

Simply download the latest release binary, `scummatlas.exe`. Put it somewhere
near your SCUMM games.

Once downloaded, open up a terminal/command line, and run it as follows (using
_Indiana Jones and the Fate of Atlantis_ as an example directory):

`scummatlas.exe -gamedir "D:/Path/To/Games/ATLANTIS" -outputdir "out/" -serve`

This will generate all of the files for the webpage, as well as start a local
web server so you can access them easily. Press Ctrl-C to stop the web server.

## Building

This fork has been migrated to a Go module.

```bash
git clone git@github.com:warriorstar-orion/scummatlas.git
cd scummatlas
go run cmd/scummatlas/main.go -gamedir "D:/Path/To/Games/ATLANTIS/" -outputdir "out/" 
```

## Original README

Scumm games parser and HTML Atlas generator written in Golang.

The aim of this software is to provide an easy to understand, testable, and organised software that unpacks games using the SCUMM engine.

It creates a bunch of HTML files that software that aim to show the inner aspects of how these beloved games were implemented.

TODO
- Improve parsing local and object scripts
- ~~Parse costumes~~
- ~~Have a game detector~~
- Add ordering capabilities to table view
- Use c64 font in some places like titles
