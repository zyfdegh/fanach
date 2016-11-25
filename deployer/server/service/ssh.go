package service

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func keyConfig(user, key string) *ssh.ClientConfig {
	signer, err := ssh.ParsePrivateKey([]byte(key))
	if err != nil {
		log.Printf("parse private key error: %v", err)
		return nil
	}
	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}
}

func passwdConfig(user, password string) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
}

func executeCmd(config *ssh.ClientConfig, host, port, cmd string) (output string) {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), config)
	if err != nil {
		log.Printf("ssh dial error: %v", err)
		return
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Printf("new session error: %v", err)
		return
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(cmd)
	if err != nil {
		log.Printf("run cmd error: %v", err)
		return
	}
	return b.String()
}

func scp(config *ssh.ClientConfig, host, port, src, dest string) {

}

func executeScript(config *ssh.ClientConfig, host, port, scriptPath string) {

}
