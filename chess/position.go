package chess

import "math/bits"

// side is an enum of either "white" or "black", and can represent either a
// player or a piece colour. It can also be used as an array index for `[2]T`
// style data structures.
type side byte

const (
	white side = 0
	black side = 1
)

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

	// flags stores flags specific to each side in each nibble. See functions
	// named `*Mask` for details.
	flags uint8

	// stm is the side to move.
	stm side

	// TODO: Also have to manage move list somehow for repetition.
}

func queenCastleMask(s side) uint8 {
	return 0x01 << (s * 4)
}

func kingCastleMask(s side) uint8 {
	return 0x02 << (s * 4)
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

	// TODO: additional assertions
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
		flags: 0x00 |
			queenCastleMask(white) |
			queenCastleMask(black) |
			kingCastleMask(white) |
			kingCastleMask(black),
		stm: white,
	}
	pos.assertInvariants()
	return pos
}
