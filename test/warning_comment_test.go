package rule_test

import (
	"go/ast"
	"testing"

	"github.com/mgechev/revive/lint"
	"github.com/mgechev/revive/rule"
)

// TestWarningCommentRule tests the warning-comment rule.
func TestWarningCommentRule(t *testing.T) {
	// Initialize the rule
	r := &rule.WarningCommentRule{}

	// Test that the rule has the expected name
	if got, want := r.Name(), "warning-comment"; got != want {
		t.Errorf("r.Name() = %q, want %q", got, want)
	}

	// Setup a basic test file with various comments
	testFile := &lint.File{
		Name: "test.go",
		AST: makeASTWithComments([]string{
			"// Regular comment",
			"//TODO: Fix this later",
			"// todo: Another thing to fix",
			"// FIXME: This is a bug",
			"/* TODO: Multiline\n\t\t\t   comment with a warning */",
			"/* Regular multiline\n\t\t\t   comment */",
			"//go:generate command",
			"// XXX: Deprecated code",
		}),
	}

	// Apply the rule to the test file
	failures := r.Apply(testFile, nil)

	// We should have 5 failures (TODO, todo, FIXME, TODO in multiline, XXX)
	wantFailures := 5
	if got := len(failures); got != wantFailures {
		t.Errorf("Expected %d failures, got %d", wantFailures, got)
	}

	// Check that each failure has the expected message
	expectedMsg := "warning comment detected, consider resolving the issue"
	for i, failure := range failures {
		if failure.Failure != expectedMsg {
			t.Errorf("failure[%d].Failure = %q, want %q", i, failure.Failure, expectedMsg)
		}
	}
}

// Helper function to create an AST with comments
func makeASTWithComments(commentTexts []string) *ast.File {
	// This is a simplified version for the test
	// In real usage, you'd parse actual Go code
	file := &ast.File{}
	
	var comments []*ast.CommentGroup
	
	for _, text := range commentTexts {
		commentGroup := &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Text: text,
				},
			},
		}
		comments = append(comments, commentGroup)
	}
	
	file.Comments = comments
	return file
}