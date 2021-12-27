package configs

type Configs struct {
	Port  string `json:"port"`
	GRPCPort string `json:"grpc_port"`
	RedisAddr string `json:"redis_addr"`
}

func NewConfig() *Configs {
	return &Configs{}
}
