package manager

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"net/url"
	"time"

	"github.com/hangyan/citadel"
)

func getTLSConfig(caCert, sslCert, sslKey []byte) (*tls.Config, error) {
	// TLS config
	var tlsConfig tls.Config
	tlsConfig.InsecureSkipVerify = true
	certPool := x509.NewCertPool()

	certPool.AppendCertsFromPEM(caCert)
	tlsConfig.RootCAs = certPool
	cert, err := tls.X509KeyPair(sslCert, sslKey)
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	return &tlsConfig, nil
}

func setEngineClient(docker *citadel.Engine, tlsConfig *tls.Config) error {
	var tc *tls.Config
	u, err := url.Parse(docker.Addr)
	if err != nil {
		return err
	}

	if u.Scheme == "https" {
		tc = tlsConfig
	}

	return docker.Connect(tc)
}

func getEngineMetrics(docker *citadel.Engine) error {
	info, err := docker.Info()
	if err != nil {
		return err
	}
	docker.Cpus = float64(info.NCPU)
	docker.Memory = float64(info.MemTotal) / 1024 / 1024
	return nil
}

func generateId(n int) string {
	hash := sha256.New()
	hash.Write([]byte(time.Now().String()))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr[:n]
}
