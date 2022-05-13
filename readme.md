## tilecat in go

tiny tool to merge images into a single image with custom patterns of rows and cols


## Usage
```sh
NAME:
   tilecat - concat images into a single image

USAGE:
   tilecat [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dir value          directory to retreive the tiles
   --rowCount value     number of output rows (default: 0)
   --columnCount value  number of output cols (default: 0)
   --out value          output file
   --help, -h           show help (default: false)
```

## Example

```sh
tilecat --dir ./resources --rowCount 1 --columnCount 4 --out ./image.png
```

## the result

![merged image](https://github.com/apotox/tilecat/blob/master/result/image.png?raw=true)
