package cli

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const LINUX_PREFIX = "/usr/local/bin/"
const WINDOWS_PREFIX = "C:\\Program Files\\"
const CLI_URL = "https://repo.maven.apache.org/maven2/io/trino/trino-cli/410/trino-cli-410-executable.jar"

type CLI struct {
	// Path: pkg/cli/cli.go
	OS              string // Target Operating System to install the Trino CLI on. Options: mac, linux, or windows.
	DestinationFile string
}

func NewCli(os string) *CLI {
	destFilename := "trino"
	// identify the OS
	if os == "mac" || os == "linux" {
		destFilename = LINUX_PREFIX + destFilename
	} else if os == "windows" {
		destFilename = LINUX_PREFIX + destFilename
	} else {
		panic("Invalid OS")
	}

	return &CLI{
		OS:              os,
		DestinationFile: destFilename,
	}
}
func (c *CLI) Install() {

	// Create the destination file
	destFile, err := os.Create(c.DestinationFile)
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	// Download the file
	resp, err := http.Get(CLI_URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Copy the file contents to the destination file
	_, err = io.Copy(destFile, resp.Body)
	if err != nil {
		panic(err)
	}

	// Get the absolute path of the destination file
	absPath, err := filepath.Abs(c.DestinationFile)
	if err != nil {
		panic(err)
	}

	// Check if Java is installed
	cmd := exec.Command("java", "-version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Java is not installed
		fmt.Println("Error: Java is not installed.")
		os.Exit(1)
	}
	// Java is installed
	fmt.Println(string(output))

	err = os.Chmod(absPath, 0700)
	if err != nil {
		log.Fatal(err)
	}

	// Print the path of the renamed file
	println("Downloaded and renamed file to:", absPath)

}
func (c *CLI) Uninstall() {}
