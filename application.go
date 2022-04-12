package main

import (
	"github.com/klovercloud-ci/config"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func main() {
	config.InitEnvironmentVariables()
	cli()
}
func cli() {
	cmd := &cobra.Command{
		Use:          "kcpctl",
		Short:        "Cli to use kloverCloud platform apis!",
		Version:      "v1",
		SilenceUsage: true,
	}
	cmd.AddCommand(GenerateTokenCommand())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func GenerateTokenCommand() *cobra.Command {
	return &cobra.Command{
		Use:       "generate-jwt",
		Short:     "Generate token with public and private key",
		ValidArgs: []string{},
		RunE: func(cmd *cobra.Command, args []string) error {
			type AgentStruct struct {
				Name string `json:"name"`
			}
			var agent AgentStruct
			for idx, _ := range args {
				if strings.Contains(strings.ToLower(args[idx]), "client=") {
					strs := strings.Split(strings.ToLower(args[idx]), "=")
					if len(strs) > 1 {
						agent.Name = strs[1]
					}
				}
			}
			token, err := Jwt{}.GenerateToken(1000000, agent)
			if err != nil {
				cmd.Println("[ERROR]: ", err.Error())
			}
			cmd.Println("token: ", token)
			return nil
		},
	}
}
