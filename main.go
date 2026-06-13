package main

import (
	"bytes"
	"fmt"
)

func main() {
	var output bytes.Buffer
	
	mandaXandaoSeFuderXVezes(10, &output)

	fmt.Println("Execução finalizada")
}

func mandaXandaoSeFuderXVezes(n int, out *bytes.Buffer) {
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, "Vai se fuder careca")
	}
}
