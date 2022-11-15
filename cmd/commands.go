package cmd

import (
	"fmt"
	"github.com/fiqrikm18/cli/tools/pkg/models"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cliTools [log dir]",
	Short: "Copy log file into file with format json or text.",
	Run:   runCommand,
}

var output string
var outputType string

func Execute() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Need log directory path parameters.")
	}

	rootCmd.Flags().StringVarP(&output, "output", "o", homeDir+"/", "choose directory where you want place the output file")
	rootCmd.Flags().StringVarP(&outputType, "type", "t", "text", "output file type which json or text")

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runCommand(cmd *cobra.Command, args []string) {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		log.Fatal(err)
	}

	content, err := readFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	switch outputType {
	case "text":
		err := content.WriteTextFile(output)
		if err != nil {
			log.Fatal(err)
		}
		break
	case "json":
		err := content.WriteJsonFile(output)
		if err != nil {
			log.Fatal(err)
		}
		break
	default:
		log.Fatal("File format not supported")
	}
}

func readFile(path string) (*models.Log, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	logData := models.NewLog()
	logData.Content = string(content)

	return logData, nil
}
