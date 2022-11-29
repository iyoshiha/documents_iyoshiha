package main

import (
	"bytes"
	"fmt"
)

func main(){
	var out *bytes.Buffer = &bytes.Buffer{}
	fmt.Fprintf(out, "majika %s\n", "ok")
	fmt.Fprintf(out, "majika %s\n", "ok")
	fmt.Fprintf(out, "majika %s\n", "ok")
	fmt.Printf(out.String())
}
