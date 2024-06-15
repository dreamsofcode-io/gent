package message

import (
	"context"
	"fmt"

	gogptscript "github.com/gptscript-ai/go-gptscript"
)

func Generate(ctx context.Context) (string, error) {
	tools := gogptscript.ToolDefs{
		gogptscript.ToolDef{
			Tools: []string{"gitstatus", "sys.abort"},
			Instructions: `
      Create well formed git commit message based of off the currently staged file
      contents. The message should convey why something was changed and not what
      changed. Use the well known format that has the prefix chore, fix, etc.

      Only include changes to source files for the programming languages, shell configurations files, documentation such as readme and other .mds, and any changes to package management file. Exclude
      any lock or sum files.

      Do not use markdown format for the output.

			For the first line of the commit message, this must be constrained to 50 characters as a
			maximum.

      If there are no changes abort.`,
		},
		{
			Name: "gitstatus",
			Instructions: `
      #!/bin/sh

      git diff --staged`,
		},
	}

	client := &gogptscript.Client{}
	client.Complete()

	run, err := client.Evaluate(ctx, gogptscript.Opts{}, tools)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate: %w", err)
	}

	return run.Text()
}
