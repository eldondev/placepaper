package main
import (
	"io/ioutil"
	"rsc.io/qr"
	"strings"
	"os"
	"log"
)

func main() {
	encodes := strings.Join(os.Args[1:], " ")
	log.Printf("Encoding %s", encodes)
	c, err := qr.Encode(encodes, qr.L)
	if err != nil {
		log.Fatal(err)
	}
	pngdat := c.PNG()
	ioutil.WriteFile("x.png", pngdat, 0666)
}

