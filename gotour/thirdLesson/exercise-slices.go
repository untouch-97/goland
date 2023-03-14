package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Create a 2D slice with the given dimensions
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
	}

	// Assign a grayscale value to each pixel
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			// Calculate the pixel value using your chosen function
			// and assign it to the corresponding element of the slice
			pic[y][x] = uint8((x + y) / 2) //(x+y)/2, x*y, and x^y
		}
	}

	// Return the completed 2D slice
	return pic
}

func main() {
	pic.Show(Pic)
}
