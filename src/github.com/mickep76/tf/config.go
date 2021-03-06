package main

import (
	"fmt"
)

// CfgDefault contains input configuration defaults.
type CfgDefault struct {
	EtcdHost      *string
	EtcdPort      *int64
	HTTPHeader    *string
	HTTPFormat    *string
	MySQLUser     *string
	MySQLPassword *string
	MySQLHost     *string
	MySQLPort     *int64
	MySQLDatabase *string
}

// CfgInput contains input configuration.
type CfgInput struct {
	Name          *string
	Type          *string
	Path          *string
	EtcdHost      *string
	EtcdPort      *int64
	EtcdDir       *string
	HTTPUrl       *string
	HTTPHeader    *string
	HTTPFormat    *string
	MySQLUser     *string
	MySQLPassword *string
	MySQLHost     *string
	MySQLPort     *int64
	MySQLDatabase *string
	MySQLQuery    *string
}

// GetDefaults gets input defaults from the config file.
func GetDefaults(defs map[string]interface{}) (CfgDefault, error) {
	var d CfgDefault
	for k, v := range defs {
		switch k {
		case "etcd_host":
			s := v.(string)
			d.EtcdHost = &s
		case "etcd_port":
			n := v.(int64)
			d.EtcdPort = &n
		case "http_header":
			s := v.(string)
			d.HTTPHeader = &s
		case "http_format":
			s := v.(string)
			d.HTTPFormat = &s
		case "mysql_user":
			s := v.(string)
			d.MySQLUser = &s
		case "mysql_password":
			s := v.(string)
			d.MySQLPassword = &s
		case "mysql_host":
			s := v.(string)
			d.MySQLHost = &s
		case "mysql_port":
			n := v.(int64)
			d.MySQLPort = &n
		case "mysql_database":
			s := v.(string)
			d.MySQLDatabase = &s
		default:
			return CfgDefault{}, fmt.Errorf("Invalid configuration key \"%v\" in [defaults]", k)
		}
	}
	return d, nil
}

// GetInput gets inputs from the configuration file.
func GetInput(name string, inp map[string]interface{}, d CfgDefault) (CfgInput, error) {
	var i CfgInput

	if d.EtcdHost != nil {
		i.EtcdHost = d.EtcdHost
	}
	if d.EtcdPort != nil {
		i.EtcdPort = d.EtcdPort
	} else {
		n := int64(2379)
		i.EtcdPort = &n
	}
	if d.HTTPHeader != nil {
		i.HTTPHeader = d.HTTPHeader
	} else {
		s := "Accept: application/json"
		i.HTTPHeader = &s
	}
	if d.HTTPFormat != nil {
		i.HTTPFormat = d.HTTPFormat
	} else {
		s := "JSON"
		i.HTTPFormat = &s
	}
	if d.MySQLUser != nil {
		i.MySQLUser = d.MySQLUser
	}
	if d.MySQLPassword != nil {
		i.MySQLPassword = d.MySQLPassword
	}
	if d.MySQLHost != nil {
		i.MySQLHost = d.MySQLHost
	}
	if d.MySQLPort != nil {
		i.MySQLPort = d.MySQLPort
	} else {
		n := int64(3306)
		i.MySQLPort = &n
	}
	if d.MySQLDatabase != nil {
		i.MySQLDatabase = d.MySQLDatabase
	}

	i.Name = &name
	for k, v := range inp {
		switch k {
		case "name":
			s := v.(string)
			i.Name = &s
		case "type":
			s := v.(string)
			i.Type = &s
		case "path":
			s := v.(string)
			i.Path = &s
		case "etcd_host":
			s := v.(string)
			i.EtcdHost = &s
		case "etcd_port":
			n := v.(int64)
			i.EtcdPort = &n
		case "etcd_dir":
			s := v.(string)
			i.EtcdDir = &s
		case "http_url":
			s := v.(string)
			i.HTTPUrl = &s
		case "http_header":
			s := v.(string)
			i.HTTPHeader = &s
		case "http_format":
			s := v.(string)
			i.HTTPFormat = &s
		case "mysql_user":
			s := v.(string)
			i.MySQLUser = &s
		case "mysql_password":
			s := v.(string)
			i.MySQLPassword = &s
		case "mysql_host":
			s := v.(string)
			i.MySQLHost = &s
		case "mysql_port":
			n := v.(int64)
			i.MySQLPort = &n
		case "mysql_database":
			s := v.(string)
			i.MySQLDatabase = &s
		case "mysql_query":
			s := v.(string)
			i.MySQLQuery = &s
		default:
			return CfgInput{}, fmt.Errorf("Invalid configuration key \"%v\" in [inputs.%v]", k, name)
		}
	}

	switch *i.Type {
	case "file":
		if i.Path == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"file\" you need to specify \"path\"", name)
		}
	case "etcd":
		if i.EtcdHost == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"etcd\" you need to specify \"etcd_host\"", name)
		}
		if i.EtcdPort == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"etcd\" you need to specify \"etcd_port\"", name)
		}
	case "http":
		if i.HTTPUrl == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"http\" you need to specify \"http_url\"", name)
		}
		if i.HTTPHeader == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"http\" you need to specify \"http_header\"", name)
		}
		if i.HTTPFormat == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"http\" you need to specify \"http_format\"", name)
		}
	case "mysql":
		if i.MySQLUser == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"mysql\" you need to specify \"mysql_user\"", name)
		}
		if i.MySQLPassword == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"mysql\" you need to specify \"mysql_password\"", name)
		}
		if i.MySQLHost == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"mysql\" you need to specify \"mysql_host\"", name)
		}
		if i.MySQLDatabase == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"mysql\" you need to specify \"mysql_database\"", name)
		}
		if i.MySQLQuery == nil {
			return CfgInput{}, fmt.Errorf("For input [inputs.%v] type \"mysql\" you need to specify \"mysql_query\"", name)
		}
	default:
		return CfgInput{}, fmt.Errorf("Unknown type \"%v\" for input [inputs.%v]", *i.Type, *i.Name)
	}

	return i, nil
}
