package auth

type AuthConfig struct {
	LifespanHours string `json:"lifespanHours"`
	Secret        string `json:"secret"`
}