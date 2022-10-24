package cronet

import (
	"crypto/x509"
	"unsafe"
)

// #include <webtx_c.h>
import "C"

type verifier struct {
	ptr unsafe.Pointer
}

func CreatePinnedCertVerifier(cert *x509.Certificate, insecureSkipVerify bool) (*verifier, error) {
	var skip int8 = 0
	if insecureSkipVerify {
		skip = 1
	}
	var ptr unsafe.Pointer
	if cert != nil {
		ptr = C.WebTx_Create_PinnedCertVerifier(
			(*C.uint8_t)(unsafe.Pointer(&cert.Raw[0])),
			C.uint64_t(len(cert.Raw)),
			C.uint8_t(skip),
		)
	} else {
		ptr = C.WebTx_Create_PinnedCertVerifier(
			nil, C.uint64_t(0), C.uint8_t(skip),
		)
	}

	v := &verifier{
		ptr: ptr,
	}

	return v, nil
}
