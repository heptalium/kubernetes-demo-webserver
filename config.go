package main

type Config struct {
	Backend struct {
		CheckInterval uint `default:"15" split_words:"true"`
		URL           string
	}
	ConfigFile string `split_words:"true"`
	LogFile    string `split_words:"true"`
	Port       uint16 `default:"80"`
	StartDelay uint   `default:"5" split_words:"true"`
}
