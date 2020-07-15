package config

type Config struct {
	CarUrl     string `long:"car-url" description:"Car service url"`
	PredictUrl string `long:"predict-url" description:"Predict service url"`
}
