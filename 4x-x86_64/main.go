package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Running 4x")
	Execute(exec.Command("rtl_sdr", "-f", "137100000", "-g", "40", "-s", "2500000", "-n", "250000000", "NOAA-Image.dat"))

}
