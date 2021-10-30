# cleanup folder
A simple CLI use to cleanup old folder

# Building

```shell script
$ go build
$ ./cleanup
Cleanup is a CLI application to remove old folder by max number of hours of exist.

Usage:
  cleanup [flags]
  cleanup [command]

Available Commands:
  folder    Remove old folder by max number of hours of exist.
  help      Help about any command.

Flags:
  -h, --help            help for cleanup
  -v, --version         version for cleanup

Use "cleanup [command] --help" for more information about a command.
```

# Folder

## Help

```shell script
$ ./cleanup folder --help

Usage:
  cleanup folder [max number of hours (integer) the old folder to be cleanup] [target folder] [flags]

Flags:
  -h, --help          help for folder
```
## Running

```shell script
$ ./cleanup folder 24 /home/dactoan/upload
2 old item found 
remove item /home/dactoan/upload/Screenshot from 2021-10-05 19-10-41.png 
remove item /home/dactoan/upload/page
cleanup /home/dactoan/upload completed

$ ./cleanup folder 24 /home/dactoan/download
0 old item found
```
