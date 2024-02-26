package internal

import "os"

// Read data in specified path
func Read(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("Error occured while reading file")
	}
	return string(data)
}

// Write data in specified path
func Write(data, filename string) {
	err := os.WriteFile(filename, []byte(data), 0777)
	if err != nil {
		panic("Error occured while writing file")
	}
}
