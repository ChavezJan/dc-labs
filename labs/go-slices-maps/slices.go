package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	var image [][]uint8

	for i := 0; i < dy; i++ {
		var row []uint8
		for j := 0; j < dx; j++ {
			pixel := uint8((i + j) / 2)
			row = append(row, pixel)
		}
		image = append(image, row)
	}
	return image

}

func main() {
	pic.Show(Pic)
}
