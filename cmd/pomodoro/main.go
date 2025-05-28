package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"pomodoro/internal/timer"
	"pomodoro/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	workFlag  = flag.Int("w", 25, "Work session duration in minutes")
	breakFlag = flag.Int("b", 5, "Break session duration in minutes")
	helpFlag  = flag.Bool("h", false, "Show help message")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Pomodoro Timer\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s                 # Default: 25min work, 5min break\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -w 30 -b 10     # 30min work, 10min break\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -w 45           # 45min work, 5min break (default)\n", os.Args[0])
	}

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	// Validate inputs
	if *workFlag <= 0 {
		fmt.Fprintf(os.Stderr, "Error: Work duration must be positive, got %d\n", *workFlag)
		os.Exit(1)
	}
	if *breakFlag <= 0 {
		fmt.Fprintf(os.Stderr, "Error: Break duration must be positive, got %d\n", *breakFlag)
		os.Exit(1)
	}

	// Create timer configuration
	config := timer.Config{
		WorkDuration:  time.Duration(*workFlag) * time.Minute,
		BreakDuration: time.Duration(*breakFlag) * time.Minute,
	}

	// Create and run the application
	model := ui.NewModel(config)
	program := tea.NewProgram(
		model,
		tea.WithAltScreen(),
	)

	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
