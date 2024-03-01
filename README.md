# Infer

## Overview

Infer is a declarative domain-specific language (DSL) and an Inference-Driven-Development (IDD) test suite for any codebase. Infer allows developers to assert inferences about their code using HCL syntax. See below for examples.

**Note**: Infer is **experimental** pre-release software, and as such all interfaces are unstable and subject to change. Please **exercise caution!**

## Features

- **Inference Testing**: Use a LLM (large language model) to assess your code through inferences.
- **Inference-Driven Development (IDD)**: Inspired by behavior-driven development (BDD), write your tests in plain natural language. Unlike BDD, don't worry about implementation: all Infer tests are self-contained inference assertions, requiring no code to implement.
- **Language Agnostic**: Infer can be applied to any codebase, regardless of language.
- **HCL Syntax**: Define your tests with the expressive and well-structured HashiCorp Configuration Language.
- **OpenAI Integration**: Integrates with OpenAI API, allowing the use of any OpenAI-compatible model provider backend.
- **CLI Flexibility**: Run your tests and configure Infer directly from the command line.
- **Parallel Execution**: Optimize test execution time with built-in support for parallel processing.

## Example: Solving the Halting Problem

To demonstrate how Infer can be applied to analyze code behavior, consider a Python script that contains a non-terminating loop. We can tag this specific section of the code and use Infer to assert whether the loop halts.

### Tagging Code for Infer


In your Python script (`example.py`), you might have a loop like this, which you'd like to test:

```python
# Infer: Suspiciously loopy code
while True:
  print("This loop will run forever.")
# EndInfer: Suspiciously loopy code
```

Similarly, you would use the native comment format of whatever language you are working in. For example, in JavaScript, you would use the following:

```javascript
// Infer: Suspiciously loopy code
while (true) {
  console.log("This loop will run forever.");
}
// EndInfer: Suspiciously loopy code
```

In any programming language, we can use comments to tag code snippets with `Infer` and `EndInfer` to indicate that we want to assert inferences about the code inside the tagged section.

### Inferfile Syntax

In the `Inferfile`, using an [HCL](https://github.com/hashicorp/hcl) syntax, you can define inference assertions related to this code section:

```hcl
file "example.py" {
  tag "Suspiciously loopy code" {
    infer {
      assert = "it should eventually halt"
      model = "gpt4-turbo"
      count = 5                # Check assertion 5 times
      threshold = 0.8          # Require 80% for success
    }

    infer {
      assert     = "it should not introduce any security vulnerabilities"
      model      = "gpt4-turbo"
      count      = 1            # Default
      threshold  = 1.0          # Default
    }
  }
}
```

This HCL snippet in the Inferfile specifies that we are making an inference about the `Suspiciously loopy code` tagged section of the code. The assertion here infers whether the loop within the tagged section will halt, demonstrating how you might structure assertions in your Inferfile.

## More examples
See our very own `Inferfile` for more examples of inference tests on this codebase.

## Installation

Build the application from source:
```
go build ./cmd/infer/infer.go -o infer
```
## Usage

### Command-Line Interface

```plaintext
Usage of ./infer:
Commands:
  validate: Validate the syntax of the Inferfile.
  infer: Run inference tests (default).
Options:
  -f Path to the Inferfile (Default: Inferfile)
  -help Show help information (Default: false)
  -openai-api-key string
        OpenAI API key (Default: environment variable $OPENAI_API_KEY)
  -openai-api-url OpenAI API URL (Default: https://api.openai.com/v1)
  -parallel-threads Number of parallel threads to run (Default: 1)
  -v Enable verbose output (Default: false)
  ```

## Build Instructions

Ensure you have Go installed and follow these steps to build Infer from source:

```sh
git clone https://github.com/yourusername/infer.git
cd infer
go build -o infer cmd/infer/main.go
./infer --help
```

## Unit Testing

Run `go test ./...` to execute the unit tests.

## Contributing

Contributions are welcome! Please refer to our contribution guidelines for how to propose changes, submit issues, or add features.

## License

Infer is provided under the GNU General Public License (Version 3) by Inferret.io.

## Support

For support, bug reports, or feature requests, please file an issue through the GitHub issue tracker.
