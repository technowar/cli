package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Read(filename string) (string, error) {
	file, err := os.Open("/tmp/" + filename)

	defer file.Close()

	if err != nil {
		return "", errors.New("File not found")
	}

	read := bufio.NewReader(file)

	content, _, err := read.ReadLine()

	if err != nil {
		return "", errors.New("File not found")
	}

	return string(content), nil
}

func Write(token Parameter, filename string) {
	file, err := os.Create("/tmp/" + filename)

	defer file.Close()

	if err != nil {
		panic(err)
	}

	write := bufio.NewWriter(file)
	_, err = write.WriteString(token.(string))

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nFile created in /tmp/%s", filename)

	write.Flush()
}
