package common

import (
	"bufio"
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Printf("encountered error: %v\n", err)
	os.Exit(1)
}

func Input(name string) ([]string, error) {
	res := []string{}
	f, err := os.Open(name)
	if err != nil {
		return res, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return res, err
	}
	return res, nil
}
