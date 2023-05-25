package main

import (
  "fmt"
  "flag"
  "os"

  "ggit/internal/cmdconfig"
  "ggit/pkg/commands"
)

func main() {
  cmdFlags := flag.NewFlagSet("ggit", flag.ExitOnError)
  cliConfig := cmdconfig.ParseFlags()

  if cliConfig.HelpFlag {
    cmdconfig.PrintHelp()
  }

  if len(os.Args) < 2 {
    cmdconfig.PrintHelp()
  }

  command := os.Args[1]
  cmdFlags.Parse(os.Args[2:])

  switch command {
  case "st":
    if cliConfig.HelpFlag {
      cmdFlags.PrintDefaults()
    } else {
      commands.Status()
    }
  default:
    fmt.Printf("Unknown command: %s\n", command)
  }
}
