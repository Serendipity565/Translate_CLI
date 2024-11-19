package api

type Translator interface {
	Translate(text, sourceLang, targetLang string) (string, error)
}
