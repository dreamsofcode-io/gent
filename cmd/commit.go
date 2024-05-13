/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dreamsofcode-io/gent/internal/message"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Creates a commit with a given commit message",
	Long:  `Creates a commit message and commits for you if acceptable`,
	RunE: func(cmd *cobra.Command, args []string) error {
		directory, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %w", err)
		}

		r, err := git.PlainOpen(directory)
		if err != nil {
			return fmt.Errorf("failed to open repository: %w", err)
		}

		w, err := r.Worktree()
		if err != nil {
			return fmt.Errorf("failed to get worktree: %w", err)
		}

		msg, err := message.Generate(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to generate a message: %w", err)
		}

		fmt.Printf("Generated Commit Message:\n\n")
		fmt.Println(msg)
		fmt.Printf("\nProceed with commit message? (Y/n): ")

		var confirmation string

		_, err = fmt.Scanln(&confirmation)
		if err != nil {
			return fmt.Errorf("failed to scan input: %w", err)
		}

		if confirmation != "" && strings.ToLower(confirmation) != "y" {
			fmt.Println("Aborting commit. No changes were made.")
			return nil
		}

		commit, err := w.Commit(msg, &git.CommitOptions{})
		if err != nil {
			return fmt.Errorf("failed to commit: %w", err)
		}

		obj, err := r.CommitObject(commit)
		if err != nil {
			return fmt.Errorf("failed to get commit obj: %w", err)
		}

		fmt.Println(obj)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

}
