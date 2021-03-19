package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

// GetServerCreds 获取服务端证书配置
func GetServerCreds() credentials.TransportCredentials {
	certificate, err := tls.LoadX509KeyPair("config/server.pem", "config/server.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("config/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
}

// GetClientCreds 獲取客戶端證書
func GetClientCreds() credentials.TransportCredentials {
	certificate, err := tls.LoadX509KeyPair("config/client.pem", "config/client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("config/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate}, // 客户端证书
		RootCAs:      certPool,
		ServerName:   "localhost",
	})
}
