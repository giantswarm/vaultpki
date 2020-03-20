package vaultpki

import (
	"github.com/giantswarm/microerror"
	vaultapi "github.com/hashicorp/vault/api"

	"github.com/giantswarm/vaultpki/key"
)

func (p *VaultPKI) BackendExists(ID string) (bool, error) {
	_, err := p.GetBackend(ID)
	if IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, microerror.Mask(err)
	}

	return true, nil
}

func (p *VaultPKI) CAExists(ID string) (bool, error) {
	_, err := p.GetCACertificate(ID)
	if IsNotFound(err) {
		return false, nil
	} else if err != nil {
		return false, microerror.Mask(err)
	}

	return true, nil
}

func (p *VaultPKI) GetBackend(ID string) (*vaultapi.MountOutput, error) {
	mounts, err := p.vaultClient.Sys().ListMounts()
	if IsNoVaultHandlerDefined(err) {
		return nil, microerror.Maskf(notFoundError, "PKI backend for ID '%s'", ID)
	} else if err != nil {
		return nil, microerror.Mask(err)
	}

	mountOutput, ok := mounts[key.ListMountsPath(ID)]
	if !ok || mountOutput.Type != MountType {
		return nil, microerror.Maskf(notFoundError, "PKI backend for ID '%s'", ID)
	}

	return mountOutput, nil
}

// GetCACertificate returns the public key of the root CA of the PKI backend
// associated to the given ID, if any.
func (p *VaultPKI) GetCACertificate(ID string) (CertificateAuthority, error) {
	secret, err := p.vaultClient.Logical().Read(key.ReadCAPath(ID))
	if IsNoVaultHandlerDefined(err) {
		return DefaultCertificateAuthority(), microerror.Maskf(notFoundError, "root CA for ID '%s'", ID)
	} else if err != nil {
		return DefaultCertificateAuthority(), microerror.Mask(err)
	}

	// If the secret is nil, the CA has not been generated.
	if secret == nil {
		return DefaultCertificateAuthority(), microerror.Maskf(notFoundError, "root CA for ID '%s'", ID)
	}

	var crt string
	{
		v, ok := secret.Data["certificate"]
		if !ok {
			return DefaultCertificateAuthority(), microerror.Maskf(executionFailedError, "certificate missing")
		}
		crt, ok = v.(string)
		if !ok {
			return DefaultCertificateAuthority(), microerror.Maskf(executionFailedError, "certificate must be string")
		}
	}

	certificateAuthority := CertificateAuthority{
		Certificate: crt,
	}

	return certificateAuthority, nil
}

func (p *VaultPKI) ListBackends() (map[string]*vaultapi.MountOutput, error) {
	mounts, err := p.vaultClient.Sys().ListMounts()
	if err != nil {
		return nil, microerror.Mask(err)
	}

	backends := map[string]*vaultapi.MountOutput{}
	for k, v := range mounts {
		if !key.IsMountPath(k) {
			continue
		}

		backends[k] = v
	}

	return backends, nil
}
