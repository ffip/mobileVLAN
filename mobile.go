package mobile

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	sm2_edch "github.com/emmansun/gmsm/ecdh"
	cfg "github.com/ffip/vlan/config"
	"github.com/ffip/vlan/lib/utils/cert"
	"github.com/ffip/vlan/lib/utils/logs/logger"
	"golang.org/x/crypto/curve25519"
)

// The CIDR type represents an IP address with its corresponding subnet mask.
// @property {string} Ip - The "Ip" property represents the IP address in the CIDR notation. It
// specifies the network address and the number of significant bits in the subnet mask.
// @property {string} MaskCIDR - The MaskCIDR property represents the subnet mask in CIDR notation.
// CIDR stands for Classless Inter-Domain Routing and is a method used to allocate IP addresses and
// define network boundaries. The MaskCIDR property specifies the number of bits in the subnet mask.
// For example, a MaskCIDR
// @property {int} MaskSize - The MaskSize property represents the size of the network mask in bits. It
// indicates the number of bits in the network portion of the IP address.
// @property {string} Network - The "Network" property represents the network address of the CIDR
// (Classless Inter-Domain Routing) block.
type CIDR struct {
	Ip       string
	MaskCIDR string
	MaskSize int
	Network  string
}

// The above code defines two types, Validity and RawCert, which are used to represent the validity and
// raw certificate information respectively.
// @property {bool} Valid - The "Valid" property is a boolean value that indicates whether the
// certificate is valid or not. If it is set to true, it means the certificate is valid. If it is set
// to false, it means the certificate is not valid.
// @property {string} Reason - The "Reason" property is a string that represents the reason for the
// validity of the certificate. It can be used to provide additional information or explanation about
// the validity status of the certificate.
type Validity struct {
	Valid  bool
	Reason string
}

// The RawCert type represents a certificate with its raw data, parsed certificate object, and validity
// information.
// @property {string} RawCert - A string that represents the raw certificate data. This is typically in
// the form of a PEM-encoded string.
// @property Cert - The Cert property is a pointer to a certificate object of type cert.Certificate.
// @property {Validity} Validity - The Validity property represents the validity period of the
// certificate. It typically includes the start and end dates of the certificate's validity.
type RawCert struct {
	RawCert  string
	Cert     *cert.Certificate
	Validity Validity
}

// The KeyPair type represents a pair of public and private keys.
// @property {string} PublicKey - The PublicKey property is a string that represents the public key of
// a cryptographic key pair. It is typically used for encryption or verifying digital signatures.
// @property {string} PrivateKey - The PrivateKey property is a string that represents the private key
// of a cryptographic key pair. It is typically used for encryption and decryption operations.
type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

// The function `GetConfigSetting` retrieves a specific setting from a given configuration data string.
func GetConfigSetting(configData string, setting string) string {
	// We don't want to leak the config into the system logs
	l := logger.New(1000)
	l.SetOutput(io.Discard)

	c := cfg.NewC(l)
	c.LoadString(configData)
	return c.GetString(setting, "")
}

// The function `ParseCIDR` takes a CIDR string, parses it, and returns the IP address, mask CIDR, mask
// size, and network address.
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

// The function `GenerateKeyPair` generates a key pair for a specified elliptic curve and returns it as
// a JSON string.
func GenerateKeyPair(curve string) (result string, err error) {
	var pub, priv []byte
	switch curve {
	case "25519", "X25519":
		pub, priv = x25519Keypair()
	case "P256":
		pub, priv = p256Keypair()
	case "SM2", "GM":
		pub, priv = sm2Keypair()
	default:
		err = fmt.Errorf("invalid curve: %s", curve)
		return
	}

	kp := KeyPair{}
	kp.PublicKey = string(cert.MarshalX25519PublicKey(pub))
	kp.PrivateKey = string(cert.MarshalX25519PrivateKey(priv))

	rawJson, err := json.Marshal(kp)
	if err != nil {
		return result, err
	}

	return string(rawJson), nil
}

// The function generates a key pair for the X25519 elliptic curve Diffie-Hellman algorithm.
func x25519Keypair() ([]byte, []byte) {
	privkey := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, privkey); err != nil {
		panic(err)
	}

	pubkey, err := curve25519.X25519(privkey, curve25519.Basepoint)
	if err != nil {
		panic(err)
	}

	return pubkey, privkey
}

// The function generates a P256 key pair and returns the public and private keys as byte arrays.
func p256Keypair() ([]byte, []byte) {
	privkey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	pubkey := privkey.PublicKey()
	return pubkey.Bytes(), privkey.Bytes()
}

// The function `sm2Keypair` generates a pair of SM2 elliptic curve key (public and private) and
// returns them as byte arrays.
func sm2Keypair() ([]byte, []byte) {
	privkey, err := sm2_edch.P256().GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	pubkey := privkey.PublicKey()
	return pubkey.Bytes(), privkey.Bytes()
}

// The function `VerifyCertAndKey` verifies if a given certificate and private key match.
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
