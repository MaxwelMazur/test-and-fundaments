package hello

func hello(name, language string) string {

	if len(name) < 1 {
		name = "world"
		language = ""
	}

	return greetingPrefix(language) + ", " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "portuguese":
		return "Olá"
	case "german":
		return "Hallo"
	case "spanish":
		return "Hola"
	default:
		return "Hello"
	}
}
