package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	name    string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "my-cli-app",
	Short: "A simple Cobra CLI application",
	Long:  `This is a simple Cobra CLI application. It has two subcommands: create and list.`,
}

func init() {
	// Create the create subcommand
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create something",
		Long:  `This subcommand creates something.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Creating something named %s...\n", name)
		},
	}

	// Add the --name flag to the create subcommand
	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "The name of the thing to create")

	// Add the create subcommand to the root command
	rootCmd.AddCommand(createCmd)

	// Create the list subcommand
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List something",
		Long:  `This subcommand lists something.`,
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				fmt.Println("Listing something verbosely...")
			} else {
				fmt.Println("Listing something...")
			}
		},
	}

	// Add the --verbose flag to the list subcommand
	listCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	// Add the list subcommand to the root command
	rootCmd.AddCommand(listCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
