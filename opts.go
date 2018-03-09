package libcni

import (
	cnilibrary "github.com/containernetworking/cni/libcni"
)

type ConfigOptions func(c *libcni) error

// WithInterfacePrefix sets the prefix for network interfaces
// e.g. eth or wlan
func WithInterfacePrefix(prefix string) ConfigOptions {
	return func(c *libcni) error {
		c.prefix = prefix
		return nil
	}
}

func WithPluginDir(dirs []string) ConfigOptions {
	return func(c *libcni) error {
		c.pluginDirs = append(c.pluginDirs, dirs...)
		return nil
	}
}

func WithPluginConfDir(dir string) ConfigOptions {
	return func(c *libcni) error {
		c.pluginConfDir = dir
		return nil
	}
}

func WithLoNetwork() ConfigOptions {
	return func(c *libcni) error {
		loConfig, _ := cnilibrary.ConfListFromBytes([]byte(`{
"cniVersion": "0.3.1",
"name": "cni-loopback",
"plugins": [{
  "type": "loopback"
}]
}`))
		c.networks = append(c.networks, &Network{
			cni:    c.cniConfig,
			config: loConfig,
			ifName: "lo",
		})
		return nil
	}
}

func WithMinNetworkCount(count int) ConfigOptions {
	return func(c *libcni) error {
		c.networkCount = count
		return nil
	}
}

//TODO: Should we support direct network configs?
/*
func WithConf(byte []bytes) ConfigOptions {
	return func(c *config) error {
			c.networks=
	}
}

func WithConfFile(fileName string) ConfigOptions {
	return func(c *config) error {

	}
}

func WithConfList(byte []bytes) ConfigOptions {
	return func(c *config) error {
      c.
	}
}
func WithConfListFile(files []string) ConfigOptions {
	return func(c *config) error {

	}
}
*/