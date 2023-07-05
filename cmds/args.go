package cmds

import (
	"flag"
	"fmt"
	"gexcel/biz/excelx"
	"log"
	"time"
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
	now := time.Now()
	fmt.Println("当前时间：", now.Format("2006-01-02 15:04:05"), "，生成excel成功")
}
