package channel

type Mode string

const (
	InputOnlyMode Mode = "input-only"
	OutputOnly    Mode = "output-only"
	BothMode      Mode = "both"
)

func (m Mode) IsOutputMode() bool {
	return m == OutputOnly || m == BothMode
}
func (m Mode) IsInputMode() bool {
	return m == InputOnlyMode || m == BothMode
}
