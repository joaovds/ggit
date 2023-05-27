package commands

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

type Entry struct {
  StatusEntryWorktree git.StatusCode
  File                string
  StatusDescription   string
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

  entryGroups := HandleWorktreeStatus(status)

  PrintStatus(entryGroups)
}

func HandleWorktreeStatus(status git.Status) map[string][]Entry {
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
      StatusDescription:   "",
    }

    switch statusEntryWorktree {
    case git.Added:
      entry.StatusDescription = "Added"
      adddeds = append(adddeds, entry)

    case git.Untracked:
      entry.StatusDescription = "Untracked"
      untrackeds = append(untrackeds, entry)

    case git.Copied:
      entry.StatusDescription = "Copied"
      modifieds = append(modifieds, entry)

    case git.Modified:
      entry.StatusDescription = "Modified"
      modifieds = append(modifieds, entry)

    case git.Renamed:
      entry.StatusDescription = "Renamed"
      modifieds = append(modifieds, entry)

    case git.UpdatedButUnmerged:
      entry.StatusDescription = "Unmerged"
      unmergeds = append(unmergeds, entry)

    default:
      entry.StatusDescription = "Unknown"
      unknowns = append(unknowns, entry)
    }
  }

  entryGroups := make(map[string][]Entry)

  entryGroups["Modifieds"] = modifieds
  entryGroups["Untrackeds"] = untrackeds
  entryGroups["Addeds"] = adddeds
  entryGroups["Unmergeds"] = unmergeds
  entryGroups["Unknowns"] = unknowns

  return entryGroups
}

func PrintStatus(entryGroups map[string][]Entry) {
  for key := range entryGroups {
    if len(entryGroups[key]) == 0 { continue }

    blue := color.New(color.FgBlue).SprintFunc()
    
    fmt.Println(fmt.Sprintf("%s\n", blue(key)))

    for _, entryGroup := range entryGroups[key] {
      red := color.New(color.FgRed).SprintFunc()

      fmt.Println(red(fmt.Sprintf("%s - %s", entryGroup.StatusDescription, entryGroup.File)))
    }

    println("")
  }
}

