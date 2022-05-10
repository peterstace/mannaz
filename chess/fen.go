package chess

import (
	"fmt"
	"strings"
)

func (p *Position) FEN() string {
	var sb strings.Builder

	occ := p.side[white] | p.side[black]
	for ri := rank(7); ri >= 0; ri-- {
		r := rankBB(ri)
		var emptyCount int
		for fi := file(0); fi < 8; fi++ {
			sq := r & fileBB(fi)
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

	if p.dp {
		sb.WriteRune(p.dpFile.algebraic())
		sb.WriteRune([2]rune{'6', '3'}[p.stm])
	} else {
		sb.WriteRune('-')
	}

	fmt.Fprintf(&sb, " %d %d", p.hmc, p.fmc)

	return sb.String()
}
