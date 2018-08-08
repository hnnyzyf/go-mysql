package hash

//the hash method for old password method
func Hash(password []byte) []uint32 {
	var add uint32 = 7
	result := []uint32{0x50305735, 0x12345671}

	for i := range password {
		// skip spaces and tabs in password
		if password[i] == ' ' || password[i] == '\t' {
			continue
		}
		tmp := uint32(password[i])
		result[0] ^= (((result[0] & 63) + add) * tmp) + (result[0] << 8)
		result[1] += (result[1] << 8) ^ result[0]
		add += tmp
	}
	// Remove sign bit (1<<31)-1)
	result[0] &= 0x7FFFFFFF
	result[1] &= 0x7FFFFFFF

	return result
}

type seed struct {
	seed1 uint32
	seed2 uint32
}

const maxSeed = 0x3FFFFFFF

// Pseudo random number generator
func NewSeed(seed1, seed2 uint32) *seed {
	return &seed{
		seed1: seed1 % maxSeed,
		seed2: seed2 % maxSeed,
	}
}

//Next return next seed
func (r *seed) Next() byte {
	r.seed1 = (r.seed1*3 + r.seed2) % maxSeed
	r.seed2 = (r.seed1 + r.seed2 + 33) % maxSeed
	return byte(uint64(r.seed1) * 31 / maxSeed)
}
