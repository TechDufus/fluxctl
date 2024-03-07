package config

type Context struct {
	// Cluster is the name of the cluster for this context
	// TODO: This may not be needed. The only cluster / server will be flux explorer api
	// unless we will use for tracking prod / testnet
	Cluster string `json:"cluster"`
	// AuthInfo is the name of the authInfo for this context
	AuthInfo string `json:"user"`
	// Namespace is the default namespace to use on unspecified requests
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// Cluster contains information about how to communicate with a kubernetes cluster
type Cluster struct {
	// Server is the address of the kubernetes cluster (https://hostname:port).
	Server string `json:"server"`
	// TLSServerName is used to check server certificate. If TLSServerName is empty, the hostname used to contact the server is used.
	// +optional
	TLSServerName string `json:"tls-server-name,omitempty"`
	// InsecureSkipTLSVerify skips the validity check for the server's certificate. This will make your HTTPS connections insecure.
	// +optional
	InsecureSkipTLSVerify bool `json:"insecure-skip-tls-verify,omitempty"`
	// CertificateAuthority is the path to a cert file for the certificate authority.
	// +optional
	CertificateAuthority string `json:"certificate-authority,omitempty"`
	// CertificateAuthorityData contains PEM-encoded certificate authority certificates. Overrides CertificateAuthority
	// +optional
	CertificateAuthorityData []byte `json:"certificate-authority-data,omitempty"`
	// ProxyURL is the URL to the proxy to be used for all requests made by this
	// client. URLs with "http", "https", and "socks5" schemes are supported.  If
	// this configuration is not provided or the empty string, the client
	// attempts to construct a proxy configuration from http_proxy and
	// https_proxy environment variables. If these environment variables are not
	// set, the client does not attempt to proxy requests.
	//
	// socks5 proxying does not currently support spdy streaming endpoints (exec,
	// attach, port forward).
	// +optional
	ProxyURL string `json:"proxy-url,omitempty"`
}

// AuthInfo contains information that describes identity information.  This is use to tell the kubernetes cluster who you are.
type AuthInfo struct {
	// ClientCertificate is the path to a client cert file for TLS.
	// +optional
	ClientCertificate string `json:"client-certificate,omitempty"`
	// ClientCertificateData contains PEM-encoded data from a client cert file for TLS. Overrides ClientCertificate
	// +optional
	ClientCertificateData []byte `json:"client-certificate-data,omitempty"`
	// ClientKey is the path to a client key file for TLS.
	// +optional
	ClientKey string `json:"client-key,omitempty"`
	// ClientKeyData contains PEM-encoded data from a client key file for TLS. Overrides ClientKey
	// +optional
	ClientKeyData []byte `json:"client-key-data,omitempty" datapolicy:"security-key"`
	// Token is the bearer token for authentication to the kubernetes cluster.
	// +optional
	Token string `json:"token,omitempty" datapolicy:"token"`
	// TokenFile is a pointer to a file that contains a bearer token (as described above).  If both Token and TokenFile are present, Token takes precedence.
	// +optional
	TokenFile string `json:"tokenFile,omitempty"`
	// Username is the username for basic authentication to the cluster.
	// +optional
	Username string `json:"username,omitempty"`
	// Password is the password for basic authentication to the cluster.
	// +optional
	Password string `json:"password,omitempty" datapolicy:"password"`
}

type Config struct {
	// Clusters is a map of referencable names to cluster configs
	Clusters map[string]*Cluster `json:"clusters"`
	// AuthInfos is a map of referencable names to user configs
	AuthInfos map[string]*AuthInfo `json:"users"`
	// Contexts is a map of referencable names to context configs
	Contexts map[string]*Context `json:"contexts"`
	// CurrentContext is the name of the context that you would like to use by default
	CurrentContext string `json:"current-context"`
}
