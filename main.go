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
	log.Printf("Got config: %+v", config)

	log.Println("Initializing SQL connection")
	url := config.Mysql.User + ":" + config.Mysql.Password + "@tcp(" + config.Mysql.Host + ":" + config.Mysql.Port + ")/" + config.Mysql.OauthTable + "?parseTime=true"

	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	cfg := osin.NewServerConfig()
	cfg.AllowGetAccessRequest = true
	cfg.AllowClientSecretInParams = true

	log.Println("Starting OAuth...")
	selectors := []string{"mail", "createTimestamp", "entryUUID", "cn"}
	transformer := LDAPTransformer{}
	ldapAuthenticator := ldap.NewLDAPAuthenticator(config.Ldap.BindDn, config.Ldap.BindPassword, config.Ldap.QueryDn, selectors, transformer)
	err = ldapAuthenticator.Connect(config.Ldap.BindUrl)
	if err != nil {
		log.Fatal(err)
	}

	oauthServer := mauth.NewOAuthServer(db, config.Mysql.OauthSchemaPrefix, cfg, &ldapAuthenticator, handleLoginLanding)

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
