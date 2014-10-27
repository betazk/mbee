mbee
====
mbee is a tool for helping develop with martini app framework.

## Requirements

- Go version >= 1.1.


Installation
===

Begin by installing `mbee` using `go get` command

	go get github.com/betazk/mbee

Then you can add bee binary to PATH environment variable in your ~/.bashrc or ~/.bash_profile file:

	export PATH=$PATH:<your_main_gopath>/bin/bee

>If you already have mbee installed, updating mbee is simple:

	go get -u github.com/betazk/mbee

Basic Commands
===

mbee now only provides three commands which can be helpful at  development. The top level commands include:

	new         create an application base on martini framework
	run         run the app which can hot compile
	version     show the mbee & go version

mbee version
===

The first command is the easiest: displaying which version of mbee and go is installed on your machine:

	$ bee version
	mbee   :0.1.0
	Go    :go version go1.3 linux/amd64

mbee new
===

Creating a new martini web application is no big deal, too.

	$ mbee new myapp
	[INFO] Creating application...
	/home/zhengkang/goPath/src/myapp/
	/home/zhengkang/goPath/src/myapp/conf/
	/home/zhengkang/goPath/src/myapp/models/
	/home/zhengkang/goPath/src/myapp/static/
	/home/zhengkang/goPath/src/myapp/static/js/
	/home/zhengkang/goPath/src/myapp/static/css/
	/home/zhengkang/goPath/src/myapp/static/img/
	/home/zhengkang/goPath/src/myapp/views/
	/home/zhengkang/goPath/src/myapp/conf/app.conf
	/home/zhengkang/goPath/src/myapp/main.go
	2014/10/27 11:12:00 [SUCC] New application successfully created!

mbee run
===

To run the application we just created, navigate to the application folder and execute mbee run.

	$ cd myapp
	$ mbee run

Help
===

If you happend to forget the usage of a command, you can always find the usage information by mbee help <command>.

For instance, to get more information about the run command:

	$ mbee help run
	usage: mbee run [appname] [watchall] [-main=*.go]

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

The idea and most of mbee are come from [bee](http://github.com/beego/bee)