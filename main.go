package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func usage(b bool) {
	if b {
		return
	}
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "\tgit-fixup [--root] [HEAD~N]\n")
	fmt.Fprintf(os.Stderr, "\tgit-fixup edit filename\n")
	os.Exit(2)
}

func edit(file string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if len(line) == 0 {
			break
		}
		space := strings.Index(line, " ")
		if i == 0 {
			lines[i] = "reword" + line[space:]
		} else {
			lines[i] = "fixup" + line[space:]
		}
	}
	out := []byte(strings.Join(lines, "\n"))
	if err := ioutil.WriteFile(file, out, 0644); err != nil {
		log.Fatal(err)
	}
}

func fixup(args []string) {
	editpath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	args = append([]string{"rebase", "-i"}, args...)
	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Env = append(os.Environ(), fmt.Sprintf("GIT_SEQUENCE_EDITOR=%s edit", editpath))
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetPrefix("git-fixup: ")
	log.SetFlags(0)
	usage(len(os.Args) >= 2)

	switch os.Args[1] {
	case "edit":
		usage(len(os.Args) == 3)
		edit(os.Args[2])
	default:
		fixup(os.Args[1:])
	}
}
