package game

type XS64 uint64

func (r XS64) Next() XS64 {
	x := uint64(r)

	x ^= x >> 12
	x ^= x << 25
	x ^= x >> 27

	return XS64(x)
}
func (r XS64) Value() uint64 {
	return uint64(r) * 2685821657736338717
}

func (r *XS64) Step() uint64 {
	val := r.Value()
	*r = r.Next()
	return val
}
