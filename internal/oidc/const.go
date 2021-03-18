package oidc

const (
	fallbackOIDCIssuer = "https://login.example.com:8080"
	wellKnownPath      = "/.well-known/openid-configuration"
	jwksPath           = "/.well-known/jwks.json"
	authorizePath      = "/api/oidc/authorize"
	tokenPath          = "/api/oidc/token" // nolint:gosec
	consentPath        = "/api/oidc/consent"
	introspectPath     = "/api/oidc/introspect"
	revokePath         = "/api/oidc/revoke"
)

var scopeDescriptions = map[string]string{
	"openid":  "Account Information",
	"email":   "Email Addresses",
	"profile": "User Profile",
	"groups":  "Group Membership",
}

var audienceDescriptions = map[string]string{}
