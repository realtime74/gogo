package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GogoFile struct {
	file *os.File
}

func (g *GogoFile) Init() string {
	hd, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fp := hd + "/.gogo"
	g.file, err = os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	return "Hello, Gogo!"
}

func (g *GogoFile) Match(pattern string) string {
	g.file.Seek(0, 0)

	scanner := bufio.NewScanner(g.file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			return line
		}
	}
	panic(fmt.Sprintf("No match found for %s", pattern))
}

func (g *GogoFile) Close() {
	if g.file != nil {
		g.file.Close()
		g.file = nil
	}
}

func main() {
	file := GogoFile{}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	file.Init()
	defer file.Close()

	// check first command line argument
	if len(os.Args) < 2 {
		fmt.Println("usage: gogo <pattern|+>")
		// bail out
		os.Exit(1)
	}
	pattern := os.Args[1]
	fmt.Println(file.Match(pattern))
}
