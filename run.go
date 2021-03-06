// Copyright 2013 bee authors
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
	"io/ioutil"
	"os"
	path "path/filepath"
	"runtime"
	"strings"
)

var cmdRun = &Command{
	UsageLine: "run [appname] [watchall] [-main=*.go] ",
	Short:     "run the app which can hot compile",
	Long: `
start the appname throw exec.Command

then start a inotify watch for current dir
										
when the file has changed mbee will auto go build and restart the app

	file changed
	     |
  check if it's go file
	     |
     yes     no
      |       |
 go build    do nothing
     |
 restart app   
`,
}

type ListOpts []string

func (opts *ListOpts) String() string {
	return fmt.Sprint(*opts)
}

func (opts *ListOpts) Set(value string) error {
	*opts = append(*opts, value)
	return nil
}

var mainFiles ListOpts

func init() {
	cmdRun.Run = runApp
	cmdRun.Flag.Var(&mainFiles, "main", "specify main go files")
}

var appname string

func runApp(cmd *Command, args []string) int {
	exit := make(chan bool)
	crupath, _ := os.Getwd()

	if len(args) == 0 || args[0] == "watchall" {
		appname = path.Base(crupath)
		ColorLog("[INFO] Uses '%s' as 'appname'\n", appname)
	} else {
		appname = args[0]
	}
	Debugf("current path:%s\n", crupath)

	var paths []string

	readAppDirectories(crupath, &paths)

	// Because monitor files has some issues, we watch current directory
	// and ignore non-go files.
	gps := GetGOPATHs()
	if len(gps) == 0 {
		ColorLog("[ERRO] Fail to start[ %s ]\n", "$GOPATH is not set or empty")
		os.Exit(2)
	}

	files := []string{}
	for _, arg := range mainFiles {
		if len(arg) > 0 {
			files = append(files, arg)
		}
	}

	NewWatcher(paths, files)

	Autobuild(files)

	for {
		select {
		case <-exit:
			runtime.Goexit()
		}
	}
	return 0
}

func readAppDirectories(directory string, paths *[]string) {
	fileInfos, err := ioutil.ReadDir(directory)
	if err != nil {
		return
	}

	useDiectory := false
	for _, fileInfo := range fileInfos {
		if strings.HasSuffix(fileInfo.Name(), "docs") {
			continue
		}
		if fileInfo.IsDir() == true && fileInfo.Name()[0] != '.' {
			readAppDirectories(directory+"/"+fileInfo.Name(), paths)
			continue
		}

		if useDiectory == true {
			continue
		}

		if path.Ext(fileInfo.Name()) == ".go" {
			*paths = append(*paths, directory)
			useDiectory = true
		}
	}

	return
}
