package cmd

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
    Use:   "info",
    Short: "display info on a specified function",
    Long:  `This subcommand is used to disaply info on a specified function`,
    Run: func(cmd *cobra.Command, args []string) {
            info(cmd, args)
    },
}

func init() {
    RootCmd.AddCommand(InfoCmd)
    InfoCmd.Flags().StringP("function", "f", "", "display info for specific function")
}

func info(cmd *cobra.Command,args []string) int {
    pwd, _ := os.Getwd()
    function, _ := cmd.Flags().GetString("function")
    if function == "" {
        return 0
    }
    readInfo(function, pwd)
    return 0
}


func readInfo(name string, pwd string) {
    err := os.Chdir(pwd + "/functions/" + name)
    if err != nil {
        panic(err)
    }
    command := exec.Command("faas", "describe", name)
    out, _ := command.Output()
    fmt.Println(string(out))
}
