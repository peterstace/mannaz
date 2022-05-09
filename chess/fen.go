package chess

import (
	"fmt"
	"math/bits"
	"strings"
)

func (p *Position) FEN() string {
	var sb strings.Builder

	occ := p.side[white] | p.side[black]
	for ri := 7; ri >= 0; ri-- {
		r := rank(ri)
		var emptyCount int
		for fi := 0; fi < 8; fi++ {
			sq := r & file(fi)
			if (occ & sq) == 0 {
				emptyCount++
				continue
			}
			piece := '0'
			switch {
			case (p.pawn & sq) != 0:
				piece = 'P'
			case (p.nite & sq) != 0:
				piece = 'N'
			case (p.bish & sq) != 0:
				piece = 'B'
			case (p.rook & sq) != 0:
				piece = 'R'
			case (p.quee & sq) != 0:
				piece = 'Q'
			case (p.king & sq) != 0:
				piece = 'K'
			}
			if (p.side[black] & sq) != 0 {
				piece += 'a' - 'A'
			}
			if emptyCount != 0 {
				emptyCount = 0
				fmt.Fprintf(&sb, "%d", emptyCount)
			}
			sb.WriteRune(piece)
		}
		if emptyCount != 0 {
			fmt.Fprintf(&sb, "%d", emptyCount)
		}
	}

	sb.WriteRune(' ')

	sb.WriteRune([2]rune{'w', 'b'}[p.stm])

	sb.WriteRune(' ')

	var anyCastle bool
	for s := white; s < black; s++ {
		if p.kcr[s] {
			anyCastle = true
			sb.WriteRune('K' + rune(s)*('a'-'A'))
		}
		if p.qcr[s] {
			anyCastle = true
			sb.WriteRune('Q' + rune(s)*('a'-'A'))
		}
	}
	if !anyCastle {
		sb.WriteRune('-')
	}

	sb.WriteRune(' ')

	//var anyEP bool
	for s := white; s < black; s++ {
		ep := p.ep[s]
		if ep != 0 {
			//anyEP = true
			file := bits.TrailingZeros8(ep)
			flippedPawns := flipDiagA1H8(p.pawn & p.side[s])
			flippedPawns >>= 8 * file
			rank := bits.TrailingZeros8(uint8(flippedPawns))
			sb.WriteRune(rune(file) + 'a')
			sb.WriteRune(rune(rank) + '0') // must be 3 or 6?
			continue
		}
	}

	return sb.String()
}
