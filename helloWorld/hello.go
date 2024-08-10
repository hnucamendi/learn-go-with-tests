package main

import "fmt"

type Message struct {
	Language Lang
	Prefix   MessagePrefix
}

type Lang string
type MessagePrefix string

const (
	SpanishLang    Lang          = "SPANISH"
	EnglishLang    Lang          = "ENGLISH"
	FrenchLang     Lang          = "FRENCH"
	JapaneseLang   Lang          = "JAPANESE"
	OtherLang      Lang          = "OTHER"
	englishPrefix  MessagePrefix = "Hello"
	spanishPrefix  MessagePrefix = "Hola"
	japanesePrefix MessagePrefix = "こんにちは"
	frenchPrefix   MessagePrefix = "Bonjour"
)

func buildString(p MessagePrefix, s string) string {
	return fmt.Sprintf("%s %s!", p, s)
}

func English() *Message {
	return &Message{
		Language: EnglishLang,
		Prefix:   englishPrefix,
	}
}
func Spanish() *Message {
	return &Message{
		Language: SpanishLang,
		Prefix:   spanishPrefix,
	}
}
func French() *Message {
	return &Message{
		Language: FrenchLang,
		Prefix:   frenchPrefix,
	}
}
func Japanese() *Message {
	return &Message{
		Language: JapaneseLang,
		Prefix:   japanesePrefix,
	}
}

func Config(l Lang) *Message {
	switch l {
	case SpanishLang:
		return Spanish()
	case FrenchLang:
		return French()
	case JapaneseLang:
		return Japanese()
	default:
		return English()
	}
}

func (m *Message) Hello(s string) string {
	if s == "" {
		s = "World"
	}

	return buildString(m.Prefix, s)
}

func main() {
	english := Config(EnglishLang)
	spanish := Config(SpanishLang)
	french := Config(FrenchLang)
	japanese := Config(JapaneseLang)

	fmt.Println(english.Hello("David"))  // Hello David!
	fmt.Println(spanish.Hello("David"))  // Hola David!
	fmt.Println(french.Hello("David"))   // Bonjour David!
	fmt.Println(japanese.Hello("David")) // こんにちは David!
}
