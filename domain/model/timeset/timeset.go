package timeset

// Setting ...
type Setting struct {
	Work       uint `json:"work"`
	ShortBreak uint `json:"short_break"`
	LongBreak  uint `json:"long_break"`
	Times      uint `json:"times"`
}
