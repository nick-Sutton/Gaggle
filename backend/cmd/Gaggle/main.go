package main

import (
	"fmt"
	"os"
	"github.com/nick-Sutton/Gaggle/backend/internal/logo"
)

func main() {
	logo.PrintConsoleLogo()

	var path string
	fmt.Println("Please enter the path to your CSV file:")
	fmt.Scanln(&path)

	file, err := os.Open(path)
    if err != nil {
		fmt.Println("Error opening file:", err)
        os.Exit(1)
    }

	fmt.Println("Your file was succcessfully opened.")
	defer file.Close()
}
