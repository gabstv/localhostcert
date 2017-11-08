# localhostcert
```bash
openssl req -x509 -nodes -sha256 -days 3650 -newkey rsa:2048 -keyout localhost.key -out localhost.cert
```