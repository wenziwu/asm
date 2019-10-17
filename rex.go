package asm

// REX is used to generate a REX prefix.
// w, r, x and b can only be set to either 0 or 1.
func REX(w, r, x, b byte) byte {
	return 64 | (w << 3) | (r << 2) | (x << 1) | b
}
