package commands

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

type Entry struct {
  StatusEntryWorktree git.StatusCode
  File                string
}

func Status() {
  repo, err := git.PlainOpen(".")

  if err != nil {
    fmt.Println("Error opening repository: ", err)
    os.Exit(1)
  }

  worktree, err := repo.Worktree()
  if err != nil {
    fmt.Println("Error getting worktree: ", err)
    os.Exit(1)
  }

  status, err := worktree.Status()
  if err != nil {
    fmt.Println("Error getting status: ", err)
    os.Exit(1)
  } 

  entries := handleWorktreeStatus(status)

  for _, entry := range entries {
    fmt.Println(entry.StatusEntryWorktree, entry.File)
  }
}

func handleWorktreeStatus(status git.Status) []Entry {
  entries := make([]Entry, 0)

  for file, statusEntry := range status {
    entry := Entry{
      StatusEntryWorktree: statusEntry.Worktree,
      File:                file,
    }

    entries = append(entries, entry)
  }

  return entries
}

// func handleWorktreeStatus(worktreeStatus git.StatusCode) ([]string, []string) {
//   untracked := make([]string, 0)
//   modified := make([]string, 0)

//   switch worktreeStatus {
//   case git.Added:
//     

//   case git.Modified:
//     fmt.Println("Arquivo modificado")

//   case git.Unmodified:
//     fmt.Println("Arquivo não modificado")

//   default:
//     fmt.Println("Status desconhecido")
//   }
// }

