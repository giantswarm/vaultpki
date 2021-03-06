package vaultpki

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	vaultapi "github.com/hashicorp/vault/api"
)

const (
	// MountType is the mount type used to mount a PKI backend in Vault.
	MountType = "pki"
)

type Config struct {
	Logger      micrologger.Logger
	VaultClient *vaultapi.Client

	CATTL            string
	CommonNameFormat string
}

func DefaultConfig() Config {
	return Config{
		Logger:      nil,
		VaultClient: nil,

		CATTL:            "",
		CommonNameFormat: "",
	}
}

type VaultPKI struct {
	logger      micrologger.Logger
	vaultClient *vaultapi.Client

	caTTL            string
	commonNameFormat string
}

func New(config Config) (*VaultPKI, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Logger must not be empty")
	}
	if config.VaultClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.VaultClient must not be empty")
	}

	if config.CATTL == "" {
		return nil, microerror.Maskf(invalidConfigError, "config.CATTL must not be empty")
	}
	if config.CommonNameFormat == "" {
		return nil, microerror.Maskf(invalidConfigError, "config.CommonNameFormat must not be empty")
	}

	p := &VaultPKI{
		logger:      config.Logger,
		vaultClient: config.VaultClient,

		caTTL:            config.CATTL,
		commonNameFormat: config.CommonNameFormat,
	}

	return p, nil
}
