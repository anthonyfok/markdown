package parser

import (
	"bytes"
)

func (p *Parser) caption(data, caption []byte) ([]byte, int) {
	if !bytes.HasPrefix(data, caption) {
		return nil, 0
	}
	j := len(caption)
	data = data[j:]
	end := p.linesUntilEmpty(data)

	return data[:end], end + j
}

// linesUntilEmpty scans lines up to the first empty line.
func (p *Parser) linesUntilEmpty(data []byte) int {
	line, i := 0, 0

	for line < len(data) {
		i++

		// find the end of this line
		for i < len(data) && data[i-1] != '\n' {
			i++
		}

		if p.isEmpty(data[line:i]) == 0 {
			line = i
			continue
		}

		break
	}
	return i
}
