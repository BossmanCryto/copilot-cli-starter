package domains

import "fmt"

// ExampleWorkflow demonstrates a domain-specific shortcut
type ExampleWorkflow struct{}

// Execute runs the example workflow
func (w *ExampleWorkflow) Execute() error {
	fmt.Println("Example workflow executed")
	return nil
}
