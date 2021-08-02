package protocol

type VarInt int32

func VarIntSize(v int) int {
	if (v & 0xFFFFFF80) == 0 {
		return 1
	}

	if (v & 0xFFFFC000) == 0 {
		return 2
	}

	if (v & 0xFFE00000) == 0 {
		return 3
	}

	if (v & 0xF0000000) == 0 {
		return 4
	}

	return 5
}