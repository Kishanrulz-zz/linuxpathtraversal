package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const root = "/"

//init initializes program from root directory
func init() {
	_ = os.Chdir(root)
}
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {

		command := make([]string, 2)
		text, _ := reader.ReadString('\n')

		if strings.Contains(text, " ") {
			command = strings.Split(text, " ")
		} else {
			command[0] = text
		}

		cmd := strings.Trim(command[0], "\n")
		action := strings.Trim(command[1], "\n")

		switch cmd {
		case "pwd":
			dir := pwd()
			fmt.Println(dir)

		case "mkdir":
			dir := pwd()
			dirExists, err := exists(fmt.Sprintf(dir + "/" + action))
			if err != nil {
				fmt.Println(err)
			}
			if dirExists {
				fmt.Println("Sorry, directory already exists")
			} else {
				if err := os.Mkdir(action, 0777); err != nil {
					fmt.Println("Invalid path")
				}
				fmt.Println("Directory created successfully")
			}

		case "ls":
			dir := pwd()
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				fmt.Println(err)
			}
			for _, file := range files {
				fmt.Println(file.Name(), file.Mode().Perm())
			}

		case "ll":
			dir := pwd()
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				fmt.Println(err)
			}
			for _, file := range files {
				fmt.Println(file.Name(), file.Mode().Perm(), file.ModTime())
			}

		case "cd":
			if err := os.Chdir(action); err != nil {
				fmt.Println(err)
			}
			dir := pwd()
			fmt.Println(dir)

		case "rm":
			if err := os.Remove(action); err != nil {
				fmt.Println("Sorry unable to delete")
			}
			fmt.Println("Deleted successfully")

		case "clear":
			_ = os.Chdir(root)
			fmt.Println(pwd())
		}
	}
}

//pwd returns the present working directory
func pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

//exists returns true if the directory exists else false
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
