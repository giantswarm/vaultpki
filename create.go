package vaultpki

import (
	"fmt"

	"github.com/giantswarm/microerror"
	vaultapi "github.com/hashicorp/vault/api"

	"github.com/giantswarm/vaultpki/key"
)

func (p *VaultPKI) CreateBackend(ID string) error {
	k := key.MountPKIPath(ID)
	v := &vaultapi.MountInput{
		Config: vaultapi.MountConfigInput{
			MaxLeaseTTL: p.caTTL,
		},
		Description: fmt.Sprintf("PKI backend for ID '%s'", ID),
		Type:        MountType,
	}

	err := p.vaultClient.Sys().Mount(k, v)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (p *VaultPKI) createNewCA(ID string, exported bool) (CertificateAuthority, error) {
	k := key.WriteCAPath(ID, exported)
	v := map[string]interface{}{
		"common_name": key.CommonName(ID, p.commonNameFormat),
		"ttl":         p.caTTL,
	}

	secret, err := p.vaultClient.Logical().Write(k, v)
	if err != nil {
		return DefaultCertificateAuthority(), microerror.Mask(err)
	}

	var certificate string
	{
		value, ok := secret.Data["certificate"]
		if !ok {
			return DefaultCertificateAuthority(), microerror.Maskf(executionFailedError, "certificate missing")
		}
		certificate, ok = value.(string)
		if !ok {
			return DefaultCertificateAuthority(), microerror.Maskf(executionFailedError, "certificate must be string")
		}
	}

	var privateKey string
	{
		if exported {
			value, ok := secret.Data["private_key"]
			if !ok {
				return DefaultCertificateAuthority(), microerror.Maskf(executionFailedError, "private key missing")
			}
			privateKey, ok = value.(string)
			if !ok {
				return DefaultCertificateAuthority(), microerror.Maskf(executionFailedError, "private key must be string")
			}
		} else {
			privateKey = ""
		}
	}

	return CertificateAuthority{
		Certificate: certificate,
		PrivateKey:  privateKey,
	}, nil
}

func (p *VaultPKI) CreateCA(ID string) (CertificateAuthority, error) {
	return p.createNewCA(ID, true)
}

func (p *VaultPKI) CreateCAWithPrivateKey(ID string) (CertificateAuthority, error) {
	return p.createNewCA(ID, true)
}
