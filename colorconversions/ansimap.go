package asciiart

func AnsiColorMap() map[string]string {
	return map[string]string{
		"cyan":          "\033[36m",
		"grey":          "\033[37m",
		"yellow":        "\033[33m",
		"black":         "\x1b[30m",
		"red":           "\x1b[31m",
		"green":         "\x1b[32m",
		"blue":          "\x1b[34m",
		"magenta":       "\x1b[35m",
		"white":         "\x1b[37m",
		"pink":          "\x1b[38;5;205m", // ANSI code for pink
		"brown":         "\x1b[38;5;94m",  // ANSI code for brown
		"orange":        "\x1b[38;5;214m", // ANSI code for orange
		"lime":          "\x1b[38;5;10m",
		"brightBlack":   "\x1b[90m",
		"brightRed":     "\x1b[91m",
		"brightGreen":   "\x1b[92m",
		"brightYellow":  "\x1b[93m",
		"brightBlue":    "\x1b[94m",
		"brightMagenta": "\x1b[95m",
		"brightCyan":    "\x1b[96m",
		"brightWhite":   "\x1b[97m",
		"resetCode":     "\033[0m",
	}
}
