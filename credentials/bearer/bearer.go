package bearer

import (
	"context"
	"fmt"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	mdKey          = "authorization"
	mdValuePrefix  = "Bearer "
	mdValueScanFmt = mdValuePrefix + "%s"
)

type Token string

func (t Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	ri, _ := credentials.RequestInfoFromContext(ctx)
	if err := credentials.CheckSecurityLevel(ri.AuthInfo, credentials.PrivacyAndIntegrity); err != nil {
		return nil, fmt.Errorf("unable to transfer bearer.Token PerRPCCredentials: %v", err)
	}
	return map[string]string{
		mdKey: mdValuePrefix + string(t),
	}, nil
}

func (t Token) RequireTransportSecurity() bool {
	return true
}

func NewPerRPCCredentials(token string) credentials.PerRPCCredentials {
	return Token(token)
}

func TokenFromContext(c context.Context) (Token, error) {
	var t Token
	mv := metadata.ValueFromIncomingContext(c, mdKey)
	if mv == nil || len(mv) == 0 {
		return t, fmt.Errorf("bearer credential metadata key not found in context")
	}
	_, err := fmt.Sscanf(mv[0], mdValueScanFmt, &t)
	return t, err
}
