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
		t := make([]byte, 0)
		for j := i; j < n; j++ {
			i = j
			t = append(t, text[j])
			if text[j] == ';' {
				break
			}
		}
		if v, ok := m[string(t)]; ok {
			result = append(result, v)
		} else {
			for _, v := range t {
				result = append(result, v)
			}
		}
	}

	return string(result)
}

func main() {
	text := "x &gt; y &amp;&amp; x &lt; y is always false"
	fmt.Println("Output: ", entityParser(text))
}
