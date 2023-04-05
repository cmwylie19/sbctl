package cli

import (
	"io"
	"log"
	"net/http"
	"os"

	"os/exec"
	"path/filepath"

	"github.com/cmwylie19/sbctl/pkg/utils"
)

const LINUX_PREFIX = "/usr/local/bin/"
const WINDOWS_PREFIX = "C:\\Program Files\\"
const CLI_URL = "https://repo.maven.apache.org/maven2/io/trino/trino-cli/410/trino-cli-410-executable.jar"

type CLI struct {
	// Path: pkg/cli/cli.go
	OS              string // Target Operating System to install the Trino CLI on. Options: mac, linux, or windows.
	DestinationFile string
	Logger          utils.Log // Logger Interface
}

func NewCli(os string, logger utils.Log) *CLI {
	destFilename := "trino"
	// identify the OS
	if os == "mac" || os == "linux" {
		destFilename = LINUX_PREFIX + destFilename
	} else if os == "windows" {
		destFilename = destFilename + ""
	} else {
		logger.Printf("sbctl: invalid os!\nsbctl: try 'sbctl cli -h' for more information.\n")
	}

	return &CLI{
		OS:              os,
		DestinationFile: destFilename,
		Logger:          logger,
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
		c.Logger.Errorf("Error while downloading the Trino CLI: %v", err)
	}

	// Get the absolute path of the destination file
	absPath, err := filepath.Abs(c.DestinationFile)
	if err != nil {
		panic(err)
	}

	// Check if Java is installed
	cmd := exec.Command("java", "-version")
	_, java_err := cmd.CombinedOutput()
	if java_err != nil {
		// Java is not installed
		c.Logger.Errorf("Error: Java is not installed.\nThe CLI requires a Java virtual machine available on the path. It can be used with Java version 8 and higher.\nSee https://starburstdata.github.io/latest/client/cli.html for more information.")
		// fmt.Println("Error: Java is not installed.")
		os.Exit(1)
	}

	err = os.Chmod(absPath, 0700)
	if err != nil {
		log.Fatal(err)
	}

	// Print the path of the renamed file
	c.Logger.Infof("Downloaded and renamed file to: %s\n\nSuccess!\n\nUsage:\n trino --version", absPath)

}
func (c *CLI) Uninstall() {
	err := os.Remove(c.DestinationFile)

	if err != nil {
		c.Logger.Errorf("Error while deleting the Trino CLI: %v", err)
		return
	}

	c.Logger.Infof("Binary file deleted successfully!")
}
