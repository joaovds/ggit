package commands

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

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

  fmt.Println(status)
}
