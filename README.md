# Arduino Projects Book

in order to run in NVim run we use [tinygo.vim](https://github.com/sago35/tinygo.vim)

## Usage

```
:TinygoTarget
```

select:

```
arduino
```

Run:

note you will have to find the device serial port in /dev (short for device) dirctory. It will be along the lines of `tty.usbmodem1401`.

```
tinygo flash -target=arduino -path=/dev/tty.usbmodem1401 /path/to/code
```
