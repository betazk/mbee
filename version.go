package main

import (
	"fmt"
	"log"
	"os/exec"
)

var cmdVersion = &Command{
	UsageLine: "version",
	Short:     "show the mbee version",
	Long: `
show the mbee version                  

bee version
    mbee: 0.1.0
`,
}

func init() {
	cmdVersion.Run = versionCmd
}

func versionCmd(cmd *Command, args []string) int {
	fmt.Println("bee   :" + version)
	goversion, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Go    :" + string(goversion))
	return 0
}
