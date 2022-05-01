package engine

/*
 The index in each square is the position in a uint64 (with 0 being the LSB and
 63 being the MSB):

         ┌────┬────┬────┬────┬────┬────┬────┬────┐
 Rank 8  │ 56 │ 57 │ 58 │ 59 │ 60 │ 61 │ 62 │ 63 │
         ├────┼────┼────┼────┼────┼────┼────┼────┤
 Rank 7  │ 48 │ 49 │ 50 │ 51 │ 52 │ 53 │ 54 │ 55 │
         ├────┼────┼────┼────┼────┼────┼────┼────┤
 Rank 6  │ 40 │ 41 │ 42 │ 43 │ 44 │ 45 │ 46 │ 47 │
         ├────┼────┼────┼────┼────┼────┼────┼────┤
 Rank 5  │ 32 │ 33 │ 34 │ 35 │ 36 │ 37 │ 38 │ 39 │
         ├────┼────┼────┼────┼────┼────┼────┼────┤
 Rank 4  │ 24 │ 25 │ 26 │ 27 │ 28 │ 29 │ 30 │ 31 │
         ├────┼────┼────┼────┼────┼────┼────┼────┤
 Rank 3  │ 16 │ 17 │ 18 │ 19 │ 20 │ 21 │ 22 │ 23 │
         ├────┼────┼────┼────┼────┼────┼────┼────┤
 Rank 2  │  8 │  9 │ 10 │ 11 │ 12 │ 13 │ 14 │ 15 │
         ├────┼────┼────┼────┼────┼────┼────┼────┤
 Rank 1  │  0 │  1 │  2 │  3 │  4 │  5 │  6 │  7 │
         └────┴────┴────┴────┴────┴────┴────┴────┘
   File:    a    b    c    d    e    f    g    h
*/

const (
	rank1 uint64 = 0xff << (8 * iota)
	rank2
	rank3
	rank4
	rank5
	rank6
	rank7
	rank8

	rank12 = rank1 | rank2
	rank18 = rank1 | rank8
	rank27 = rank2 | rank7
	rank78 = rank7 | rank8
)

const (
	fileA uint64 = 0x0101010101010101 << iota
	fileB
	fileC
	fileD
	fileE
	fileF
	fileG
	fileH

	fileAH = fileA | fileH
	fileBG = fileB | fileG
	fileCF = fileC | fileF
)

type side byte

const (
	white side = 0
	black side = 1
)

type Position struct {
	side [2]uint64
	pawn uint64
	nite uint64
	bish uint64
	rook uint64
	quee uint64
	king uint64
	// TODO: half-move clock for the 50 move rule
	// TODO: castling rights
	// TODO: En passant
	// TODO: side to move
	// TODO: Also have to manage move list somehow for repetition.
}

func InitialPosition() Position {
	return Position{
		side: [2]uint64{rank12, rank78},
		pawn: rank27,
		nite: rank18 & fileBG,
		bish: rank18 & fileCF,
		rook: rank18 & fileAH,
		quee: rank18 & fileD,
		king: rank18 & fileE,
	}
}
