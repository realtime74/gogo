package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type GogoFile struct {
	Home string
	Path string

	file *os.File
}

func (g *GogoFile) Init() {
	hd, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	g.Home = hd
	g.Path = hd + "/.gogo"
	g.file, err = os.OpenFile(g.Path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	return
}

func (g *GogoFile) Append(dir string) {
	g.file.Seek(0, 2) // seek to end of file
	info, err := g.file.Stat()
	if err != nil {
		panic(err)
	}
	if info.Size() > 0 {
		g.file.WriteString("\n")
	}
	g.file.WriteString(dir)
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

func _edit(file GogoFile) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}
	fmt.Println("Opening", file.Path, "with", editor)
	cmd := exec.Command(editor, file.Path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
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
		fmt.Println(file.Home)
		return
	}
	pattern := os.Args[1]
	switch pattern {
	case "-e":
		_edit(file)
		return
	case "~":
		fmt.Println(file.Home)
		return
	case "-":
		fmt.Println("-")
		return
	case "+":
		cwd, _ := os.Getwd()
		file.Append(cwd)
		fmt.Println(cwd)
		return
	}
	dir := file.Match(pattern)
	fmt.Println(dir)
}
