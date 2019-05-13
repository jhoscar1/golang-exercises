package main

import "fmt"

func main() {
	friends := map[string][]string{
		"oscar_jason":    {"friends", "movies", "chicken fingers"},
		"herren_josh":    {"musicals", "movies", "teaching"},
		"mansfield_june": {"ceramics", "movies", "savvy"},
	}

	friends["steinberg_henry"] = []string{"printing", "games", "movies"}

	delete(friends, "oscar_jason")

	for _, v := range friends {
		for ind, item := range v {
			fmt.Println(ind, item)
		}
	}
}
