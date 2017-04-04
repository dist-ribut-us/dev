package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

var dirs = []string{
	"beacon",
	"bufpool",
	"crypto",
	"dht",
	"errors",
	"ipc",
	"log",
	"merkle",
	"message",
	"overlay",
	"packeter",
	"pool",
	"prog",
	"rnet",
	"serial",
}

func main() {
	clear()
	gopath := os.Getenv("GOPATH")
	chdir(gopath, "src", "github.com", "dist-ribut-us")

	print("find . -name '*.go' | xargs wc -l | grep total")

	passing := true
	for _, dir := range dirs {
		chdir(dir)
		summary, output := test()
		passing = passing && summary == "Pass"
		fmt.Print("== ", dir, " : ", summary, " ==\n", output)
		print("git status --porcelain")
		print("golint `find $d -maxdepth 1 -mindepth 1 -name '*.go' -a ! -name '*.pb.go'`")
		chdir("..")
	}

	if passing {
		print("./build.sh")
		print("cat doNext.txt")
	}
}

func test() (string, string) {
	testOutput := run("go test")
	if testOutput[:4] == "PASS" {
		return "Pass", ""
	}
	return "Fail", testOutput
}

func run(cmd string) string {
	out, _ := exec.Command("sh", "-c", cmd).CombinedOutput()
	return string(out)
}

func print(cmd string) {
	fmt.Print(run(cmd))
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func chdir(strs ...string) {
	os.Chdir(path.Join(strs...))
}
