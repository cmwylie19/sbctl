package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cmwylie19/sbctl/pkg/utils"
)

func TestCliInstall(t *testing.T) {
	// Create a temporary directory for the test
	tempDir := t.TempDir()

	// Create a new CLI instance and set the destination file to be inside the temporary directory
	cli := NewCli("mac", utils.Logger{Debug: true})
	cli.DestinationFile = filepath.Join(tempDir, "trino")

	// Call the Install method
	cli.Install()

	// Check that the destination file exists and has the correct permissions
	info, err := os.Stat(cli.DestinationFile)
	if err != nil {
		t.Errorf("Expected destination file to exist, but got error: %v", err)
	}
	if info.Mode().Perm() != 0700 {
		t.Errorf("Expected destination file to have mode 0700, but got mode %v", info.Mode().Perm())
	}
}
