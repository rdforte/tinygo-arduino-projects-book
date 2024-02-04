# Arduino Projects Book

in order to run in NVim run we use [tinygo.vim](https://github.com/sago35/tinygo.vim)

## Usage

```
:TinyGoSetTarget arduino
```

Run:

note you will have to find the device serial port in /dev (short for device) dirctory. It will be along the lines of `tty.usbmodem1401`.
If it does not show try unplugging the usb cable and then plugging it back in.

```
tinygo flash -target=arduino -port=/dev/tty.usbmodem1401 Project00/main.go
```
