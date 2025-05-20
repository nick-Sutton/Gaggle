package main

import (
	"fmt"
	// "log"
	// "net/http"
	"os"

	// "github.com/nick-Sutton/Gaggle/backend/api"
	"github.com/nick-Sutton/Gaggle/backend/internal/logo"
)

func main() {
	logo.PrintConsoleLogo()

	// // setup router
	// r := api.SetupRouter()
	// log.Println("Server listening on http://localhost:3000")
	// http.ListenAndServe(":3000", r)

	// get file from user
	var path string
	fmt.Println("Please enter the path to your CSV file:")
	fmt.Scanln(&path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Println("Your file was succcessfully opened.")
}
