package localhostcert

import (
	"bytes"
	"io"
	"io/ioutil"
)

var (
	cert = `-----BEGIN CERTIFICATE-----
MIIDoDCCAoigAwIBAgIJAJ9sb6/Dn9m0MA0GCSqGSIb3DQEBCwUAMD4xGzAZBgNV
BAoTEkxvY2FsaG9zdCBEZXYgQ2VydDELMAkGA1UECxMCSVQxEjAQBgNVBAMTCUxP
Q0FMSE9TVDAeFw0xNzExMDgxMzU1MDVaFw0yNzExMDYxMzU1MDVaMD4xGzAZBgNV
BAoTEkxvY2FsaG9zdCBEZXYgQ2VydDELMAkGA1UECxMCSVQxEjAQBgNVBAMTCUxP
Q0FMSE9TVDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAK+qhSB0BrYq
0e0/AHJCfIYavzH7HjnR5LtxHNd1dINKoK241exB2i14Qj6puTSYglbVRqEBcDgq
WykuKo5uItk2b+fB+n1C/BcgKG+rAlWQOzziLT4RT8yzghdGoI+w9gT6wpwoW+SR
0Ew6EqSQy9+NxOlARldXngR4O/P4hUTsEz9JYWFMQqMwp82YrDtEMdW4UyBYxkxa
DzHC8q7gTuTSJYG86paKWvINmQjs6mIZWfUXzUEyHPrVQMUotW8fHeyg8Edz09Dx
pu93RZ161W1WbjCXBxyjH6ZIWoDj2tTpIjJ2MkCbx+Z1nNA0eoiaruMNOy90Q29Q
Tjr/kZp2E0sCAwEAAaOBoDCBnTAdBgNVHQ4EFgQUdMBf/0UnlAXYyww0LvFerpVn
J7owbgYDVR0jBGcwZYAUdMBf/0UnlAXYyww0LvFerpVnJ7qhQqRAMD4xGzAZBgNV
BAoTEkxvY2FsaG9zdCBEZXYgQ2VydDELMAkGA1UECxMCSVQxEjAQBgNVBAMTCUxP
Q0FMSE9TVIIJAJ9sb6/Dn9m0MAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQAD
ggEBAAymhUwQ7eA2ywnvL6V/JTXCKLjBerhoPSwuMTDoJGCl76/vglHi0+EO6m+n
FeIVF8K5dXRzM0b+VXENOCVugdDTvtqZ1wNkBLlhmb3pxfmRn4EG7yGVSraA4Vwh
i0pQWtF8lv2e2X28KAetEbjBLCYBtW1r5hXDXrKEd9IT/W65oCkp6PrZCJxm6vHr
4xuFi4vc7icpWKZ+kCABefTDIS5qgHVEhQjUgol5PT0y7Yg8VKBpWk8o4pdSHWRZ
yRA/dwSy51RQhnbgN/QaEsWbMzMW7i+YPewwjpnw2EvxADI3ORs5YOdHNYD+/tuH
DEiyzOevWltkDFA34a9LYHwlLIQ=
-----END CERTIFICATE-----
`
	key = `-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEAr6qFIHQGtirR7T8AckJ8hhq/MfseOdHku3Ec13V0g0qgrbjV
7EHaLXhCPqm5NJiCVtVGoQFwOCpbKS4qjm4i2TZv58H6fUL8FyAob6sCVZA7POIt
PhFPzLOCF0agj7D2BPrCnChb5JHQTDoSpJDL343E6UBGV1eeBHg78/iFROwTP0lh
YUxCozCnzZisO0Qx1bhTIFjGTFoPMcLyruBO5NIlgbzqlopa8g2ZCOzqYhlZ9RfN
QTIc+tVAxSi1bx8d7KDwR3PT0PGm73dFnXrVbVZuMJcHHKMfpkhagOPa1OkiMnYy
QJvH5nWc0DR6iJqu4w07L3RDb1BOOv+RmnYTSwIDAQABAoIBAQCWfd5ZBC1vyNVZ
i5zFRkJJ3QYpOnoVjMSI9ImB6CTuCXQIAA58vZm9VIZkCqEY9wLKgyJj8siBxX+d
kuhwZthCAAn0oVEIGOfApBMgP3/bb5ngOeAKiWg09SAg8qfBEhhiAbXRdB3tfiHV
+/ZWt5mBJoIoVaEtWGmOaLbr9t+tP6p6pJRMcP0pzZJkzQelM9PbGV2Jx/kbnC+E
iGIVx/kKIQRz8bmo1G5VVqPV99E3OEaawZlU0nkycJHa5Hw1xkazTg6LIfgkgU3g
EZou3oFhbhtEj78wsY6c+f8hN3JQEzlhqasFx2e+jWKqDimslDncqEuKdccNAD1H
quulmjdRAoGBANinOT0M9vYzwOoVH0+4EKZlJGWo9aIZFtCZKANa6hWh1NviaTSu
AkwYWsH4TZdtQSNEnNKLjCvPVjU0OZMuy7i0mZQ2+D5F9aKnnePSFUeQRsOz36i1
mXH/RWaqx5Tl56cI+aW+DLHmnmwRU4D6AQZGkDLarf5OBf+2rcgHdNPpAoGBAM+R
spL8n3xFcGM+CFgz3ccl+hLsSNeODXMS97rADnfPY5f/dhUzr/Q20Lka/hX50kqe
jQ/5Rdo9bioDYQvfGBq9RWmHMK/BBxFIcudBz78ilYdbVvvporrSnMes6oskSX5i
36rZAeY/JCfl35zccCCpgDgETN4lvTePQdFzbPETAoGBALn/lpxbLQphlBVi3ObL
1z+DRaQhUgBAGd4sHrYCr0SzEtNTpY6cdUxu/DvauIJwQaPNm/UF6OkTr6ctluBN
JIkGQ3ODXcvYs/FYhay5B5vQuW/6VOG6RmogBa0GoSGr3x0AD5PDfZKdsxSEK3Rn
Lcn9en6uTwIsaeoHI1q/TAVBAoGBAIjxkSVUskwk6/8t4AeBQKKxNYnihjRgrhLr
wrFdIK5/DzArBNb8IOw4dgCKBHSvO8SrUlnlleZkfRgO3qocaCTMFs2GueJasbT/
XX+hddSelpSU+JB/FjO33GQIez2NlUdjKDnprk0f/1SrXpp4/Skvaz6J58TwlSPg
ygXL52YtAoGBANMAzkkrHdTIXmfPYcRGLmg8xLFnoOPD4q7LMUZTNbFqvkxWoCFn
60xr0R0W6Xeg/jHAiErzxSbENdOet9bxHolZ9I1xoylFhFUipDmUivkQEPTAXATp
JUYjhFT5w50zQMPpo7SCp+l0C7oL9OBGFFY/z3UxRKyqBeU6Da0bs0le
-----END RSA PRIVATE KEY-----
`
)

func GetCert() string {
	return cert
}

func GetKey() string {
	return key
}

func GetCertStream() io.Reader {
	b := new(bytes.Buffer)
	b.Write([]byte(cert))
	return b
}

func GetKeyStream() io.Reader {
	b := new(bytes.Buffer)
	b.Write([]byte(key))
	return b
}

func GetCertTempFilePath() (string, error) {
	tcertfile, err := ioutil.TempFile("", "localhostcert_tempcert_")
	if err != nil {
		return "", err
	}
	path := tcertfile.Name()
	tcertfile.Write([]byte(cert))
	return path, tcertfile.Close()
}

func GetKeyTempFilePath() (string, error) {
	tkeyfile, err := ioutil.TempFile("", "localhostcert_tempkey_")
	if err != nil {
		return "", err
	}
	path := tkeyfile.Name()
	tkeyfile.Write([]byte(key))
	return path, tkeyfile.Close()
}
