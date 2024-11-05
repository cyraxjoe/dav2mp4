# dav2mp4

Video converter from dahua video format to mp4 (among others).

The implementation relies on the closed source dahua play api (bundled in this repo), which in general tends to be faster
and more efficient than using ffmpeg for this particular use case (dramatically more efficient, like 0.01% the time that full transcode in ffmpeg takes).


## Installation

Currently the recommended way to build and install this program is using [Nix](https://nixos.org/guides/install-nix.html).

Alternatively, there should be windows installer in the releases section.


To install it:

    nix-env  -f ./ -i

Or to test it in a nix-shell:

    nix-shell


### Compatibility

Linux & Windows only, I might consider to support other platforms as they become relevant to the project.

## CLI interface
    dav2mp4 [options]  <input-dav-file-or-dir> <output-file-or-dir>
    Options:
            -v, --version: Show the version number and exit.
            -h, --help: Show this help message and exit.
            -f, --format: Video output format, supported: mp4 (default), raw, avi and asf.
            -b, --batch-mode: Use the positional arguments as input and output directories.
