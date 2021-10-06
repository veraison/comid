// Copyright 2021 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0

package corim

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testECKey = []byte(`{
		"kty": "EC",
		"crv": "P-256",
		"x": "MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4",
		"y": "4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM",
		"d": "870MB6gfuTJ4HtUnUvYMyJpr5eUZNP4Bk43bVdj3eAE",
		"use": "enc",
		"kid": "1"
	  }`)
)

func TestSignedCorim_FromCOSE_ok(t *testing.T) {
	/*
	   18(
	     [
	       / protected h'a10126' / << {
	         / alg / 1: -7, / ECDSA 256 /
	         / content-type / 3: "application/rim+cbor",
	         / issuer-key-id / 4: "meriadoc.brandybuck@buckland.example",
	         / corim-meta / 8: h'a200a1006941434d45204c74642e01a101c11a5fad2056'
	       } >>,
	       / unprotected / {},
	       / payload / << {
	         0: "test corim id",
	         1: [
	           h'D901FAA40065656E2D474201A1005043BBE37F2E614B33AED353CFF1428B160281A3006941434D45204C74642E01D8207468747470733A2F2F61636D652E6578616D706C65028300010204A1008182A100A300D90227582061636D652D696D706C656D656E746174696F6E2D69642D303030303030303031016441434D45026A526F616452756E6E657283A200D90258A30162424C0465322E312E30055820ACBB11C7E4DA217205523CE4CE1A245AE1A239AE3C6BFD9E7871F7E5D8BAE86B01A102818201582087428FC522803D31065E7BCE3CF03FE475096631E5E07BBD7A0FDE60C4CF25C7A200D90258A3016450526F540465312E332E35055820ACBB11C7E4DA217205523CE4CE1A245AE1A239AE3C6BFD9E7871F7E5D8BAE86B01A10281820158200263829989B6FD954F72BAAF2FC64BC2E2F01D692D4DE72986EA808F6E99813FA200D90258A3016441526F540465302E312E34055820ACBB11C7E4DA217205523CE4CE1A245AE1A239AE3C6BFD9E7871F7E5D8BAE86B01A1028182015820A3A5E715F0CC574A73C3F9BEBB6BC24F32FFD5B67B387244C2C909DA779A1478'
	         ]
	       } >>,
	       / signature / h'deadbeef'
	     ]
	   )
	*/
	tv := []byte{
		0xd2, 0x84, 0x58, 0x59, 0xa4, 0x01, 0x26, 0x03, 0x74, 0x61, 0x70, 0x70,
		0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x69, 0x6d,
		0x2b, 0x63, 0x62, 0x6f, 0x72, 0x04, 0x78, 0x24, 0x6d, 0x65, 0x72, 0x69,
		0x61, 0x64, 0x6f, 0x63, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x79, 0x62,
		0x75, 0x63, 0x6b, 0x40, 0x62, 0x75, 0x63, 0x6b, 0x6c, 0x61, 0x6e, 0x64,
		0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x08, 0x57, 0xa2, 0x00,
		0xa1, 0x00, 0x69, 0x41, 0x43, 0x4d, 0x45, 0x20, 0x4c, 0x74, 0x64, 0x2e,
		0x01, 0xa1, 0x01, 0xc1, 0x1a, 0x5f, 0xad, 0x20, 0x56, 0xa0, 0x59, 0x01,
		0xb8, 0xa2, 0x00, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x20, 0x63, 0x6f, 0x72,
		0x69, 0x6d, 0x20, 0x69, 0x64, 0x01, 0x81, 0x59, 0x01, 0xa3, 0xd9, 0x01,
		0xfa, 0xa4, 0x00, 0x65, 0x65, 0x6e, 0x2d, 0x47, 0x42, 0x01, 0xa1, 0x00,
		0x50, 0x43, 0xbb, 0xe3, 0x7f, 0x2e, 0x61, 0x4b, 0x33, 0xae, 0xd3, 0x53,
		0xcf, 0xf1, 0x42, 0x8b, 0x16, 0x02, 0x81, 0xa3, 0x00, 0x69, 0x41, 0x43,
		0x4d, 0x45, 0x20, 0x4c, 0x74, 0x64, 0x2e, 0x01, 0xd8, 0x20, 0x74, 0x68,
		0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x63, 0x6d, 0x65, 0x2e,
		0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x02, 0x83, 0x00, 0x01, 0x02,
		0x04, 0xa1, 0x00, 0x81, 0x82, 0xa1, 0x00, 0xa3, 0x00, 0xd9, 0x02, 0x27,
		0x58, 0x20, 0x61, 0x63, 0x6d, 0x65, 0x2d, 0x69, 0x6d, 0x70, 0x6c, 0x65,
		0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x69, 0x64,
		0x2d, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x01, 0x64,
		0x41, 0x43, 0x4d, 0x45, 0x02, 0x6a, 0x52, 0x6f, 0x61, 0x64, 0x52, 0x75,
		0x6e, 0x6e, 0x65, 0x72, 0x83, 0xa2, 0x00, 0xd9, 0x02, 0x58, 0xa3, 0x01,
		0x62, 0x42, 0x4c, 0x04, 0x65, 0x32, 0x2e, 0x31, 0x2e, 0x30, 0x05, 0x58,
		0x20, 0xac, 0xbb, 0x11, 0xc7, 0xe4, 0xda, 0x21, 0x72, 0x05, 0x52, 0x3c,
		0xe4, 0xce, 0x1a, 0x24, 0x5a, 0xe1, 0xa2, 0x39, 0xae, 0x3c, 0x6b, 0xfd,
		0x9e, 0x78, 0x71, 0xf7, 0xe5, 0xd8, 0xba, 0xe8, 0x6b, 0x01, 0xa1, 0x02,
		0x81, 0x82, 0x01, 0x58, 0x20, 0x87, 0x42, 0x8f, 0xc5, 0x22, 0x80, 0x3d,
		0x31, 0x06, 0x5e, 0x7b, 0xce, 0x3c, 0xf0, 0x3f, 0xe4, 0x75, 0x09, 0x66,
		0x31, 0xe5, 0xe0, 0x7b, 0xbd, 0x7a, 0x0f, 0xde, 0x60, 0xc4, 0xcf, 0x25,
		0xc7, 0xa2, 0x00, 0xd9, 0x02, 0x58, 0xa3, 0x01, 0x64, 0x50, 0x52, 0x6f,
		0x54, 0x04, 0x65, 0x31, 0x2e, 0x33, 0x2e, 0x35, 0x05, 0x58, 0x20, 0xac,
		0xbb, 0x11, 0xc7, 0xe4, 0xda, 0x21, 0x72, 0x05, 0x52, 0x3c, 0xe4, 0xce,
		0x1a, 0x24, 0x5a, 0xe1, 0xa2, 0x39, 0xae, 0x3c, 0x6b, 0xfd, 0x9e, 0x78,
		0x71, 0xf7, 0xe5, 0xd8, 0xba, 0xe8, 0x6b, 0x01, 0xa1, 0x02, 0x81, 0x82,
		0x01, 0x58, 0x20, 0x02, 0x63, 0x82, 0x99, 0x89, 0xb6, 0xfd, 0x95, 0x4f,
		0x72, 0xba, 0xaf, 0x2f, 0xc6, 0x4b, 0xc2, 0xe2, 0xf0, 0x1d, 0x69, 0x2d,
		0x4d, 0xe7, 0x29, 0x86, 0xea, 0x80, 0x8f, 0x6e, 0x99, 0x81, 0x3f, 0xa2,
		0x00, 0xd9, 0x02, 0x58, 0xa3, 0x01, 0x64, 0x41, 0x52, 0x6f, 0x54, 0x04,
		0x65, 0x30, 0x2e, 0x31, 0x2e, 0x34, 0x05, 0x58, 0x20, 0xac, 0xbb, 0x11,
		0xc7, 0xe4, 0xda, 0x21, 0x72, 0x05, 0x52, 0x3c, 0xe4, 0xce, 0x1a, 0x24,
		0x5a, 0xe1, 0xa2, 0x39, 0xae, 0x3c, 0x6b, 0xfd, 0x9e, 0x78, 0x71, 0xf7,
		0xe5, 0xd8, 0xba, 0xe8, 0x6b, 0x01, 0xa1, 0x02, 0x81, 0x82, 0x01, 0x58,
		0x20, 0xa3, 0xa5, 0xe7, 0x15, 0xf0, 0xcc, 0x57, 0x4a, 0x73, 0xc3, 0xf9,
		0xbe, 0xbb, 0x6b, 0xc2, 0x4f, 0x32, 0xff, 0xd5, 0xb6, 0x7b, 0x38, 0x72,
		0x44, 0xc2, 0xc9, 0x09, 0xda, 0x77, 0x9a, 0x14, 0x78, 0x44, 0xde, 0xad,
		0xbe, 0xef,
	}

	var actual SignedCorim
	err := actual.FromCOSE(tv)

	assert.Nil(t, err)
}

func TestSignedCorim_FromCOSE_fail_no_tag(t *testing.T) {
	// a single null byte is sufficient to test this condition because the tag
	// is the very first thing we stumble upon
	tv := []byte{0xf6}
	var actual SignedCorim
	err := actual.FromCOSE(tv)

	assert.EqualError(t, err, "failed CBOR decoding for COSE-Sign1 signed CoRIM: cbor: wrong tag number 0")
}
func TestSignedCorim_FromCOSE_fail_corim_bad_cbor(t *testing.T) {
	/*
		18(
		  [
		    / protected / << {
		      / alg / 1: -7, / ECDSA 256 /
		      / content-type / 3: "application/rim+cbor",
		      / corim-meta / 8: h'a200a1006941434d45204c74642e01a101c11a5fad2056'
		    } >>,
		    / unprotected / {},
		    / payload / h'badcb030',
		    / signature / h'deadbeef'
		  ]
		)
	*/
	tv := []byte{
		0xd2, 0x84, 0x58, 0x32, 0xa3, 0x01, 0x26, 0x03, 0x74, 0x61, 0x70, 0x70,
		0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x69, 0x6d,
		0x2b, 0x63, 0x62, 0x6f, 0x72, 0x08, 0x57, 0xa2, 0x00, 0xa1, 0x00, 0x69,
		0x41, 0x43, 0x4d, 0x45, 0x20, 0x4c, 0x74, 0x64, 0x2e, 0x01, 0xa1, 0x01,
		0xc1, 0x1a, 0x5f, 0xad, 0x20, 0x56, 0xa0, 0x44, 0xba, 0xdc, 0xb0, 0x30,
		0x44, 0xde, 0xad, 0xbe, 0xef,
	}

	var actual SignedCorim
	err := actual.FromCOSE(tv)

	assert.EqualError(t, err, "failed CBOR decoding of unsigned CoRIM: unexpected EOF")
}

func TestSignedCorim_FromCOSE_fail_invalid_corim(t *testing.T) {
	/*
		18(
		  [
		    / protected / << {
		      / alg / 1: -7, / ECDSA 256 /
		      / content-type / 3: "application/rim+cbor",
		      / corim-meta / 8: h'a200a1006941434d45204c74642e01a101c11a5fad2056'
		    } >>,
		    / unprotected / {},
		    / payload / << {
		      0: "invalid corim"
		    } >>,
		    / signature / h'deadbeef'
		  ]
		)
	*/
	tv := []byte{
		0xd2, 0x84, 0x58, 0x32, 0xa3, 0x01, 0x26, 0x03, 0x74, 0x61, 0x70, 0x70,
		0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x69, 0x6d,
		0x2b, 0x63, 0x62, 0x6f, 0x72, 0x08, 0x57, 0xa2, 0x00, 0xa1, 0x00, 0x69,
		0x41, 0x43, 0x4d, 0x45, 0x20, 0x4c, 0x74, 0x64, 0x2e, 0x01, 0xa1, 0x01,
		0xc1, 0x1a, 0x5f, 0xad, 0x20, 0x56, 0xa0, 0x50, 0xa1, 0x00, 0x6d, 0x69,
		0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x63, 0x6f, 0x72, 0x69, 0x6d,
		0x44, 0xde, 0xad, 0xbe, 0xef,
	}

	var actual SignedCorim
	err := actual.FromCOSE(tv)

	assert.EqualError(t, err, "failed validation of unsigned CoRIM: tags validation failed: no tags")
}

func TestSignedCorim_FromCOSE_fail_no_content_type(t *testing.T) {
	/*
	   18(
	     [
	       / protected / << {
	         / alg / 1: -7 / ECDSA 256 /
	       } >>,
	       / unprotected / {},
	       / payload / << {
	         0: "test corim id",
	         1: [ h'cafecafe' ]
	       } >>,
	       / signature / h'deadbeef'
	     ]
	   )
	*/
	tv := []byte{
		0xd2, 0x84, 0x43, 0xa1, 0x01, 0x26, 0xa0, 0x57, 0xa2, 0x00, 0x6d, 0x74,
		0x65, 0x73, 0x74, 0x20, 0x63, 0x6f, 0x72, 0x69, 0x6d, 0x20, 0x69, 0x64,
		0x01, 0x81, 0x44, 0xca, 0xfe, 0xca, 0xfe, 0x44, 0xde, 0xad, 0xbe, 0xef,
	}
	var actual SignedCorim
	err := actual.FromCOSE(tv)

	assert.EqualError(t, err, "processing COSE headers: missing mandatory content type")
}

func TestSignedCorim_FromCOSE_fail_unexpected_content_type(t *testing.T) {
	/*
	   18(
	     [
	       / protected / << {
	         / alg / 1: -7, / ECDSA 256 /
	         / content-type / 3: "application/cbor"
	       } >>,
	       / unprotected / {},
	       / payload / << {
	         0: "test corim id",
	         1: [ h'cafecafe' ]
	       } >>,
	       / signature / h'deadbeef'
	     ]
	   )
	*/
	tv := []byte{
		0xd2, 0x84, 0x55, 0xa2, 0x01, 0x26, 0x03, 0x70, 0x61, 0x70, 0x70, 0x6c,
		0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x62, 0x6f, 0x72,
		0xa0, 0x57, 0xa2, 0x00, 0x6d, 0x74, 0x65, 0x73, 0x74, 0x20, 0x63, 0x6f,
		0x72, 0x69, 0x6d, 0x20, 0x69, 0x64, 0x01, 0x81, 0x44, 0xca, 0xfe, 0xca,
		0xfe, 0x44, 0xde, 0xad, 0xbe, 0xef,
	}
	var actual SignedCorim
	err := actual.FromCOSE(tv)

	assert.EqualError(t, err, `processing COSE headers: expecting content type "application/rim+cbor", got "application/cbor" instead`)
}

func unsignedCorimFromCBOR(t *testing.T, cbor []byte) *UnsignedCorim {
	var unsignedCorim UnsignedCorim

	err := unsignedCorim.FromCBOR(cbor)
	require.Nil(t, err)
	require.Nil(t, unsignedCorim.Valid())

	return &unsignedCorim
}

func metaGood(t *testing.T) *Meta {
	var (
		name     = "ACME Ltd."
		notAfter = time.Date(2021, time.October, 0, 0, 0, 0, 0, time.UTC)
	)

	m := NewMeta().
		SetSigner(name, nil).
		SetValidity(notAfter, nil)
	require.NotNil(t, m)

	return m
}

func TestSignedCorim_SignVerify_ok(t *testing.T) {
	signer, err := SignerFromJWK(testECKey)
	require.NoError(t, err)

	var SignedCorimIn SignedCorim

	SignedCorimIn.UnsignedCorim = *unsignedCorimFromCBOR(t, testGoodUnsignedCorim)
	SignedCorimIn.Meta = *metaGood(t)

	cbor, err := SignedCorimIn.Sign(signer)
	assert.Nil(t, err)

	var SignedCorimOut SignedCorim

	fmt.Printf("signed-corim: %x\n", cbor)

	err = SignedCorimOut.FromCOSE(cbor)
	assert.Nil(t, err)

	err = SignedCorimOut.Verify(signer.Verifier().PublicKey)
	assert.Nil(t, err)
}

func TestSignedCorim_SignVerify_fail_tampered(t *testing.T) {
	signer, err := SignerFromJWK(testECKey)
	require.NoError(t, err)

	var SignedCorimIn SignedCorim

	SignedCorimIn.UnsignedCorim = *unsignedCorimFromCBOR(t, testGoodUnsignedCorim)

	cbor, err := SignedCorimIn.Sign(signer)
	assert.Nil(t, err)

	var SignedCorimOut SignedCorim

	fmt.Printf("signed-corim: %x", cbor)

	// Flip the last byte in the signature field
	cbor[len(cbor)-1] ^= 0xff

	// Since we don't modify the Sign1 payload structurally, decoding the COSE
	// envelope is still OK...
	err = SignedCorimOut.FromCOSE(cbor)
	assert.Nil(t, err)

	// ... but the signature verification fails
	err = SignedCorimOut.Verify(signer.Verifier().PublicKey)
	assert.EqualError(t, err, "verification failed ecdsa.Verify")
}

func TestSignedCorim_Sign_fail_bad_corim(t *testing.T) {
	signer, err := SignerFromJWK(testECKey)
	require.NoError(t, err)

	var SignedCorimIn SignedCorim

	emptyCorim := NewUnsignedCorim()
	require.NotNil(t, emptyCorim)

	SignedCorimIn.UnsignedCorim = *emptyCorim

	_, err = SignedCorimIn.Sign(signer)
	assert.EqualError(t, err, "failed validation of unsigned CoRIM: empty id")
}

func TestSignedCorim_Sign_fail_no_signer(t *testing.T) {
	var SignedCorimIn SignedCorim

	emptyCorim := NewUnsignedCorim()
	require.NotNil(t, emptyCorim)

	SignedCorimIn.UnsignedCorim = *emptyCorim

	_, err := SignedCorimIn.Sign(nil)
	assert.EqualError(t, err, "nil signer")
}
