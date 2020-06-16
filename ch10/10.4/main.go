package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

/*
	Construct a tool that reports the set of all packages in the workspace that
	transitively depend on the packages specified by the arguments.

	Hint: you will need to run `go list` twice, once for the initial packages
	and once for all packages. You may want to parse its JSON output using the
	encoding/json package (§4.5)
*/

var defaultArgs = []string{"list", "-f", "'{{ join .Imports \" \"}}'"}

func main() {
	buf := bytes.NewBuffer([]byte{})
	args := defaultArgs[:]
	if len(os.Args) < 2 {
		args = append(args, os.Args[1:]...)
	}
	// get imports for current workspace OR supplied packages
	err := runCommand(args, buf)
	if err != nil {
		log.Fatalf("%v", err)
	}
	imports := bufOutput(buf)

	// get transitive imports from initial imports
	args = append(defaultArgs[:], imports...)
	err = runCommand(args, buf)
	if err != nil {
		log.Fatalf("%v", err)
	}
	transitiveImports := bufOutput(buf)

	output := bytes.NewBuffer([]byte{})
	for _, v := range transitiveImports {
		output.WriteString(v)
		output.WriteString("\n")
	}
	output.WriteTo(os.Stdout)
}

func runCommand(args []string, buf *bytes.Buffer) error {
	buf.Reset()
	c := exec.Command("go", args...)
	c.Stdout = buf
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}

// bufOutput removes newlines and quotations from buf's output and
// returns a slice of sorted unique strings
func bufOutput(buf *bytes.Buffer) []string {
	s := strings.TrimSuffix(buf.String(), "\n")
	s = strings.ReplaceAll(s, "'", "")
	o := removeDups(strings.Fields(s))
	sort.Strings(o)
	return o
}

func removeDups(s []string) []string {
	out := []string{}
	seen := map[string]struct{}{}
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			out = append(out, v)
			seen[v] = struct{}{}
		}
	}
	return out
}
