//  XX X      X              XXX     XXX             XXX      XX      X       X      XXX
//  X X X           XXXX    X   X   X   X   X XX        X    X       XXX                X
//  X X X     X     X   X   XXXXX   X       XX  X    XXXX   XXXX      X       X      XXXX
//  X   X     X     X   X   X       X   X   X       X   X    X        X       X     X   X
//  X   X     X     X   X    XXXX    XXX    X        XXXX    X         X      X      XXXX

package rgbmatrix

import "github.com/pbnjay/pixfont"

var Font *pixfont.PixFont

func init() {
	charMap := map[int32]uint16{48: 0x0, 49: 0x1, 50: 0x2, 51: 0x3, 52: 0x18, 53: 0x19, 54: 0x1a, 55: 0x1b, 56: 0x30, 57: 0x31, 65: 0x32, 66: 0x33, 67: 0x48, 68: 0x49, 69: 0x4a, 70: 0x4b, 71: 0x60, 72: 0x61, 73: 0x62, 74: 0x63, 75: 0x78, 76: 0x79, 77: 0x7a, 78: 0x7b, 79: 0x90, 80: 0x91, 81: 0x92, 82: 0x93, 83: 0xa8, 84: 0xa9, 85: 0xaa, 86: 0xab, 87: 0xc0, 88: 0xc1, 89: 0xc2, 90: 0xc3, 97: 0xd8, 98: 0xd9, 99: 0xda, 100: 0xdb, 101: 0xf0, 102: 0xf1, 103: 0xf2, 104: 0xf3, 105: 0x108, 106: 0x109, 107: 0x10a, 108: 0x10b, 109: 0x120, 110: 0x121, 111: 0x122, 112: 0x123, 113: 0x138, 114: 0x139, 115: 0x13a, 116: 0x13b, 117: 0x150, 118: 0x151, 119: 0x152, 120: 0x153, 121: 0x168, 122: 0x169}
	data := []uint32{0x1c1e081c, 0x20200c22, 0x1808082a, 0x20020822, 0x1c3e3e1c, 0x0, 0x3e1c3e3c, 0x20020222, 0x101e1e3e, 0x8222020, 0x81c1e20, 0x0, 0x1e1c1c1c, 0x22222222, 0x1e3e3c1c, 0x22222022, 0x1e221c1c, 0x0, 0x3e3e1e1c, 0x2022222, 0xe0e2202, 0x2022222, 0x23e1e1c, 0x0, 0x201c223c, 0x20082202, 0x20083e3a, 0x22082222, 0x1c1c221c, 0x0, 0x22220222, 0x26360212, 0x2a2a020e, 0x32220212, 0x22223e22, 0x0, 0x1e1c1e1c, 0x22222222, 0x1e221e22, 0x22120222, 0x222c021c, 0x0, 0x22223e3c, 0x22220802, 0x2222081c, 0x14220820, 0x81c081e, 0x0, 0x3e222222, 0x20141422, 0x808082a, 0x2081436, 0x3e082222, 0x0, 0x201c021c, 0x20221a20, 0x2c02263c, 0x22222222, 0x3c1c1e3c, 0x0, 0x23c181c, 0x2220422, 0x1a221e3e, 0x223c0402, 0x2220043c, 0x1e0000, 0x2022008, 0x20a0000, 0x2062008, 0xa2008, 0x122208, 0x1c00, 0x16, 0x1a1c1e2a, 0x2222222a, 0x1e222222, 0x21c2222, 0x2000000, 0x83c0000, 0x1c021a2c, 0x81c2622, 0x820023c, 0x101e0220, 0x20, 0x22000000, 0x14222222, 0x82a2222, 0x142a1422, 0x223c083c, 0x0, 0x0, 0x3e22, 0x1022, 0x43c, 0x3e20, 0x1e}
	Font = pixfont.NewPixFont(7, 6, charMap, data)
	Font.SetVariableWidth(false)
}
