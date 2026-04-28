package game

//import ("fmt")

func contains(slice []uint8, v uint8) bool {
	for _, x := range slice {
		if x == v {
			return true
		}
	}
	return false
}

func min(a, b uint8) uint8 {
	if a < b {
		return a
	} else {
		return b
	}
}
