package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// functions, of cours, need to be declared before the endpoint
func tellMyHistory(name string, age int) {
	fmt.Printf("Hello, my name is %v and I am %v", name, age)
}

func main() {
	name := "vitor"
	var age int = 30
	var float_age float64 = 1995.00 / float64(age)
	hasColor := true
	fmt.Printf("Hello %v, this is %T", name, name)
	fmt.Printf("\nyou are %v years old", age)
	fmt.Printf("===============\n")
	fmt.Printf(
		"My float age is %.2f and its type is %T, and hasColor: %v is %T",
		float_age,
		float_age,
		hasColor,
		hasColor)

	// reassing variables
	hasColor = false
	fmt.Printf("\nyou are %v years old", hasColor)

	// constraints
	const minuteLife int = 60
	now_minute := time.Now().Minute()

	// control flow
	if minuteLife <= now_minute && minuteLife > 45 {
		fmt.Printf("\nThe current minnute is: %v", now_minute)
		now_minute += 60
		fmt.Printf("\nNow its value is %v", now_minute)
	} else if now_minute >= 30 {
		now_minute += 30
	} else {
		now_minute -= 60
	}

	fmt.Printf("\nNow its value is %v", now_minute)

	// if the variable is only used on if statment (syntx)
	if now_minute_asyntx := time.Now().Minute(); minuteLife >= now_minute_asyntx {
		fmt.Printf("New Syntx")
		fmt.Printf("now_minute_asyntx: %v is only accessible in this scope", now_minute_asyntx)
	}

	// executing the function
	tellMyHistory("Vitor", 30)

	// executing cmd
	// wait for exec.Command to complete
	args := strings.Split("/c dir", " ")
	cmd := exec.Command("cmd.exe", args...)

	cmd.Dir = filepath.Join("C://", "Windows")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run: %s failed: %s\n", err, err)
	}
}
