package sgogen

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Generator struct {
	cmd      string
	isDryRun bool
	paths    []string
}

func New(cmd string, isDryRun bool, paths []string) *Generator {
	if len(paths) == 0 {
		paths = []string{"."}
	}

	return &Generator{
		cmd:      cmd,
		isDryRun: isDryRun,
		paths:    paths,
	}
}

func (g *Generator) Gen() error {
	if g.cmd == "" {
		return nil
	}

	args := []string{"generate", "-n"}
	args = append(args, g.paths...)

	r, err := executeCmd("go", args...)
	if err != nil {
		return fmt.Errorf("failed to executeCmd: %w", err)
	}
	for _, c := range strings.Split(r, "\n") {
		if strings.HasPrefix(c, g.cmd) {
			fmt.Printf("[execute command]%s\n", c)

			if !g.isDryRun {
				args := strings.Split(c, " ")
				e, err := executeCmd(args[0], args[1:]...)
				if err != nil {
					return fmt.Errorf("failed to executeCmd: %w", err)
				}

				if e == "" {
					fmt.Println("[executed]")
				} else {
					fmt.Printf("[results] %s", e)
				}
			}
		}
	}
	return nil
}

func executeCmd(cmd string, args ...string) (string, error) {
	var stdout bytes.Buffer
	c := exec.Command(cmd, args...)
	c.Stderr = &stdout
	c.Stdout = &stdout
	if err := c.Run(); err != nil {
		return "", fmt.Errorf("failed to %s: %s", cmd, &stdout)
	}

	return stdout.String(), nil
}
