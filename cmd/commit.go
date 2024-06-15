/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"

	"github.com/dreamsofcode-io/gent/internal/message"
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

		status, err := w.Status()
		if err != nil {
			return fmt.Errorf("failed to get git status, %w", err)
		}

		if status.IsClean() {
			return fmt.Errorf("Git worktree is clean, no need to run gent")
		}

		msg, err := message.Generate(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to generate a message: %w", err)
		}

		command := exec.Command("git", "commit", "-em", msg)
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout

		err = command.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to start proc:", err)
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
