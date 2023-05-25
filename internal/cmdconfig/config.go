package cmdconfig

import (
  "flag"
  "fmt"
  "os"
)

type CLIConfig struct {
  HelpFlag bool
}

func ParseFlags() *CLIConfig {
  helpFlag := flag.Bool("help", false, "Show help")

  flag.Parse()

  return &CLIConfig{
    HelpFlag: *helpFlag,
  }
}

func PrintHelp() {
  fmt.Println("Usage: ggit <command> [<args>]")
  os.Exit(1)
}
