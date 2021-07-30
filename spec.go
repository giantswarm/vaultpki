package vaultpki

import (
	vaultapi "github.com/hashicorp/vault/api"
)

type CertificateAuthority struct {
	Certificate string
	PrivateKey  string
}

func DefaultCertificateAuthority() CertificateAuthority {
	return CertificateAuthority{
		Certificate: "",
		PrivateKey:  "",
	}
}

type Interface interface {
	BackendExists(ID string) (bool, error)
	CAExists(ID string) (bool, error)
	CreateBackend(ID string) error
	CreateCA(ID string) (CertificateAuthority, error)
	CreateCAWithPrivateKey(ID string) (CertificateAuthority, error)
	DeleteBackend(ID string) error
	GetBackend(ID string) (*vaultapi.MountOutput, error)
	GetCACertificate(ID string) (CertificateAuthority, error)
	ListBackends() (map[string]*vaultapi.MountOutput, error)
	UpdateBackend(ID string) error
}
