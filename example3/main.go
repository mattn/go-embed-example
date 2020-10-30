package main

import (
	"embed"
	"io"
	"log"
	"os"
	"path"
)

//go:embed static
var local embed.FS

func main() {
	fis, err := local.ReadDir("static")
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fis {
		in, err := local.Open(path.Join("static", fi.Name()))
		if err != nil {
			log.Fatal(err)
		}
		out, err := os.Create("embed-" + path.Base(fi.Name()))
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(out, in)
		out.Close()
		in.Close()
		log.Println("exported", "embed-"+path.Base(fi.Name()))
	}
}
