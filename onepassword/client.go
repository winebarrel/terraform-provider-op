package onepassword

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Client struct {
	command string
}

func NewClient(cmd string) (*Client, error) {
	err := exec.Command(cmd, "--version").Run()

	if err != nil {
		return nil, fmt.Errorf("failed to execute '%s --version': %w", cmd, err)
	}

	client := &Client{
		command: cmd,
	}

	return client, nil
}

func (client *Client) GetItem(idOrTitle string, options ...string) (*Item, error) {
	cmdArgs := []string{client.command, idOrTitle, "--format", "json"}
	cmdArgs = append(cmdArgs, options...)

	cmd := exec.Command(cmdArgs[0], cmdArgs[:1]...)
	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Run()

	if err != nil {
		return nil, fmt.Errorf("failed to get 1password item: %w: %s", err, errbuf.String())
	}

	item := &Item{}
	err = json.Unmarshal(outbuf.Bytes(), item)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal 1password item: %w", err)
	}

	return item, nil
}
