file "/tmp/mycode.go" {
  tag "OpenAI client" {
    infer { 
      assert    = "Does not contain any hard-coded credentials." 
      model     = "gpt-3.5-turbo"
      count     = 5                # Check assertion 5 times
      threshold = 0.8              # Require 80% for success
    }
  }
}
