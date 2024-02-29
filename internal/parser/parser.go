package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type InferConfiguration struct {
	Files []File `hcl:"file,block"`
}

type File struct {
	Path string `hcl:"path,label"`
	Tags []Tag  `hcl:"tag,block"`
}

type Tag struct {
	Name       string      `hcl:"name,label"`
	Inferences []Inference `hcl:"infer,block"`
	Code       string      `hcl:"code,optional"` // Exclude from HCL parsing, but use for internal processing
}

type Inference struct {
	Assertion   string  `hcl:"assert"`
	Model       string  `hcl:"model"`
	Count       int     `hcl:"count,optional"`
	Threshold   float64 `hcl:"threshold,optional"`
	MaxTokens   int     `hcl:"max_tokens,optional"`
	Temperature float64 `hcl:"temperature,optional"`
	Tag_Name    string
}

func ParseInferfile(filename string) (*InferConfiguration, error) {
	parser := hclparse.NewParser()
	hclFile, diags := parser.ParseHCLFile(filename)
	if diags.HasErrors() {
		return nil, diags
	}

	var config InferConfiguration
	diags = gohcl.DecodeBody(hclFile.Body, nil, &config)
	if diags.HasErrors() {
		return nil, diags
	}

	// Check if the files exist and note the line number of the Inferfile that pointed to that file.
	syntaxBody, ok := hclFile.Body.(*hclsyntax.Body)
	if !ok {
		return nil, fmt.Errorf("file body is not hclsyntax.Body")
	}
	for _, file := range config.Files {
		if _, err := os.Stat(file.Path); os.IsNotExist(err) {
			for _, block := range syntaxBody.Blocks {
				if block.Type == "file" && len(block.Labels) > 0 && block.Labels[0] == file.Path {
					return nil, fmt.Errorf("file not found: %s (Inferfile line: %d)", file.Path, block.DefRange().Start.Line)
				}
			}
		}
	}

	// After checking for file existence, read and attach the code for each tag.
	for i, file := range config.Files {
		err := attachCodeToTags(&file)
		if err != nil {
			return nil, err
		}
		config.Files[i] = file
	}

	return &config, nil
}

func attachCodeToTags(file *File) error {
	// Read the entire source file content
	source, err := ioutil.ReadFile(file.Path)
	if err != nil {
		return err
	}

	// Convert the file content into a slice of lines
	lines := strings.Split(string(source), "\n")

	// Iterate over each tag in the file
	for i := range file.Tags {
		tag := &file.Tags[i] // Get a reference to the tag to modify it directly
		var tagBuilder strings.Builder
		var inTagBlock bool

		// Iterate over each line in the source file
		for _, line := range lines {
			if strings.Contains(line, "Infer: "+tag.Name) {
				inTagBlock = true // Start of tag block
				continue
			}
			if strings.Contains(line, "EndInfer: "+tag.Name) {
				inTagBlock = false // End of tag block
				break
			}
			if inTagBlock {
				tagBuilder.WriteString(line + "\n") // Collect the lines within the tag block
			}
		}

		// Update the tag's Code field with the collected code block
		tag.Code = tagBuilder.String()
	}

	return nil
}
