package cmd

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var all, verbose bool

var BuildCmd = &cobra.Command{
    Use:   "build",
    Short: "build a specified function",
    Long:  `This subcommand is used to build a specified function`,
    Run: func(cmd *cobra.Command, args []string) {
            build(cmd, args)
    },
}

func init() {
    RootCmd.AddCommand(BuildCmd)
    BuildCmd.Flags().BoolVarP(&all, "all", "a", false, "build all functions")
    BuildCmd.Flags().StringP("function", "f", "", "build specific function")
    BuildCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose mode")
}

func build(cmd *cobra.Command,args []string) int {
    pwd, _ := os.Getwd()
    if all {
        files, err := ioutil.ReadDir("./functions")
        if err != nil {
            log.Fatal(err)
        }
        for _, f := range files {
            readAndExecute(f.Name(), pwd)
        }
        return 0
    }
    function, _ := cmd.Flags().GetString("function")
    if function == "" {
        return 0
    }
    readAndExecute(function, pwd)
    return 0
}


func readAndExecute(name string, pwd string) {
    err := os.Chdir(pwd + "/functions/" + name)
    if err != nil {
        panic(err)
    }
    command := exec.Command("faas", "build", "-f", name + ".yml")
    if verbose {
        out, _ := command.Output()
        fmt.Println(string(out))
    }
    fmt.Println("Built " + name + " successfully!")
}
