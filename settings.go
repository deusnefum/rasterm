package rasterm

type Settings struct {
	EscapeTmux    bool
	X, Y          int
	Rows, Columns int
	MoveCursor    bool
}
