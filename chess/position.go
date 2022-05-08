package chess

import (
	"fmt"
	"math/bits"
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
	// side bitboards indicate the colour of a piece occupying a square.
	side [2]uint64

	// piece specific bitboards show if a square is occupied by a piece.
	pawn uint64
	nite uint64
	bish uint64
	rook uint64
	quee uint64
	king uint64

	// hmc (half move clock) counts the number of (half) moves since a pawn has
	// been captured captured or moved.
	hmc int16

	// ep show en passant rights for each file.
	ep [2]uint8

	// qcr shows queen-side castle rights for each side.
	qcr [2]bool

	// kcr shows king-side castle rights for each side.
	kcr [2]bool

	// stm is the side to move.
	stm side
}

func (p *Position) assertInvariants() {
	if p.side[0]&p.side[1] != 0 {
		panic("side bitboards overlap")
	}

	var ones int
	for _, bb := range []uint64{p.pawn, p.nite, p.bish, p.rook, p.quee, p.king} {
		ones += bits.OnesCount64(bb)
	}
	if ones != bits.OnesCount64(p.side[0]|p.side[1]) {
		panic("piece bitboards overlap")
	}

	if p.pawn|p.nite|p.bish|p.rook|p.quee|p.king != p.side[0]|p.side[1] {
		panic("piece bitboards and side bitboards don't match")
	}

	if p.ep[white] & ^uint8((p.pawn&p.side[white]&rank4)<<24) != 0 {
		panic("white en passant rights set but no pawn")
	}
	if p.ep[black] & ^uint8((p.pawn&p.side[black]&rank5)<<32) != 0 {
		panic("black en passant rights set but no pawn")
	}

	majorPieceRank := [2]uint64{rank1, rank8}
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
				requireRook := majorPieceRank[s] & castle.rookFile
				haveRook := p.side[s] & p.rook
				if (requireRook & haveRook) == 0 {
					panic(fmt.Sprintf(
						"%s %s-side castle rights but %s's rook not present",
						s, castle.name, castle.name,
					))
				}
			}
		}
		if p.qcr[s] || p.kcr[s] {
			requireKing := majorPieceRank[s] & fileE
			haveKing := p.side[s] & p.king
			if (requireKing & haveKing) == 0 {
				panic(fmt.Sprintf("%s castle rights "+
					"but king not at origin", s))
			}
		}
	}
}

//func (p *Position) FEN() string {
//	for ri := 7; ri >= 0; ri-- {
//		r := rank(ri)
//		for fi := 0; fi < 8; fi++ {
//			sq := r & file(fi)
//		}
//	}
//}

func InitialPosition() Position {
	pos := Position{
		side: [2]uint64{rank12, rank78},
		pawn: rank27,
		nite: rank18 & fileBG,
		bish: rank18 & fileCF,
		rook: rank18 & fileAH,
		quee: rank18 & fileD,
		king: rank18 & fileE,
		hmc:  0,
		ep:   [2]uint8{},
		qcr:  [2]bool{true, true},
		kcr:  [2]bool{true, true},
		stm:  white,
	}
	pos.assertInvariants()
	return pos
}
