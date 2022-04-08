package pagination

func Page[T any](tar []*T, index uint, size uint) ([]*T, uint) {
	var l uint
	if l = uint(len(tar)); tar == nil || l == 0 {
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
