package main

import (
	"./wkhtmltopdf"
	"fmt"
	"io/ioutil"
)

func main() {
	// global settings: http://www.cs.au.dk/~jakobt/libwkhtmltox_0.10.0_doc/pagesettings.html#pagePdfGlobal
	gs := wkhtmltopdf.NewGolbalSettings()
	gs.Set("outputFormat", "pdf")
	gs.Set("out", "")
	gs.Set("orientation", "Portrait")
	gs.Set("colorMode", "Color")
	gs.Set("size.paperSize", "A4")
	//gs.Set("load.cookieJar", "myjar.jar")
	// object settings: http://www.cs.au.dk/~jakobt/libwkhtmltox_0.10.0_doc/pagesettings.html#pagePdfObject
	os := wkhtmltopdf.NewObjectSettings()
	os.Set("page", "http://www.slashdot.org")
	os.Set("load.debugJavascript", "false")
	//os.Set("load.jsdelay", "1000") // wait max 1s
	os.Set("web.enableJavascript", "false")
	os.Set("web.enablePlugins", "false")
	os.Set("web.loadImages", "true")
	os.Set("web.background", "true")

	c := gs.NewConverter()
	c.Add(os)
	//c.AddHtml(os, "<html><body><h3>HELLO</h3><p>World</p></body></html>")

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
	c.Finished = func(c *wkhtmltopdf.Converter, s int) {
		fmt.Printf("Finished: %d\n", s)
	}
	c.Convert()

	payload, length := c.Payload()

	ioutil.WriteFile("out.pdf", payload, 0644)

	fmt.Printf("Length: %d\n", length)

	fmt.Printf("Got error code: %d\n", c.ErrorCode())
}
