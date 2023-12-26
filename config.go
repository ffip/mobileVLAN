package mobile

// The above type represents a configuration structure with various fields for different settings.
// @property {Sync} Sync - The `Sync` property is of type `Sync` and is used for configuring
// synchronization settings.
// @property {PKI} PKI - PKI stands for Public Key Infrastructure. It is a set of roles, policies, and
// procedures needed to create, manage, distribute, use, store, and revoke digital certificates and
// manage public-key encryption. In the context of the `config` struct, the `PKI` property represents
// the configuration
// @property Points - The `Points` property is a map where the keys are strings and the values are
// slices of strings. It is used to store a collection of points, where each point is represented by a
// string.
// @property {Tower} Tower - The `Tower` property represents the configuration for the tower component.
// It may include settings such as the tower's address, port, and authentication credentials.
// @property {Listen} Listen - The `Listen` property is a configuration for specifying the listening
// address and port for the server. It typically includes the following fields:
// @property {Punchy} Punchy - The `Punchy` property is a struct that contains configuration options
// related to the Punchy service. It may include settings such as the Punchy server address, port, and
// other relevant parameters.
// @property {SSH} ssh - The `SSH` property represents the configuration for the SSH daemon. It
// includes settings such as the SSH port, allowed users, and authentication methods.
// @property {Proxy} Proxy - The `Proxy` property in the `config` struct represents the configuration
// for the proxy settings. It includes settings such as the proxy type, address, and authentication
// credentials.
// @property {Tun} Tun - The `Tun` property in the `config` struct represents the configuration for the
// tunnel interface. It includes settings such as the IP address, subnet mask, and MTU (Maximum
// Transmission Unit) for the tunnel interface.
// @property {Logging} Logging - The `Logging` property is a struct that contains configuration options
// for logging. It may include properties such as log level, log file path, log rotation settings, etc.
// @property {Stats} Stats - The `Stats` property is a struct that contains configuration options
// related to statistics and monitoring. It may include properties such as:
// @property {Handshakes} Handshakes - The `Handshakes` property is a struct that contains
// configuration options related to handshakes. It may include properties such as `Timeout`,
// `MaxRetries`, `RetryInterval`, etc. These properties define how handshakes are handled in the
// application.
// @property {Timers} Timers - The `Timers` property is a struct that contains various timer
// configurations. It is used to define the timing settings for different operations within the
// application.
// @property {PSK} PSK - The `PSK` property in the `config` struct represents the Pre-Shared Key
// configuration. It is used for authentication and encryption purposes in a network communication. The
// `PSK` struct may contain additional fields that define the specific configuration for the Pre-Shared
// Key.
// @property {Firewall} Firewall - The `Firewall` property is a configuration for firewall settings. It
// specifies rules and settings related to network traffic filtering and security.
// @property {string} Cipher - The `Cipher` property is a string that specifies the encryption cipher
// to be used. It is used to encrypt and decrypt data during communication.
type config struct {
	Sync       Sync                `json:"sync,omitempty" yaml:"sync,omitempty"`
	PKI        PKI                 `json:"pki,omitempty" yaml:"pki,omitempty"`
	Points     map[string][]string `json:"points,omitempty" yaml:"points,omitempty"`
	Tower      Tower               `json:"tower,omitempty" yaml:"tower,omitempty"`
	Listen     Listen              `json:"listen,omitempty" yaml:"listen,omitempty"`
	Punchy     Punchy              `json:"punchy,omitempty" yaml:"punchy,omitempty"`
	SSH        ssh                 `json:"ssh,omitempty" yaml:"ssh,omitempty"`
	Proxy      Proxy               `json:"proxy,omitempty" yaml:"proxy,omitempty"`
	Tun        Tun                 `json:"tun,omitempty" yaml:"tun,omitempty"`
	Logging    Logging             `json:"logging,omitempty" yaml:"logging,omitempty"`
	Stats      Stats               `json:"stats,omitempty" yaml:"stats,omitempty"`
	Handshakes Handshakes          `json:"handshakes,omitempty" yaml:"handshakes,omitempty"`
	Timers     Timers              `json:"timers,omitempty" yaml:"timers,omitempty"`
	PSK        PSK                 `json:"psk,omitempty" yaml:"psk,omitempty"`
	Firewall   Firewall            `json:"firewall,omitempty" yaml:"firewall,omitempty"`
	Cipher     string              `json:"cipher,omitempty" yaml:"cipher,omitempty"`
}

// The Sync type represents a synchronization configuration with various properties.
// @property {bool} Enable - A boolean value indicating whether synchronization is enabled or not.
// @property {bool} Persistent - The `Persistent` property indicates whether the synchronization should
// be persistent or not. If set to `true`, the synchronization will persist even after the application
// restarts. If set to `false`, the synchronization will only occur during the current session and will
// not persist after the application restarts.
// @property {string} Interval - The `Interval` property specifies the time interval at which the
// synchronization should occur. It is a string value that represents a duration, such as "1h" for 1
// hour or "30s" for 30 seconds.
// @property {string} Source - The `Source` property represents the source from where the data will be
// synchronized. It could be a file path, a database connection string, or any other source of data.
// @property {string} Store - The `Store` property is used to specify the location where the
// synchronized data will be stored.
// @property {string} Addition - The "Addition" property is an optional field that can be used to
// provide additional information or configuration for the synchronization process.
type Sync struct {
	Enable     bool   `json:"enable,omitempty" yaml:"enable,omitempty"`
	Persistent bool   `json:"persistent,omitempty" yaml:"persistent,omitempty"`
	Interval   string `json:"interval,omitempty" yaml:"interval,omitempty"`
	Source     string `json:"source,omitempty" yaml:"source,omitempty"`
	Store      string `json:"store,omitempty" yaml:"store,omitempty"`
	Addition   string `json:"addition,omitempty" yaml:"addition,omitempty"`
}

// The ExpiryCheck type represents an expiry check with optional fields for enabling/disabling, time
// left, and log interval.
// @property {bool} Enabled - The "Enabled" property is a boolean value that indicates whether the
// expiry check is enabled or not. If it is set to true, it means the expiry check is enabled. If it is
// set to false, it means the expiry check is disabled.
// @property {string} TimeLeft - The `TimeLeft` property represents the remaining time until expiry. It
// is a string type and is used to store the time left until expiration in a human-readable format,
// such as "2 days", "1 hour", etc.
// @property {string} LogInterval - The `LogInterval` property is a string that represents the interval
// at which the expiry check should log information.
type ExpiryCheck struct {
	Enabled     bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	TimeLeft    string `json:"time_left,omitempty" yaml:"time_left,omitempty"`
	LogInterval string `json:"log_interval,omitempty" yaml:"log_interval,omitempty"`
}

// The PKI type represents a set of properties related to Public Key Infrastructure.
// @property {string} CA - The CA property represents the Certificate Authority (CA) used for the
// Public Key Infrastructure (PKI). The CA is responsible for issuing and managing digital
// certificates.
// @property {string} Cert - The `Cert` property in the `PKI` struct represents the certificate file
// used in a Public Key Infrastructure (PKI) system. It is a string that contains the path or content
// of the certificate file.
// @property {string} Key - The `Key` property in the `PKI` struct represents the private key
// associated with the PKI (Public Key Infrastructure). It is used for cryptographic operations such as
// signing and decrypting data.
// @property {[]string} Blocklist - The `Blocklist` property is a list of strings that represents a
// blocklist of certificates. These certificates are considered invalid and will be rejected or ignored
// by the PKI system.
// @property {bool} DisconnectInvalid - The `DisconnectInvalid` property is a boolean value that
// determines whether to disconnect clients with invalid certificates. If set to `true`, clients with
// invalid certificates will be disconnected. If set to `false`, clients with invalid certificates will
// not be disconnected.
// @property {ExpiryCheck} ExpiryCheck - ExpiryCheck is a struct that contains properties related to
// checking the expiry of certificates in the PKI (Public Key Infrastructure).
type PKI struct {
	CA                string      `json:"ca,omitempty" yaml:"ca,omitempty"`
	Cert              string      `json:"cert,omitempty" yaml:"cert,omitempty"`
	Key               string      `json:"key,omitempty" yaml:"key,omitempty"`
	Blocklist         []string    `json:"blocklist,omitempty" yaml:"blocklist,omitempty"`
	DisconnectInvalid bool        `json:"disconnect_invalid,omitempty" yaml:"disconnect_invalid,omitempty"`
	ExpiryCheck       ExpiryCheck `json:"expiry_check,omitempty" yaml:"expiry_check,omitempty"`
}

// The DNS type represents DNS configuration settings including enable/disable, address, port,
// interval, mirror, and records.
// @property {bool} Enable - A boolean value indicating whether the DNS is enabled or not.
// @property {string} Addr - The `Addr` property in the `DNS` struct represents the IP address or
// hostname of the DNS server.
// @property {int} Port - The `Port` property in the `DNS` struct represents the port number on which
// the DNS server will listen for incoming DNS requests.
// @property {int} Interval - The `Interval` property represents the time interval in seconds at which
// the DNS should be checked or updated.
// @property {string} Mirror - The "Mirror" property in the DNS struct represents the URL or address of
// a mirror server. A mirror server is a duplicate server that contains the same content as the
// original server. It is used as a backup in case the original server becomes unavailable.
// @property Records - The "Records" property is a map that stores DNS records. Each record consists of
// a key-value pair, where the key is the domain name and the value is the corresponding IP address.
type DNS struct {
	Enable   bool              `json:"enable,omitempty" yaml:"enable,omitempty"`
	Addr     string            `json:"addr,omitempty" yaml:"addr,omitempty"`
	Port     int               `json:"port,omitempty" yaml:"port,omitempty"`
	Interval int               `json:"interval,omitempty" yaml:"interval,omitempty"`
	Mirror   string            `json:"mirror,omitempty" yaml:"mirror,omitempty"`
	Records  map[string]string `json:"records,omitempty" yaml:"records,omitempty"`
}

// The Tower type represents a configuration for a service with DNS settings, interval, detection
// points, remote allow lists, local allow lists, and advertise addresses.
// @property {bool} Service - The `Service` property is a boolean value that indicates whether the
// tower is a service or not. If it is set to `true`, it means the tower is a service. If it is set to
// `false` or omitted, it means the tower is not a service.
// @property {DNS} DNS - The `DNS` property is a struct that contains information related to DNS
// configuration. It may include properties such as `Nameservers`, `SearchDomains`, `Options`, etc.
// @property {int} Interval - The `Interval` property in the `Tower` struct represents the time
// interval in seconds at which certain actions or checks should be performed. It specifies the
// frequency at which the tower should perform its tasks or operations.
// @property DetectionPoint - The `DetectionPoint` property is a nested map that represents the
// detection points for the tower. It has the following structure:
// @property RemoteAllowList - The `RemoteAllowList` property is a map where the keys are strings
// representing remote addresses and the values are booleans indicating whether the remote address is
// allowed or not. It is used to specify a list of remote addresses that are allowed to access the
// tower.
// @property RemoteAllowRanges - The `RemoteAllowRanges` property is a map that contains ranges of
// remote IP addresses that are allowed. Each key in the map represents a range of IP addresses, and
// the corresponding value is a map that contains specific IP addresses within that range that are
// allowed.
// @property LocalAllowList - The `LocalAllowList` property is a map where the keys are strings and the
// values can be of any type. It is used to specify a list of allowed local addresses.
// @property {[]string} AdvertiseAddrs - AdvertiseAddrs is a slice of strings that represents the
// addresses that the Tower should advertise for incoming connections.
type Tower struct {
	Service           bool                       `json:"service,omitempty" yaml:"service,omitempty"`
	DNS               DNS                        `json:"dns,omitempty" yaml:"dns,omitempty"`
	Interval          int                        `json:"interval,omitempty" yaml:"interval,omitempty"`
	DetectionPoint    map[string]map[string]any  `json:"detection_point,omitempty" yaml:"detection_point,omitempty"`
	RemoteAllowList   map[string]bool            `json:"remote_allow_list,omitempty" yaml:"remote_allow_list,omitempty"`
	RemoteAllowRanges map[string]map[string]bool `json:"remote_allow_ranges,omitempty" yaml:"remote_allow_ranges,omitempty"`
	LocalAllowList    map[string]any             `json:"local_allow_list,omitempty" yaml:"local_allow_list,omitempty"`
	AdvertiseAddrs    []string                   `json:"advertise_addrs,omitempty" yaml:"advertise_addrs,omitempty"`
}

// The `Listen` type represents the configuration for a listening server.
// @property {string} Addr - The `Addr` property represents the address on which the server should
// listen for incoming connections. It is a string value.
// @property {int} Port - The `Port` property represents the port number on which the server will
// listen for incoming connections.
// @property {int} Batch - The `Batch` property specifies the number of messages that can be processed
// in a single batch. It determines how many messages can be read from the input buffer and processed
// together before sending them to the output buffer.
// @property {int} ReadBuffer - The `ReadBuffer` property specifies the size of the buffer used for
// reading data from the network connection. It determines how much data can be read at once from the
// connection.
// @property {int} WriteBuffer - The `WriteBuffer` property specifies the size of the write buffer in
// bytes. It determines the amount of data that can be written to the network connection at once.
// @property {string} SendRecvError - The `SendRecvError` property is a string that represents an error
// message related to sending or receiving data. It is used to handle any errors that occur during the
// communication process.
// @property {int} Routines - The `Routines` property in the `Listen` struct represents the number of
// goroutines that will be used to handle incoming requests or connections. Goroutines are lightweight
// threads managed by the Go runtime, and they are used to achieve concurrency in Go programs. By
// specifying the number of routines, you can
type Listen struct {
	Addr          string `json:"addr,omitempty" yaml:"addr,omitempty"`
	Port          int    `json:"port,omitempty" yaml:"port,omitempty"`
	Batch         int    `json:"batch,omitempty" yaml:"batch,omitempty"`
	ReadBuffer    int    `json:"read_buffer,omitempty" yaml:"read_buffer,omitempty"`
	WriteBuffer   int    `json:"write_buffer,omitempty" yaml:"write_buffer,omitempty"`
	SendRecvError string `json:"send_recv_error,omitempty" yaml:"send_recv_error,omitempty"`
	Routines      int    `json:"routines,omitempty" yaml:"routines,omitempty"`
}

// The Punchy type is a struct that represents a configuration for a feature called Punchy, with
// various properties such as enable/disable, frequency, response settings, delay settings, and
// preferred ranges.
// @property {bool} Enable - A boolean value indicating whether the "Punchy" feature is enabled or not.
// @property {string} Frequency - The `Frequency` property in the `Punchy` struct represents the
// frequency at which the punchy feature is enabled. It is a string value that can be set to different
// values depending on the desired frequency, such as "daily", "weekly", "monthly", etc.
// @property {bool} Respond - The "Respond" property is a boolean value that determines whether the
// Punchy feature should respond to incoming requests or not. If set to true, Punchy will respond to
// incoming requests. If set to false, Punchy will not respond to incoming requests.
// @property {string} Delay - The `Delay` property in the `Punchy` struct represents the delay between
// each punch. It is of type string and is used to specify the duration of the delay.
// @property {string} RespondDelay - The `RespondDelay` property is a string that represents the delay
// before responding to a punch. It is used in the `Punchy` struct.
// @property {[]string} PreferredRanges - PreferredRanges is a slice of strings that represents a list
// of preferred ranges. Each string in the slice represents a preferred range.
type Punchy struct {
	Enable          bool     `json:"enable,omitempty" yaml:"enable,omitempty"`
	Frequency       string   `json:"frequency,omitempty" yaml:"frequency,omitempty"`
	Respond         bool     `json:"respond,omitempty" yaml:"respond,omitempty"`
	Delay           string   `json:"delay,omitempty" yaml:"delay,omitempty"`
	RespondDelay    string   `json:"respond_delay,omitempty" yaml:"respond_delay,omitempty"`
	PreferredRanges []string `json:"preferred_ranges,omitempty" yaml:"preferred_ranges,omitempty"`
}

// The Users type represents a user with a name and a list of keys.
// @property {string} Name - The `Name` property is a string that represents the name of a user. It is
// tagged with `json:"name,omitempty"` and `yaml:"name,omitempty"`, which means that when the struct is
// serialized to JSON or YAML, the `Name` field will be included only if it is not
// @property {[]string} Keys - The "Keys" property is a slice of strings. It represents a collection of
// keys associated with a user.
type Users struct {
	Name string   `json:"name,omitempty" yaml:"name,omitempty"`
	Keys []string `json:"keys,omitempty" yaml:"keys,omitempty"`
}

// The ssh type represents the configuration for an SSH server, including its enabled status, port
// number, encryption key, and a list of users.
// @property {bool} Enabled - The "Enabled" property is a boolean value that indicates whether ssh is
// enabled or not. If it is set to true, ssh is enabled. If it is set to false, ssh is disabled.
// @property {int} Port - The `Port` property represents the port number on which the ssh service is
// running.
// @property {string} PointKey - The "PointKey" property in the ssh struct represents the SSH public
// key used for authentication.
// @property {[]Users} Users - The `Users` property is an array of `Users` objects.
type ssh struct {
	Enabled  bool    `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Port     int     `json:"port,omitempty" yaml:"port,omitempty"`
	PointKey string  `json:"point_key,omitempty" yaml:"point_key,omitempty"`
	Users    []Users `json:"users,omitempty" yaml:"users,omitempty"`
}

// The Socks5 type represents a configuration for a Socks5 proxy server.
// @property {string} Addr - The `Addr` property represents the address of the SOCKS5 server. It is a
// string that specifies the IP address or hostname of the server.
// @property {int} Port - The `Port` property is an integer that represents the port number for the
// SOCKS5 server.
// @property {string} User - The "User" property in the Socks5 struct represents the username used for
// authentication when connecting to the SOCKS5 proxy server.
// @property {string} Password - The `Password` property is a string that represents the password used
// for authentication when connecting to the SOCKS5 server.
type Socks5 struct {
	Addr     string `json:"addr,omitempty" yaml:"addr,omitempty"`
	Port     int    `json:"port,omitempty" yaml:"port,omitempty"`
	User     string `json:"user,omitempty" yaml:"user,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

// The Forward type is used to define a forwarding configuration with protocol, local address, and
// remote address.
// @property {string} Proto - The `Proto` property is a string that represents the protocol used for
// forwarding. It can be any valid protocol such as TCP, UDP, HTTP, etc.
// @property {string} Local - The "Local" property in the "Forward" struct represents the local address
// or endpoint to which the incoming network traffic will be forwarded.
// @property {string} Remote - The "Remote" property is a string that represents the remote address or
// destination for the forward. It is used to specify the address or hostname of the remote server or
// device that the forward should be directed to.
type Forward struct {
	Proto  string `json:"proto,omitempty" yaml:"proto,omitempty"`
	Local  string `json:"local,omitempty" yaml:"local,omitempty"`
	Remote string `json:"remote,omitempty" yaml:"remote,omitempty"`
}

// The Proxy type is a struct that contains slices of Socks5 and Forward structs.
// @property {[]Socks5} Socks5 - The `Socks5` property is an array of `Socks5` objects. Each `Socks5`
// object represents a SOCKS5 proxy configuration.
// @property {[]Forward} Forward - The `Forward` property is an array of `Forward` objects. Each
// `Forward` object represents a forwarding configuration.
type Proxy struct {
	Socks5  []Socks5  `json:"socks5,omitempty" yaml:"socks5,omitempty"`
	Forward []Forward `json:"forward,omitempty" yaml:"forward,omitempty"`
}

// The "Routes" type represents a network route with an optional maximum transmission unit (MTU) value.
// @property {int} Mtu - The `Mtu` property represents the Maximum Transmission Unit (MTU) for a
// network interface. It specifies the maximum size of a data packet that can be transmitted over the
// network.
// @property {string} Route - The `Route` property is a string that represents a specific route. It is
// used to define the destination for network traffic.
type Routes struct {
	Mtu   int    `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	Route string `json:"route,omitempty" yaml:"route,omitempty"`
}

// The RouteTable type represents a route with its associated properties.
// @property {string} Route - The "Route" property represents the destination network or IP address for
// the route. It specifies where the traffic should be directed to.
// @property {string} Via - The "Via" property in the RouteTable struct represents the next hop or
// gateway IP address for the specified route. It is used to determine the next network hop for
// forwarding packets.
// @property {int} Mtu - The "Mtu" property in the RouteTable struct represents the Maximum
// Transmission Unit (MTU) value for the route. MTU is the maximum size of a data packet that can be
// transmitted over a network. It specifies the size of the largest protocol data unit (PDU) that can
// be
// @property {int} Metric - The "Metric" property in the RouteTable struct represents the metric value
// associated with the route. It is used to determine the priority of the route when multiple routes
// are available for a destination. A lower metric value indicates a higher priority.
// @property {bool} Enable - The "Enable" property is a boolean value that indicates whether the route
// is enabled or disabled. If set to true, the route is enabled. If set to false, the route is
// disabled.
type RouteTable struct {
	Route  string `json:"route,omitempty" yaml:"route,omitempty"`
	Via    string `json:"via,omitempty" yaml:"via,omitempty"`
	Mtu    int    `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	Metric int    `json:"metric,omitempty" yaml:"metric,omitempty"`
	Enable bool   `json:"enable,omitempty" yaml:"enable,omitempty"`
}

// The type Tun represents a network tunnel configuration.
// @property {bool} Enable - The `Enable` property is a boolean value that indicates whether the Tun
// interface is enabled or not. If set to `true`, the Tun interface is enabled. If set to `false`, the
// Tun interface is disabled.
// @property {string} Dev - The `Dev` property in the `Tun` struct represents the network device name
// for the tunnel interface. It is used to specify the name of the network interface that the tunnel
// should be created on.
// @property {bool} DropLocalBroadcast - The `DropLocalBroadcast` property is a boolean value that
// determines whether local broadcast packets should be dropped or not. If set to `true`, local
// broadcast packets will be dropped. If set to `false`, local broadcast packets will be allowed.
// @property {bool} DropMulticast - The `DropMulticast` property is a boolean value that determines
// whether to drop multicast packets. If set to `true`, multicast packets will be dropped. If set to
// `false`, multicast packets will be allowed.
// @property {int} TxQueue - The `TxQueue` property in the `Tun` struct represents the size of the
// transmit queue for the network interface. It specifies the number of packets that can be queued for
// transmission before they are sent out.
// @property {int} Mtu - The `Mtu` property in the `Tun` struct represents the Maximum Transmission
// Unit. It specifies the maximum size of a packet that can be transmitted over the network.
// @property {[]Routes} Routes - The `Routes` property is a slice of `Routes` structs. Each `Routes`
// struct represents a network route and contains the following properties:
// @property {[]RouteTable} RouteTable - The `RouteTable` property is a slice of `RouteTable` structs.
type Tun struct {
	Enable             bool         `json:"enable,omitempty" yaml:"enable,omitempty"`
	Dev                string       `json:"dev,omitempty" yaml:"dev,omitempty"`
	DropLocalBroadcast bool         `json:"drop_local_broadcast,omitempty" yaml:"drop_local_broadcast,omitempty"`
	DropMulticast      bool         `json:"drop_multicast,omitempty" yaml:"drop_multicast,omitempty"`
	TxQueue            int          `json:"tx_queue,omitempty" yaml:"tx_queue,omitempty"`
	Mtu                int          `json:"mtu,omitempty" yaml:"mtu,omitempty"`
	Routes             []Routes     `json:"routes,omitempty" yaml:"routes,omitempty"`
	RouteTable         []RouteTable `json:"route_table,omitempty" yaml:"route_table,omitempty"`
}

// The Logging type is a struct that represents logging configuration with various properties such as
// level, language, format, file path, maximum size, maximum backups, and maximum age.
// @property {string} Level - The level property represents the logging level, which determines the
// severity of the logged messages. Common levels include "debug", "info", "warning", "error", and
// "fatal".
// @property {string} Language - The `Language` property represents the programming language used for
// logging.
// @property {string} Format - The `Format` property specifies the format in which the log messages
// will be written. It can be a string value representing the desired format, such as "text", "json",
// "xml", etc.
// @property {string} FilePath - The `FilePath` property is used to specify the file path where the log
// files will be stored.
// @property {int} MaxSize - MaxSize is the maximum size in megabytes that a log file can reach before
// it is rotated and a new log file is created.
// @property {int} MaxBackups - MaxBackups is the maximum number of backup log files to retain. When
// the log file reaches its maximum size, a new log file is created and the oldest log file is deleted
// if the number of backup log files exceeds the maximum number of backups specified.
// @property {int} MaxAge - The `MaxAge` property represents the maximum number of days to retain log
// files before they are automatically deleted.
type Logging struct {
	Level      string `json:"level,omitempty" yaml:"level,omitempty"`
	Language   string `json:"lang,omitempty" yaml:"lang,omitempty"`
	Format     string `json:"format,omitempty" yaml:"format,omitempty"`
	FilePath   string `json:"file_path,omitempty" yaml:"file_path,omitempty"`
	MaxSize    int    `json:"max_size,omitempty" yaml:"max_size,omitempty"`
	MaxBackups int    `json:"max_backups,omitempty" yaml:"max_backups,omitempty"`
	MaxAge     int    `json:"max_age,omitempty" yaml:"max_age,omitempty"`
}

// The above type represents statistics related to a server configuration.
// @property {string} Type - The "Type" property is a string that represents the type of the stats. It
// could be used to differentiate between different types of stats.
// @property {string} Listen - The "Listen" property specifies the network address and port on which
// the server should listen for incoming connections.
// @property {string} Path - The `Path` property represents the file path where the stats will be
// stored or retrieved from.
// @property {string} NameSpace - The `NameSpace` property is used to specify the namespace for the
// stats. It is an optional field and can be used to categorize or group the stats based on a specific
// namespace.
// @property {string} Extention - The `Extention` property is used to specify the file extension for
// the stats. It is an optional property and can be omitted if not needed.
// @property {string} Prefix - The "Prefix" property is used to specify a prefix that will be added to
// the generated metrics. It is an optional property and can be omitted if not needed.
// @property {string} Protocol - The "Protocol" property in the Stats struct represents the protocol
// used for communication. It can be used to specify the type of protocol being used, such as HTTP,
// TCP, UDP, etc.
// @property {string} Server - The "Server" property represents the server address or hostname. It is
// used to specify the server where the stats will be collected or sent to.
// @property {bool} MessageMetrics - MessageMetrics is a boolean property that indicates whether
// message metrics should be enabled or not. If set to true, it means that the system will collect and
// track metrics related to messages being processed.
// @property {bool} TowerMetrics - The `TowerMetrics` property is a boolean value that indicates
// whether tower metrics are enabled or not. If it is set to `true`, tower metrics are enabled. If it
// is set to `false` or omitted, tower metrics are disabled.
type Stats struct {
	Type           string `json:"type,omitempty" yaml:"type,omitempty"`
	Listen         string `json:"listen,omitempty" yaml:"listen,omitempty"`
	Path           string `json:"path,omitempty" yaml:"path,omitempty"`
	NameSpace      string `json:"name_space,omitempty" yaml:"name_space,omitempty"`
	Extention      string `json:"extention,omitempty" yaml:"extention,omitempty"`
	Prefix         string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Protocol       string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Server         string `json:"server,omitempty" yaml:"server,omitempty"`
	MessageMetrics bool   `json:"message_metrics,omitempty" yaml:"message_metrics,omitempty"`
	TowerMetrics   bool   `json:"tower_metrics,omitempty" yaml:"tower_metrics,omitempty"`
}

// The Handshakes type represents a set of parameters related to handshakes.
// @property {string} TryInterval - The TryInterval property specifies the interval at which handshake
// attempts should be made. It is represented as a string.
// @property {int} Retries - The "Retries" property specifies the number of times a handshake should be
// retried if it fails initially.
// @property {int} TriggerBuffer - The `TriggerBuffer` property is an integer that represents the
// number of handshakes that can be triggered before the system starts limiting the number of
// handshakes.
// @property {bool} ChurnLimiting - ChurnLimiting is a boolean property that indicates whether churn
// limiting is enabled or not. Churn limiting is a technique used to prevent excessive retries or
// failures in a system. When churn limiting is enabled, the system will limit the number of retries or
// failures that can occur within a certain period of time
// @property {int} ChurnNumFailures - ChurnNumFailures is the number of consecutive handshake failures
// that will trigger churn limiting.
// @property {string} ChurnPeriod - ChurnPeriod is a string property that represents the period of time
// during which churn is measured. It is used in the context of churn limiting, which is a technique
// used to prevent excessive retries or failures in a system.
type Handshakes struct {
	TryInterval      string `json:"try_interval,omitempty" yaml:"try_interval,omitempty"`
	Retries          int    `json:"retries,omitempty" yaml:"retries,omitempty"`
	TriggerBuffer    int    `json:"trigger_buffer,omitempty" yaml:"trigger_buffer,omitempty"`
	ChurnLimiting    bool   `json:"churn_limiting,omitempty" yaml:"churn_limiting,omitempty"`
	ChurnNumFailures int    `json:"churn_num_failures,omitempty" yaml:"churn_num_failures,omitempty"`
	ChurnPeriod      string `json:"churn_period,omitempty" yaml:"churn_period,omitempty"`
}

// The Timers type represents intervals for connection alive and pending deletion.
// @property {int} ConnectionAliveInterval - The `ConnectionAliveInterval` property represents the
// interval (in seconds) at which the connection status is checked to ensure it is still alive.
// @property {int} PendingDeletionInterval - The `PendingDeletionInterval` property represents the
// interval (in seconds) at which pending deletions are checked. It determines how often the system
// will check for any pending deletions and take appropriate actions.
type Timers struct {
	ConnectionAliveInterval int `json:"connection_alive_interval,omitempty" yaml:"connection_alive_interval,omitempty"`
	PendingDeletionInterval int `json:"pending_deletion_interval,omitempty" yaml:"pending_deletion_interval,omitempty"`
}

// The PSK type represents a pre-shared key with a mode and keys.
// @property {string} Mode - The `Mode` property is a string that represents the mode of the PSK
// (Pre-Shared Key). It can be used to specify the encryption algorithm or the security protocol being
// used.
// @property {any} Keys - The "Keys" property is of type "any", which means it can hold any type of
// value. It is used to store the keys associated with the PSK (Pre-Shared Key) object.
type PSK struct {
	Mode string `json:"mode,omitempty" yaml:"mode,omitempty"`
	Keys any    `json:"keys,omitempty" yaml:"keys,omitempty"`
}

// The Conntrack type represents a connection tracking configuration with timeout values for TCP, UDP,
// and default connections.
// @property {string} TCPTimeout - The `TCPTimeout` property represents the timeout value for TCP
// connections in the `Conntrack` struct.
// @property {string} UDPTimeout - The `UDPTimeout` property represents the timeout value for UDP
// connections in the `Conntrack` struct. It specifies the duration after which an idle UDP connection
// will be considered expired and removed from the connection tracking table.
// @property {string} DefaultTimeout - The `DefaultTimeout` property in the `Conntrack` struct
// represents the default timeout value for connection tracking. It specifies the amount of time after
// which an idle connection will be closed if no activity is detected.
type Conntrack struct {
	TCPTimeout     string `json:"tcp_timeout,omitempty" yaml:"tcp_timeout,omitempty"`
	UDPTimeout     string `json:"udp_timeout,omitempty" yaml:"udp_timeout,omitempty"`
	DefaultTimeout string `json:"default_timeout,omitempty" yaml:"default_timeout,omitempty"`
}

// The above type represents an outbound connection with properties such as port, protocol, and
// endpoint.
// @property {string} Port - The `Port` property represents the port number for the outbound
// connection. It is a string type and is used to specify the port number for the outbound connection.
// @property {string} Proto - The "Proto" property in the Outbound struct represents the protocol used
// for outbound connections. It can be used to specify the communication protocol, such as TCP or UDP.
// @property {string} Point - The "Point" property in the Outbound struct represents the destination
// point for the outbound connection.
type Outbound struct {
	Port  string `json:"port,omitempty" yaml:"port,omitempty"`
	Proto string `json:"proto,omitempty" yaml:"proto,omitempty"`
	Point string `json:"point,omitempty" yaml:"point,omitempty"`
}

// The above type represents an inbound connection with port, protocol, point, and groups attributes.
// @property {string} Port - The "Port" property represents the port number for the inbound connection.
// It is a string type and is used to specify the port on which the connection will be made.
// @property {string} Proto - The "Proto" property in the Inbound struct represents the protocol used
// for the inbound connection. It can be a string value such as "tcp", "udp", or "http".
// @property {string} Point - The "Point" property in the Inbound struct represents the specific
// endpoint or destination for the inbound traffic. It could be an IP address, domain name, or any
// other identifier that specifies where the traffic should be directed to.
// @property {[]string} Groups - The "Groups" property is a slice of strings that represents the groups
// associated with the inbound object. It is used to categorize or group inbound objects together.
type Inbound struct {
	Port   string   `json:"port,omitempty" yaml:"port,omitempty"`
	Proto  string   `json:"proto,omitempty" yaml:"proto,omitempty"`
	Point  string   `json:"point,omitempty" yaml:"point,omitempty"`
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`
}

// The Firewall type represents a firewall configuration with outbound and inbound rules.
// @property {string} OutboundAction - The `OutboundAction` property specifies the default action to be
// taken for outbound traffic. It can have values like "allow", "deny", or "reject".
// @property {string} InboundAction - The `InboundAction` property specifies the action to be taken for
// inbound traffic. It determines whether to allow or block incoming connections.
// @property {Conntrack} Conntrack - Conntrack is a struct that represents the connection tracking
// configuration for the firewall. It contains various properties related to connection tracking, such
// as timeout values and tracking modes.
// @property {[]Outbound} Outbound - The `Outbound` property is a slice of `Outbound` structs. It
// represents the rules for outgoing network connections. Each `Outbound` struct contains information
// such as the source IP address, destination IP address, protocol, and action to be taken for the
// connection.
// @property {[]Inbound} Inbound - The `Inbound` property is an array of `Inbound` objects. Each
// `Inbound` object represents a rule for incoming network traffic. It specifies the action to be taken
// for incoming traffic, such as allowing or blocking it.
type Firewall struct {
	OutboundAction string     `json:"outbound_action,omitempty" yaml:"outbound_action,omitempty"`
	InboundAction  string     `json:"inbound_action,omitempty" yaml:"inbound_action,omitempty"`
	Conntrack      Conntrack  `json:"conntrack,omitempty" yaml:"conntrack,omitempty"`
	Outbound       []Outbound `json:"outbound,omitempty" yaml:"outbound,omitempty"`
	Inbound        []Inbound  `json:"inbound,omitempty" yaml:"inbound,omitempty"`
}

// It returns a pointer to a config object.
// The function `newConfig()` returns a new instance of the `config` struct with default values.
func newConfig() *config {
	mtu := 1300
	return &config{
		PKI: PKI{
			Blocklist: []string{},
		},
		Points: map[string][]string{},
		Tower: Tower{
			DNS: DNS{Interval: 60},
		},
		Listen: Listen{
			Addr:  "0.0.0.0",
			Port:  0,
			Batch: 64,
		},
		Punchy: Punchy{
			Enable: true,
			Delay:  "1s",
		},
		Cipher: "aes",
		SSH: ssh{
			Users: []Users{},
		},
		Tun: Tun{
			Dev:                "vlan",
			DropLocalBroadcast: true,
			DropMulticast:      true,
			TxQueue:            500,
			Mtu:                mtu,
			Routes:             []Routes{},
			RouteTable:         []RouteTable{},
		},
		Logging: Logging{
			Level:  "info",
			Format: "text",
		},
		Stats: Stats{},
		Handshakes: Handshakes{
			TryInterval: "100ms",
			Retries:     20,
		},
		Firewall: Firewall{
			Conntrack: Conntrack{
				TCPTimeout:     "120h",
				UDPTimeout:     "3m",
				DefaultTimeout: "10m",
			},
			Outbound: []Outbound{
				{
					Port:  "any",
					Proto: "any",
					Point: "any",
				},
			},
			Inbound: []Inbound{},
		},
	}
}
