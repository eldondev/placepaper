package main
import (
	"rsc.io/qr"
	"strings"
	"os"
	"log"
	//"os/exec"
	"image"
	"golang.org/x/image/draw"
	"golang.org/x/image/bmp"
)

func main() {
	encodes := strings.Join(os.Args[1:], " ")
	log.Printf("Encoding %s", encodes)
	c, err := qr.Encode(encodes, qr.L)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("monocolor.bmp", os.O_RDWR|os.O_CREATE, 0666)
	scale := image.NewGray(image.Rect(0,0,176,176))
	draw.NearestNeighbor.Scale(scale, scale.Bounds(), c.Image(), image.Rect(0,0, c.Size, c.Size), draw.Src, nil)
	fit := image.NewGray(image.Rect(0,0,264,176))
	draw.Draw(fit, fit.Bounds(), image.White, fit.Bounds().Min, draw.Src)
	draw.Draw(fit, scale.Bounds(), scale, scale.Bounds().Min, draw.Src)
	bmp.Encode(file, fit)
	log.Printf("%+v", c.Image().Bounds())
	exec.Command("python", "main.py")
}

