// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package terminalstdin represents the imported interface "wasi:cli/terminal-stdin@0.2.0".
package terminalstdin

import (
	terminalinput "github.com/Lay3rLabs/wavs-wasi/go/wasi/cli/terminal-input"
	"go.bytecodealliance.org/cm"
)

// TerminalInput represents the imported type alias "wasi:cli/terminal-stdin@0.2.0#terminal-input".
//
// See [terminalinput.TerminalInput] for more information.
type TerminalInput = terminalinput.TerminalInput

// GetTerminalStdin represents the imported function "get-terminal-stdin".
//
//	get-terminal-stdin: func() -> option<terminal-input>
//
//go:nosplit
func GetTerminalStdin() (result cm.Option[TerminalInput]) {
	wasmimport_GetTerminalStdin(&result)
	return
}
