package tokenizer

import (
	"io"
	"sort"
)

type LineEndTracker struct {
	r           io.Reader
	lineOffsets []int
	totalRead   int
	sawCR       bool
}

func NewLineEndTracker(r io.Reader) *LineEndTracker {
	return &LineEndTracker{
		r:           r,
		lineOffsets: []int{0},
	}
}

func (trk *LineEndTracker) Read(p []byte) (int, error) {
	n, err := trk.r.Read(p)
	for i := 0; i < n; i++ {
		b := p[i]
		if trk.sawCR {
			trk.sawCR = false
			if b == '\n' {
				trk.lineOffsets = append(trk.lineOffsets, trk.totalRead+i+1)
				continue
			} else {
				trk.lineOffsets = append(trk.lineOffsets, trk.totalRead+i)
			}
		}

		switch b {
		case '\n':
			trk.lineOffsets = append(trk.lineOffsets, trk.totalRead+i+1)
		case '\r':
			trk.sawCR = true
		}
	}

	if trk.sawCR && err == io.EOF {
		trk.lineOffsets = append(trk.lineOffsets, trk.totalRead+n)
		trk.sawCR = false
	}

	trk.totalRead += n
	return n, err
}

func (ptr *LineEndTracker) LineCol(pos int) (line, col int) {
	idx := sort.Search(len(ptr.lineOffsets), func(i int) bool {
		return ptr.lineOffsets[i] > pos
	})

	if idx == 0 {
		return 1, pos + 1
	}

	line = idx
	col = pos - ptr.lineOffsets[idx-1] + 1
	return
}
