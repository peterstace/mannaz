package chess

import (
	"fmt"
	"strings"
)

func (p *Position) FEN() string {
	var sb strings.Builder

	for ri := rank(7); ri >= 0; ri-- {
		r := rankBB(ri)
		var emptyCount int
		for fi := file(0); fi < 8; fi++ {
			sq := r & fileBB(fi)

			var piece rune
			for _, itm := range []struct {
				bb    uint64
				piece rune
			}{
				{p.pawn[white], 'P'},
				{p.nite[white], 'N'},
				{p.bish[white], 'B'},
				{p.rook[white], 'R'},
				{p.quee[white], 'Q'},
				{p.king[white], 'K'},
				{p.pawn[black], 'p'},
				{p.nite[black], 'n'},
				{p.bish[black], 'b'},
				{p.rook[black], 'r'},
				{p.quee[black], 'q'},
				{p.king[black], 'k'},
			} {
				if sq&itm.bb != 0 {
					piece = itm.piece
				}
			}
			if piece == 0 {
				emptyCount++
				continue
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
