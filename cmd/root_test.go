package cmd

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	c, err = root.ExecuteC()
	return c, buf.String(), err
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func TestHelloCommandExists(t *testing.T) {

	c, _, err := executeCommandC(rootCmd, "hello")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if c.Name() != "hello" {
		t.Errorf(`invalid command returned from ExecuteC: expected "child"', got %q`, c.Name())
	}

}
