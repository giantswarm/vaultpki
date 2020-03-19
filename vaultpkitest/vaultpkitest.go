package vaultpkitest

import (
	"github.com/giantswarm/vaultpki"
	vaultapi "github.com/hashicorp/vault/api"
)

type VaultPKITest struct {
}

func New() *VaultPKITest {
	return &VaultPKITest{}
}

func (p *VaultPKITest) BackendExists(ID string) (bool, error) {
	return false, nil
}

func (p *VaultPKITest) CAExists(ID string) (bool, error) {
	return false, nil
}

func (p *VaultPKITest) CreateBackend(ID string) error {
	return nil
}

func (p *VaultPKITest) CreateCA(ID string) (vaultpki.CertificateAuthority, error) {
	return vaultpki.DefaultCertificateAuthority(), nil
}

func (p *VaultPKITest) CreateCAWithPrivateKey(ID string) (vaultpki.CertificateAuthority, error) {
	return vaultpki.DefaultCertificateAuthority(), nil
}

func (p *VaultPKITest) DeleteBackend(ID string) error {
	return nil
}

func (p *VaultPKITest) GetBackend(ID string) (*vaultapi.MountOutput, error) {
	return nil, nil
}

func (p *VaultPKITest) GetCACertificate(ID string) (vaultpki.CertificateAuthority, error) {
	return vaultpki.DefaultCertificateAuthority(), nil
}

func (p *VaultPKITest) ListBackends() (map[string]*vaultapi.MountOutput, error) {
	return nil, nil
}
