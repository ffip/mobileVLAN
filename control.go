package mobileHiPer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ffip/hiper"
	"net"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/ffip/hiper"

	hc "github.com/ffip/hiper/config"
	"github.com/ffip/hiper/iputil"
	"github.com/ffip/hiper/udp"
)

type HiPer struct {
	c *hiper.Control
	l *logrus.Logger
}

func init() {
	// Reduces memory utilization according to https://twitter.com/felixge/status/1355846360562589696?s=20
	runtime.MemProfileRate = 0
}

func NewHiPer(configData string, key string, logFile string, tunFd int) (*HiPer, error) {
	// GC more often, largely for iOS due to extension 15mb limit
	debug.SetGCPercent(20)

	yamlConfig, err := RenderConfig(configData, key)
	if err != nil {
		return nil, err
	}

	l := logrus.New()
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	l.SetOutput(f)

	c := hc.NewC(l)
	err = c.LoadString(yamlConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	//TODO: inject our version
	ctrl, err := hiper.Main(c, false, "", l, &tunFd)
	if err != nil {
		switch v := err.(type) {
		case util.ContextualError:
			v.Log(l)
			return nil, v.Unwrap()
		default:
			l.WithError(err).Error("Failed to start")
			return nil, err
		}
	}

	return &HiPer{ctrl, l}, nil
}

func (h *HiPer) Log(v string) {
	h.l.Println(v)
}

func (h *HiPer) Start() {
	h.c.Start()
}

func (h *HiPer) ShutdownBlock() {
	h.c.ShutdownBlock()
}

func (h *HiPer) Stop() {
	h.c.Stop()
}

func (h *HiPer) Rebind(reason string) {
	h.l.Debugf("Rebinding UDP listener and updating towers due to %s", reason)
	h.c.RebindUDPServer()
}

func (h *HiPer) ListHostmap(pending bool) (string, error) {
	hosts := h.c.ListHostmap(pending)
	b, err := json.Marshal(hosts)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (h *HiPer) GetHostInfoByEndpoint(endpoint string, pending bool) (string, error) {
	b, err := json.Marshal(h.c.GetHostInfoByEndpoint(stringIpToInt(endpoint), pending))
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (h *HiPer) CloseTunnel(endpoint string) bool {
	return h.c.CloseTunnel(stringIpToInt(endpoint), false)
}

func (h *HiPer) SetRemoteForTunnel(endpoint string, addr string) (string, error) {
	udpAddr := udp.NewAddrFromString(addr)
	if udpAddr == nil {
		return "", errors.New("could not parse udp address")
	}

	b, err := json.Marshal(h.c.SetRemoteForTunnel(stringIpToInt(endpoint), *udpAddr))
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (h *HiPer) Sleep() {
	if closed := h.c.CloseAllTunnels(true); closed > 0 {
		h.l.WithField("tunnels", closed).Info("Sleep called, closed non Tower tunnels")
	}
}

func stringIpToInt(ip string) iputil.Endpoint {
	return iputil.IP2Endpoint(net.ParseIP(ip))
}
