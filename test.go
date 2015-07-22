package main

import (
	"fmt"
	"github.com/hailocab/wkhtmltopdf-go/wkhtmltopdf"
)

func main() {
	gs := wkhtmltopdf.NewGolbalSettings()
	gs.Set("out", "test.pdf")
	gs.Set("load.cookieJar", "myjar.jar")
	os := wkhtmltopdf.NewObjectSettings()
	//os.Set("page", "file:///home/jimmy/libwk/invoice.html")
	os.Set("page", "http://www.google.se")

	c := gs.NewConverter()
	c.Add(os)

	c.ProgressChanged = func(c *wkhtmltopdf.Converter, b int) {
		fmt.Printf("Progress: %d\n", b)
	}
	c.Error = func(c *wkhtmltopdf.Converter, msg string) {
		fmt.Printf("error: %s\n", msg)
	}
	c.Warning = func(c *wkhtmltopdf.Converter, msg string) {
		fmt.Printf("error: %s\n", msg)
	}
	c.Phase = func(c *wkhtmltopdf.Converter) {
		fmt.Printf("Phase\n")
	}
	c.Convert()

	fmt.Printf("Got error code: %d\n", c.ErrorCode())
}
