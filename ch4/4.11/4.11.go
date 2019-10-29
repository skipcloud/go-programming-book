package main

import (
	"fmt"
	"gobook/ch4/github2"
	"log"
	"os"
)

/*
 * Build a tool that lets users create, read, update, and close Github issues from
 * the command line, invokinng their preferred text editor when substancial input
 * is required.
 */

func main() {
	if len(os.Args) == 1 {
		fmt.Fprint(os.Stderr, "not enough input\n")
		os.Exit(1)
	}
	verb := os.Args[1]

	switch verb {
	case "search": // read
		resp, _ := github2.SearchIssues(os.Args[2:])
		for _, item := range resp.Items {
			fmt.Printf("%8d\t%.20s...\t%15s\n", item.Number, item.Title, item.User.Login)
		}
	case "create":
		var owner, repo string

		fmt.Printf("Please enter the repo owner: ")
		_, err := fmt.Scanln(&owner)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Please enter the repo name: ")
		_, err = fmt.Scanln(&repo)
		if err != nil {
			log.Fatal(err)
		}

		resp := github2.CreateIssue(owner, repo)
		fmt.Println(resp)
	default:
		fmt.Fprint(os.Stderr, "%s is not a recognised command: usage create/read/update/close\n", verb)
		os.Exit(1)
	}
}

// func openFileInEditor(filename string) error {
// 	const DefaultEditor = "vim"

// 	editor := os.Getenv("EDITOR")
// 	if editor == "" {
// 		editor = DefaultEditor
// 	}

// 	path, err := exec.LookPath(editor)
// 	if err != nil {
// 		return err
// 	}

// 	cmd := exec.Command("vim", "/tmp/userInput.txt")
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	err := cmd.Run()
// }

// func captureInputFromEditor() ([]byte, error) {
// 	file, err := ioutil.TempFile("", "*")
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	filename := file.Name()
// 	defer os.Remove(filename)

// 	if err = file.Close(); err != nil {
// 		return []byte{}, err
// 	}

// 	if err = openFileInEditor(filename); err != nil {
// 		return []byte{}, err
// 	}

// 	bytes, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	return bytes, nil

// }
