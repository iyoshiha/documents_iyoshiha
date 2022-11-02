package main

import(
	"os"
)

func main() {
	fd, e := os.Create("./tmp/ok.txt")
	if e!=nil{
		os.Exit(1)
	}
	var s []byte = []byte("abc\nefg\n")
	fd.Write(s)
	fd.Close()
}
