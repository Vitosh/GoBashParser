package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	fun "./fun"
	lib "./lib"
)

func main() {

	fmt.Print(fun.ShowCatsWelcome())

	lib.CreatePath()
	logFile, err := os.Create(lib.LogPath)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Printf("Log file is created...")

	fmt.Println(information())

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v\n", err)
			continue
		}

		cmd = strings.ToLower(strings.TrimSpace(cmd))
		runCommand(cmd)
	}
}

func runCommand(cmd string) {
	// Run bash cmd

	log.Printf("Running command: %s\n", cmd)
	args := strings.Fields(cmd)

	switch args[0] {
	case "ls":
		listDirectory(".")
	case "cd":
		if len(args) < 2 {
			log.Printf("Error running cd command: missing argument\n")
			return
		}
		changeDirectory(args[1])
	case "mkdir":
		if len(args) < 2 {
			log.Printf("Error running mkdir command: missing argument\n")
			return
		}
		makeDirectory(args[1])
	case "pwd":
		printWorkingDirectory()
	case "rm":
		if len(args) < 2 {
			log.Printf("Error running rm command: missing argument\n")
			return
		}
		removeFile(args[1])
	case "touch":
		if len(args) < 2 {
			log.Printf("Error running touch command: missing argument\n")
			return
		}
		createFile(args[1])
	case "curl":
		if len(args) < 2 {
			log.Printf("Error running curl command: missing argument\n")
			return
		}
		fetchURL(args[1], os.Stdout)
	case "log":
		displayLog()
	case "catsneeded", "hidden":
		fmt.Println(fun.ShowCatsWelcome())
		fmt.Println(fun.ShowCatsGoodbye())
	case "info":
		fmt.Println(information())
	case "exit":
		msg := "Exit successful :)"

		fmt.Println(msg)
		fmt.Println(fun.ShowCatsGoodbye())
		fmt.Printf("Press enter to exit!")

		log.Print(msg)
		log.Printf("Thank you for using our services!")

		//Just wait for key
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')
		os.Exit(0)

	default:
		msg := fmt.Sprintf("Unknown command: %s\n", args[0])
		log.Print(msg)
		fmt.Print(msg)
	}
}

func information() string {
	// Print available commands

	info := new(strings.Builder)
	info.WriteString("Welcome to the bash parser program!\n")
	info.WriteString("=============================================================+\n")
	info.WriteString("=============================================================+\n")
	info.WriteString("----------------Available commands:--------------------------+\n")
	info.WriteString("ls 	- List directory contents                            |\n")
	info.WriteString("cd 	- Change directory                                   |\n")
	info.WriteString("mkdir 	- Create a new directory                             |\n")
	info.WriteString("pwd 	- Print the current working directory                |\n")
	info.WriteString("rm 	- Remove file                                        |\n")
	info.WriteString("touch 	- Create file                                        |\n")
	info.WriteString("curl 	- Displays the html of a web resource                |\n")
	info.WriteString("=============================================================+\n")
	info.WriteString("log 	- Shows data in log file                             |\n")
	info.WriteString("info 	- See this menu once more                            |\n")
	info.WriteString("exit 	- Exit the program                                   |\n")
	info.WriteString("=============================================================+\n")
	info.WriteString("=============================================================+\n")
	info.WriteString("Thank you for using our program, Your choice is awaited!")

	log.Printf("Info accessed")
	return info.String()
}

func displayLog() {
	// Display log file
	file, err := os.Open(lib.LogPath)
	if err != nil {
		log.Printf("Error opening log file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading log file: %v\n", err)
	}

	log.Print("Log accessed")
}

func listDirectory(path string) {
	// Get list of files and directories in path
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		log.Printf("Error running ls command with path %s: %v\n", path, err)
		fmt.Printf("Error! Please, check log file:\n %s", lib.LogPath)
		return
	}

	msg := fmt.Sprintf("Output of ls command with path %s:\n%s\n", path, files)
	log.Print(msg)
	fmt.Print(msg)
}

func changeDirectory(path string) {
	// Change working directory
	err := os.Chdir(path)
	if err != nil {
		log.Printf("Error running cd command with path %s: %v\n", path, err)
		fmt.Printf("Error! Please, check log file:\n %s", lib.LogPath)
		return
	}
	msg := fmt.Sprintf("Changed directory to %s\n", path)
	log.Print(msg)
	fmt.Print(msg)
}

func makeDirectory(path string) {
	// Create directory
	err := os.Mkdir(path, os.ModePerm)

	if err != nil {
		log.Printf("Error running mkdir command with path %s: %v\n", path, err)
		fmt.Printf("Error! Please, check log file:\n %s", lib.LogPath)
		return
	}

	msg := fmt.Sprintf("New directory created ->%s\n", path)
	log.Print(msg)
	fmt.Print(msg)
}

func printWorkingDirectory() {
	// Displays working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error running pwd command: %v\n", err)
		fmt.Printf("Error! Please, check log file:\n %s", lib.LogPath)
		return
	}

	msg := fmt.Sprintf("pwd - [%s]", wd)
	log.Print(msg)
	fmt.Println(wd)
}

func fetchURL(url string, w io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	output := string(body[:1000]) + "\n... only the first 1000 characters are displayed."
	fmt.Fprintf(w, "Curl of %s:\n %s\n", url, output)

	return nil
}
func createFile(path string) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		log.Printf("Error running touch command with path %s: %v\n", path, err)
		fmt.Printf("Error! Please, check log file:\n %s\n", lib.LogPath)
		return nil, err
	}
	msg := fmt.Sprintf("Created file %s\n", path)
	log.Print(msg)
	fmt.Print(msg)
	return f, nil
}

func removeFile(path string) {
	// Remove file
	err := os.Remove(path)
	if err != nil {
		log.Printf("Error running rm command for [%s]: %v\n", path, err)
		fmt.Printf("Error! Please, check log file:\n %s\n", lib.LogPath)
		return
	}

	msg := fmt.Sprintf("Removed file [%s]\n", path)
	log.Print(msg)
	fmt.Print(msg)
}
