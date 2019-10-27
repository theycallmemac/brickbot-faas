package cmd

import (
    "fmt"
    "os/exec"

    "github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
    Use:   "test",
    Short: "test a specified function",
    Long:  `This subcommand is used to test a specified function`,
    Run: func(cmd *cobra.Command, args []string) {
            test(cmd, args)
    },
}

func init() {
    RootCmd.AddCommand(TestCmd)
    TestCmd.Flags().StringP("function", "f", "", "test specific function")
    TestCmd.Flags().StringP("input", "i", "", "give input as a request body")
}

func test(cmd *cobra.Command,args []string) int {
    url := "faas.jamesmcdermott.ie/function/"
    function, _ := cmd.Flags().GetString("function")
    input, _ := cmd.Flags().GetString("input")
    if function == "" {
        return 0
    }
    readAndTest(function, input, url)
    return 0
}


func readAndTest(name string, input string, url string) {
    command := exec.Command("curl", url + name, "--data", input)
    out, _ := command.Output()
    fmt.Println(string(out))
}
