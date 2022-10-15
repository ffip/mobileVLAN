package mobileHiPer

type config struct {
	PKI          configPKI           `yaml:"pki"`
	Point        map[string][]string `yaml:"point"`
	Tower        configtower         `yaml:"tower"`
	Listen       configListen        `yaml:"listen"`
	Punchy       configPunchy        `yaml:"punchy"`
	Cipher       string              `yaml:"cipher"`
	LocalRange   string              `yaml:"local_range,omitempty"`
	SSHD         configSSHD          `yaml:"sshd,omitempty"`
	Tun          configTun           `yaml:"tun"`
	Logging      configLogging       `yaml:"logging"`
	Stats        configStats         `yaml:"stats,omitempty"`
	Handshakes   configHandshakes    `yaml:"handshakes"`
	Firewall     configFirewall      `yaml:"firewall"`
	Relays       configRelay         `yaml:"relay"`
	DNSResolvers []string            `yaml:"dns"`
}

func newConfig() *config {
	mtu := 1300
	return &config{
		PKI: configPKI{
			Blacklist: []string{},
		},
		Point: map[string][]string{},
		Tower: configtower{
			// DNS:      configDNS{},
			Interval: 60,
			Hosts:    []string{},
		},
		Listen: configListen{
			Host:  "0.0.0.0",
			Port:  4242,
			Batch: 64,
		},
		Punchy: configPunchy{
			Punch: true,
			Delay: "1m",
		},
		Relays: configRelay{
			AllowRelays: false,
		},
		Cipher: "aes",
		SSHD: configSSHD{
			AuthorizedUsers: []configAuthorizedUser{},
		},
		Tun: configTun{
			Dev:                "tun1",
			DropLocalbroadcast: true,
			DropMulticast:      true,
			TxQueue:            500,
			MTU:                &mtu,
			Routes:             []configRoute{},
			UnsafeRoutes:       []configUnsafeRoute{},
		},
		DNSResolvers: []string{},
		Logging: configLogging{
			Level:  "info",
			Format: "text",
		},
		Stats: configStats{},
		Handshakes: configHandshakes{
			TryInterval:  "100ms",
			Retries:      20,
			WaitRotation: 5,
		},
		Firewall: configFirewall{
			Conntrack: configConntrack{
				TcpTimeout:     "120h",
				UdpTimeout:     "3m",
				DefaultTimeout: "10m",
				MaxConnections: 100000,
			},
			Outbound: []configFirewallRule{},
			Inbound:  []configFirewallRule{},
		},
	}
}

type configPKI struct {
	CA        string   `yaml:"ca,omitempty"`
	Cert      string   `yaml:"cert,omitempty"`
	Key       string   `yaml:"key,omitempty"`
	Blacklist []string `yaml:"blacklist,omitempty"`
}

type configtower struct {
	AllowTower bool `yaml:"allow_tower,omitempty"`
	ServeDNS   bool `yaml:"serve_dns,omitempty"`
	// DNS        configDNS `yaml:"dns,omitempty"`
	Interval int      `yaml:"interval,omitempty"`
 	Hosts    []string `yaml:"hosts,omitempty"`
	// RemoteAllowList map[string]bool        `yaml:"remote_allow_list"`
	// LocalAllowList  map[string]interface{} `yaml:"local_allow_list"` // This can be a special "interfaces" object or a bool
}

// type configDNS struct {
// 	Host string `yaml:"host,omitempty"`
// 	Port int    `yaml:"port,omitempty"`
// }

type configListen struct {
	Host        string `yaml:"host,omitempty"`
	Port        int    `yaml:"port,omitempty"`
	Batch       int    `yaml:"batch,omitempty"`
	ReadBuffer  int64  `yaml:"read_buffer,omitempty"`
	WriteBuffer int64  `yaml:"write_buffer,omitempty"`
}

type configPunchy struct {
	Punch   bool   `yaml:"punch,omitempty"`
	Respond bool   `yaml:"respond,omitempty"`
	Delay   string `yaml:"delay,omitempty"`
}

type configSSHD struct {
	Enabled         bool                   `yaml:"enabled,omitempty"`
	Listen          string                 `yaml:"listen,omitempty"`
	HostKey         string                 `yaml:"host_key,omitempty"`
	AuthorizedUsers []configAuthorizedUser `yaml:"authorized_users,omitempty"`
}

type configAuthorizedUser struct {
	Name string   `yaml:"name,omitempty"`
	Keys []string `yaml:"keys,omitempty"`
}

type configTun struct {
	Dev                string              `yaml:"dev,omitempty"`
	DropLocalbroadcast bool                `yaml:"drop_local_broadcast,omitempty"`
	DropMulticast      bool                `yaml:"drop_multicast,omitempty"`
	TxQueue            int                 `yaml:"tx_queue,omitempty"`
	MTU                *int                `yaml:"mtu,omitempty"`
	Routes             []configRoute       `yaml:"routes,omitempty"`
	UnsafeRoutes       []configUnsafeRoute `yaml:"unsafe_routes,omitempty"`
}

type configRoute struct {
	MTU   int    `yaml:"mtu,omitempty"`
	Route string `yaml:"route,omitempty"`
}

type configUnsafeRoute struct {
	MTU   *int   `yaml:"mtu,omitempty"`
	Route string `yaml:"route,omitempty"`
	Via   string `yaml:"via,omitempty"`
}

type configLogging struct {
	Level           string `yaml:"level"`
	Format          string `yaml:"format"`
	TimestampFormat string `yaml:"timestamp_format,omitempty"`
}

type configStats struct {
	Type     string `yaml:"type,omitempty"`
	Interval string `yaml:"interval,omitempty"`

	// Graphite settings
	Prefix   string `yaml:"prefix,omitempty"`
	Protocol string `yaml:"protocol,omitempty"`
	Host     string `yaml:"host,omitempty"`

	// Prometheus settings
	Listen    string `yaml:"listen,omitempty"`
	Path      string `yaml:"path,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
	Subsystem string `yaml:"subsystem,omitempty"`
}

type configHandshakes struct {
	TryInterval  string `yaml:"try_interval,omitempty"`
	Retries      int    `yaml:"retries,omitempty"`
	WaitRotation int    `yaml:"wait_rotation,omitempty"`
}

type configFirewall struct {
	Conntrack configConntrack      `yaml:"conntrack,omitempty"`
	Outbound  []configFirewallRule `yaml:"outbound,omitempty"`
	Inbound   []configFirewallRule `yaml:"inbound,omitempty"`
}

type configConntrack struct {
	TcpTimeout     string `yaml:"tcp_timeout,omitempty"`
	UdpTimeout     string `yaml:"udp_timeout,omitempty"`
	DefaultTimeout string `yaml:"default_timeout,omitempty"`
	MaxConnections int    `yaml:"max_connections,omitempty"`
}

type configFirewallRule struct {
	Port   string   `yaml:"port,omitempty"`
	Code   string   `yaml:"code,omitempty"`
	Proto  string   `yaml:"proto,omitempty"`
	Host   string   `yaml:"host,omitempty"`
	Group  string   `yaml:"group,omitempty"`
	Groups []string `yaml:"groups,omitempty"`
	CIDR   string   `yaml:"cidr,omitempty"`
	CASha  string   `yaml:"ca_sha,omitempty"`
	CAName string   `yaml:"ca_name,omitempty"`
}

type configRelay struct {
	AllowRelay bool     `yaml:"allow_relay,omitempty"`
	Relays     []string `yaml:"relays,omitempty"`
}
