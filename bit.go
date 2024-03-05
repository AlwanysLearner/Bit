package Bit

import "errors"

type Bits struct {
	Size int
	Mem  []uint8
}

func NewBits(size int) *Bits {
	return &Bits{size, make([]uint8, size/8+size%8*1)}
}

func (bits *Bits) SetOne(index int) error {
	if index < 0 || index >= bits.Size {
		return errors.New("inde out of range")
	}
	bits.Mem[index/8] |= uint8(1 << (8 - index%8))
	return nil
}

func (bits *Bits) SetZero(index int) error {
	if index < 0 || index >= bits.Size {
		return errors.New("inde out of range")
	}
	bits.Mem[index/8] &= uint8(255 ^ (1 << (8 - index%8)))
	return nil
}
