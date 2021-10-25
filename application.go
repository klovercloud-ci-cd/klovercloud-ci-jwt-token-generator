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
		Short:        "Generate Jwt token from keys!",
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
		RunE: func(cmd *cobra.Command, args []string) error {
			token, err :=	Jwt{}.GenerateToken(10000,nil)
			if err != nil {
				cmd.Println("[ERROR]: ", err.Error())
			}
			cmd.Println("token: ", token)
			return nil
		},
	}
}