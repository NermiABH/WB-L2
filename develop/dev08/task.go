package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		switch command[0] {
		case "cd":
			Cd(command[1:])
		case "pwd":
			Pwd()
		case "echo":
			Echo(command[1:])
		case "kill":
			Kill(command[1:])
		case "ps":
			Ps()
		}
	}
}

func Cd(args []string) {
	if len(args) != 1 {
		fmt.Println("cd: ", "must be only 1 argument")
		return
	}
	if err := os.Chdir(args[0]); err != nil {
		fmt.Println("cd: ", err)
	}
}

func Pwd() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("pwd: ", err)
		return
	}
	fmt.Print(path)
}

func Echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func Kill(args []string) {
	for _, v := range args {
		if pid, err := strconv.Atoi(v); err != nil {
			fmt.Println("kill: process id ", pid, "is not correct")
		} else {
			if process, err := os.FindProcess(pid); err != nil {
				fmt.Println("kill: ", err)
			} else if err := process.Kill(); err != nil {
				fmt.Println("kill: ", err)
			}
		}
	}
}

func Ps() {
	processes, err := ps.Processes()
	if err != nil {
		fmt.Println("ps: ", err)
		return
	}
	for _, process := range processes {
		fmt.Printf("Process name: %v process id: %v\n", process.Executable(), process.Pid())
	}
}
