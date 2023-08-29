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

type Bulk struct {
	c      *entry.Control
	l      *logger.Logger
	config *cfg.C
}

func init() {
	// Reduces memory utilization according to https://twitter.com/felixge/status/1355846360562589696?s=20
	runtime.MemProfileRate = 0
}

func NewBulk(configData string, key string, logFile string, tunFd int) (*Bulk, error) {
	// GC more often, largely for iOS due to extension 15mb limit
	debug.SetGCPercent(20)

	yamlConfig, err := RenderConfig(configData, key)
	if err != nil {
		return nil, err
	}

	l := logger.New(1000)
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	l.SetOutput(f)

	c := cfg.NewC(l)
	err = c.LoadString(yamlConfig)
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

func (n *Bulk) Log(v string) {
	n.l.Println(v)
}

func (n *Bulk) Start() {
	n.c.Start()
}

func (n *Bulk) ShutdownBlock() {
	n.c.ShutdownBlock()
}

func (n *Bulk) Stop() {
	n.c.Stop()
}

func (n *Bulk) Rebind(reason string) {
	n.l.Debug("Rebinding UDP listener and updating towers due to %s", reason)
	n.c.RebindUDPServer()
}

func (n *Bulk) Reload(configData string, key string) error {
	n.l.Info("Reloading Nebula")
	yamlConfig, err := RenderConfig(configData, key)
	if err != nil {
		return err
	}

	return n.config.ReloadConfigString(yamlConfig)
}

func (n *Bulk) ListHostmap(pending bool) (string, error) {
	hosts := n.c.ListProcessesPoints(pending)
	b, err := json.Marshal(hosts)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (n *Bulk) GetPointInfoByEndpoint(endpoint string, pending bool) (string, error) {
	b, err := json.Marshal(n.c.GetpointByEndpoint(stringIpToInt(endpoint), pending))
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (n *Bulk) CloseTunnel(endpoint string) bool {
	return n.c.CloseTunnel(stringIpToInt(endpoint), false)
}

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

func (n *Bulk) Sleep() {
	if closed := n.c.CloseAllTunnels(true); closed > 0 {
		n.l.Echo().WithField("tunnels", closed).Info("Sleep called, closed non tower tunnels")
	}
}

func stringIpToInt(ip string) iputil.Endpoint {
	return iputil.IP2Endpoint(net.ParseIP(ip))
}
