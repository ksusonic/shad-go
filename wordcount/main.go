//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFiles(args []string) ([]string, error) {
	var res []string
	for _, arg := range args {
		file, err := os.OpenFile(arg, os.O_RDONLY, 0755)
		if err != nil {
			return nil, err
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			res = append(res, scanner.Text())
		}

		if err = scanner.Err(); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func main() {
	m := make(map[string]int)

	fileData, err := readFiles(os.Args[1:])
	if err != nil {
		panic(err)
	}
	for _, d := range fileData {
		_, ok := m[d]
		if ok {
			m[d]++
		} else {
			m[d] = 1
		}
	}

	for key, value := range m {
		if value > 1 {
			fmt.Printf("%d\t%s\n", value, key)
		}
	}
}
