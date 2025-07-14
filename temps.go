package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("test")

	cmd := exec.Command("bash", "-c", "vcgencmd measure_temp")
	// fmt.Println(cmd.StdoutPipe())
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println("Output:", string(output))

}
