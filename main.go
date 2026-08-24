package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func svgBox(lines []string) string {
	h := len(lines) + 2
	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="560" height="%d" viewBox="0 0 560 %d">`, 40+h*24, 40+h*24)
	b.WriteString(`<rect width="100%" height="100%" fill="#0d1117"/>`)
	fmt.Fprintf(&b, `<rect x="16" y="16" width="528" height="%d" rx="8" fill="#161b22" stroke="#30363d"/>`, 24+h*24)
	for i, ln := range lines {
		fmt.Fprintf(&b, `<text x="36" y="%d" font-family="monospace" font-size="18" fill="#7ee787">%s</text>`, 56+i*24, esc(ln))
	}
	b.WriteString(`</svg>`)
	return b.String()
}

func esc(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;")
	return r.Replace(s)
}

func main() {
	outFlag := flag.String("out", "profile.svg", "output svg file")
	spec := flag.String("s", "> hi\n> welcome to my profile\ntype help for commands", "multi-line terminal content")
	flag.Parse()

	lines := strings.Split(*spec, "\n")
	svg := svgBox(lines)
	if err := os.WriteFile(*outFlag, []byte(svg), 0644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("wrote %s (%d lines)\n", *outFlag, len(lines))
}