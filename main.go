package main

import "go-tool/zip"

func main() {
	zip.DeZip("./data/admin.jar", "./tmp/")
	//CreateDirFile("./data/META-INF/admin.txt")
}
