package sanitizer

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	htmlTagRE      = regexp.MustCompile(`(?i)<\s*(script|img|iframe|object|embed|form|input|textarea|select|style|link|meta|base|applet)[^>]*>`)
	eventHandlerRE = regexp.MustCompile(`(?i)\bon\w+\s*=`)
	jsURLRE        = regexp.MustCompile(`(?i)^\s*(javascript|data|vbscript)\s*:`)
	sqlInjectionRE = regexp.MustCompile(`(?i)(?:DROP\s+TABLE|UNION\s+SELECT|INSERT\s+INTO|DELETE\s+FROM|--\s|;\s*--|/\*.*\*/|'\s*OR\s+'1'\s*=\s*'1)`)
	htmlEntityRE   = regexp.MustCompile(`&#x?[0-9a-fA-F]+;`)
	urlEncodeRE    = regexp.MustCompile(`%[0-9a-fA-F]{2}`)
)

func Sanitize(input string) error {
	if strings.TrimSpace(input) == "" {
		return nil
	}

	decoded := decodeInput(input)

	if htmlTagRE.MatchString(decoded) {
		return fmt.Errorf("input contains disallowed HTML tags")
	}
	if eventHandlerRE.MatchString(decoded) {
		return fmt.Errorf("input contains disallowed event handlers")
	}
	if jsURLRE.MatchString(decoded) {
		return fmt.Errorf("input contains disallowed URL schemes")
	}
	if sqlInjectionRE.MatchString(decoded) {
		return fmt.Errorf("input contains disallowed SQL patterns")
	}

	return nil
}

func decodeInput(input string) string {
	result := input

	result = urlEncodeRE.ReplaceAllStringFunc(result, func(s string) string {
		var b byte
		fmt.Sscanf(s, "%%%02x", &b)
		return string(b)
	})

	result = htmlEntityRE.ReplaceAllStringFunc(result, func(s string) string {
		s = strings.TrimPrefix(s, "&#")
		s = strings.TrimSuffix(s, ";")
		var code rune
		if strings.HasPrefix(s, "x") || strings.HasPrefix(s, "X") {
			fmt.Sscanf(s[1:], "%x", &code)
		} else {
			fmt.Sscanf(s, "%d", &code)
		}
		return string(code)
	})

	return result
}
