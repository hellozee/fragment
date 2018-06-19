/*Package flipper  The sole purpose of this package is to flip the image while
it being saved by the renderer.
  Shamelessly copied from https://github.com/disintegration/imaging to avoid having
  any dependecies other than Go's standard library, only for flipping the generated
  image.
*/
package flipper

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

type scanner struct {
	image   image.Image
	w, h    int
	palette []color.NRGBA
}

func newScanner(img image.Image) *scanner {
	s := &scanner{
		image: img,
		w:     img.Bounds().Dx(),
		h:     img.Bounds().Dy(),
	}
	if img, ok := img.(*image.Paletted); ok {
		s.palette = make([]color.NRGBA, len(img.Palette))
		for i := 0; i < len(img.Palette); i++ {
			s.palette[i] = color.NRGBAModel.Convert(img.Palette[i]).(color.NRGBA)
		}
	}
	return s
}

func (s *scanner) scan(x1, y1, x2, y2 int, dst []uint8) {
	j := 0
	b := s.image.Bounds()
	x1 += b.Min.X
	x2 += b.Min.X
	y1 += b.Min.Y
	y2 += b.Min.Y
	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			r16, g16, b16, a16 := s.image.At(x, y).RGBA()
			switch a16 {
			case 0xffff:
				dst[j+0] = uint8(r16 >> 8)
				dst[j+1] = uint8(g16 >> 8)
				dst[j+2] = uint8(b16 >> 8)
				dst[j+3] = 0xff
			case 0:
				dst[j+0] = 0
				dst[j+1] = 0
				dst[j+2] = 0
				dst[j+3] = 0
			default:
				dst[j+0] = uint8(((r16 * 0xffff) / a16) >> 8)
				dst[j+1] = uint8(((g16 * 0xffff) / a16) >> 8)
				dst[j+2] = uint8(((b16 * 0xffff) / a16) >> 8)
				dst[j+3] = uint8(a16 >> 8)
			}
			j += 4
		}
	}
}

func parallel(start, stop int, fn func(<-chan int)) {
	count := stop - start
	if count < 1 {
		return
	}

	procs := runtime.GOMAXPROCS(0)
	if procs > count {
		procs = count
	}

	c := make(chan int, count)
	for i := start; i < stop; i++ {
		c <- i
	}
	close(c)

	var wg sync.WaitGroup
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(c)
		}()
	}
	wg.Wait()
}

//FlipV  Function to flip an image horizontally
func FlipV(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dstW := src.w
	dstH := src.h
	rowSize := dstW * 4
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel(0, dstH, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcY := dstH - dstY - 1
			src.scan(0, srcY, src.w, srcY+1, dst.Pix[i:i+rowSize])
		}
	})
	return dst
}
