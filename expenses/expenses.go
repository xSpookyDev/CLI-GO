package expenses

//Average promedia los gastos
func Average(expns ...float32) float32 {

	return Sum(expns...) / float32(len(expns))
}

//Sum Hace la suma de los gastos
func Sum(expns ...float32) float32 {
	var sum float32
	for _, exp := range expns {
		sum += exp
	}

	return sum
}

//Max Funcion que indica el maximo de los gastos
func Max(expns ...float32) float32 {
	var max float32
	for _, exp := range expns {
		if exp > max {
			max = exp
		}
	}
	return max
}

//Min Devuelve el minimo de los gastos
func Min(expns ...float32) float32 {

	if len(expns) == 0 {
		return 0
	}
	var min float32 = expns[0]
	for _, exp := range expns {
		if exp < min {
			min = exp
		}
	}
	return min
}
