package generator

import (
	"strings"
	"text/template"
	"unicode"
)

var codeTemplate = template.Must(template.New("code").Funcs(template.FuncMap{
	"capitalize": capitalize,
}).Parse(`// Code generated by generator; DO NOT EDIT.
package {{.PackageName}}

{{/* New 関数の生成 */}}
func New{{.Name}}({{range .Fields}}{{if .Name}}{{.Name}} {{.Type}}, {{end}}{{end}}) *{{.Name}} {
    return &{{.Name}}{
        {{range .Fields}}{{if .Name}}{{.Name}}: {{.Name}},
        {{end}}{{end}}}
}

{{/* Getter メソッドの生成 */}}
{{range .Fields}}{{if .Name}}
func (s *{{$.Name}}) {{capitalize .Name}}() {{.Type}} {
    return s.{{.Name}}
}
{{end}}{{end}}
`))

var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}

	words := splitIntoWords(s)

	for i, word := range words {
		upperWord := strings.ToUpper(word)
		if commonInitialisms[upperWord] {
			words[i] = upperWord
		} else {
			words[i] = toTitle(word)
		}
	}

	return strings.Join(words, "")
}

func uncapitalize(s string) string {
	if s == "" {
		return ""
	}

	words := splitIntoWords(s)

	for i, word := range words {
		upperWord := strings.ToUpper(word)
		if commonInitialisms[upperWord] {
			if i == 0 {
				words[i] = strings.ToLower(upperWord)
			} else {
				words[i] = upperWord
			}
		} else {
			if i == 0 {
				words[i] = strings.ToLower(word)
			} else {
				words[i] = toTitle(word)
			}
		}
	}

	return strings.Join(words, "")
}

func splitIntoWords(s string) []string {
	var words []string
	var word []rune
	for i, r := range s {
		if i > 0 && (unicode.IsUpper(r) || unicode.IsDigit(r)) && (unicode.IsLower(rune(s[i-1])) || unicode.IsDigit(rune(s[i-1]))) {
			words = append(words, string(word))
			word = []rune{r}
		} else {
			word = append(word, r)
		}
	}
	if len(word) > 0 {
		words = append(words, string(word))
	}
	return words
}

func toTitle(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToUpper(rune(s[0]))) + strings.ToLower(s[1:])
}
