package rule

import (
	"regexp"
	"strings"

	"github.com/mgechev/revive/lint"
)

// Define regex pattern for warning comments at package level
// Match TODO, FIXME, XXX at the beginning of a comment (case-insensitive)
var warningPattern = regexp.MustCompile(`(?i)^[\s*]*(TODO|FIXME|XXX)`)

// WarningCommentRule lints comments that start with warning markers like TODO, FIXME, XXX.
type WarningCommentRule struct{}

// Apply applies the rule to given file.
func (*WarningCommentRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure

	// Check all comments in the file
	for _, commentGroup := range file.AST.Comments {
		for _, comment := range commentGroup.List {
			commentText := comment.Text

			// Skip if it's a directive comment (like //go:generate)
			if isDirectiveComment(commentText) {
				continue
			}

			// For single line comments (// style)
			if strings.HasPrefix(commentText, "//") {
				commentText = strings.TrimPrefix(commentText, "//")
				if warningPattern.MatchString(commentText) {
					failures = append(failures, lint.Failure{
						Node:       comment,
						Confidence: 1,
						Category:   lint.FailureCategoryStyle,
						Failure:    "warning comment detected, consider resolving the issue",
					})
				}
				continue
			}

			// For multi-line comments (/* style */)
			if strings.HasPrefix(commentText, "/*") {
				commentText = strings.TrimPrefix(commentText, "/*")
				commentText = strings.TrimSuffix(commentText, "*/")
				
				// Check first line of the multi-line comment
				lines := strings.Split(commentText, "\n")
				if len(lines) > 0 && warningPattern.MatchString(lines[0]) {
					failures = append(failures, lint.Failure{
						Node:       comment,
						Confidence: 1,
						Category:   lint.FailureCategoryStyle,
						Failure:    "warning comment detected, consider resolving the issue",
					})
				}
			}
		}
	}

	return failures
}

// Name returns the rule name.
func (*WarningCommentRule) Name() string {
	return "warning-comment"
}
