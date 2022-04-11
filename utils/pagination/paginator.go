package pagination

// Page Tar is the object's point slice, index is the page index, size is the page capacity,
// return goal object's point slice and the total page number.
func Page[T any](tar []*T, index uint32, size uint32) ([]*T, uint32) {
	var l uint32
	if l = uint32(len(tar)); tar == nil || l == 0 {
		return []*T{}, 0
	}
	ps := l / size
	if l%size > 0 {
		ps += 1
	}
	if index > ps {
		index = ps
	}
	start := (index - 1) * size
	end := start + size
	if end > l {
		end = l
	}
	return tar[start:end], ps
}
