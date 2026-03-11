package charset

import (
	"bufio"
	"io"
	"strings"
)

func SearchString(charset string, text string, w io.Writer) error {
	return Search(charset, strings.NewReader(text), w)
}

func Search(charset string, r io.Reader, w io.Writer) error {
	if charset == "" {
		_, err := io.ReadAll(r)
		return err
	}

	scanner := bufio.NewScanner(r)
	scanner.Buffer(nil, 1024*1024*1024)

	required := make(map[rune]struct{})
	for _, r := range charset {
		required[r] = struct{}{}
	}

	targetCount := len(required)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < len(charset) {
			continue
		}

		foundInLine := make(map[rune]struct{}, targetCount)

		for _, char := range line {
			if _, ok := required[char]; ok {
				foundInLine[char] = struct{}{}
			}

			if len(foundInLine) == targetCount {
				break
			}
		}

		if len(foundInLine) == targetCount {
			_, err := w.Write([]byte(line + "\n"))
			if err != nil {
				return err
			}
		}
	}

	return scanner.Err()
}
