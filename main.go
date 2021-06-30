
package main

import (
	"github.com/signintech/gopdf"
	"log"
)

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	err := pdf.AddTTFFont("test1", "./supermarket.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("test1", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "asdfasdf  น้ำ  พี่ ใส่  พื้นที่")
	pdf.WritePdf("hello.pdf")

}

