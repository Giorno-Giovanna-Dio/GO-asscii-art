package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mbndr/figlet4go"
)

func main() {
	// CLI flags
	text := flag.String("text", "Go ASCII!", "Text to convert into ASCII art")
	font := flag.String("font", "standard", "Font style (e.g. standard, slant, big, shadow, banner3-D)")
	flag.Parse()

	renderer := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontName = *font

	result, err := renderer.RenderOpts(*text, options)
	if err != nil {
		log.Fatalf("Error rendering text: %v", err)
	}

	fmt.Println(result)
}
