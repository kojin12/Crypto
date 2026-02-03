package checkresult

type DataResult struct {
	Ema50Result      int
	Ema200Result     int
	Trend15Result    int
	Trend60Result    int
	RsiResult        int
	MACDResult       int
	AtrResult        int
	VolumeResult     int
	PatternsResult   int
	ADXResult        int
	ATRPercentResult int
}

func GetIndicators(dataRes map[string]int) DataResult {
	var result DataResult

	result.Ema50Result = dataRes["Ema50Result"]
	result.Ema200Result = dataRes["Ema200Result"]
	result.Trend15Result = dataRes["Trend15Result"]
	result.Trend60Result = dataRes["Trend60Result"]
	result.RsiResult = dataRes["RsiResult"]
	result.MACDResult = dataRes["MACDResult"]
	result.AtrResult = dataRes["AtrResult"]
	result.VolumeResult = dataRes["VolumeResult"]
	result.PatternsResult = dataRes["PatternsResult"]
	result.ADXResult = dataRes["ADXResult"]
	result.ATRPercentResult = dataRes["ATRPercentResult"]

	return result
}
