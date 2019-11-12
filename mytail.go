package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func fileToTextLines(fileName string) []string {
	var textLines []string
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textLines = append(textLines, scanner.Text())
	}

	return textLines
}

func pickTextLinesFromEnd(textLines []string, n int) []string {
	top := len(textLines) - n
	if top < 0 {
		top = 0
	}
	return textLines[top:]
}

func pickTextLinesFromTop(textLines []string, n int) []string {
	end := n
	if end > len(textLines) {
		end = len(textLines)
	}
	return textLines[:end]
}

func pickTextLinesFromRandom(textLines []string, n int) []string {
	end := n
	for i := len(textLines) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		textLines[i], textLines[j] = textLines[j], textLines[i]
	}

	if end > len(textLines) {
		end = len(textLines)
	}

	return textLines[:end]
}

func getTextLinesWithOptionO(o string, textLines []string, n int) []string {
	switch o {
	case "b":
		textLines = pickTextLinesFromEnd(textLines, n)
	case "f":
		textLines = pickTextLinesFromTop(textLines, n)
	case "r":
		textLines = pickTextLinesFromRandom(textLines, n)
	}
	return textLines
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [OPTIONS] ARGS...
Options\n`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	var (
		n = flag.Int("n", 10, "How many lines to read from the end")
		o = flag.String("o", "b", "Read in reverse option, b: back(default), f: front, r: random shuffle")
		okazu = flag.Bool("okazu", false, "I'm hungry! Yes, let's decide the side dish!")
	)
	flag.Parse()
	args := flag.Args()

	if *okazu == true {
		// おかずランキング
		urls, menus, imageUrls :=  ebaraFoodScraping()
		urls = getTextLinesWithOptionO(*o, urls, *n)
		menus = getTextLinesWithOptionO(*o, menus, *n)
		imageUrls = getTextLinesWithOptionO(*o, imageUrls, *n)

		for j := range menus {
			fmt.Println(menus[j] + ": " + urls[j])
			fmt.Println(imageUrls[j])
		}
	} else {
		// tailコマンド
		for i := range args {
			fmt.Fprintf(os.Stdout, "==> %s <==\n", args[i])
			textLines := fileToTextLines(args[i])
			textLines = getTextLinesWithOptionO(*o, textLines, *n)

			for j := range textLines {
				fmt.Println(textLines[j])
			}
			fmt.Println()
		}
	}
}
