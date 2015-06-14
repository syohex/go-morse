package morse

import (
	"bytes"
	"log"
	"strings"
)

const (
	// EN means English
	EN = iota
)

var english = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	".": ".-.-.-",
	",": "--..--",
	"/": "-...-",
	":": "---...",
	"'": ".----.",
	"-": "-....-",
	"?": "..--..",
	"!": "..--.",
	"@": "...-.-",
	"+": ".-.-.",
	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
}

var japanese = map[string]string{
	"ア": "--.--",
	"イ": ".-",
	"ウ": "..-",
	"エ": "-.---",
	"オ": ".-...",

	"カ": ".-..",
	"キ": "-.-..",
	"ク": "...-",
	"ケ": "-.--",
	"コ": "----",

	"サ": "-.-.-",
	"シ": "--.-.",
	"ス": "---.-",
	"セ": ".---.",
	"ソ": "---.",

	"タ": "-.",
	"チ": "..-.",
	"ツ": ".--.",
	"テ": ".-.--",
	"ト": "..-..",

	"ナ": ".-.",
	"ニ": "-.-.",
	"ヌ": "....",
	"ネ": "--.-",
	"ノ": "..--",

	"ハ": "-...",
	"ヒ": "--..-",
	"フ": "--..",
	"ヘ": ".",
	"ホ": "-..",

	"マ": "-..-",
	"ミ": "..-.-",
	"ム": "-",
	"メ": "-...-",
	"モ": "-..-.",

	"ヤ": ".--",
	"ユ": "-..--",
	"ヨ": "--",

	"ラ": "...",
	"リ": "--.",
	"ル": "-.--.",
	"レ": "---",
	"ロ": ".-.-",

	"ワ": "-.-",
	"ヲ": ".---",
	"ン": ".-.-.",

	"゛": "..",
	"゜": "..--.",
	"̄": ".--.-",
	"、": ".-.-.-",
	"。": ".-.-..",
}

func findMap(lang int) map[string]string {
	switch lang {
	case EN:
		return english
	}

	log.Fatalln("Invalid Language\n")
	return nil
}

// Encode coverts from string to morse code
func Encode(input string, lang int) string {
	table := findMap(lang)

	buf := bytes.NewBufferString("")
	for _, c := range []rune(input) {
		code, ok := table[string(c)]
		if !ok {
			continue
		}

		buf.WriteString(code)
		buf.WriteString(" ")
	}

	return strings.TrimSpace(buf.String())
}

func reverseDictionary(dict map[string]string) map[string]string {
	rev := make(map[string]string)
	for k, v := range dict {
		rev[v] = k
	}

	return rev
}

// Decode decodes from morse code to original statement
func Decode(input string, lang int) string {
	table := reverseDictionary(findMap(lang))

	buf := bytes.NewBufferString("")
	for _, c := range strings.Split(input, " ") {
		v, ok := table[c]
		if !ok {
			continue
		}

		buf.WriteString(v)
	}

	return strings.TrimSpace(buf.String())
}
