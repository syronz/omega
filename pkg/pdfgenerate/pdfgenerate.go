package pdfgenerate

import (
	"fmt"
	"lab/pkg/random"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePdf(outputDirPath string, inputHtml string, patientName string) (path string, err error) {

	path = ""
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return path, err
	}

	f, err := os.Open(inputHtml)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return path, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return path, err
	}

	randomString := random.String(4)
	newFileName := fmt.Sprintf(`%v/%v-%v.pdf`, outputDirPath, patientName, randomString)
	err = pdfg.WriteFile(newFileName)
	if err != nil {
		return path, err
	}
	return newFileName, nil
}
