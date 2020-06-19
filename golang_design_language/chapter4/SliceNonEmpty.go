package main

import "fmt"

func main() {
	dara := []string{"dsd", "", "ASd"}
	fmt.Println(nonempty(dara)) //[dsd ASd]
	fmt.Println(dara)           //[dsd ASd ASd]

	dara2 := []string{"dsd", "", "ASd"}
	fmt.Println(nonempty2(dara2)) //[dsd ASd]
	fmt.Println(dara2)            //[dsd ASd ASd]
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
