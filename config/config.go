package config

import (
	"bookstore-go/pkg/mlog"
	"bookstore-go/pkg/utils"
	"crypto/sha256"
	"fmt"
	"github.com/mohae/deepcopy"
	"go.uber.org/zap"
	"os"
	"reflect"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

type ConfigInfo struct {
	Server Server `yaml:"server,omitempty"`
	Mysql  Mysql  `yaml:"mysql,omitempty"`
	Redis  Redis  `yaml:"redis,omitempty"`

	// to check whether the configuration content has changed
	checksum string
}

type Server struct {
	Host string `yaml:"host,omitempty"`
	Port int    `yaml:"port,omitempty"`
}

type Mysql struct {
	Host         string `yaml:"host,omitempty"`
	Port         int    `yaml:"port,omitempty"`
	User         string `yaml:"user,omitempty"`
	Password     string `yaml:"password,omitempty"`
	DatabaseName string `yaml:"databaseName,omitempty"`
}

type Redis struct {
	Host     string `yaml:"host,omitempty"`
	Port     int    `yaml:"port,omitempty"`
	Password string `yaml:"password,omitempty"`
	DB       int    `yaml:"db,omitempty"`
}

// *********************************************************************************************************************

type Config interface {
	Value() *ConfigInfo
	AddListener(listener func(oldCfg, newCfg *ConfigInfo))
	RemoveListener(listener func(oldCfg, newCfg *ConfigInfo))
}

type configImpl struct {
	config      *ConfigInfo
	configMutex sync.RWMutex
	listeners   []func(oldCfg, newCfg *ConfigInfo)
}

// Value return the pointer to the ConfigInfo structure
func (c *configImpl) Value() *ConfigInfo {
	c.configMutex.RLock()
	defer c.configMutex.RUnlock()
	return c.config
}

func (c *configImpl) AddListener(listenHandler func(oldCfg, newCfg *ConfigInfo)) {
	c.listeners = append(c.listeners, listenHandler)
}

func (c *configImpl) RemoveListener(listenHandler func(oldCfg, newCfg *ConfigInfo)) {
	for i, listener := range c.listeners {
		if reflect.ValueOf(listener).Pointer() == reflect.ValueOf(listenHandler).Pointer() {
			c.listeners = append(c.listeners[:i], c.listeners[i+1:]...)
			break
		}
	}
}

func NewConfig(path string) (Config, error) {
	configImpl := &configImpl{
		config:    &ConfigInfo{},
		listeners: make([]func(oldCfg, newCfg *ConfigInfo), 0),
	}

	// step 1: load config file and update config struct
	cfg, err := configImpl.loadConfig(path)
	if err != nil {
		return nil, fmt.Errorf("load config from %v failed, err: %v", path, err)
	}

	if err := configImpl.updateConfig(cfg); err != nil {
		return nil, err
	}

	// step 2: monitor local file content changes, and dynamically update the config
	utils.GoHandleLoopWithRecover(func() {
		configImpl.watchConfigChanges(path)
	})

	return configImpl, nil
}

// loadConfig Parse the locally saved config file
func (c *configImpl) loadConfig(configPath string) (*ConfigInfo, error) {
	contentBytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	cfg := &ConfigInfo{}
	if err = yaml.Unmarshal(contentBytes, cfg); err != nil {
		return nil, err
	}

	c.completeDefaultConfig(cfg)

	cfg.checksum = fmt.Sprintf("%x", sha256.Sum256(contentBytes))

	return cfg, nil
}

// updateConfig updates the current configuration with the new one.
// If a sub-configuration in the old config is not nil/empty and the corresponding
// sub-configuration in the new config is nil/empty, the old sub-configuration is preserved.
// Otherwise, the new configuration values override the old ones.
// The operation is thread-safe using a mutex lock.
func (c *configImpl) updateConfig(cfg *ConfigInfo) error {
	if cfg == nil {
		return fmt.Errorf("cfg is nil")
	}

	c.configMutex.Lock()
	defer c.configMutex.Unlock()

	baseCfg := deepcopy.Copy(c.config).(*ConfigInfo)

	cfgYAML, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(cfgYAML, baseCfg)
	if err != nil {
		return err
	}

	baseCfg.checksum = cfg.checksum
	c.config = baseCfg

	return nil
}

// watchConfigChanges Periodically check config file changes
func (c *configImpl) watchConfigChanges(configPath string) {
	for {
		time.Sleep(time.Second * 5)
		loadCfg, err := c.loadConfig(configPath)
		if err != nil {
			mlog.Error("failed to load config", zap.String("configPath", configPath), zap.Error(err))
			continue
		}

		// Check if the config has changed
		if loadCfg.checksum != c.config.checksum {
			mlog.Info("config changed", zap.Any("old config", c.config), zap.Any("new config", loadCfg))
			oldCfg := deepcopy.Copy(c.config).(*ConfigInfo)
			if err := c.updateConfig(loadCfg); err != nil {
				mlog.Error("local file has changed, but update config failed", zap.Error(err))
				continue
			}
			c.notifyListeners(oldCfg, c.config)
		}
	}
}

// notifyListeners notify all listeners
func (c *configImpl) notifyListeners(oldCfg, newCfg *ConfigInfo) {
	if len(c.listeners) == 0 {
		return
	}

	oldCfgSnapshot := deepcopy.Copy(oldCfg).(*ConfigInfo)
	newCfgSnapshot := deepcopy.Copy(newCfg).(*ConfigInfo)

	for _, listener := range c.listeners {
		listener(oldCfgSnapshot, newCfgSnapshot)
	}
}

func (c *configImpl) completeDefaultConfig(cfg *ConfigInfo) {
	if cfg.Server.Host == "" {
		cfg.Server.Host = "0.0.0.0"
	}
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 19101
	}
}
