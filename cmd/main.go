package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	information()

	reader := bufio.NewReader(os.Stdin)
	information()

	for {
		fmt.Print("Enter a command: ")
		text, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalf("Failed to read input: %s", err)
		}

		text = strings.TrimSpace(text)

		switch parts := strings.Split(text, " "); parts[0] {
		case "ls":
			err = listDirectory(parts)
		case "cd":
			err = changeDirectory(parts)
		case "mkdir":
			err = makeDirectory(parts)
		case "pwd":
			err = printWorkingDirectory()
		case "info":
			information()
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Error: unknown command")
		}
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}
}

func information(){
	fmt.Println("Welcome to the bash parser program!")
	fmt.Println("=============================================================")
	fmt.Println("Available commands:")
	fmt.Println("ls - List directory contents")
	fmt.Println("cd - Change directory")
	fmt.Println("mkdir - Create a new directory")
	fmt.Println("pwd - Print the current working directory")
	fmt.Println("log - Open a log file")
	fmt.Println("exit - Exit the program")
	fmt.Println("=============================================================")
	fmt.Println("Thank you for using our program, Your choice is awaited!")
	
}

func listDirectory(parts []string) error {
	if len(parts) != 1 {
		return errors.New("ls command does not take arguments")
	}

	files, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to list directory: %w", err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}

func changeDirectory(parts []string) error {
	if len(parts) != 2 {
		return errors.New("cd command takes one argument")
	}

	err := os.Chdir(parts[1])
	if err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	return nil
}

func makeDirectory(parts []string) error {
	if len(parts) != 2 {
		return errors.New("mkdir command takes one argument")
	}

	err := os.Mkdir(parts[1], 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	return nil
}

func printWorkingDirectory() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	fmt.Println(wd)
	return nil
}
