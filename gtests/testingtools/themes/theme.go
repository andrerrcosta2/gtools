// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package themes

type Theme struct {
	Clock    string
	Error    string
	Info     string
	Question string
	Success  string
	Warning  string
}

var Default = Theme{
	Clock:    "⏱",
	Error:    "✘",
	Info:     "ℹ",
	Question: "?",
	Success:  "✔",
	Warning:  "⚠",
}

var Color = Theme{
	Clock:    "⏰",
	Error:    "❌",
	Info:     "ℹ️",
	Question: "❓",
	Success:  "✔️",
	Warning:  "⚠️",
}
