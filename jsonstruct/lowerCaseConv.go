package jsonstruct

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"unicode"
	"unicode/utf8"
)

// Regexp definitions
var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)

type conventionalMarshallerSnakeCase struct {
	Value interface{}
}

func (c conventionalMarshallerSnakeCase) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)

	fmt.Println("marshalled", string(marshalled))

	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			fmt.Println("match", string(match))
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)

	return converted, err
}

type conventionalMarshallerCamelCase struct {
	Value interface{}
}

func (c conventionalMarshallerCamelCase) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)

	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			// Empty keys are valid JSON, only lowercase if we do not have an
			// empty key.
			fmt.Println("match", match[1:])
			if len(match) > 2 {
				// Decode first rune after the double quotes
				r, width := utf8.DecodeRune(match[1:])
				fmt.Println("r", r)
				fmt.Println("width", width)
				r = unicode.ToLower(r)
				utf8.EncodeRune(match[1:width+1], r)
			}
			return match
		},
	)

	return converted, err
}

type ProductTokopedia struct {
	id     string
	name   string
	images []ProductTokopediaImage
}

type ProductTokopediaImage struct {
	id       string
	name     string
	fileName string
}

func New3() {
	productTokopedia := ProductTokopedia{
		id:   "id-001",
		name: "botol",
		images: []ProductTokopediaImage{
			{
				id:       "id-001-001",
				name:     "botol img 1",
				fileName: "botolimg.png",
			},
			{
				id:       "id-001-002",
				name:     "botol img 2",
				fileName: "botolimg2.png",
			},
		},
	}

	encoded, _ := json.MarshalIndent(conventionalMarshallerSnakeCase{productTokopedia}, "", "  ")

	fmt.Println(string(encoded))
}
