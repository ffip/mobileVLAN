package mobile

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	cfg "github.com/ffip/vlan/config"
	"github.com/ffip/vlan/entry"
	"github.com/ffip/vlan/lib/utils/caution"
	"github.com/ffip/vlan/lib/utils/cert"
	"github.com/ffip/vlan/lib/utils/logs/logger"
	"github.com/ffip/vlan/lib/yaml"
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

func RenderConfig(configData string, key string) (string, error) {
	var d m

	err := json.Unmarshal([]byte(configData), &d)
	if err != nil {
		return "", err
	}

	// If this is a managed config, go ahead and return it
	// checkConfig(c, l.WithField("extention", "sync").WithField("function", "config"))
	// go func() {
	// 	for {
	// 		time.Sleep(time.Duration(c.GetDuration("sync.interval", 15*time.Minute)))
	// 		go checkConfig(c, l.WithField("extention", "sync").WithField("function", "config"))
	// 	}
	// }()

	// Otherwise, build the config
	cfg := newConfig()
	cfg.PKI.CA, _ = d["ca"].(string)
	cfg.PKI.Cert, _ = d["cert"].(string)
	cfg.PKI.Key = key

	i, _ := d["port"].(float64)
	cfg.Listen.Port = int(i)

	cfg.Cipher, _ = d["cipher"].(string)
	// Log verbosity is not required
	if val, _ := d["logVerbosity"].(string); val != "" {
		cfg.Logging.Level = val
	}

	i, _ = d["lhDuration"].(float64)
	cfg.Tower.Interval = int(i)

	if i, ok := d["mtu"].(float64); ok {
		mtu := int(i)
		cfg.Tun.Mtu = mtu
	}

	points := d["points"].(map[string]any)
	for nebIp, mapping := range points {
		hosts := mapping.([]any)

		realHosts := make([]string, len(hosts))

		for i, h := range hosts {
			realHosts[i] = h.(string)
		}

		cfg.Points[nebIp] = realHosts
	}

	if routeTable, ok := d["routeTable"].([]any); ok {
		cfg.Tun.RouteTable = make([]RouteTable, len(routeTable))
		for i, r := range routeTable {
			rawRoute := r.(map[string]any)
			route := &cfg.Tun.RouteTable[i]
			route.Route = rawRoute["route"].(string)
			route.Via = rawRoute["via"].(string)
		}
	}

	finalConfig, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}

	return string(finalConfig), nil
}

func TestConfig(configData string, key string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	yamlConfig, err := RenderConfig(configData, key)
	if err != nil {
		return err
	}

	// We don't want to leak the config into the system logs
	l := logger.New(1000)
	l.SetOutput(bytes.NewBuffer([]byte{}))

	c := cfg.NewC(l)
	err = c.LoadString(yamlConfig)
	if err != nil {
		return fmt.Errorf("failed to load config: %s", err)
	}

	_, err = entry.Main(c, true, "", l, nil)
	if err != nil {
		caution.LogWithContextIfNeeded("Failed to start", err, l)
		return err
	}
	return nil
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
