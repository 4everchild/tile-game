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
