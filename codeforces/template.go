package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func flush()                            { out.Flush() }
func scan(a ...interface{})             { fmt.Fscan(in, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(in, f, a...) }
func print(a ...interface{})            { fmt.Fprintln(out, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(out, f, a...) }
func assert(f bool) {
	if !f {
		panic("gg")
	}
}

func slove(cases int) {
}

func main() {
	defer flush()

	cases := 0
	scan(&cases)
	for i := 0; i < cases; i++ {
		slove(i)
	}
}
