package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var isAlphaFunc = regexp.MustCompile(`^[A-Za-z]+$`).MatchString

// testValidity returns true if given string match with a sequence of numbers followed by dash followed by text, eg: 23-ab-48-caba-56-haha
func testValidity(str string) bool {
	arrStr := strings.Split(str, "-")
	if len(arrStr)%2 != 0 {
		return false
	}

	for i, v := range arrStr {
		if (i+1)%2 == 0 {
			if !isAlphaFunc(v) {
				return false
			}
		} else {
			num, err := strconv.Atoi(v)
			if err != nil || num < 0 {
				return false
			}
		}
	}

	return true
}

// averageNumber returns the average number from all the numbers by given string
func averageNumber(str string) (int, error) {
	if !testValidity(str) {
		return 0, fmt.Errorf("value is not valid %s", str)
	}

	var sum, total int
	arrStr := strings.Split(str, "-")
	for i, v := range arrStr {
		if (i+1)%2 != 0 {
			n, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}
			sum = sum + n
			total = total + 1
		}
	}

	return (sum / total), nil
}

// wholeStory returns a text that is composed from all the text words separated by spaces
func wholeStory(str string) (string, error) {
	if !testValidity(str) {
		return "", fmt.Errorf("value is not valid %s", str)
	}

	var sl []string
	arrStr := strings.Split(str, "-")
	for i, v := range arrStr {
		if (i+1)%2 == 0 {
			sl = append(sl, v)
		}
	}
	return strings.Join(sl, " "), nil
}

// storyStats returns shortest word, longest word, the average word length and list of all words from the story that have the length the same as the average length
func storyStats(str string) (shortest string, longest string, average int, list []string, err error) {
	if !testValidity(str) {
		return shortest, longest, average, list, fmt.Errorf("value is not valid %s", str)
	}

	var (
		sumw   int
		totalw int
		wlist  []string
	)

	arrStr := strings.Split(str, "-")
	for i, v := range arrStr {
		if (i+1)%2 == 0 {
			// longest validation
			if len(longest) < len(v) {
				longest = v
			}

			// shortest validation
			if i == 1 {
				shortest = v
			}
			if len(shortest) > len(v) {
				shortest = v
			}

			sumw = sumw + len(v)
			totalw = totalw + 1

			wlist = append(wlist, v)
		}
	}

	average = sumw / totalw

	for _, v := range wlist {
		if len(v) == average {
			list = append(list, v)
		}
	}

	return shortest, longest, average, list, err
}

// generate generates random correct strings if the flag valid parameter is true
func generate() string {
	valid := flag.Bool("valid", false, "return valid string when value is true")
	flag.Parse()

	if *valid {
		matchPattern := "%d-%s-%d-%s"
		return fmt.Sprintf(matchPattern, randInt(100), randStringRunes(5), randInt(100), randStringRunes(5))
	}

	return fmt.Sprintf("%s-%d-%s-%d", randStringRunes(4), randInt(4), randStringRunes(4), randInt(4))
}

func randInt(n int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	return rand.New(s1).Intn(n)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	str := generate()
	fmt.Println(str)
}
