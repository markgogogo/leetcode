package convertthetemperature2469

func convertTemperature(celsius float64) []float64 {
	var ans []float64
	ans = append(ans, celsius+273.15, celsius*1.8+32)
	return ans
}
