package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Hi there! This is OCP Vacancy service (ocp-vacancy-api)")

	readFileFunc := func(path string) error {
		f, err := os.Open(path)
		if err != nil {
			return errors.Wrap(err, "can not open file")
		}
		defer f.Close()
		return nil
	}

	for i := 0; i < 5; i++ {
		if err := readFileFunc("config.cfg"); err != nil {
			fmt.Println(err)
			break
		}
	}
}
