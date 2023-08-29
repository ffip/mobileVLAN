package mobile

import (
	"encoding/json"
	"fmt"
	"strings"

	cfg "github.com/ffip/vlan/config"
	"github.com/ffip/vlan/lib/utils/algo/suid"
	hc "github.com/ffip/vlan/lib/utils/http/client"
	"github.com/ffip/vlan/lib/utils/logs/logger"
	"github.com/ffip/vlan/lib/yaml"
)

func checkConfig(c *cfg.C, l *logger.Logger) {
	src := c.GetString("sync.source", "")
	if src == "" {
		l.Debug("Check config skipped, uuid or auth url is empty")
	} else {
		_, err := suid.ParseBase58([]byte(src))
		switch {
		case err == nil:
			src = fmt.Sprintf("https://cert.mcer.cn/%s.yml", src)
			fallthrough
		case strings.HasPrefix(src, "http://"), strings.HasPrefix(src, "https://"):
			l.Info("Checking config")
		default:
			l.Warn("Unsupported config source: %s", src)
			return
		}
		req := hc.NewRequest().SetUrl(src).Do()
		if req.GetStatusCode() != 200 || req.GetBodyString() == "" {
			l.Warn("Sync config failed, auth failed")
		} else {
			l.Debug("Get config success")
			tc := cfg.NewC(l)
			if err := tc.LoadString(req.GetBodyString()); err != nil {
				l.Echo().WithError(err).Warn("Sync config failed")
			} else {
				l.Info("Sync config success")
			}
			if tc.GetString("pki.ca", "") != "" {
				tc.SetPath(c.GetPath())
				tc.SaveConfig()
			}
			c.ReloadConfig()
		}
	}

	addition := c.GetString("sync.addition", "")
	if addition == "" {
		l.Debug("Check addition config skipped, addition is empty")
		return
	} else {
		l.Info("Checking addition config")
		req := hc.NewRequest().SetUrl(addition).Do()
		if req.GetStatusCode() != 200 || req.GetBodyString() == "" {
			l.Warn("Sync addition config failed, remote addition config is empty")
		} else {
			l.Debug("Get addition config success")
			var m map[string]any
			err := yaml.Unmarshal([]byte(req.GetBodyString()), &m)
			if err != nil {
				err = json.Unmarshal([]byte(req.GetBodyString()), &m)
				if err != nil {
					return
				}
			}
			if err != nil {
				l.Echo().WithError(err).Warn("Sync addition config failed")
				return
			} else {
				cfg.MergeSettings(c.Settings, m)
				l.Info("Sync addition config success")
			}
			if c.GetString("pki.ca", "") != "" {
				c.SaveConfig()
			}
			c.ReloadConfig()
		}
	}
}
