package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"

)

func main(){
	//colorPalette    string
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		w, h 				   = 1024, 1024
	)
	//image out
    img := image.NewRGBA(image.Rect(0, 0, w, h))
    for pixelY :=0; pixelY < h; pixelY++ {
    	//y
    	y := float64(pixelY)/h * (ymax-ymin) + ymin
    	    //x
    		for pixelX :=0; pixelX< w; pixelX++ {
    			x := float64(pixelX) / w * (xmax - xmin) + xmin
    			//complex num have to sum
    			z := complex(x, y)
    			//call mandelbrot function to image
    			img.Set(pixelX, pixelY, mandelbrot(z))
				//image.Set(ix, iy, uint32ToRgba(color))
			}
	}

    //create fractal image
    png.Encode(os.Stdout, img)


}
//mandelbrot
func mandelbrot(z complex128) color.Color {

	//num iterations
	const NITERS = 210
	//increase differents
	const CONSTRAST = 2
	//num 0 default
	var num complex128

	//iteration
	for i := uint8(0); i < NITERS; i++ {
		// min 2 iterations + z
		num = num*num + z
		//check num > 2 iterations
		if cmplx.Abs(num) > 2 {
			//gray
			//return color.Gray{255 - i * CONSTRAST }

			//color
			//colors := interpolateColors(&colorPalette, colorStep)
			//color1 := color.RGBA{0x00, 0x04, 0x0f, 0xff}
			//color2 := color.RGBA{0x00, 0x04, 0x0f, 0xff}
			//color := linearInterpolation(rgbaToUint(color1),  rgbaToUint(color2), uint32(iteration))
			//return color.RGBA{0x94 - i * CONSTRAST, 0x19 - i * CONSTRAST , 0x6a - i * CONSTRAST, 0xff - i }
			//return color.RGBA{0x24 - i * CONSTRAST, 0x99 - i * CONSTRAST , 0x6a - i * CONSTRAST, 0xff - i }
			return color.RGBA{14 -i * CONSTRAST , 35 -i * CONSTRAST, 15 -i * CONSTRAST, 255}


		}
	}
	//area black
	return color.Black

}


//COLOR UTILS

var Gameboy = []color.Color{
	color.RGBA{14, 55, 15, 255},
	color.RGBA{47, 97, 48, 255},
	color.RGBA{138, 171, 25, 255},
	color.RGBA{154, 187, 27, 255},
}
func linearInterpolation(c1, c2, mu uint32) uint32 {
	return c1*(1-mu) + c2*mu
}
func rgbaToUint(color color.RGBA) uint32 {
	r, g, b, a := color.RGBA()
	r /= 0xff
	g /= 0xff
	b /= 0xff
	a /= 0xff
	return uint32(r)<<24 | uint32(g)<<16 | uint32(b)<<8 | uint32(a)
}

func uint32ToRgba(col uint32) color.RGBA {
	r := col >> 24 & 0xff
	g := col >> 16 & 0xff
	b := col >> 8 & 0xff
	a := 0xff
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}