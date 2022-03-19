package editor

type PosConv struct {
	newLines []int
}

func NewPosConv(input string) *PosConv {
	pc := &PosConv{}
	for i, s := range input {
		if s == '\n' {
			pc.newLines = append(pc.newLines, i)
		}
	}
	return pc
}

func (pc PosConv) ToAbsolute(line, column int) int {
	if line < 2 {
		return column
	}
	return pc.newLines[line-2] + column + 1
}

func (pc PosConv) ToRelative(abs int) (int, int) {
	for i := len(pc.newLines) - 1; i >= 0; i -= 1 {
		column := abs - pc.newLines[i]
		if column >= 1 {
			return i + 2, column - 1
		}
	}
	return 1, abs
}
