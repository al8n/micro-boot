package jwt

import (
	"errors"
	bootflag "github.com/al8n/micro-boot/flag"
	"github.com/al8n/micro-boot/internal/utils"
	"github.com/golang-jwt/jwt"
	"io"
	"os"
	"strings"
	"time"
)

var (
	defaultJWTPrefix = "jwt"
	defaultIssuer = ""
	defaultIssuerUsage = "set the issuer for JWT standard claims"
	defaultAudience = ""
	defaultAudienceUsage = "set the audience for JWT standard claims"
	defaultId = ""
	defaultIdUsage = "set the jti for JWT standard claims"
	defaultExpiresAt int64 = 0
	defaultExpiresAtUsage = "set the exp for JWT standard claims"
	defaultIssuedAt int64 = 0
	defaultIssuedAtUsage = "set the isa for JWT standard claims"
	defaultNotBefore int64 = 0
	defaultNotBeforeUsage = "set the nbf for JWT standard claims"
	defaultSubject = ""
	defaultSubjectUsage = "set the sub for JWT standard claims"
	defaultExpiration time.Duration = 0
	defaultExpirationUsage = "set the expiration for a JWT token, will be used if exp is not set"
	defaultPrivateKey = ""
	defaultPrivateKeyUsage = "set the private key to generate JWT token, if private key file is not provided. Otherwise, overwritten by private key file content"
	defaultPrivateKeyFile = ""
	defaultPrivateKeyFileUsage = "set the private key file to generate JWT token"
	defaultPublicKey = ""
	defaultPublicKeyUsage = "set the public key to generate JWT token, if public key file is not provided. Otherwise, overwritten by public key file content"
	defaultPublicKeyFile = ""
	defaultPublicKeyFileUsage = "set the public key file to generate JWT token"
	defaultMethod = ""
	defaultMethodUsage = "set the algorithm to sign the JWT token"
)

// Config is the configuration for JWT standard claims and some other extension fields.
type Config struct {
	// Structured version of Claims Section, as referenced at
	// https://tools.ietf.org/html/rfc7519#section-4.1
	// See examples for how to use this with your own claim types
	Audience  string `json:"aud,omitempty" yaml:"aud,omitempty"`

	// The "exp" (expiration time) claim identifies the expiration time on
	// or after which the JWT MUST NOT be accepted for processing.  The
	// processing of the "exp" claim requires that the current date/time
	// MUST be before the expiration date/time listed in the "exp" claim.
	// Implementers MAY provide for some small leeway, usually no more than
	// a few minutes, to account for clock skew.  Its value MUST be a number
	// containing a NumericDate value.  Use of this claim is OPTIONAL.
	ExpiresAt int64  `json:"exp,omitempty" yaml:"exp,omitempty"`

	// The "jti" (JWT ID) claim provides a unique identifier for the JWT.
	// The identifier value MUST be assigned in a manner that ensures that
	// there is a negligible probability that the same value will be
	// accidentally assigned to a different data object; if the application
	// uses multiple issuers, collisions MUST be prevented among values
	// produced by different issuers as well.  The "jti" claim can be used
	// to prevent the JWT from being replayed.  The "jti" value is a case-
	// sensitive string.  Use of this claim is OPTIONAL.
	Id        string `json:"jti,omitempty" yaml:"jti,omitempty"`

	// The "iat" (issued at) claim identifies the time at which the JWT was
	// issued.  This claim can be used to determine the age of the JWT.  Its
	// value MUST be a number containing a NumericDate value.  Use of this
	// claim is OPTIONAL.
	IssuedAt  int64  `json:"iat,omitempty" yaml:"iat,omitempty"`

	// The "iss" (issuer) claim identifies the principal that issued the
	// JWT.  The processing of this claim is generally application specific.
	// The "iss" value is a case-sensitive string containing a StringOrURI
	// value.  Use of this claim is OPTIONAL.
	Issuer    string `json:"iss,omitempty" yaml:"iss,omitempty"`

	// The "nbf" (not before) claim identifies the time before which the JWT
	// MUST NOT be accepted for processing.  The processing of the "nbf"
	// claim requires that the current date/time MUST be after or equal to
	// the not-before date/time listed in the "nbf" claim.  Implementers MAY
	// provide for some small leeway, usually no more than a few minutes, to
	// account for clock skew.  Its value MUST be a number containing a
	// NumericDate value.  Use of this claim is OPTIONAL.
	NotBefore int64  `json:"nbf,omitempty" yaml:"nbf,omitempty"`

	// The "sub" (subject) claim identifies the principal that is the
	// subject of the JWT.  The claims in a JWT are normally statements
	// about the subject.  The subject value MUST either be scoped to be
	// locally unique in the context of the issuer or be globally unique.
	// The processing of this claim is generally application specific.  The
	// "sub" value is a case-sensitive string containing a StringOrURI
	// value.  Use of this claim is OPTIONAL.
	Subject   string `json:"sub,omitempty" yaml:"sub,omitempty"`

	// Expiration is the expiration for a JWT token, which is used when the ExpiresAt field is 0
	Expiration time.Duration `json:"expiration" yaml:"expiration"`

	// PrivateKeyFile is used to generate JWT token, if this field is specified,
	// PrivateKey will be overwrite by the file content.
	PrivateKeyFile string `json:"private-key-file" yaml:"private-key-file"`

	// PrivateKey is used to generate JWT token
	PrivateKey string `json:"private-key" yaml:"private-key"`

	// PublicKeyFile is used to generate JWT token, if this field is specified,
	// PublicKey will be overwrite by the file content.
	PublicKeyFile string `json:"public-key-file" yaml:"public-key-file"`

	// PublicKey is used to generate JWT token
	PublicKey string `json:"public-key" yaml:"public-key"`

	// Method is the algorithm used to sign the JWT token
	Method string `json:"method" yaml:"method"`

	// CustomBindFlagsFunc defines custom bind flags behaviour for structure,
	// if CustomBindFlagsFunc is nil, default  bind flags behaviour will be used
	CustomBindFlagsFunc func(fs *bootflag.FlagSet) `json:"-" yaml:"-"`

	// CustomParseFunc defines custom parse behaviour for structure,
	// if CustomParseFunc is nil, default parse behaviour will be used
	CustomParseFunc func() (err error) `json:"-" yaml:"-"`
}

func (c *Config) SetAudience(aud string) {
	c.Audience = aud
}

func (c *Config) SetExpiresAt(exp int64) {
	c.ExpiresAt = exp
}

func (c *Config) SetId(id string) {
	c.Id = id
}

func (c *Config)  SetIssueAt(isa int64) {
	c.IssuedAt = isa
}

func (c *Config)  SetIssuer(iss string) {
	c.Issuer = iss
}

func (c *Config)  SetNotBefore(nbf int64) {
	c.NotBefore = nbf
}

func (c *Config)  SetSubject(sub string) {
	c.Subject = sub
}

func (c Config) Standardize() (claims jwt.StandardClaims) {
	var (
		now time.Time
	)

	claims = jwt.StandardClaims{}

	if c.Audience != "" {
		claims.Audience = c.Audience
	}

	if c.Id != "" {
		claims.Id = c.Id
	}

	if c.Issuer != "" {
		claims.Issuer = c.Issuer
	}

	if c.NotBefore != 0 {
		claims.NotBefore = c.NotBefore
	}

	if c.Subject != "" {
		claims.Subject = c.Subject
	}

	if c.IssuedAt == 0 {
		claims.IssuedAt = now.Unix()
	}

	if c.ExpiresAt == 0 {
		if c.Expiration != 0 {
			claims.ExpiresAt = now.Add(c.Expiration).Unix()
		}
	} else {
		claims.ExpiresAt = c.ExpiresAt
	}

	return claims
}

func (c *Config) BindFlags(fs *bootflag.FlagSet)  {
	if c.CustomBindFlagsFunc != nil {
		c.CustomBindFlagsFunc(fs)
		return
	}

	fs.StringVar(
		&c.Audience,
		utils.BuildFlagName(defaultJWTPrefix, "aud"),
		defaultAudience,
		defaultAudienceUsage)

	fs.Int64Var(
		&c.ExpiresAt,
		utils.BuildFlagName(defaultJWTPrefix, "exp"),
		defaultExpiresAt,
		defaultExpiresAtUsage)

	fs.Int64Var(
		&c.IssuedAt,
		utils.BuildFlagName(defaultJWTPrefix, "iat"),
		defaultIssuedAt,
		defaultIssuedAtUsage)

	fs.StringVar(
		&c.Issuer,
		utils.BuildFlagName(defaultJWTPrefix, "iss"),
		defaultIssuer,
		defaultIssuerUsage)

	fs.StringVar(
		&c.Id,
		utils.BuildFlagName(defaultJWTPrefix, "id"),
		defaultId,
		defaultIdUsage)

	fs.Int64Var(
		&c.NotBefore,
		utils.BuildFlagName(defaultJWTPrefix, "nbf"),
		defaultNotBefore,
		defaultNotBeforeUsage)

	fs.StringVar(
		&c.Subject,
		utils.BuildFlagName(defaultJWTPrefix, "sub"),
		defaultSubject,
		defaultSubjectUsage)

	fs.StringVar(
		&c.PrivateKey,
		utils.BuildFlagName(defaultJWTPrefix, "private-key"),
		defaultPrivateKey,
		defaultPrivateKeyUsage)

	fs.StringVar(
		&c.PrivateKeyFile,
		utils.BuildFlagName(defaultJWTPrefix, "private-key-file"),
		defaultPrivateKeyFile,
		defaultPrivateKeyFileUsage)

	fs.StringVar(
		&c.PublicKey,
		utils.BuildFlagName(defaultJWTPrefix, "public-key"),
		defaultPublicKey,
		defaultPublicKeyUsage)

	fs.StringVar(
		&c.PublicKeyFile,
		utils.BuildFlagName(defaultJWTPrefix, "public-key-file"),
		defaultPublicKeyFile,
		defaultPublicKeyFileUsage)

	fs.StringVar(
		&c.Method,
		utils.BuildFlagName(defaultJWTPrefix, "method"),
		defaultMethod,
		defaultMethodUsage)
	
	fs.DurationVar(
		&c.Expiration,
		utils.BuildFlagName(defaultJWTPrefix, "expiration"),
		defaultExpiration,
		defaultExpirationUsage)
}

func (c *Config) Parse() (err error) {
	if c.CustomParseFunc != nil {
		return c.CustomParseFunc()
	}

	if c.PrivateKeyFile != "" {
		if f, err := os.OpenFile(c.PrivateKeyFile, os.O_RDONLY, 0444); err != nil {
			return err
		} else {
			if content, err := io.ReadAll(f); err != nil {
				return err
			} else {
				c.PrivateKey = string(content)
			}
		}
	}

	if c.PublicKeyFile != "" {
		if f, err := os.OpenFile(c.PublicKeyFile, os.O_RDONLY, 0444); err != nil {
			return err
		} else {
			if content, err := io.ReadAll(f); err != nil {
				return err
			} else {
				c.PublicKey = string(content)
			}
		}
	}

	if c.Method != "" {
		switch strings.ToUpper(c.Method) {
		case "HS256", "HS384", "HS512", "PS256", "PS384", "PS512", "RS256", "RS384", "RS512", "ES256", "ES384", "ES512":
			break
		default:
			return errors.New("invalid JWT signature method")
		}
	}
	
	c.Standardize()

	return nil
}