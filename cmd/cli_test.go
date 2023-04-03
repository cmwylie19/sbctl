package cmd

import (
	"testing"

	"github.com/cmwylie19/sbctl/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetCliCommand(t *testing.T) {
	cmd := getCliCommand(utils.Logger{Debug: true})
	assert.NotNil(t, cmd)

	flags := cmd.Flags()
	assert.NotNil(t, flags)
}
