package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

func TestSetupAllCmd_NoFlags(t *testing.T) {

	var setupAllCmd = getSetupAllCmd()

	setupAllCmdFlags(setupAllCmd)

	_, err := execute(t, setupAllCmd)

	assert.Error(t, err, "Should thrown an error about the required flags")
}
