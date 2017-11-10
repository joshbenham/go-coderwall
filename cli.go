package main

import (
	"fmt"

	"github.com/ttacon/chalk"
)

// StyleError styles Errors
func StyleError(value string) error {
	red := chalk.Red.NewStyle().
		WithTextStyle(chalk.Bold).
		Style

	fmt.Println(red(value))

	return fmt.Errorf(value)
}

// StyleHeading styles Heading
func StyleHeading(value string) {
	red := chalk.Green.NewStyle().
		WithTextStyle(chalk.Bold).
		Style

	fmt.Println(red(value))
}
