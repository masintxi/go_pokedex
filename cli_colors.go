package main // or whatever your package name is

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorReset  = "\033[0m"
)

func fmtSuccess(msg string) string {
	return colorGreen + msg + colorReset
}

func fmtFail(msg string) string {
	return colorRed + msg + colorReset
}

func fmtAction(msg string) string {
	return colorYellow + msg + colorReset
}

func alternateColor(text string, useColor bool) string {
	if useColor {
		return colorCyan + text + colorReset
	}
	return text
}
