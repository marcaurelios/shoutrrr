// Code generated by "shoutrrr-gen "; DO NOT EDIT.
package generic

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
	ContentType string     `key:"contenttype" `
	DisableTLS  bool       `key:"disabletls" `
	Host        string     `url:"host" `
	Password    string     `url:"password" `
	Path        string     `url:"path" `
	Query       url.Values `url:"query" `
	Template    string     `key:"template" `
	Title       string     `key:"title" `
	User        string     `url:"user" `
}

type configProp int

const (
	propContentType configProp = 0
	propDisableTLS  configProp = 1
	propHost        configProp = 2
	propPassword    configProp = 3
	propPath        configProp = 4
	propQuery       configProp = 5
	propTemplate    configProp = 6
	propTitle       configProp = 7
	propUser        configProp = 8
	propCount                  = 9
)

var propInfo = types.ConfigPropInfo{
	PropNames: []string{
		"ContentType",
		"DisableTLS",
		"Host",
		"Password",
		"Path",
		"Query",
		"Template",
		"Title",
		"User",
	},

	// Note that propKeys may not align with propNames, as a property can have no or multiple keys
	Keys: []string{
		"contenttype",
		"disabletls",
		"template",
		"title",
	},

	DefaultValues: []string{
		"application/json",
		"No",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	},

	PrimaryKeys: []int{
		0,
		1,
		-1,
		-1,
		-1,
		-1,
		2,
		3,
		-1,
	},

	KeyPropIndexes: map[string]int{
		"contenttype": 0,
		"disabletls":  1,
		"template":    6,
		"title":       7,
	},
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
		User:     conf.UserInfoOrNil(url.UserPassword(config.User, config.Password)),
		Host:     config.Host,
		Path:     conf.JoinPath(string(config.Path)),
		RawQuery: conf.QueryValues(config).Encode(),
		Scheme:   Scheme,
	}
}

func (config *Config) CustomQueryVars() url.Values {
	return config.Query
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
	updates[int(propHost)] = configURL.Hostname()
	if pwd, found := configURL.User.Password(); found {
		updates[int(propPassword)] = pwd
	}
	updates[int(propUser)] = configURL.User.Username()
	updates[int(propPath)] = configURL.Path
	customQuery := url.Values{}

	for key, value := range configURL.Query() {

		if propIndex, found := propInfo.PropIndexFor(key); found {
			updates[propIndex] = value[0]
		} else {
			customQuery.Set(conf.UnescapeCustomQueryKey(key), value[0])
		}
	}
	updates[int(propQuery)] = customQuery.Encode()

	err := config.Update(updates)
	if err != nil {
		return err
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
		case propContentType:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.ContentType = val
			}
		case propDisableTLS:
			if val, err := conf.ParseToggleValue(value); err != nil {
				last_err = err
			} else {
				config.DisableTLS = val
			}
		case propHost:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Host = val
			}
		case propPassword:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Password = val
			}
		case propPath:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Path = val
			}
		case propQuery:
			if val, err := url.ParseQuery(value); err != nil {
				last_err = err
			} else {
				config.Query = val
			}
		case propTemplate:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Template = val
			}
		case propTitle:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.Title = val
			}
		case propUser:
			if val, err := conf.ParseTextValue(value); err != nil {
				last_err = err
			} else {
				config.User = val
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
	case propContentType:
		return conf.FormatTextValue(config.ContentType)
	case propDisableTLS:
		return conf.FormatToggleValue(config.DisableTLS)
	case propHost:
		return conf.FormatTextValue(config.Host)
	case propPassword:
		return conf.FormatTextValue(config.Password)
	case propPath:
		return conf.FormatTextValue(config.Path)
	case propQuery:
		return config.Query.Encode()
	case propTemplate:
		return conf.FormatTextValue(config.Template)
	case propTitle:
		return conf.FormatTextValue(config.Title)
	case propUser:
		return conf.FormatTextValue(config.User)
	default:
		return ""
	}
}
