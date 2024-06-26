package message

import (
	"context"
	"fmt"

	gogptscript "github.com/gptscript-ai/go-gptscript"
)

var gitStatusTool = gogptscript.ToolDef{
	Name: "gitstatus",
	Instructions: `
  #!/bin/sh

  git diff --staged`,
}

var normalInstruction = `
Create well formed git commit message based of off the currently staged file
contents. The message should convey why something was changed and not what
changed. Use the well known format that has the prefix chore, fix, etc. Additionally
add in some emojis just for fun.

Only include changes to source files for the programming languages, shell configurations files, documentation such as readme and other .mds, and any changes to package management file. Exclude
any lock or sum files.

Do not use markdown format for the output.

For the first line of the commit message, this must be constrained to 40 characters as a
maximum and use additional lines for any further context.

If there are no changes abort.
`

var cursedInstruction = `
Create well formed git commit message using only emojis based of off the currently staged file
contents. The message should convey why something was changed and not what
changed but using only emojis. Do not use any other text characters.

Only include changes to source files for the programming languages, shell configurations files, documentation such as readme and other .mds, and any changes to package management file. Exclude
any lock or sum files.

Do not use markdown format for the output.

For the first line of the commit message, this must be constrained to 40 characters as a
maximum and use additional lines for any further context.

If there are no changes abort.
`

func Generate(ctx context.Context, cursed bool) (string, error) {
	instruction := normalInstruction
	if cursed {
		instruction = cursedInstruction
	}

	tools := gogptscript.ToolDefs{
		gogptscript.ToolDef{
			Tools:        []string{"gitstatus", "sys.abort"},
			Instructions: instruction,
		},
		gitStatusTool,
	}

	client := &gogptscript.Client{}
	client.Complete()

	run, err := client.Evaluate(ctx, gogptscript.Opts{}, tools)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate: %w", err)
	}

	return run.Text()
}
