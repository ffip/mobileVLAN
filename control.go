package mobile

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"runtime"
	"runtime/debug"

	cfg "github.com/ffip/vlan/config"
	"github.com/ffip/vlan/entry"
	"github.com/ffip/vlan/lib/network/iputil"
	"github.com/ffip/vlan/lib/network/udp"
	"github.com/ffip/vlan/lib/utils/caution"
	"github.com/ffip/vlan/lib/utils/logs/logger"
)

// The Bulk type is a struct that contains a control, logger, and configuration.
// @property c - A pointer to an instance of the `Control` struct.
// @property l - A pointer to a logger.Logger object.
// @property config - The `config` property is a pointer to an object of type `cfg.C`.
type Bulk struct {
	c      *entry.Control
	l      *logger.Logger
	config *cfg.C
}

func init() {
	// Reduces memory utilization according to https://twitter.com/felixge/status/1355846360562589696?s=20
	runtime.MemProfileRate = 0
}

// The function `NewBulk` creates a new instance of the `Bulk` struct with the provided configuration
// data, log file path, and tunnel file descriptor.
func NewBulk(configData string, logFile string, tunFd int) (*Bulk, error) {
	// GC more often, largely for iOS due to extension 15mb limit
	debug.SetGCPercent(20)

	l := logger.New(1000)
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	l.SetOutput(f)

	c := cfg.NewC(l)
	err = c.LoadString(configData)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	//TODO: inject our version
	ctrl, err := entry.Main(c, false, "", l, &tunFd)
	if err != nil {
		caution.LogWithContextIfNeeded("Failed to start", err, l)
		return nil, err
	}

	return &Bulk{ctrl, l, c}, nil
}

// The `Log` function is a method of the `Bulk` struct. It takes a string `v` as a parameter and logs
// the string using the logger associated with the `Bulk` instance.
func (n *Bulk) Log(v string) {
	n.l.Println(v)
}

// The `Start()` method of the `Bulk` struct is used to start the execution of the `Control` instance
// associated with the `Bulk` instance. It calls the `Start()` method of the `Control` instance, which
// starts the main event loop and begins handling network traffic.
func (n *Bulk) Start() {
	n.c.Start()
}

// The `ShutdownBlock()` method of the `Bulk` struct is used to block the execution of the program
// until all active connections are closed. It calls the `ShutdownBlock()` method of the `Control`
// instance associated with the `Bulk` instance, which waits for all active connections to be closed
// before returning. This method is typically used when gracefully shutting down the program to ensure
// that all network connections are properly closed before exiting.
func (n *Bulk) ShutdownBlock() {
	n.c.ShutdownBlock()
}

// The `Stop()` method of the `Bulk` struct is used to stop the execution of the `Control` instance
// associated with the `Bulk` instance. It calls the `Stop()` method of the `Control` instance, which
// stops the main event loop and terminates the handling of network traffic.
func (n *Bulk) Stop() {
	n.c.Stop()
}

// The `Rebind` method of the `Bulk` struct is used to rebind the UDP listener and update the towers.
// It takes a `reason` string as a parameter, which is used for logging purposes. Inside the method, it
// calls the `RebindUDPServer` method of the `Control` instance associated with the `Bulk` instance.
// This method rebinds the UDP listener and updates the towers, which can be useful in scenarios where
// the network configuration has changed or there is a need to refresh the network connections.
func (n *Bulk) Rebind(reason string) {
	n.l.Debug("Rebinding UDP listener and updating towers due to %s", reason)
	n.c.RebindUDPServer()
}

// The `Reload` method of the `Bulk` struct is used to reload the configuration of the `Bulk` instance.
// It takes a `configData` string as a parameter, which represents the new configuration data. Inside
// the method, it calls the `ReloadConfigString` method of the `cfg.C` instance associated with the
// `Bulk` instance. This method reloads the configuration using the provided `configData` string. If
// there is an error during the reloading process, the method returns an error.
func (n *Bulk) Reload(configData string) error {
	n.l.Info("Reloading Nebula")

	return n.config.ReloadConfigString(configData)
}

// The `ListPendingPoints` method of the `Bulk` struct is used to retrieve a list of pending points. It
// takes a boolean parameter `pending` which indicates whether to list pending points or not.
func (n *Bulk) ListPendingPoints(pending bool) (string, error) {
	points := n.c.ListProcessesPoints(pending)
	b, err := json.Marshal(points)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// The `GetPointInfoByEndpoint` method of the `Bulk` struct is used to retrieve information about a
// specific network point (endpoint). It takes two parameters: `endpoint` which is the IP address of
// the network point, and `pending` which indicates whether to include pending points or not.
func (n *Bulk) GetPointInfoByEndpoint(endpoint string, pending bool) (string, error) {
	endpointInt := stringIpToInt(endpoint)
	b, err := json.Marshal(n.c.GetpointByEndpoint(endpointInt, pending))
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// The `CloseTunnel` method of the `Bulk` struct is used to close a tunnel associated with a specific
// endpoint. It takes an `endpoint` string parameter, which represents the IP address of the network
// point.
func (n *Bulk) CloseTunnel(endpoint string) bool {
	return n.c.CloseTunnel(stringIpToInt(endpoint), false)
}

// The `SetRemoteForTunnel` method of the `Bulk` struct is used to set the remote address for a tunnel
// associated with a specific endpoint. It takes two parameters: `endpoint` which is the IP address of
// the network point, and `addr` which is the remote address to set for the tunnel.
func (n *Bulk) SetRemoteForTunnel(endpoint string, addr string) (string, error) {
	udpAddr := udp.NewAddrFromString(addr)
	if udpAddr == nil {
		return "", errors.New("could not parse udp address")
	}

	b, err := json.Marshal(n.c.SetRemoteForTunnel(stringIpToInt(endpoint), *udpAddr))
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// The `Sleep()` method of the `Bulk` struct is used to put the program to sleep. It closes all
// non-tower tunnels and logs the number of closed tunnels. This method is typically used when the
// program needs to temporarily pause its execution or go into a sleep mode.
func (n *Bulk) Sleep() {
	if closed := n.c.CloseAllTunnels(true); closed > 0 {
		n.l.Echo().WithField("tunnels", closed).Info("Sleep called, closed non tower tunnels")
	}
}

func stringIpToInt(ip string) iputil.Endpoint {
	return iputil.IP2Endpoint(net.ParseIP(ip))
}
