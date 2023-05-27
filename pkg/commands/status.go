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

  entryGroups := handleWorktreeStatus(status)

  for _, entryGroup := range entryGroups {
    fmt.Println(entryGroup)
  }
}

func handleWorktreeStatus(status git.Status) [][]Entry {
  untrackeds := make([]Entry, 0)
  modifieds := make([]Entry, 0)
  adddeds := make([]Entry, 0)
  unmergeds := make([]Entry, 0)
  unknowns := make([]Entry, 0)

  for file, statusEntry := range status {
    statusEntryWorktree := statusEntry.Worktree

    entry := Entry{
      StatusEntryWorktree: statusEntryWorktree,
      File:                file,
    }

    switch statusEntryWorktree {
    case git.Added:
      adddeds = append(adddeds, entry)

    case git.Untracked:
      untrackeds = append(untrackeds, entry)

    case git.Modified, git.Renamed, git.Copied:
      modifieds = append(modifieds, entry)

    case git.UpdatedButUnmerged:
      unmergeds = append(unmergeds, entry)

    default:
      unknowns = append(unknowns, entry)
    }
  }

  entryGroups := [][]Entry{
    untrackeds,
    modifieds,
    adddeds,
    unmergeds,
    unknowns,
  }

  return entryGroups
}

