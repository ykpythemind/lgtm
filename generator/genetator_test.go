package generator

import (
	_ "image/jpeg"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	src, err := os.Open("testdata/manners.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create("testdata/dest.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer dst.Close()

	gen := NewGenerator(Option{FontSize: 200})

	err = gen.Generate(src, dst)
	if err != nil {
		t.Fatal(err)
	}
}
