package main

import (
	"fmt"
	"os/exec"
)

func main() {

	// ==== Basic

	// cmd := exec.Command("echo", "Hello World")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	//	return
	// }

	// fmt.Println("Output:", string(output))

	// ==== Suing grep example

	// cmd := exec.Command("grep", "foo")

	// // Set input for the command
	// cmd.Stdin = strings.NewReader("foo is good\nfood is good\nbar is good\nbaz is good\n")

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Output:", string(output))

	// ==== Using start and wait methods

	// cmd := exec.Command("sleep", "2")

	// // Start the command
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("Error starting:", err)
	// 	return
	// }

	// fmt.Println("Running in the background")

	// // Waiting
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println("Error waiting:", err)
	// 	return
	// }

	// fmt.Println("Process is complete")

	// ===== Killing process

	// cmd := exec.Command("sleep", "60")

	// // Start the command
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println("Error starting:", err)
	// 	return
	// }

	// time.Sleep(2 * time.Second)
	// err = cmd.Process.Kill()
	// if err != nil {
	// 	fmt.Println("Error killing process:", err)
	// 	return
	// }

	// fmt.Println("Process was killed before finished")

	// ==== Using pipelines

	// cmd := exec.Command("grep", "foo")
	// pr, pw := io.Pipe()

	// cmd.Stdin = pr

	// go func() {
	// 	defer pw.Close()
	// 	pw.Write([]byte("food is good\nbar\nbaz"))
	// }()

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Output:", string(output))

	// ==== Using Combined Output

	cmd := exec.Command("ls", "-l")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", string(output))
}
