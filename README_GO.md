# dav2mp4 (Go Port)

This is a port of the `dav2mp4` project from Nim to Go, including a native GUI.

## Prerequisites

- Go 1.18+
- GCC/Clang (for cgo)
- Linux or Windows
- Dahua Play SDK (bundled in `vendor/linux/dhplay` or `vendor/windows/dhplay`)

### Linux Dependencies

You need OpenGL and X11 development headers for the GUI (Fyne).

On Ubuntu/Debian:
```bash
sudo apt-get install libgl1-mesa-dev xorg-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev libxv-dev
```

## Building

### CLI

```bash
go build ./cmd/dav2mp4
```

### GUI

```bash
go build ./cmd/dav2mp4-gui
```

## Running

The application requires `libdhplay.so` (Linux) or `play.dll` (Windows) to be in the library path.

### Linux

```bash
export LD_LIBRARY_PATH=$(pwd)/vendor/linux/dhplay
./dav2mp4 -h
./dav2mp4-gui
```

## Usage

### CLI

```
dav2mp4 [options] <input-dav-file-or-dir> <output-file-or-dir>
Options:
  -f, --format <format>     Video output format (mp4, raw, avi, asf) [default: mp4]
  -b, --batch-mode          Use positional arguments as input and output directories
  -v, --version             Show version number and exit
  -h, --help                Show this help message and exit
```

### GUI

Run `dav2mp4-gui` to launch the graphical interface. Select input file and output directory, choose format, and click Convert.
