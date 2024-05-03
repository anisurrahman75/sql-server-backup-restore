## Generate CA private key (ca.key):
```bash
$ openssl genrsa -out ca.key 2048

$ openssl req -new -x509 -days 365 -key ca.key -subj "/CN=backup.local" -out ca.crt
```

## Create a server CSR (server.csr):
```bash
$ openssl req -newkey rsa:2048 -nodes -keyout server.key -subj "/CN=backup.local" -out server.csr
```

## Sign the server CSR with the CA certificate and key to generate the server certificate:

```bash
$ openssl x509 -req -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt
```

> Now, use `server.key` and `server.crt`
