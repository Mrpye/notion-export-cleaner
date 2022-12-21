package cmd

import (
	"errors"

	"github.com/Mrpye/notion-export-cleaner/notion"
	"github.com/spf13/cobra"
)

const CODE_BLOCK = "```"

func Clean_Command() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "clean [zip file] [output directory]",
		Short: "Command to read a notion.io export zip file and clean the uuid from the file names",
		Long: `
Description:
Command to read a notion.io export zip file and clean the uuid from the file names

Example Command:
` + CODE_BLOCK + `
hauler job run "test install" "uninstall"
` + CODE_BLOCK + `

		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			//********************************
			//Check for the required arguments
			//********************************
			if len(args) < 1 {
				return errors.New("missing zip file")
			}
			if len(args) < 2 {
				return errors.New("missing output directory")
			}
			//*************************
			//Lets clean the file names
			//*************************
			_, err := notion.UnzipCleanFileNames(args[0], args[1])
			return err

		},
	}

	return cmd
}

func init() {
	rootCmd.AddCommand(Clean_Command())
}
