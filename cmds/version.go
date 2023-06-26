package cmds

import "fmt"

var Ver string
var Date string
var Git string

func getVersion() {
	// -ldflags "-X console/cmds.Ver=0.0.1 -X console/cmds.Date=2023-05-16 -X console/cmds.Git=$(git describe --tags --abbrev=0)"
	// -ldflags "-X console/cmds.Ver=0.0.1 -X console/cmds.Date=2023-05-16 -X console/cmds.Git=$(git rev-parse --short HEAD)"
	fmt.Println("Version  : ", Ver)
	fmt.Println("Date     : ", Date)
	fmt.Println("Git      : ", Git)
}
