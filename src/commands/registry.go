package commands

import "noteworthy/src/framework"

// Command represents command function
type Command func(framework.Context)

// Registry is a global holder for all commands
var Registry = make(map[string]Command)

// Register adds all commands to the Registry
func Register() {
	Registry["play"] = Play
}

// GetCommand returns Command
func GetCommand(cmd string) Command {
	return Registry[cmd]
}

// IsNil checks if Command is nil
func (c Command) IsNil() bool {
	return c == nil
}
