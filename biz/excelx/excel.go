package excelx

import (
	"fmt"
	"gexcel/biz/util"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"strconv"
)

type ExcelToLine struct {
	SrcFile  string
	DistFile string
}

func NewExcelToLine(srcFile string) *ExcelToLine {
	exeDir, _ := util.GetExeDir()
	distFile := filepath.Join(exeDir, "excel", "dist.xlsx")

	return &ExcelToLine{SrcFile: srcFile, DistFile: distFile}
}

func (e *ExcelToLine) readExcel() ([]SrcExcel, error) {
	var data = make([]SrcExcel, 0)

	f, err := excelize.OpenFile(e.SrcFile)
	if err != nil {
		return data, err
	}
	defer func() {
		_ = f.Close()
	}()
	rows, err := f.GetRows("Sheet1")

	for index, row := range rows {
		if index == 0 || len(row) != 2 {
			continue
		}
		data = append(data, SrcExcel{
			Date:   row[0],
			Number: row[1],
		})
		//fmt.Println(row)
	}
	return data, nil

}

func (e *ExcelToLine) createLineExcel(excelData []SrcExcel) error {
	f := excelize.NewFile()

	// 工作表
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"A1": "日期",
		"B1": "体温",
	}
	for i, d := range excelData {
		number, errx := strconv.ParseFloat(d.Number, 64)
		if errx != nil {
			return errx
		}

		i += 2
		Ai := fmt.Sprintf("A%d", i)
		Bi := fmt.Sprintf("B%d", i)
		data[Ai] = d.Date
		data[Bi] = number
	}

	for k, v := range data {
		_ = f.SetCellValue("Sheet1", k, v)
	}

	countNumber := len(excelData) + 1
	print(00, countNumber)

	chartOption := &excelize.Chart{
		Type: excelize.Line,
		Format: excelize.GraphicOptions{
			OffsetX: 62,
		},
		Title: excelize.ChartTitle{
			Name: "2023年体温记录",
		},
		XAxis: excelize.ChartAxis{
			Font: excelize.Font{
				Color: "#333",
			},
			MajorGridLines: true,
		},
		YAxis: excelize.ChartAxis{
			Font: excelize.Font{
				Color: "#333",
			},
			MajorGridLines: true,
			MajorUnit:      0.1,
		},
		Series: []excelize.ChartSeries{
			{
				Name:       "Sheet1!$B$1",
				Categories: fmt.Sprintf("Sheet1!$A$2:$A$%d", countNumber),
				Values:     fmt.Sprintf("Sheet1!$B$2:$B$%d", countNumber),
			},
		},
	}

	err = f.AddChart("Sheet1", "C1", chartOption)
	if err != nil {
		return err
	}
	// 将工作表设置为默认选中
	f.SetActiveSheet(index)

	// 保存XLSX文件
	return f.SaveAs(e.DistFile)
}

func (e *ExcelToLine) Run() error {
	data, err := e.readExcel()
	if err != nil {
		return err
	}
	//fmt.Println(data)
	return e.createLineExcel(data)
}
