file "./internal/api/openai.go" {
  tag "OpenAI client" {
    infer { 
      assert    = "This code does connect to OpenAI." 
      model     = "gpt-3.5-turbo"
      count     = 1
      threshold = 1.0
    }

    infer { 
      assert    = "Does this code configure the OpenAI API URL?"
      model     = "gpt-3.5-turbo"
      count     = 1
      threshold = 1.0
    }

    infer {
      assert    = "This code does not contain any hard-coded API credentials."
      model     = "gpt-3.5-turbo"
      count     = 1
      threshold = 1.0
    }

    infer {
      assert    = "Does this code limit the API request max tokens as specified by inference.max_tokens?"
      model     = "gpt-3.5-turbo"
      count     = 1
      threshold = 1.0
    }
  }
}

file "./cmd/infer/infer.go" {
  tag "command line arguments" {
    infer {
      assert    = "This code accepts an argument for parallel threads." 
      model     = "gpt-3.5-turbo"
      count     = 1
      threshold = 1.0
    }

    infer {
      assert    = "Does this code accept an argument for an OpenAI API URL?"
      model     = "gpt-3.5-turbo"
      count     = 1
      threshold = 1.0
    }
  }

  tag "execution" {
    infer {
      assert    = "This code runs parallel threads."
      model     = "gpt-3.5-turbo"
      count     = 5
      threshold = 1.0
    }
  }
}

file "./internal/executor/executor.go" {
  tag "Execute" {
    infer {

      assert    = "Does this code run each inference multiple times, as specified by inference.count?"
      model     = "gpt-3.5-turbo"
      count     = 5
      threshold = 1.0
    }
  }
}
