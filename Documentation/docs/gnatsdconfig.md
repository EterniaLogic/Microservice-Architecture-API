# GNATsd Config

#### http://nats.io/ - Server download location

sync1i must be set in the /etc/hosts file on the private interface. `*.*.*.*` sync1i.


#### Temporary server

./server -d config.cfg

#### Service

By placing the gnatsd server at `/home/goserv/server`, it will host a gnatsd server on a service if you use the service config from server configuration with "/home/goserv/server -d config.cfg"


``` YAML
# TLS config file

port: 28129 # port to listen for client connections
net: sync1i # net interface to listen

#http_port: 9090 # HTTP monitoring port

# Authorization for client connections
authorization {
  user:     a3cqEAdz1om
  password: GOXxG3HamsCZtu2LlEW
  timeout:  0.5
}

tls {
  cert_file:  "./certs/server-cert.pem"
  key_file:   "./certs/server-key.pem"
  ca_file:    "./certs/ca.pem"
  timeout:    0.5
}

# logging options
debug:   false
trace:   true
logtime: true
log_file: "/home/goserv/gnatsd.log"

# pid file
pid_file: "/tmp/gnatsd.pid"

# Some system overides
# max_connections
max_connections: 10000

# maximum protocol control line
max_control_line: 512

# maximum payload
max_payload: 65536

# slow consumer threshold
max_pending_size: 10000000

```


#### Keygen for gnatsd Server

``` Bash
rm *.csr *.pem *.srl *.cnf

openssl genrsa -aes256 -out ca-key.pem 4096
openssl req -new -x509 -days 99999 -key ca-key.pem -sha256 -out ca.pem
openssl genrsa -out server-key.pem 4096
openssl req -subj "/CN=sync1i" -sha256 -new -key server-key.pem -out server.csr

echo subjectAltName = IP:10.132.162.38,IP:127.0.0.1 > extfile.cnf
openssl x509 -req -days 99999 -sha256 -in server.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile extfile.cnf

rm -v server.csr
```


#### Keygen for nats.io Client

``` Bash
openssl genrsa -out client_key.pem 4096
openssl req -subj '/CN=client' -new -key client_key.pem -out client.csr
echo extendedKeyUsage=clientAuth > extfile.cnf
openssl x509 -req -days 90000 -sha256 -in client.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out client_cert.pem -extfile extfile.cnf

rm -v client.csr extfile.cnf
```