package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	agents := [][]string{jb, mp}

	for _, agent := range agents {
		fmt.Println(agent)
		for _, details := range agent {
			fmt.Println(details)
		}
	}
}
