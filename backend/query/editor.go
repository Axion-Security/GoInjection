package query

import (
	"fmt"
	"strings"
)

type Editor struct {
	Type      string
	TypeValue string
	Repeat    int
}

// Edit Returns: (string: Payload)
func (e *Editor) Edit(payload string) string {
	var tempPayload strings.Builder
	for i := 0; i < e.Repeat; i++ {
		if i > 0 {
			tempPayload.WriteString(",")
		}
		switch strings.ToLower(e.Type) {
		case "null":
			tempPayload.WriteString("NULL")
		case "string":
			tempPayload.WriteString(fmt.Sprintf("'%s'", e.TypeValue))
		case "concat":
			tempPayload.WriteString(fmt.Sprintf("CONCAT(%s)", e.TypeValue))
		case "custom":
			tempPayload.WriteString(e.TypeValue)
		}
	}

	return strings.ReplaceAll(payload, "§", tempPayload.String())
}
