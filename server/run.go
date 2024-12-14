package server

import (
	"os"
	"strconv"
)

var _serverControl control

func Run() {
	readEnvironment()
	buildDatabase()
	_serverControl.boot()
}

func readEnvironment() {
	const (
		modeEnvVar    = "DIVELOG_MODE"
		dbPathEnvVar  = "DIVELOG_DBFILE_PATH"
		dnsNameEnvVar = "DIVELOG_DNS_NAME"
		portEnvVar    = "DIVELOG_PORT"
	)

	mode := os.Getenv(modeEnvVar)
	trace(_env, "%s = %q", modeEnvVar, mode)
	if mode == "dev" {
		_serverControl.localAPI = true
		_serverControl.encryptedTraffic = false
		_serverControl.endpoint = "localhost:8072"
	} else if mode == "" || mode == "prod" {
		dnsName := os.Getenv(dnsNameEnvVar)
		trace(_env, "%s = %q", dnsNameEnvVar, dnsName)
		if dnsName == "" {
			trace(_error, "%s is empty or undefined", dnsNameEnvVar)
			os.Exit(1)
		}

		port := os.Getenv(portEnvVar)
		trace(_env, "%s = %q", portEnvVar, port)
		if port == "" {
			port = "443"
		} else {
			if num, err := strconv.Atoi(port); err != nil || num < 1 || num > 65535 {
				trace(_error, "value of %s is invalid or is not a valid TCP port number", portEnvVar)
				os.Exit(1)
			}
		}

		_serverControl.encryptedTraffic = true
		_serverControl.endpoint = dnsName + ":" + port
	} else {
		trace(_error, "value of %s is invalid", modeEnvVar)
		os.Exit(1)
	}

	_inmemDatabase.Metadata.Source = os.Getenv(dbPathEnvVar)
	trace(_env, "%s = %q", dbPathEnvVar, _inmemDatabase.Metadata.Source)
	if _inmemDatabase.Metadata.Source == "" {
		trace(_error, "%s is empty or undefined", dbPathEnvVar)
		os.Exit(1)
	}
}
