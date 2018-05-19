package main

import (
	"errors"
	"flag"
)

type CLIParameters struct {
	StartServer  *bool
	AddClient    *bool
	RevokeClient *bool

	ClientId     *string
	ClientSecret *string
	RedirectUri  *string

	ConfigPath *string
}

func showDefaults() {
	flag.PrintDefaults()
}

func handleCLIParameters(params *CLIParameters) (err error) {
	params.StartServer = flag.Bool("start-server", false, "Starts the webserver if set.")
	params.AddClient = flag.Bool("add-client", false, "Add the specified ClientId and Secret.")
	params.RevokeClient = flag.Bool("revoke-client", false, "Revokes the ClientId.")
	params.ClientId = flag.String("client-id", "", "The new ClientId to be added or revoked.")
	params.ClientSecret = flag.String("client-secret", "", "The new ClientSecret.")
	params.RedirectUri = flag.String("redirect-uri", "", "The RedirectUri.")
	params.ConfigPath = flag.String("config", "", "Path to config file in ini format.")

	flag.Parse()

	// Validate CLI values
	if !(*params.StartServer) && !(*params.AddClient) && !(*params.RevokeClient) {
		err = errors.New("You need to specify StartServer, AddClient or RevokeClient!")
	}

	if *params.ConfigPath == "" {
		err = errors.New("No config Path given.")
	}

	if *(params.StartServer) && *(params.AddClient) || *(params.StartServer) && *(params.RevokeClient) {
		err = errors.New("You can not add/revoke a client and start the server!")
	}

	if *(params.AddClient) && *(params.RevokeClient) {
		err = errors.New("Can not revoke and add at the same time!")
	}

	if *(params.AddClient) {
		if *(params.ClientId) == "" {
			err = errors.New("Invalid ClientId!")
		}

		if *(params.ClientSecret) == "" {
			err = errors.New("Invalid ClientSecret!")
		}
	}

	if *(params.RevokeClient) {
		if *(params.ClientId) == "" {
			err = errors.New("Invalid ClientId!")
		}
	}

	return err
}
