package ssstuck

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/ssh"
)

// Config is the basic structure to hold required startup variables
type Config struct {
	Port int
}

// CheckConfig validates the input config
func CheckConfig(config Config) error {
	if config.Port < 1 || config.Port > 65535 {
		return fmt.Errorf("port(%d) please specify a port between 1-65535", config.Port)
	}
	return nil
}

// Serve is the main entrypoint to the program
func Serve(config Config) {
	err := CheckConfig(config)
	if err != nil {
		log.Panicf("invalid setting: %s", err)
	}
	serverConfig := getServer()
}

func getServer() ssh.ServerConfig {
	serverConfig := &ssh.ServerConfig{
		PublicKeyCallback: authKey,
		PasswordCallback:  authPassword,
	}

	keyBytes, _ := rsa.GenerateKey(rand.Reader, 4096)
	key, _ := ssh.NewSignerFromSigner(keyBytes)
	serverConfig.AddHostKey(key)

	return *serverConfig
}

func authKey(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
	log.WithFields(log.Fields{
		"type":   "public key",
		"ip":     conn.RemoteAddr(),
		"key":    key,
		"client": string(conn.ClientVersion()),
	}).Info("connection attempt")
	return nil, fmt.Errorf("invalid credentials")
}

func authPassword(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	log.WithFields(log.Fields{
		"type":     "user-password",
		"ip":       conn.RemoteAddr(),
		"user":     string(conn.User()),
		"password": string(password),
		"client":   string(conn.ClientVersion()),
	}).Info("connection attempt")
	return nil, fmt.Errorf("invalid credentials")
}
