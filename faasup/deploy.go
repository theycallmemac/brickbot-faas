package cmd

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "os/exec"

    "github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
    Use:   "deploy",
    Short: "deploy a specified function",
    Long:  `This subcommand is used to deploy a specified function`,
    Run: func(cmd *cobra.Command, args []string) {
            deploy(cmd, args)
    },
}

func init() {
    RootCmd.AddCommand(DeployCmd)
    DeployCmd.Flags().BoolVarP(&all, "all", "a", false, "build all functions")
    DeployCmd.Flags().StringP("function", "f", "", "build specific function")
    DeployCmd.Flags().BoolVar(&verbose, "verbose", false, "enable verbose mode")
}

func deploy(cmd *cobra.Command,args []string) int {
    pwd, _ := os.Getwd()
    if all {
        files, err := ioutil.ReadDir("./functions")
        if err != nil {
            log.Fatal(err)
        }
        for _, f := range files {
            readAndDeploy(f.Name(), pwd)
        }
        return 0
    }
    function, _:= cmd.Flags().GetString("function")
    if function == "" {
        return 0
    }
    readAndDeploy(function, pwd)
    return 0
}


func readAndDeploy(name string, pwd string) {
    err := os.Chdir(pwd + "/functions/" + name)
    if err != nil {
        panic(err)
    }
    command := exec.Command("faas", "deploy", "-f", name + ".yml")
    if verbose {
        out, _ := command.Output()
        fmt.Println(string(out))
    }
    fmt.Println("Deployed " + name + " successfully!")
}
