package main

import (
	"os"
	"github.com/otiai10/gosseract"
	"github.com/sosedoff/gofpdf"
)

func main() {
	// 打开PDF文件
	pdf := gofpdf.New("P", "mm", "A4", "input.pdf")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	for i := 1; i <= pdf.PageCount(); i++ {
		// 读取PDF的每一页内容并添加到Word文档中
		text, _ := pdf.GetPageText(i)
		for _, line := range text {
			gosseract.ConvertTextToImage(line, "output.docx")
		}
	}
}