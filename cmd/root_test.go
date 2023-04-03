package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cmwylie19/sbctl/cmd"
	"github.com/cmwylie19/sbctl/pkg/utils"
)

func TestGetRootCommand(t *testing.T) {
	cmd := cmd.GetRootCommand(utils.Logger{Debug: true})

	// Test that the command has the expected properties
	assert.Equal(t, "sbctl", cmd.Use)

}
