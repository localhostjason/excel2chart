package cmds

import (
	"flag"
	"gexcel/biz/excelx"
	"log"
)

func Run() {
	defaultFile := defaultExcelSrcFile()
	srcExcelFile := flag.String("p", defaultFile, "default src excel")

	flag.Parse()

	ex := excelx.NewExcelToLine(*srcExcelFile)
	err := ex.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
