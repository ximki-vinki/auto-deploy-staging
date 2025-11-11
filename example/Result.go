package example

import (
	"fmt"
	"os"
)

type Result struct {
	ProjectID  string
	PipelineID int
	URL        string
	Error      error
}

func printResults(results []Result) (successCount, errorCount int) {
	for _, r := range results {
		if r.Error != nil {
			fmt.Printf("❌ [%s] Error: %v\n", r.ProjectID, r.Error)
			errorCount++
		} else {
			fmt.Printf("✅ [%s] Pipeline %d: %s\n", r.ProjectID, r.PipelineID, r.URL)
			successCount++
		}
	}
	fmt.Printf("\nSuccess: %d, Errors: %d\n", successCount, errorCount)
	return successCount, errorCount
}

func exitIfErrors(successCount, errorCount int) {
	if errorCount > 0 {
		os.Exit(1)
	}
}
