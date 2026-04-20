package main

import (
	"fmt"
	"os"

	"github.com/greeschenko/vim9gorn"
)

func main() {
	fmt.Println("vim9gorn - Vim9 plugin generator")

	if len(os.Args) < 2 {
		fmt.Println("Usage: vim9gorn <command>")
		fmt.Println("Commands:")
		fmt.Println("  fetch   - Fetch external plugins")
		fmt.Println("  update - Update external plugins")
		fmt.Println("  build  - Generate vim config")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "fetch":
		fmt.Println("Fetching plugins...")
		manager := vim9gorn.NewPluginManager()
		if err := manager.FetchAll(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Done!")

	case "update":
		fmt.Println("Updating plugins...")
		manager := vim9gorn.NewPluginManager()
		if err := manager.UpdateAll(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Done!")

	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}
