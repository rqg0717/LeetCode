package main

import "fmt"

func entityParser(text string) string {
	n := len(text)
	m := make(map[string]byte)
	m["&quot;"] = '"'
	m["&apos;"] = '\''
	m["&amp;"] = '&'
	m["&gt;"] = '>'
	m["&lt;"] = '<'
	m["&frasl;"] = '/'

	result := make([]byte, 0)

	for i := 0; i < n; i++ {
		if text[i] != '&' {
			result = append(result, text[i])
			continue
		}
		special := make([]byte, 0)
		for j := i; j < n; j++ {
			i = j
			special = append(special, text[j])
			if text[j] == ';' {
				break
			}
		}
		if v, ok := m[string(special)]; ok {
			result = append(result, v)
		} else {
			for _, v := range special {
				result = append(result, v)
			}
		}
	}

	return string(result)
}

func main() {
	text := "&amp; is an HTML entity but &ambassador; is not."
	fmt.Println("Output: ", entityParser(text))
}
