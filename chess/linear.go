package chess

// file represents a file from file A (0) through to file H (7).
type file int8

// algebraic gives the algebraic file name 'a' through 'h'.
func (f file) algebraic() rune {
	return rune(f) + 'a'
}

// rank represents a rank from rank 1 (0) through to rank 8 (7).
type rank int8
