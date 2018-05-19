package main

import (
	"database/sql"
	"log"

	"github.com/RangelReale/osin"
	ldap "github.com/zonradkuse/go-ldap-authenticator"
	mauth "github.com/zonradkuse/oauth-authenticator"
)

func main() {
	var cli CLIParameters
	err := handleCLIParameters(&cli)

	if err != nil {
		showDefaults()
		log.Fatal(err)
	}

	config := parseConfig(*cli.ConfigPath)

	log.Println("Initializing SQL connection")
	url := config.Mysql["user"] + ":" + config.Mysql["password"] + "@tcp(" + config.Mysql["host"] + ":" + config.Mysql["port"] + ")/" + config.Mysql["oauthTable"] + "?parseTime=true"

	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	cfg := osin.NewServerConfig()
	cfg.AllowGetAccessRequest = true
	cfg.AllowClientSecretInParams = true

	log.Println("Starting OAuth...")
	ldapAuthenticator := ldap.NewLDAPAuthenticator(config.Ldap["bindDn"], config.Ldap["bindPassword"], config.Ldap["queryDn"])
	ldapAuthenticator.Connect(config.Ldap["bindUrl"])
	oauthServer := mauth.NewOAuthServer(db, config.Mysql["oauthSchemaPrefix"], cfg, ldapAuthenticator, handleLoginLanding)

	if *cli.StartServer {
		startServer(&oauthServer)
	}

	if *cli.AddClient {
		oauthServer.CreateClient(*cli.ClientId, *cli.ClientSecret, *cli.RedirectUri)
	}

	if *cli.RevokeClient {
		oauthServer.RemoveClient(*cli.ClientId)
	}
}
