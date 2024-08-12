package cmd

import (
	"bytes"
	"io"
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := CreateVersionCmd("v1.0")
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "version: v1.0" {
		t.Fatalf("expected \"%s\" got \"%s\"", "version: v1.0", string(out))
	}
}
