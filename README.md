# Go Bash Parser Program 

This program is a command-line interface (CLI) tool that allows users to execute various commands within a shell environment.

## Features
* Basic file manipulation commands (`cd`, `ls`, `touch`, `rm`, `mkdir`, `pwd`)
* Ability to fetch the HTML of a web resource using the `curl` command
* Log file for keeping track of executed commands
* Fun cat pictures upon request :cat::dog::elf:

## How to use
1. Clone the repository
2. Open the terminal and navigate to the directory where the repository was cloned
3. Run the command go build to build the program
4. Run the command ./GoBashParser to start the program
5. Enter any command of your choice and hit <kbd>Enter</kbd>

## Available Commands
* `ls` - List directory contents
* `cd` - Change directory
* `mkdir` - Create a new directory
* `pwd` - Print the current working directory
* `rm` - Remove file
* `touch` - Create file
* `curl` - Displays the HTML of a web resource
* `log` - Shows data in log file
* `catsneeded` - Show cat pictures to make you smile :)
* `info` - See available commands again
* `exit` - Exit the program

### Examples
```
> ls
# list the contents of the current directory

> cd my-folder
# change the current working directory to my-folder

> mkdir my-new-folder
# create a new directory named my-new-folder

> pwd
# print the current working directory

> rm my-file.txt
# remove the file named my-file.txt

> touch new-file.txt
# create a new file named new-file.txt

> curl https://www.google.com
# fetch the first 1000 chars of the HTML content of https://www.google.com

> log
# display the program log

> catsneeded
# display some ASCII art of cats

> info
# display the list of available commands

> exit
# exit the program
```

## Note
This program is for educational purposes only.
It is not meant to be used in a production environment.

## Possible Next Steps

### Bugs fixing
* When the directory is changed, the `log` command does not function:
`2023/02/15 14:14:04 Error opening log file: open ../logs/log_2023-02-15T121305.7452648Z.log: The system cannot find the path specified.`

### Quick wins
* Improve error handling and error messages
* Increase test coverage
* Add more commands to the parser
* Add support for running scripts, reading commands from a file and executing them
* Replace `ioutil` with `os` everywhere, as **`io/ioutil` has been deprecated since Go 1.16: As of Go 1.16, the same functionality is now provided by package io or package os, and those implementations should be preferred in new code. See the specific function documentation for details.  (SA1019)**

### Not so quick wins
* Implement the program as a web service instead of a CLI
* Add support for environment variables
* Add support for piping commands - `func runPipe(cmd1, cmd2 string)`
* Add support for tab completion
* Create a Dockerfile to build and run the program inside a container
* Add support for scripts or aliases so users can create their own custom commands

### Thank you for your attention!:smile::school::dancer:
