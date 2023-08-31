package mobile

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	cfg "github.com/ffip/vlan/config"
	"github.com/ffip/vlan/lib/utils/cert"
	"github.com/ffip/vlan/lib/utils/logs/logger"
	"golang.org/x/crypto/curve25519"
)

type m map[string]any

type CIDR struct {
	Ip       string
	MaskCIDR string
	MaskSize int
	Network  string
}

type Validity struct {
	Valid  bool
	Reason string
}

type RawCert struct {
	RawCert  string
	Cert     *cert.Certificate
	Validity Validity
}

type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

func GetConfigSetting(configData string, setting string) string {
	// We don't want to leak the config into the system logs
	l := logger.New(1000)
	l.SetOutput(io.Discard)

	c := cfg.NewC(l)
	c.LoadString(configData)
	return c.GetString(setting, "")
}

func ParseCIDR(cidr string) (*CIDR, error) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	size, _ := ipNet.Mask.Size()

	return &CIDR{
		Ip:       ip.String(),
		MaskCIDR: fmt.Sprintf("%d.%d.%d.%d", ipNet.Mask[0], ipNet.Mask[1], ipNet.Mask[2], ipNet.Mask[3]),
		MaskSize: size,
		Network:  ipNet.IP.String(),
	}, nil
}

// Returns a JSON representation of 1 or more certificates
func ParseCerts(rawStringCerts string) (string, error) {
	var certs []RawCert
	var c *cert.Certificate
	var err error
	rawCerts := []byte(rawStringCerts)

	for {
		c, rawCerts, err = cert.UnmarshalCertificateFromPEM(rawCerts)
		if err != nil {
			return "", err
		}

		rawCert, err := c.MarshalToPEM()
		if err != nil {
			return "", err
		}

		rc := RawCert{
			RawCert: string(rawCert),
			Cert:    c,
			Validity: Validity{
				Valid: true,
			},
		}

		if c.Expired(time.Now()) {
			rc.Validity.Valid = false
			rc.Validity.Reason = "Certificate is expired"
		}

		if rc.Validity.Valid && c.Details.IsCA && !c.CheckSignature(c.Details.PublicKey) {
			rc.Validity.Valid = false
			rc.Validity.Reason = "Certificate signature did not match"
		}

		certs = append(certs, rc)

		if rawCerts == nil || strings.TrimSpace(string(rawCerts)) == "" {
			break
		}
	}

	rawJson, err := json.Marshal(certs)
	if err != nil {
		return "", err
	}

	return string(rawJson), nil
}

func GenerateKeyPair() (string, error) {
	pub, priv, err := x25519Keypair()
	if err != nil {
		return "", err
	}

	kp := KeyPair{}
	kp.PublicKey = string(cert.MarshalX25519PublicKey(pub))
	kp.PrivateKey = string(cert.MarshalX25519PrivateKey(priv))

	rawJson, err := json.Marshal(kp)
	if err != nil {
		return "", err
	}

	return string(rawJson), nil
}

func x25519Keypair() ([]byte, []byte, error) {
	var pubkey, privkey [32]byte
	if _, err := io.ReadFull(rand.Reader, privkey[:]); err != nil {
		return nil, nil, err
	}
	curve25519.ScalarBaseMult(&pubkey, &privkey)
	return pubkey[:], privkey[:], nil
}

func VerifyCertAndKey(rawCert string, pemPrivateKey string) (bool, error) {
	rawKey, _, err := cert.UnmarshalX25519PrivateKey([]byte(pemPrivateKey))
	if err != nil {
		return false, fmt.Errorf("error while unmarshaling private key: %s", err)
	}

	cert, _, err := cert.UnmarshalCertificateFromPEM([]byte(rawCert))
	if err != nil {
		return false, fmt.Errorf("error while unmarshaling cert: %s", err)
	}

	if err = cert.VerifyPrivateKey(cert.Details.Curve, rawKey); err != nil {
		return false, err
	}

	return true, nil
}
