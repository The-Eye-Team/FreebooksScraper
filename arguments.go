package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func parseArgs(args []string) {
	// Create new parser object
	parser := argparse.NewParser("FreebooksScraper", "Scraper for freebooks.com")

	// Create flags
	input := parser.String("i", "input", &argparse.Options{
		Required: false,
		Help:     "Input link"})

	output := parser.String("o", "output", &argparse.Options{
		Required: false,
		Default:  "./Downloads",
		Help:     "Output directory"})

	randomUA := parser.Flag("", "random-ua", &argparse.Options{
		Required: false,
		Help:     "Randomize user agent on request"})

	concurrency := parser.Int("j", "concurrency", &argparse.Options{
		Required: false,
		Default:  2,
		Help:     "Concurrent jobs for download"})

	startID := parser.Int("", "start-id", &argparse.Options{
		Required: false,
		Default:  1,
		Help:     "Start from this ID"})

	stopID := parser.Int("", "stop-id", &argparse.Options{
		Required: false,
		Default:  10000000,
		Help:     "Stop at this ID"})

	verbose := parser.Flag("v", "verbose", &argparse.Options{
		Required: false,
		Help:     "Display various informations"})

	// Parse input
	err := parser.Parse(args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(0)
	}

	// Fill arguments structure
	arguments.Input = *input
	arguments.Output = *output
	arguments.Concurrency = *concurrency
	arguments.RandomUA = *randomUA
	arguments.Verbose = *verbose
	arguments.StartID = *startID
	arguments.StopID = *stopID
}
