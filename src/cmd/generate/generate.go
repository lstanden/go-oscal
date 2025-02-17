package generate

import (
	"fmt"
	"log"
	"os"

	"github.com/defenseunicorns/go-oscal/src/internal/generate"
	"github.com/defenseunicorns/go-oscal/src/pkg/utils"

	"github.com/spf13/cobra"
)

var opts = &generate.BaseFlags{}

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate Golang data types from OSCAL schema",
	Long:  "Generate Golang data types from OSCAL Complete schema.",
	// Example: generateHelp,
	RunE: func(cmd *cobra.Command, componentDefinitionPaths []string) error {
		output, err := GenerateCommand(*opts)
		if err != nil {
			return err
		}
		// Write the Go struct output to either stdout or a file.
		if opts.OutputFile == "" {
			log.Println(string(output))
		} else {
			err = utils.WriteOutput(output, opts.OutputFile)
			if err != nil {
				return fmt.Errorf("failed to write output to file: %s", err)
			}
		}
		return nil
	},
}

func GenerateCommand(flags generate.BaseFlags) (output []byte, err error) {
	schemaBytes, err := os.ReadFile(flags.InputFile)
	if err != nil {
		return
	}

	tagList, err := utils.FormatTags(flags.Tags)
	if err != nil {
		return
	}

	// Generate the Go structs.
	output, err = generate.Generate(schemaBytes, flags.Pkg, tagList)
	if err != nil {
		return output, fmt.Errorf("failed to generate Go structs: %s", err)
	}

	return
}

func init() {
	GenerateCmd.Flags().StringVarP(&opts.InputFile, "input-file", "f", "", "the path to a oscal json schema file")
	GenerateCmd.Flags().StringVarP(&opts.OutputFile, "output-file", "o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
	GenerateCmd.Flags().StringVarP(&opts.Pkg, "pkg", "p", "main", "the name of the package for the generated code")
	GenerateCmd.Flags().StringVarP(&opts.Tags, "tags", "t", "json,yaml", "comma separated list of the tags to put on the struct")
}
