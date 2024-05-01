package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// buildBalloon takes a slice of strings of max width maxWidth
// prepends/appends margins on first and last line and at start/end of each line
//  and returns a string with the contents of the balloon

func buildBalloon(lines []string, maxWidth int) string {
	var borders []string
	count := len(lines)
	var result []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxWidth+2)
	bottom := " " + strings.Repeat("-", maxWidth+2)

	result = append(result, top)

	if count == 1 {
		s := fmt.Sprintf(`%s %s %s`, borders[5], lines[0], borders[6])
		result = append(result, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])
		result = append(result, s)

		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], lines[i], borders[4])
			result = append(result, s)
		}

		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[3])
		result = append(result, s)
	}
	result = append(result, bottom)
	return strings.Join(result, "\n")

}

// tabsToSpaces converts all tabs found in the strings
// found in the `lines` slice to 4 spaces, to prevent misalignments in
// counting the runes
func tabsToSpaces(lines []string) []string {
	var result []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		result = append(result, l)
	}

	return result
}

// calculateMaxWidth given a slice of string returns the length of the
// string with max width
func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}

	return w
}

// normalizeStringLength takes a slice of strings and appends to each one a number
// of spaces needed to have them all the same number of runes
func normalizeStringLength(lines []string, maxWidth int) []string {
	var result []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxWidth-utf8.RuneCountInString(l))
		result = append(result, s)
	}

	return result
}

func printFigure(name string) {
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	var stegosaurus = `         	\                      .       .
		\                    / ` + "`" + `.   .' "
		 \           .---.  <    > <    >  .---.
		  \          |    \  \ - ~ ~ - /  /    |
		_____           ..-~             ~-..-~
	   |     |   \~~~\\.'                    ` + "`" + `./~~~/
	  ---------   \__/                         \__/
	 .'  O    \     /               /       \  "
	(_____,    ` + "`" + `._.'               |         }  \/~~~/
	 ` + "`" + `----.          /       }     |        /    \__/
		   ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
			   ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
					|     /        |     /     ~-.     ` + "`" + `-. _  _  _
					|_____|        |_____|         ~ - . _ _ _ _ _>

  `

	switch name {
	case "cow":
		fmt.Println(cow)
	case "stegosaurus":
		fmt.Println(stegosaurus)
	default:
		fmt.Println("Unknown Figure")
	}
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	var lines []string
	var figure string

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		lines = append(lines, string(line))
	}

	flag.StringVar(&figure, "f", "cow", "the figure name. Valid values are `cow` and `stegosaurus`")
	flag.Parse()

	lines = tabsToSpaces(lines)
	maxWidth := calculateMaxWidth(lines)
	message := normalizeStringLength(lines, maxWidth)
	ballon := buildBalloon(message, maxWidth)

	fmt.Println(ballon)
	printFigure(figure)
	fmt.Println()

}
