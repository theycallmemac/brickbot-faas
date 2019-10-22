package cmd

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var listVar bool

var MakeCmd = &cobra.Command{
    Use:   "make",
    Short: "make a new function",
    Long:  `This subcommand is used to make a new function`,
    Run: func(cmd *cobra.Command, args []string) {
            makeFunc(cmd, args)
    },
}

func init() {
    RootCmd.AddCommand(MakeCmd)
    MakeCmd.Flags().BoolVarP(&listVar, "list", "l", false, "list available templates")
    MakeCmd.Flags().StringP("name", "n", "", "specify a name for your function")
    MakeCmd.Flags().StringP("template", "t", "", "specify a template for your function")
}

func makeFunc(cmd *cobra.Command,args []string) int {
    if listVar {
        command := exec.Command("faas", "template", "store", "list")
        out, _ := command.Output()
        fmt.Println(string(out))
        return 0;
    }
    name, _ := cmd.Flags().GetString("name")
    template, _ := cmd.Flags().GetString("template")
    pwd, _ := os.Getwd()
    pwd += "/functions/" + name
    if (name == "" || template == "") {
        return 0
    }
    os.MkdirAll(pwd, 0775)
    os.Chdir(pwd)
    command := exec.Command("faas", "new", "--lang", template, name, "--prefix", "theycallmemac")
    out, _ := command.Output()
    fmt.Println(string(out))
    os.Chdir("../")
    return 0
}

