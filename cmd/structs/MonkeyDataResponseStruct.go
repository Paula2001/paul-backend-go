package structs

type MonkeyDataResponseStruct struct {
	Success  bool
	Messages []string
	Version  string
	Data     []MonkeyDataCountriesResponseStruct
}
