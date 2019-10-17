package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)


var cfgFile string

var RootCmd = &cobra.Command{
    Use:   "faasup",
    Short: "manage the brickbot openfaas functions",
    Long: `manage the brickbot openfaas functions`,
}


func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
