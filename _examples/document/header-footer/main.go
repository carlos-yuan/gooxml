// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"log"

	"github.com/qifengzhang007/gooxml/common"
	"github.com/qifengzhang007/gooxml/document"
	"github.com/qifengzhang007/gooxml/measurement"
	"github.com/qifengzhang007/gooxml/schema/soo/wml"
)

func main() {
	var pathPre = "F:/Project/2023/word-format/使用的库代码参考/gooxml/_examples/document/header-footer/"
	doc := document.New()

	img, err := common.ImageFromFile(pathPre + "gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	hdr := doc.AddHeader()
	// We need to add a reference of the image to the header instead of to the
	// document
	iref, err := hdr.AddImage(img)
	if err != nil {
		log.Fatalf("unable to to add image to document: %s", err)
	}

	para := hdr.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	run := para.AddRun()
	run.AddTab()
	run.AddText("文档标题")

	imgInl, _ := para.AddRun().AddDrawingInline(iref)
	imgInl.SetSize(1*measurement.Inch, 1*measurement.Inch, 1)

	// Headers and footers are not immediately associated with a document as a
	// document can have multiple headers and footers for different sections.
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)

	ftr := doc.AddFooter()
	para = ftr.AddParagraph()
	para.Properties().AddTabStop(6*measurement.Inch, wml.ST_TabJcRight, wml.ST_TabTlcNone)
	run = para.AddRun()
	run.AddText("Some subtitle goes here")
	run.AddTab()
	run.AddText("Pg ")
	run.AddField(document.FieldCurrentPage)
	run.AddText(" of ")
	run.AddField(document.FieldNumberOfPages)
	doc.BodySection().SetFooter(ftr, wml.ST_HdrFtrDefault)

	lorem := `新的测试段落文本`

	for i := 0; i < 5; i++ {
		para = doc.AddParagraph()
		run = para.AddRun()
		run.AddText(lorem)
	}

	doc.SaveToFile(pathPre + "header-footer-2.docx")
}
