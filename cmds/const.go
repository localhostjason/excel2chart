package cmds

import (
	"gexcel/biz/util"
	"path/filepath"
)

func defaultExcelSrcFile() string {
	exeDir, _ := util.GetExeDir()
	return filepath.Join(exeDir, "excel", "src.xlsx")
}
