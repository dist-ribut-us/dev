package main

import (
	"bytes"
	"fmt"
	"github.com/dist-ribut-us/prog"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var dirs = []string{
	"beacon",
	"bufpool",
	"crypto",
	"dht",
	"dhtserver",
	"errors",
	"ipc",
	"ipcrouter",
	"log",
	"merkle",
	"message",
	"overlay",
	"packeter",
	"pool",
	"prog",
	"rnet",
	"serial",
	"vm",
}

var gopath string

func main() {
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		gopath = filepath.Join(prog.UserHomeDir(), "go")
	}

	app := cli.NewApp()
	app.Name = "ribut.build"
	app.Usage = "Show project status and if all tests are passing, run build"
	app.Action = ributBuild

	app.Commands = []cli.Command{
		{
			Name:   "status",
			Usage:  "Run tests, show lint and git status",
			Action: ributStatus,
		}, {
			Name:   "dev-tools",
			Usage:  "build dev tools",
			Action: ributDevTools,
		}, {
			Name:   "build",
			Usage:  "Only runs generate and build",
			Action: justBuild,
		},
		{
			Name:   "generate",
			Usage:  "Only runs generate",
			Action: justGenerate,
		},
	}

	app.Run(os.Args)
}

func ributBuild(c *cli.Context) error {
	clear()
	lineCount()
	generate()
	passing := runStatus()
	if passing {
		build()
		print("cat doNext.txt")
	}
	return nil
}

func justBuild(c *cli.Context) error {
	generate()
	build()
	return nil
}

func justGenerate(c *cli.Context) error {
	generate()
	return nil
}

func ributStatus(c *cli.Context) error {
	clear()
	lineCount()
	runStatus()
	return nil
}

func ributDevTools(c *cli.Context) error {
	print("go install github.com/dist-ribut-us/dev/generator/ribut.generator")
	print("go install github.com/dist-ribut-us/dev/ribut.build")
	return nil
}

func generate() {
	projectRoot()
	print("protoc --go_out=. pool/*.proto")
	print("protoc --go_out=. message/*.proto")
	print("ribut.generator < packeter/gen.json >packeter/gen.go")
	print("ribut.generator < overlay/gen.json >overlay/gen.go")
	print("ribut.generator < ipc/gen.json >ipc/gen.go")
	print("ribut.generator < ipcrouter/gen.json >ipcrouter/gen.go")
}

func build() {
	print("go install github.com/dist-ribut-us/pool/ribut.pool")
	print("go install github.com/dist-ribut-us/overlay/ribut.overlay")
	print("go install github.com/dist-ribut-us/dhtserver/ribut.dht")
	print("go install github.com/dist-ribut-us/beacon/ribut.beacon")
}

func projectRoot() {
	chdir(gopath, "src", "github.com", "dist-ribut-us")
}

func lineCount() {
	projectRoot()
	print("find . -name '*.go' | xargs wc -l | grep total")
}

func runStatus() bool {
	projectRoot()
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
	return passing
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func test() (string, string) {
	cmd := exec.Command("go", "test", "./...")

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	cmd.Run()
	cmd.Wait()

	if errStr := stderr.String(); errStr != "" {
		return "Error", errStr
	}

	outStr := stdout.String()
	lines := strings.Split(outStr, "\n")
	passed := true
	for _, line := range lines {
		if line != "" && strings.HasPrefix("FAIL", line) {
			passed = false
			break
		}
	}
	if passed == true {
		return "Pass", ""
	}
	return "Fail", outStr
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
