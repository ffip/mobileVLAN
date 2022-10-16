package mobileHiPer

import (
	"testing"

	hc "github.com/ffip/hiper/config"
	"github.com/sirupsen/logrus"
)

func TestParseCerts(t *testing.T) {
	jsonConfig := `{
  "name": "Debug Test - unsafe",
  "id": "be9d6756-4099-4b25-a901-9d3b773e7d1a",
  "point": {
    "6.6.6.6": {
      "tower": true,
      "destinations": [
        "120.92.140.174:65533"
      ]
    }
  },
  "unsafeRoutes": [
    {
      "route": "10.3.3.3/32",
      "via": "10.1.0.1",
      "mtu": null
    },
    {
      "route": "1.1.1.2/32",
      "via": "10.1.0.1",
      "mtu": null
    }
  ],
  "dns": [
    "6.6.6.6",
    "7.7.7.7"
  ],
  "ca": "-----BEGIN HIPER CERTIFICATE-----\nCj4KDEhpUGVyIFB1YmxpYyjGs6mXBjDG49GrBzog7+h8wZVKgdU4Fh4pwaLekH6D\nn+J8rTcgwNN7YaxcSFJAARJAIEzWZa79d+2RJ+17pay9oEehsV9coLgP72M0XZkw\nff6hHY99VsTLAiXvExd6eYyKRhcriqlr0O7BR+k6/qcqDQ==\n-----END HIPER CERTIFICATE-----\n",
  "cert": "-----BEGIN HIPER CERTIFICATE-----\nCmEKBGRlbW8SCYaMmDCAgIDwDyjn3aeaBjD3+aeaBjogsx61tqk5cAXrc1TExMlp\nCcObDPAVLU94jIhe+HzxV0NKIG7LJZr9vlShnmxQ1IMlsW0lREpZtd0bMFr3UVMv\nxoHlEkDeszjkqz37ZVkH7k3iRwsjdcSvg8bGoiQuvRLgSBizzb+pSmpoHiC7+88/\naJaOxAaYqQ0jTF/g/WhhVG+ctBIH\n-----END HIPER CERTIFICATE-----\n",
  "key": "-----BEGIN HIPER X25519 PRIVATE KEY-----\nersCnJjWfCBJl1M02Wtib/nSQccTqO+INnvOyKKsQwA=\n-----END HIPER X25519 PRIVATE KEY-----\n",
  "lhDuration": 7200,
  "port": 65533,
  "mtu": 1300,
  "cipher": "aes",
  "sortKey": 3,
  "logVerbosity": "info"
}`
	s, _ := RenderConfig(jsonConfig, "")

	config := hc.NewC(logrus.New())
	err := config.LoadString(s)

	t.Log(err)
}
