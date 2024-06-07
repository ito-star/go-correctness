package main

import (
	"regexp"
	"strings"
)

func yourFunc(input string) string {
	// Replace colon followed by incorrect
	input = strings.Replace(input, ":", ";", -1)
	// Check if input contains ";"
	if strings.Contains(input, ";") {
		// Split the string at ";"
		parts := strings.SplitN(input, ";", 2)
		prefix := strings.TrimSpace(parts[0])
		suffix := strings.TrimSpace(parts[1])

		// Use a regular expression to find the last number in the suffix
		re := regexp.MustCompile(`(\d+)\s*$`)
		matches := re.FindStringSubmatch(suffix)

		if len(matches) > 1 {
			// If a match is found, use it as the new suffix with "DOM {number}"
			suffix = "DOM " + matches[1]
		} else {
			// If no match is found, assume the suffix should start with "DOM 0"
			suffix = "DOM 0"
		}

		// Reassemble the corrected string
		return prefix + ";" + suffix
	} else {
		// If there is no ";", find "DOM" and ensure it's properly formatted
		domIndex := strings.Index(input, "DOM")
		if domIndex != -1 {
			// Extract the number following "DOM" if it exists
			substr := input[domIndex:]

			re := regexp.MustCompile(`DOM\s+(\d+)`)
			matches := re.FindStringSubmatch(substr)
			if len(matches) > 1 {
				// If a number is found, format it correctly
				return strings.TrimSpace(input[:domIndex]) + ";" + "DOM " + matches[1]
			} else {
				// If no number is found, default to "DOM 0"
				return strings.TrimSpace(input[:domIndex]) + ";" + "DOM 0"
			}
		}
	}

	// Return the input if no corrections are needed
	return input
}
