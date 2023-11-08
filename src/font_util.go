package main

import (
	"bufio"
	"strings"

	"golang.org/x/image/font"
)

func partitionText(face font.Face, text string, maxWidth int) []string {
	s := bufio.NewScanner(strings.NewReader(text))
	var lines []string
	for s.Scan() {
		currentLine := ""
		newLineWidth := 0
		words := strings.Split(s.Text(), " ")
		for _, word := range words {
			wordWidth := font.MeasureString(face, word+" ").Ceil()

			if newLineWidth+wordWidth < maxWidth {
				currentLine += word + " "
				newLineWidth += wordWidth
			} else {
				lines = append(lines, currentLine[:len(currentLine)-1])
				currentLine = word
				newLineWidth = wordWidth
			}
		}
	}
	return lines
}

func calculateMaxLines(face font.Face, maxHeight int) int {
	metrics := face.Metrics()
	return maxHeight / metrics.Height.Ceil()
}

func GroupText(face font.Face, text string, maxWidth int, maxHeight int) []string {
	lines := partitionText(face, text, maxWidth)
	result := []string{}
	maxLines := calculateMaxLines(face, maxHeight)
	for i := 0; i < len(lines); i += maxLines {
		end := i + maxLines
		if end > len(lines) {
			end = len(lines)
		}
		result = append(result, strings.Join(lines[i:end], " "))
	}
	return result
}
