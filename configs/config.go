package configs

type Configs struct {
	Port  string `json:"port"`
}

func NewConfig() *Configs {
	return &Configs{}
}
