// Copyright 2013 mbee authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"os"
	path "path/filepath"
	"strings"
)

var cmdNew = &Command{
	UsageLine: "new [appname]",
	Short:     "create an application base on martini framework",
	Long: `
create an application base on martini framework,

which in the current path with folder named [appname].

The [appname] folder has following structure:

    |- main.go
    |- conf
        |-  app.conf
    |- models
    |- tests
         |- default_test.go
	|- static
         |- js
         |- css
         |- img
    |- templates

`,
}

func init() {
	cmdNew.Run = createApp
}

func createApp(cmd *Command, args []string) int {
	curpath, _ := os.Getwd()
	if len(args) != 1 {
		ColorLog("[ERRO] Argument [appname] is missing\n")
		os.Exit(2)
	}

	gopath := os.Getenv("GOPATH")
	Debugf("gopath:%s", gopath)
	if gopath == "" {
		ColorLog("[ERRO] $GOPATH not found\n")
		ColorLog("[HINT] Set $GOPATH in your environment vairables\n")
		os.Exit(2)
	}
	haspath := false
	// appsrcpath := ""

	wgopath := path.SplitList(gopath)
	for _, wg := range wgopath {
		wg, _ = path.EvalSymlinks(path.Join(wg, "src"))

		if strings.HasPrefix(strings.ToLower(curpath), strings.ToLower(wg)) {
			haspath = true
			// appsrcpath = wg
			break
		}
	}

	if !haspath {
		ColorLog("[ERRO] Unable to create an application outside of $GOPATH(%s)\n", gopath)
		ColorLog("[HINT] Change your work directory by `cd ($GOPATH%ssrc)`\n", string(path.Separator))
		os.Exit(2)
	}

	apppath := path.Join(curpath, args[0])

	if _, err := os.Stat(apppath); os.IsNotExist(err) == false {
		fmt.Printf("[ERRO] Path(%s) has alreay existed\n", apppath)
		os.Exit(2)
	}

	fmt.Println("[INFO] Creating application...")

	os.MkdirAll(apppath, 0755)
	fmt.Println(apppath + string(path.Separator))
	os.Mkdir(path.Join(apppath, "conf"), 0755)
	fmt.Println(path.Join(apppath, "conf") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "models"), 0755)
	fmt.Println(path.Join(apppath, "models") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static"), 0755)
	fmt.Println(path.Join(apppath, "static") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "js"), 0755)
	fmt.Println(path.Join(apppath, "static", "js") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "css"), 0755)
	fmt.Println(path.Join(apppath, "static", "css") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "img"), 0755)
	fmt.Println(path.Join(apppath, "static", "img") + string(path.Separator))
	fmt.Println(path.Join(apppath, "templates") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "templates"), 0755)
	fmt.Println(path.Join(apppath, "conf", "app.conf"))
	writetofile(path.Join(apppath, "conf", "app.conf"), strings.Replace(appconf, "{{.Appname}}", args[0], -1))

	fmt.Println(path.Join(apppath, "main.go"))
	writetofile(path.Join(apppath, "main.go"), strings.Replace(maingo, "{{.tmpl}}", "`"+indexTmpl+"`", -1))

	ColorLog("[SUCC] New application successfully created!\n")
	return 0
}

var appconf = `appname = {{.Appname}}
httpport = 3000
runmode = dev
`

var indexTmpl = `<html>
			<head>
				<title>Martini</title>
				<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
			</head>
			<style type="text/css">
				body{
					background-color: #009999;
				}
				h1 {
					text-align: center;
					position: absolute;
			  		left: 450px;
			  		top: 200px;
				}
			</style>
			<body>
				<h1>Welcome to Martini's World !</h1>
			</body>
		</html>`

var maingo = `package main

import (
	"github.com/go-martini/martini"
)

var indexTmpl = {{.tmpl}}

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return indexTmpl
	})
	m.Run()
	}

`

func writetofile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}
