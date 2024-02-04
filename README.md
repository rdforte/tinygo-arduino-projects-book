# Arduino Projects Book

in order to run in NVim we use [tinygo.nvim](https://github.com/pcolladosoto/tinygo.nvim)

## Setup

[Tinygo getting started MacOS](https://tinygo.org/getting-started/install/macos/)

Note that all packages are setup as part of my .dotfiles.

## Usage

In the Project file where you are writing the source code for the arduino:

```
:TinyGoSetTarget arduino
```

This will enable us to get LSP support for Tinygo specific packages.

## Building and Deploying to Arduino

note you will have to find the device serial port in /dev (short for device) dirctory. It will be along the lines of `tty.usbmodem1401`.
If it does not show try unplugging the usb cable and then plugging it back in.

```
tinygo flash -target=arduino -port=/dev/tty.usbmodem1401 Project00/main.go
```
