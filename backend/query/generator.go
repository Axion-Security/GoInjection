package query

import "fmt"

type Query struct {
	DBMS string
}

// Build Returns: ([]string: Payloads, string: Status)
func (q *Query) Build(payload string) ([]string, string) {
	openings, okOpening := Opening[q.DBMS]
	modifiers, okModifier := Modifier[q.DBMS]

	if !okOpening || !okModifier {
		openings, okOpening = Opening["Default"]
		modifiers, okModifier = Modifier["Default"]
		if !okOpening || !okModifier {
			return []string{}, "Unsupported DBMS"
		}
	}

	var payloads []string
	for _, opening := range openings {
		for _, modifier := range modifiers {
			payloads = append(payloads, fmt.Sprintf("%s %s %s", opening, payload, modifier))
		}
	}

	return payloads, "Success"
}
