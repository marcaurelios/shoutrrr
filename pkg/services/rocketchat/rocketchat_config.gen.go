// Code generated by "shoutrrr-gen "; DO NOT EDIT.
package rocketchat

import (
	"fmt"
	"net/url"

	"github.com/containrrr/shoutrrr/pkg/conf"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// (‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾)
//  )  Props                          (
// (___________________________________)

type Config struct {
	Channel  string `url:"path3" `
	Host     string `url:"host" `
	Port     int64  `url:"port" `
	TokenA   string `url:"path1" `
	TokenB   string `url:"path2" `
	UserName string `url:"user" `
}

type configProp int

const (
	propChannel  configProp = 0
	propHost     configProp = 1
	propPort     configProp = 2
	propTokenA   configProp = 3
	propTokenB   configProp = 4
	propUserName configProp = 5
	propCount               = 6
)

var propInfo = types.ConfigPropInfo{
	PropNames: []string{
		"Channel",
		"Host",
		"Port",
		"TokenA",
		"TokenB",
		"UserName",
	},

	// Note that propKeys may not align with propNames, as a property can have no or multiple keys
	Keys: []string{},

	DefaultValues: []string{
		"",
		"",
		"",
		"",
		"",
		"",
	},

	PrimaryKeys: []int{
		-1,
		-1,
		-1,
		-1,
		-1,
		-1,
	},

	KeyPropIndexes: map[string]int{},
}

func (_ *Config) PropInfo() *types.ConfigPropInfo {
	return &propInfo
}

// (‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾)
//  )  GetURL                         (
// (___________________________________)

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL {
	return &url.URL{
		User:     conf.UserInfoOrNil(url.User(config.UserName)),
		Host:     conf.FormatHost(config.Host, config.Port),
		Path:     conf.JoinPath(string(config.TokenA), string(config.TokenB), string(config.Channel)),
		RawQuery: conf.QueryValues(config).Encode(),
		Scheme:   Scheme,
	}
}

// (‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾)
//  )  SetURL                         (
// (___________________________________)

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(configURL *url.URL) error {
	if lc, ok := (interface{})(config).(types.ConfigWithLegacyURLSupport); ok {
		configURL = lc.UpdateLegacyURL(configURL)
	}
	updates := make(map[int]string, propCount)
	if port := configURL.Port(); port != "" {
		updates[int(propPort)] = port
	}
	updates[int(propHost)] = configURL.Hostname()
	updates[int(propUserName)] = configURL.User.Username()

	pathParts := conf.SplitPath(configURL.Path)
	if len(pathParts) > 0 {
		updates[int(propTokenA)] = pathParts[0]
	}
	if len(pathParts) > 1 {
		updates[int(propTokenB)] = pathParts[1]
	}
	if len(pathParts) > 2 {
		updates[int(propChannel)] = pathParts[2]
	}
	if len(pathParts) > 3 {
		return fmt.Errorf("too many path items: %v, expected 3", len(pathParts))
	}

	for key, value := range configURL.Query() {

		if propIndex, found := propInfo.PropIndexFor(key); found {
			updates[propIndex] = value[0]
		} else if key != "title" {
			return fmt.Errorf("invalid key %q", key)
		}
	}

	err := config.Update(updates)
	if err != nil {
		return err
	}

	if config.Host == "" {
		return fmt.Errorf("host missing from config URL")
	}

	if config.TokenA == "" {
		return fmt.Errorf("tokenA missing from config URL")
	}

	if config.TokenB == "" {
		return fmt.Errorf("tokenB missing from config URL")
	}

	return nil
}

// (‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾)
//  )  Enums / Options                (
// (___________________________________)

func (config *Config) Enums() map[string]types.EnumFormatter {
	return map[string]types.EnumFormatter{}
}

// Update updates the Config from a map of it's properties
func (config *Config) Update(updates map[int]string) error {
	var last_err error
	for index, value := range updates {
		switch configProp(index) {
		case propChannel:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Channel = val
			}
		case propHost:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Host = val
			}
		case propPort:
			if val, err := conf.ParseNumberValue(value, 0); err != nil {
				last_err = err
			} else {
				config.Port = val
			}
		case propTokenA:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.TokenA = val
			}
		case propTokenB:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.TokenB = val
			}
		case propUserName:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.UserName = val
			}
		default:
			return fmt.Errorf("invalid key")
		}
		if last_err != nil {
			return fmt.Errorf("failed to set value for %v: %v", propInfo.PropNames[index], last_err)
		}
	}
	return nil
}

// Update updates the Config from a map of it's properties
func (config *Config) PropValue(prop int) string {
	switch configProp(prop) {
	case propChannel:
		return conf.FormatTextValue(config.Channel)
	case propHost:
		return conf.FormatTextValue(config.Host)
	case propPort:
		return conf.FormatNumberValue(config.Port, 0)
	case propTokenA:
		return conf.FormatTextValue(config.TokenA)
	case propTokenB:
		return conf.FormatTextValue(config.TokenB)
	case propUserName:
		return conf.FormatTextValue(config.UserName)
	default:
		return ""
	}
}
