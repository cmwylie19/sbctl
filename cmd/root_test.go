package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cmwylie19/sbctl/cmd"
)

func TestGetRootCommand(t *testing.T) {
	cmd := cmd.GetRootCommand()

	// Test that the command has the expected properties
	assert.Equal(t, "sbctl", cmd.Use)

}
