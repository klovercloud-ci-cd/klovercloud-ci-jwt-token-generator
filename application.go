package main

import (
	"github.com/klovercloud-ci/config"
	"github.com/spf13/cobra"
	"os"
)

func main(){
	config.InitEnvironmentVariables()
	cli()
}
func cli(){
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

func GenerateTokenCommand()*cobra.Command {
	return &cobra.Command{
		Use: "generate-jwt",
		Short:        "Generate token with public and private key",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, err := Jwt{}.GenerateToken(1000000, nil)
			if err != nil {
				cmd.Println("[ERROR]: ", err.Error())
			}
			cmd.Println("token: ", token)
			return nil
		},
	}
}