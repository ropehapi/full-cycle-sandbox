package main

import (
	"bytes"
	"testing"
)

func TestMandaXandaoSeFuderXVezes(t *testing.T) {
	var output bytes.Buffer

	mandaXandaoSeFuderXVezes(5, &output)

	expectedOutput :=
		"Vai se fuder Alexandre de Moraes\n" +
			"Vai se fuder Alexandre de Moraes\n" +
			"Vai se fuder Alexandre de Moraes\n" +
			"Vai se fuder Alexandre de Moraes\n" +
			"Vai se fuder Alexandre de Moraes\n"

	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s",
			expectedOutput,
			output.String())
	}
}