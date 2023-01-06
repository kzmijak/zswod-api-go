package jwt

import "github.com/kzmijak/zswod_api_go/modules/auth"

const (
	CLAIM_AUTHORIZED = "authorized"
	CLAIM_GUID       = "guid"
	CLAIM_ROLE       = "role"
	CLAIM_EXPIRED    = "exp"
)

const (
ErrConfigRequired						 = "ErrConfigRequired: Config is required to run this method"
	ErrFailedToConvertLifespan = "ErrFailedToConvertLifespan: Failed to read lifespan from config"
	ErrFailedSigningToken      = "ErrFailedSigningToken: Failed to sign token with secret from config"
)

type JwtService struct {
	cfg *auth.AuthConfig
}

func New() *JwtService {
	return &JwtService{}
}

func (s *JwtService) WithConfig(cfg auth.AuthConfig) *JwtService {
	s.cfg = &cfg
	return s
}