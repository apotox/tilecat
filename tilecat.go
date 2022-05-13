package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	tilecat := NewTileCat(nil)

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dir",
				Usage: "directory to retreive the tiles",
			},
			&cli.IntFlag{
				Name:  "rowCount",
				Usage: "number of output rows",
			},
			&cli.IntFlag{
				Name:  "columnCount",
				Usage: "number of output cols",
			},
			&cli.StringFlag{
				Name:  "out",
				Usage: "output file",
			},
		},
		Name:  "tilecat",
		Usage: "concat images into a single image",
		Action: func(c *cli.Context) error {

			options := &tileCatOptions{}

			if c.String("dir") == "" {
				return errors.New("no directory!")
			}
			options.dir = c.String("dir")

			if c.String("out") == "" {
				return errors.New("no output file!")
			}
			options.out = c.String("out")
			if c.Int("rowCount") == 0 || c.Int("columnCount") == 0 {
				return errors.New("no rowCount or columnCount!")
			}

			options.columnCount = c.Int("columnCount")
			options.rowCount = c.Int("rowCount")

			tilecat.setOptions(options)

			tilecat.save()

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
