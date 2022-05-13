package main

import (
	"errors"
	"image"
	"image/png"
	"os"

	gim "github.com/ozankasikci/go-image-merge"
)

type tileCatOptions struct {
	dir         string
	out         string
	rowCount    int
	columnCount int
}

type TileCat struct {
	options   *tileCatOptions
	allImages []*image.Image
}

func NewTileCat(options *tileCatOptions) *TileCat {
	return &TileCat{
		options:   options,
		allImages: []*image.Image{},
	}
}

func (t *TileCat) setOptions(options *tileCatOptions) error {
	t.options = options
	return nil
}

func (t *TileCat) loadImages(files []string) *image.Image {

	for _, filePath := range files {
		f, err := os.Open(filePath)
		if err == nil {
			img, _, _ := image.Decode(f)
			t.allImages = append(t.allImages, &img)
		}
	}

	return nil
}

func (t *TileCat) processImages(files []string) (*image.RGBA, error) {
	if t.options == nil {
		return nil, errors.New("no options!")
	}

	grids := []*gim.Grid{}

	for _, filePath := range files {
		grids = append(grids, &gim.Grid{
			ImageFilePath: filePath,
		})

	}

	rgba, err := gim.New(grids, t.options.columnCount, t.options.rowCount).Merge()

	return rgba, err
}

func (t *TileCat) save() error {
	if t.options == nil {
		return errors.New("no options!")
	}

	files, err := listImages(t.options.dir)

	if err != nil {
		return err
	}

	rgba, err := t.processImages(files)

	if err != nil {
		return err
	}

	dst, err := os.Create(t.options.out)
	//err = jpeg.Encode(dst, rgba, &jpeg.Options{Quality: 80})
	err = png.Encode(dst, rgba)
	dst.Close()

	return err
}
