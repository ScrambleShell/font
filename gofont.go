package font

import (
	"github.com/as/font/gomedium"
	"github.com/as/font/gomono"
	"github.com/as/font/goregular"
	"github.com/golang/freetype/truetype"
)

// New returns a GoMedium font that is cached and replaces
// non-printable characters with their hex equivalent encodings
func NewFace(size int) Face{
	return NewReplacer(NewGoMedium(size), NewHex(size), nil)
}

func NewGoMedium(size int) Face {
	return Open(truetype.NewFace(gomedium.Font, &truetype.Options{
		SubPixelsX: 64,
		SubPixelsY: 64,
		Size:       float64(size),
	}))
}
var g = gomono.Font
func NewGoMono(size int) Face {
	return NewReplacer(NewGoMedium(size), NewHex(size), nil)
	
/*
	return NewCache((truetype.NewFace(gomono.Font, &truetype.Options{
		SubPixelsX: 64,
		SubPixelsY: 64,
		Size:       float64(size),
	})))
*/
}

func NewGoRegular(size int) Face {
	return Open(truetype.NewFace(goregular.Font, &truetype.Options{
		SubPixelsX: 64,
		SubPixelsY: 64,
		Size:       float64(size),
	}))
}
