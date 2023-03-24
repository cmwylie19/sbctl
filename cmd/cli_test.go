package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCliCommand(t *testing.T) {
	cmd := getCliCommand()
	assert.NotNil(t, cmd)

	flags := cmd.Flags()
	assert.NotNil(t, flags)
}
