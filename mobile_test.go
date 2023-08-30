package mobile

import (
	"testing"

	cfg "github.com/ffip/vlan/config"
	"github.com/ffip/vlan/lib/utils/logs/logger"
)

func TestParseCerts(t *testing.T) {
	jsonConfig := `{
    "name": "Debug Test - unsafe",
    "id": "be9d6756-4099-4b25-a901-9d3b773e7d1a",
    "pki": {
        "ca": "-----BEGIN VLAN CERTIFICATE-----\nCj4KDEhpUGVyIFB1YmxpYyjGs6mXBjDG49GrBzog7+h8wZVKgdU4Fh4pwaLekH6D\nn+J8rTcgwNN7YaxcSFJAARJAIEzWZa79d+2RJ+17pay9oEehsV9coLgP72M0XZkw\nff6hHY99VsTLAiXvExd6eYyKRhcriqlr0O7BR+k6/qcqDQ==\n-----END VLAN CERTIFICATE-----\n",
        "cert": "-----BEGIN VLAN CERTIFICATE-----\nCmEKBHN3YXASCYGCnDiAgIDwDyjlgPyXBjDF49GrBzog0UtIu9+bcam6euyq4qJi\nO5PBr4pxuVc4PLWfTGhtVDdKIG7LJZr9vlShnmxQ1IMlsW0lREpZtd0bMFr3UVMv\nxoHlEkDdOgb49QHZKYfCI33ekvAvaM8VczepReCeQNg2vAmk9FXf8IpVKWTJBssA\ng42SwsBAaH1kpZlYZyqyEQxOTUUB\n-----END VLAN CERTIFICATE-----\n",
        "key": "-----BEGIN VLAN X25519 PRIVATE KEY-----\nCUvpfSbxU0EwVTT85NABo/VagsaXKiw2Uft1bF5M0hU=\n-----END VLAN X25519 PRIVATE KEY-----\n",
        "blocklist": [
            "c99d4e650533b92061b09918e838a5a0a6aaee21eed1d12fd937682865936c72"
        ],
        "disconnect_invalid": true,
        "expiry_check": {
            "enabled": true,
            "time_left": "72h",
            "log_interval": "60m"
        }
    },
    "points": {
        "6.6.6.6": [
            "120.92.140.174:35533"
        ],
        "7.7.7.7": [
            "121.62.22.148:35533"
        ],
        "7.7.1.1": [
            "ddns.xiaomckedou233.top:35533"
        ],
        "6.6.1.1": [
            "160.119.69.222:35533"
        ]
    },
    "tower": {
        "service": false,
        "dns": {
            "enable": true,
            "addr": "0.0.0.0",
            "port": 53,
            "interval": 10,
            "mirror": "1.1.1.1",
            "records": [
                {
                    "example.com": "192.168.1.113"
                }
            ]
        },
        "detection_point": {
            "10.0.10.123/24": [
                {
                    "mask": "192.168.1.123/24",
                    "port": 35533
                }
            ]
        },
        "remote_allow_list": {
            "172.16.0.0/12": false,
            "0.0.0.0/0": true,
            "10.0.0.0/8": false,
            "10.42.42.0/24": true
        },
        "remote_allow_ranges": {
            "10.42.42.0/24": {
                "192.168.0.0/16": true
            },
            "10.42.41.0/24": {
                "123.45.0.0/16": true
            }
        },
        "local_allow_list": {
            "interfaces": {
                "tun0": false,
                "docker.*": false,
                "10.0.0.0/8": true
            }
        },
        "advertise_addrs": [
            "1.1.1.1:35533",
            "1.2.3.4:0"
        ]
    },
    "listen": {
        "addr": [
            {
                ":": null
            }
        ],
        "port": 35533,
        "batch": 64,
        "read_buffer": 104857600,
        "write_buffer": 104857600,
        "send_recv_error": "always",
        "routines": 1
    },
    "punchy": {
        "enable": true,
        "frequency": "10s",
        "respond": true,
        "delay": "1s",
        "respond_delay": "5s",
        "preferred_ranges": [
            "172.16.0.0/24"
        ]
    },
    "cipher": "aes",
    "sshd": {
        "enabled": true,
        "port": 22222,
        "point_key": "/etc/vlan/ssh_point_rsa_key",
        "users": [
            {
                "name": "user1",
                "keys": [
                    "ssh-rsa xxxxx",
                    "ssh-ed25519 xxxx"
                ]
            },
            {
                "name": "user2",
                "keys": [
                    "ssh-rsa xxxxx",
                    "ssh-ed25519 xxxx"
                ]
            }
        ]
    },
    "proxy": {
        "socks5": [
            {
                "addr": "0.0.0.0",
                "port": 10800,
                "user": "username",
                "password": "password"
            }
        ],
        "forward": [
            {
                "proto": "tcp",
                "local": "0.0.0.0:3388",
                "remote": "192.168.1.105:3389"
            },
            {
                "proto": "udp",
                "local": "6.6.9.9:65534",
                "remote": "10.1.253.1:35533"
            }
        ]
    },
    "tun": {
        "enable": false,
        "dev": "vlan_network",
        "drop_local_broadcast": false,
        "drop_multicast": false,
        "tx_queue": 5000,
        "mtu": 1500,
        "routes": [
            {
                "mtu": 8800,
                "route": "10.0.0.0/16"
            }
        ],
        "route_table": [
            {
                "route": "172.16.1.0/24",
                "via": "6.6.6.99",
                "mtu": 1500,
                "metric": 100,
                "enable": true
            }
        ]
    },
    "logging": {
        "level": "info",
        "format": "text",
        "disable_timestamp": false,
        "file_path": "/var/log/vlan/vlan",
        "max_size": 20,
        "max_backups": 100,
        "max_age": 30,
        "timestamp_format": "2006-01-02T15:04:05.000Z07:00"
    },
    "stats": {
        "message_metrics": false,
        "tower_metrics": false
    },
    "handshakes": {
        "try_interval": "100ms",
        "retries": 10,
        "trigger_buffer": 64,
        "churn_limiting": true,
        "churn_num_failures": 1,
        "churn_period": "30s"
    },
    "timers": {
        "connection_alive_interval": 5,
        "pending_deletion_interval": 10
    },
    "psk": {
        "mode": "none",
        "keys": null
    },
    "firewall": {
        "outbound_action": "drop",
        "inbound_action": "drop",
        "conntrack": {
            "tcp_timeout": "12m",
            "udp_timeout": "3m",
            "default_timeout": "10m"
        },
        "outbound": [
            {
                "port": "any",
                "proto": "any",
                "point": "any"
            }
        ],
        "inbound": [
            {
                "port": "any",
                "proto": "any",
                "point": "any"
            },
            {
                "port": 443,
                "proto": "tcp",
                "groups": [
                    "laptop",
                    "home"
                ]
            }
        ]
    }
}`
	s, err := RenderConfig(jsonConfig, "")
	if err != nil {
		t.Log(err)
	}
	config := cfg.NewC(logger.New(1000))
	err = config.LoadString(s)

	t.Log(err)
}
