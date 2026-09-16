package main

import (
	"fmt"
	"os"
)

const VERSION = "0.1.0"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v":
			fmt.Printf("copilot-cli v%s\n", VERSION)
		case "--help", "-h":
			printHelp()
		default:
			fmt.Fprintf(os.Stderr, "Unknown command: %s\n", os.Args[1])
			os.Exit(1)
		}
	} else {
		printHelp()
	}
}

func printHelp() {
	fmt.Println(`Copilot CLI - AI-ready command-line tool

Usage:
  copilot-cli [command] [flags]

Commands:
  auth       Authentication and authorization
  config     Configuration management
  skills     AI agent skills
  shortcuts  Domain-specific shortcuts

Flags:
  -h, --help      Show this help message
  -v, --version   Show version

For more information, visit: https://github.com/BossmanCryto/copilot-cli-starter
`)
}
