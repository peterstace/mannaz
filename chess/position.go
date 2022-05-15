package chess

import (
	"fmt"
)

// side is an enum of either "white" or "black", and can represent either a
// player or a piece colour. It can also be used as an array index for `[2]T`
// style data structures.
type side byte

const (
	white side = 0
	black side = 1
)

func (s side) String() string {
	return [2]string{"white", "black"}[s]
}

type Position struct {
	// piece specific bitboards show if a square is occupied by a piece.
	pawn [2]uint64
	nite [2]uint64
	bish [2]uint64
	rook [2]uint64
	quee [2]uint64
	king [2]uint64

	// hmc (half move clock) counts the number of (half) moves since a pawn has
	// been captured captured or moved.
	hmc int16

	// fmc (full move clock) counts the number of (full) moves in the game so
	// far. It starts at 1 and increments whenever black completes a move.
	fmc int16

	// qcr shows queen-side castle rights for each side.
	qcr [2]bool

	// kcr shows king-side castle rights for each side.
	kcr [2]bool

	// dp indicates if a pawn was double pushed in the last move.
	dp bool

	// dpFile gives the file of a pawn that was just double pushed (otherwise
	// set to 0).
	dpFile file

	// stm is the side to move.
	stm side
}

func (p *Position) assertInvariants() {
	var occ uint64
	for s := white; s < black; s++ {
		for _, pieceBB := range [...][2]uint64{
			p.pawn, p.nite, p.bish,
			p.rook, p.quee, p.king,
		} {
			if occ&pieceBB[s] != 0 {
				panic("overlapping piece bitboards")
			}
			occ |= pieceBB[s]
		}
	}

	backRank := [2]uint64{rank1, rank8}
	for s := white; s < black; s++ {
		for _, castle := range [2]struct {
			name     string
			rookFile uint64
			rights   bool
		}{
			{"queen", fileA, p.qcr[s]},
			{"king", fileH, p.kcr[s]},
		} {
			if castle.rights {
				requireRook := backRank[s] & castle.rookFile
				if (requireRook & p.rook[s]) == 0 {
					panic(fmt.Sprintf(
						"%s %s-side castle rights but %s's rook not present",
						s, castle.name, castle.name,
					))
				}
			}
		}
		if p.qcr[s] || p.kcr[s] {
			requireKing := backRank[s] & fileE
			if (requireKing & p.king[s]) == 0 {
				panic(fmt.Sprintf("%s castle rights "+
					"but king not at origin", s))
			}
		}
	}

	if !p.dp && p.dpFile != 0 {
		panic("dp set but dpFile not zero")
	}
	if p.dp {
		r := [2]uint64{rank6, rank3}[p.stm]
		f := fileBB(p.dpFile)
		if (r & f) == 0 {
			panic("en passant pawn missing")
		}
	}
}

func InitialPosition() Position {
	pos := Position{
		pawn:   [2]uint64{rank2, rank7},
		nite:   [2]uint64{rank1 & fileBG, rank8 & fileBG},
		bish:   [2]uint64{rank1 & fileCF, rank8 & fileCF},
		rook:   [2]uint64{rank1 & fileAH, rank8 & fileAH},
		quee:   [2]uint64{rank1 & fileD, rank8 & fileD},
		king:   [2]uint64{rank1 & fileE, rank8 & fileE},
		hmc:    0,
		fmc:    1,
		qcr:    [2]bool{true, true},
		kcr:    [2]bool{true, true},
		dp:     false,
		dpFile: 0,
		stm:    white,
	}
	pos.assertInvariants()
	return pos
}
