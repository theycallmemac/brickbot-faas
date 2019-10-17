package cmd

import (
    "fmt"
    "io/ioutil"
    "log"

    "github.com/spf13/cobra"
)


var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "list usable functions",
    Long:  `This subcommand is used to list usable functions`,
    Run: func(cmd *cobra.Command, args []string) {
            list()
    },
}

func init() {
    RootCmd.AddCommand(ListCmd)
}

func list() int {
    files, err := ioutil.ReadDir("./functions")
    if err != nil {
        log.Fatal(err)
    }
    for _, f := range files {
        fmt.Println(f.Name())
    }
    return 0
}


