package utils

import "regexp"

var expandVarRegex = regexp.MustCompile(`(\\)?(\$)(\()?\{?([A-Za-z0-9_]+)?\}?`)

func ExpandVariables(v string, m map[string]string) string {
	return expandVarRegex.ReplaceAllStringFunc(v, func(s string) string {
		submatch := expandVarRegex.FindStringSubmatch(s)
		if submatch == nil {
			return s
		}
		if submatch[1] == "\\" || submatch[2] == "(" {
			return submatch[0][1:]
		} else if submatch[4] != "" {
			return m[submatch[4]]
		}
		return s
	})
}
