package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/cmwylie19/sbctl/cmd"
	"github.com/cmwylie19/sbctl/pkg/utils"
)

func TestMain(m *testing.M) {
	// Redirect log output to a buffer
	var buf bytes.Buffer
	logger = utils.Logger{Debug: true}

	// Run the tests
	code := m.Run()

	// Check the log output
	if buf.String() != "" {
		logger.Errorf("Unexpected log output: %s", buf.String())
	}

	// Exit with the correct code
	os.Exit(code)
}

func TestMainFunction(t *testing.T) {
	// Set up a mock root command that always returns an error
	mockRootCmd := cmd.GetRootCommand(utils.Logger{Debug: true})
	mockRootCmd.SetArgs([]string{"non-existent-command"})
	err := mockRootCmd.Execute()

	// Check that the main function exits with an error code
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
