// Copyright 2023-2024 Contributors to the Veraison project.
// SPDX-License-Identifier: Apache-2.0
package cots

import (
	"fmt"
	"os"

	"github.com/veraison/corim/comid"
	"github.com/veraison/swid"
)

func makeZestyEntityWithRoles(roles ...interface{}) swid.Entity {
	e := swid.Entity{
		EntityName: "Zesty Hands, Inc.",
	}

	_ = e.SetRoles(roles...)

	return e
}

func Example_encode_single_organization_keys_spki() {
	tagVersion := uint(5)
	exampleTaData := []byte{0x30, 0x59, 0x30, 0x13, 0x6, 0x7, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x2, 0x1, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x3, 0x1, 0x7, 0x3, 0x42, 0x0, 0x4, 0xad, 0x8a, 0xc, 0x1, 0xda, 0x9e, 0xda, 0x2, 0x53, 0xdc, 0x2b, 0xc2, 0x72, 0x27, 0xd9, 0xc7, 0x21, 0x3d, 0xf8, 0xdf, 0x13, 0xe8, 0x9c, 0xb9, 0xcd, 0xb7, 0xa8, 0xe4, 0xb6, 0x2d, 0x9c, 0xe8, 0xa9, 0x9a, 0x2d, 0x70, 0x5c, 0xf, 0x7f, 0x80, 0xdb, 0x65, 0xc0, 0x6, 0xd1, 0x9, 0x14, 0x22, 0xb4, 0x7f, 0xc6, 0x11, 0xcb, 0xd4, 0x68, 0x69, 0x73, 0x3d, 0x9c, 0x48, 0x38, 0x84, 0xd5, 0xfe}
	cots := NewConciseTaStore().
		SetTagIdentity("ab0f44b1-bfdc-4604-ab4a-30f80407ebcc", &tagVersion).
		AddEnvironmentGroup(
			*NewEnvironmentGroup().
				SetEnvironment(
					comid.Environment{
						Class: comid.NewClassOID("1.2.3.4.5").
							SetVendor("Worthless Sea, Inc."),
					},
				),
		).
		SetKeys(
			TasAndCas{
				Tas: []TrustAnchor{
					*NewTrustAnchor().
						SetData(exampleTaData).
						SetFormat(2),
				},
			},
		)

	cbor, err := cots.ToCBOR()
	if err == nil {
		fmt.Printf("%x\n", cbor)
	}

	json, err := cots.ToJSON()
	if err == nil {
		fmt.Printf("%s\n", string(json))
	}

	// Output:
	// a301a20050ab0f44b1bfdc4604ab4a30f80407ebcc01050281a101a100a200d86f442a0304050173576f7274686c657373205365612c20496e632e06a100818202585b3059301306072a8648ce3d020106082a8648ce3d03010703420004ad8a0c01da9eda0253dc2bc27227d9c7213df8df13e89cb9cdb7a8e4b62d9ce8a99a2d705c0f7f80db65c006d1091422b47fc611cbd46869733d9c483884d5fe
	// {"tag-identity":{"id":"ab0f44b1-bfdc-4604-ab4a-30f80407ebcc","version":5},"environments":[{"environment":{"class":{"id":{"type":"oid","value":"1.2.3.4.5"},"vendor":"Worthless Sea, Inc."}}}],"keys":{"tas":[{"format":2,"data":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAErYoMAdqe2gJT3CvCcifZxyE9+N8T6Jy5zbeo5LYtnOipmi1wXA9/gNtlwAbRCRQitH/GEcvUaGlzPZxIOITV/g=="}]}}
}

func Example_encode_multiple_organizations_keys_cert_and_ta() {
	exampleTaData1 := []byte{0x30, 0x82, 0x1, 0xbd, 0x30, 0x82, 0x1, 0x64, 0xa0, 0x3, 0x2, 0x1, 0x2, 0x2, 0x15, 0x0, 0xd0, 0x9d, 0x90, 0xbf, 0x3d, 0x52, 0x5c, 0xc7, 0x73, 0xd5, 0x22, 0xed, 0x77, 0xd5, 0x9e, 0x22, 0xbb, 0xa4, 0x5b, 0x88, 0x30, 0xa, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x4, 0x3, 0x2, 0x30, 0x3e, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x10, 0x30, 0xe, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x7, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x1d, 0x30, 0x1b, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x14, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x30, 0x1e, 0x17, 0xd, 0x32, 0x32, 0x30, 0x35, 0x31, 0x39, 0x31, 0x35, 0x31, 0x33, 0x30, 0x37, 0x5a, 0x17, 0xd, 0x33, 0x32, 0x30, 0x35, 0x31, 0x36, 0x31, 0x35, 0x31, 0x33, 0x30, 0x37, 0x5a, 0x30, 0x3e, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x10, 0x30, 0xe, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x7, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x31, 0x1d, 0x30, 0x1b, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x14, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x30, 0x59, 0x30, 0x13, 0x6, 0x7, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x2, 0x1, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x3, 0x1, 0x7, 0x3, 0x42, 0x0, 0x4, 0xe3, 0x51, 0xaa, 0x10, 0x39, 0x24, 0x7, 0xa4, 0xbd, 0x3, 0x7c, 0xa0, 0xbc, 0x11, 0x54, 0xd0, 0xe7, 0x1, 0xf0, 0x67, 0x4f, 0x39, 0xcb, 0x2c, 0x4f, 0x92, 0x10, 0x69, 0x2c, 0xeb, 0xbe, 0xec, 0x1d, 0x27, 0x97, 0x7d, 0xc5, 0x61, 0x65, 0x75, 0x1e, 0xe, 0x23, 0x7b, 0xfd, 0xfb, 0x15, 0x36, 0xe9, 0x9a, 0x88, 0x45, 0x91, 0x42, 0x96, 0x61, 0xdf, 0x35, 0xcf, 0xc0, 0xbf, 0x2b, 0x50, 0xcc, 0xa3, 0x3f, 0x30, 0x3d, 0x30, 0x1d, 0x6, 0x3, 0x55, 0x1d, 0xe, 0x4, 0x16, 0x4, 0x14, 0x1, 0x5c, 0x45, 0xc9, 0xac, 0xb0, 0x46, 0x2a, 0x71, 0x5d, 0xd7, 0x10, 0xa0, 0x78, 0xc0, 0x15, 0x49, 0xf1, 0x1, 0x3f, 0x30, 0xb, 0x6, 0x3, 0x55, 0x1d, 0xf, 0x4, 0x4, 0x3, 0x2, 0x2, 0x84, 0x30, 0xf, 0x6, 0x3, 0x55, 0x1d, 0x13, 0x1, 0x1, 0xff, 0x4, 0x5, 0x30, 0x3, 0x1, 0x1, 0xff, 0x30, 0xa, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x4, 0x3, 0x2, 0x3, 0x47, 0x0, 0x30, 0x44, 0x2, 0x20, 0xb, 0x6, 0x27, 0x40, 0x6, 0xc7, 0xe9, 0xcc, 0x6d, 0x25, 0x4c, 0xbf, 0x44, 0x87, 0xd6, 0xfa, 0x1, 0x46, 0xea, 0x9f, 0x31, 0x7e, 0x92, 0x81, 0x19, 0x6e, 0xb, 0xe9, 0xa6, 0xfb, 0xed, 0xf5, 0x2, 0x20, 0x56, 0x81, 0x3e, 0x6e, 0x11, 0xd, 0xd2, 0x3e, 0xb0, 0x48, 0xfc, 0xde, 0x3e, 0x32, 0xeb, 0x11, 0xd0, 0xfe, 0x3c, 0x48, 0x32, 0x8c, 0x72, 0x79, 0xad, 0xb0, 0x35, 0xd5, 0x23, 0xea, 0xff, 0x53}
	exampleTaData2 := []byte{0xa2, 0x82, 0x2, 0xb6, 0x30, 0x82, 0x2, 0xb2, 0x30, 0x59, 0x30, 0x13, 0x6, 0x7, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x2, 0x1, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x3, 0x1, 0x7, 0x3, 0x42, 0x0, 0x4, 0x97, 0xcf, 0x6d, 0x70, 0xd7, 0x6a, 0x30, 0x40, 0xc, 0x79, 0xf1, 0xeb, 0xab, 0x6a, 0xd6, 0x16, 0x88, 0x71, 0x24, 0x87, 0x10, 0xd4, 0xf4, 0xc1, 0x20, 0x7a, 0xda, 0xef, 0x2, 0xd1, 0x98, 0x67, 0x13, 0x67, 0x91, 0xf2, 0xc2, 0xb2, 0x9f, 0x6f, 0x2d, 0xcc, 0x3c, 0x21, 0x25, 0xc2, 0xf2, 0xd2, 0xfd, 0xee, 0x4f, 0x59, 0x9, 0x4b, 0x67, 0xaf, 0x43, 0x33, 0x2f, 0xd, 0xfb, 0x5c, 0xa4, 0x40, 0x4, 0x14, 0xf6, 0xda, 0xd1, 0xe5, 0x12, 0x8b, 0xbf, 0xd, 0xe9, 0xe9, 0x53, 0x43, 0xb3, 0x71, 0xc6, 0xf7, 0xff, 0xe7, 0xe2, 0x6e, 0x30, 0x82, 0x2, 0x3d, 0x30, 0x52, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x1a, 0x30, 0x18, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x11, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x31, 0x27, 0x30, 0x25, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x1e, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0xa0, 0x82, 0x1, 0xe5, 0x30, 0x82, 0x1, 0x8b, 0xa0, 0x3, 0x2, 0x1, 0x2, 0x2, 0x14, 0xb, 0xdc, 0x4a, 0xa0, 0x51, 0x79, 0x50, 0x3e, 0x58, 0xf2, 0x75, 0xd5, 0x52, 0x46, 0x73, 0x47, 0xbc, 0xaf, 0xb5, 0x33, 0x30, 0xa, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x4, 0x3, 0x2, 0x30, 0x52, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x1a, 0x30, 0x18, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x11, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x31, 0x27, 0x30, 0x25, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x1e, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x30, 0x1e, 0x17, 0xd, 0x32, 0x32, 0x30, 0x35, 0x31, 0x39, 0x31, 0x35, 0x31, 0x33, 0x30, 0x37, 0x5a, 0x17, 0xd, 0x33, 0x32, 0x30, 0x35, 0x31, 0x36, 0x31, 0x35, 0x31, 0x33, 0x30, 0x37, 0x5a, 0x30, 0x52, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x1a, 0x30, 0x18, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x11, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x31, 0x27, 0x30, 0x25, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x1e, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x30, 0x59, 0x30, 0x13, 0x6, 0x7, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x2, 0x1, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x3, 0x1, 0x7, 0x3, 0x42, 0x0, 0x4, 0x97, 0xcf, 0x6d, 0x70, 0xd7, 0x6a, 0x30, 0x40, 0xc, 0x79, 0xf1, 0xeb, 0xab, 0x6a, 0xd6, 0x16, 0x88, 0x71, 0x24, 0x87, 0x10, 0xd4, 0xf4, 0xc1, 0x20, 0x7a, 0xda, 0xef, 0x2, 0xd1, 0x98, 0x67, 0x13, 0x67, 0x91, 0xf2, 0xc2, 0xb2, 0x9f, 0x6f, 0x2d, 0xcc, 0x3c, 0x21, 0x25, 0xc2, 0xf2, 0xd2, 0xfd, 0xee, 0x4f, 0x59, 0x9, 0x4b, 0x67, 0xaf, 0x43, 0x33, 0x2f, 0xd, 0xfb, 0x5c, 0xa4, 0x40, 0xa3, 0x3f, 0x30, 0x3d, 0x30, 0x1d, 0x6, 0x3, 0x55, 0x1d, 0xe, 0x4, 0x16, 0x4, 0x14, 0xf6, 0xda, 0xd1, 0xe5, 0x12, 0x8b, 0xbf, 0xd, 0xe9, 0xe9, 0x53, 0x43, 0xb3, 0x71, 0xc6, 0xf7, 0xff, 0xe7, 0xe2, 0x6e, 0x30, 0xb, 0x6, 0x3, 0x55, 0x1d, 0xf, 0x4, 0x4, 0x3, 0x2, 0x2, 0x84, 0x30, 0xf, 0x6, 0x3, 0x55, 0x1d, 0x13, 0x1, 0x1, 0xff, 0x4, 0x5, 0x30, 0x3, 0x1, 0x1, 0xff, 0x30, 0xa, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x4, 0x3, 0x2, 0x3, 0x48, 0x0, 0x30, 0x45, 0x2, 0x20, 0x1d, 0xa5, 0x8b, 0xe7, 0xfa, 0x44, 0x2c, 0x6c, 0xd8, 0x49, 0xef, 0x35, 0x67, 0x22, 0x4a, 0x92, 0x3, 0xc2, 0x25, 0x15, 0x89, 0x66, 0xb2, 0x1a, 0xfd, 0x40, 0xf3, 0x19, 0x2c, 0xf3, 0x47, 0x98, 0x2, 0x21, 0x0, 0x98, 0x27, 0xb3, 0xe0, 0xa1, 0xab, 0xa2, 0x51, 0x4a, 0x39, 0x94, 0xfa, 0x6e, 0xfa, 0x9f, 0xd6, 0xc6, 0x10, 0xb8, 0x90, 0x5f, 0xbe, 0xd9, 0x3f, 0xcb, 0x52, 0x50, 0x75, 0x4b, 0xe8, 0xaa, 0x58}
	exampleTaData3 := []byte{0xa2, 0x82, 0x2, 0xd5, 0x30, 0x82, 0x2, 0xd1, 0x30, 0x59, 0x30, 0x13, 0x6, 0x7, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x2, 0x1, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x3, 0x1, 0x7, 0x3, 0x42, 0x0, 0x4, 0xcd, 0xd1, 0xfe, 0x64, 0xcf, 0x2c, 0x4, 0xcd, 0x93, 0x98, 0x6d, 0xa5, 0x76, 0xdc, 0xed, 0xcd, 0xf1, 0xc9, 0x2e, 0xdf, 0xd2, 0x68, 0x2c, 0xd8, 0xe5, 0x1c, 0xfc, 0x4, 0x9, 0xb5, 0xc3, 0xd, 0x6a, 0x74, 0x2e, 0x90, 0xe, 0xd7, 0x3d, 0xb8, 0xf3, 0xf8, 0x68, 0xcb, 0xa5, 0x16, 0xfd, 0x36, 0x4c, 0x4c, 0xf3, 0xd1, 0xff, 0xdd, 0xd7, 0xc0, 0x7b, 0x6, 0xd7, 0xa9, 0x92, 0x17, 0x2f, 0x17, 0x4, 0x14, 0x8a, 0x84, 0xcf, 0xf9, 0x80, 0x95, 0xa3, 0xbc, 0x36, 0xd6, 0xee, 0xa5, 0x18, 0xd6, 0x97, 0x8d, 0x9b, 0xd7, 0x1f, 0x60, 0x30, 0x82, 0x2, 0x5c, 0x30, 0x5c, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x1f, 0x30, 0x1d, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x16, 0x53, 0x6e, 0x6f, 0x62, 0x62, 0x69, 0x73, 0x68, 0x20, 0x41, 0x70, 0x70, 0x61, 0x72, 0x65, 0x6c, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x31, 0x2c, 0x30, 0x2a, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x23, 0x53, 0x6e, 0x6f, 0x62, 0x62, 0x69, 0x73, 0x68, 0x20, 0x41, 0x70, 0x70, 0x61, 0x72, 0x65, 0x6c, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0xa0, 0x82, 0x1, 0xfa, 0x30, 0x82, 0x1, 0x9f, 0xa0, 0x3, 0x2, 0x1, 0x2, 0x2, 0x14, 0x10, 0x1b, 0x93, 0x44, 0x65, 0xc0, 0x10, 0x45, 0x44, 0x1e, 0x1b, 0xb8, 0xc5, 0xa7, 0xc0, 0x9e, 0xa9, 0xbe, 0xa9, 0x88, 0x30, 0xa, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x4, 0x3, 0x2, 0x30, 0x5c, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x1f, 0x30, 0x1d, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x16, 0x53, 0x6e, 0x6f, 0x62, 0x62, 0x69, 0x73, 0x68, 0x20, 0x41, 0x70, 0x70, 0x61, 0x72, 0x65, 0x6c, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x31, 0x2c, 0x30, 0x2a, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x23, 0x53, 0x6e, 0x6f, 0x62, 0x62, 0x69, 0x73, 0x68, 0x20, 0x41, 0x70, 0x70, 0x61, 0x72, 0x65, 0x6c, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x30, 0x1e, 0x17, 0xd, 0x32, 0x32, 0x30, 0x35, 0x31, 0x39, 0x31, 0x35, 0x31, 0x33, 0x30, 0x38, 0x5a, 0x17, 0xd, 0x33, 0x32, 0x30, 0x35, 0x31, 0x36, 0x31, 0x35, 0x31, 0x33, 0x30, 0x38, 0x5a, 0x30, 0x5c, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x1f, 0x30, 0x1d, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x16, 0x53, 0x6e, 0x6f, 0x62, 0x62, 0x69, 0x73, 0x68, 0x20, 0x41, 0x70, 0x70, 0x61, 0x72, 0x65, 0x6c, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x31, 0x2c, 0x30, 0x2a, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x23, 0x53, 0x6e, 0x6f, 0x62, 0x62, 0x69, 0x73, 0x68, 0x20, 0x41, 0x70, 0x70, 0x61, 0x72, 0x65, 0x6c, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x30, 0x59, 0x30, 0x13, 0x6, 0x7, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x2, 0x1, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x3, 0x1, 0x7, 0x3, 0x42, 0x0, 0x4, 0xcd, 0xd1, 0xfe, 0x64, 0xcf, 0x2c, 0x4, 0xcd, 0x93, 0x98, 0x6d, 0xa5, 0x76, 0xdc, 0xed, 0xcd, 0xf1, 0xc9, 0x2e, 0xdf, 0xd2, 0x68, 0x2c, 0xd8, 0xe5, 0x1c, 0xfc, 0x4, 0x9, 0xb5, 0xc3, 0xd, 0x6a, 0x74, 0x2e, 0x90, 0xe, 0xd7, 0x3d, 0xb8, 0xf3, 0xf8, 0x68, 0xcb, 0xa5, 0x16, 0xfd, 0x36, 0x4c, 0x4c, 0xf3, 0xd1, 0xff, 0xdd, 0xd7, 0xc0, 0x7b, 0x6, 0xd7, 0xa9, 0x92, 0x17, 0x2f, 0x17, 0xa3, 0x3f, 0x30, 0x3d, 0x30, 0x1d, 0x6, 0x3, 0x55, 0x1d, 0xe, 0x4, 0x16, 0x4, 0x14, 0x8a, 0x84, 0xcf, 0xf9, 0x80, 0x95, 0xa3, 0xbc, 0x36, 0xd6, 0xee, 0xa5, 0x18, 0xd6, 0x97, 0x8d, 0x9b, 0xd7, 0x1f, 0x60, 0x30, 0xb, 0x6, 0x3, 0x55, 0x1d, 0xf, 0x4, 0x4, 0x3, 0x2, 0x2, 0x84, 0x30, 0xf, 0x6, 0x3, 0x55, 0x1d, 0x13, 0x1, 0x1, 0xff, 0x4, 0x5, 0x30, 0x3, 0x1, 0x1, 0xff, 0x30, 0xa, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x4, 0x3, 0x2, 0x3, 0x49, 0x0, 0x30, 0x46, 0x2, 0x21, 0x0, 0xb6, 0x71, 0xfe, 0x37, 0x7f, 0x73, 0xcf, 0x94, 0x23, 0xba, 0xfd, 0xdc, 0x6f, 0xe3, 0x47, 0xed, 0x22, 0xc, 0x71, 0x4e, 0x82, 0x87, 0x17, 0xbd, 0x94, 0xcc, 0x43, 0x6f, 0xef, 0xec, 0xb8, 0xca, 0x2, 0x21, 0x0, 0x81, 0xad, 0xb, 0xfa, 0xfa, 0x48, 0x66, 0x85, 0x31, 0xa3, 0xfb, 0xcc, 0x3d, 0x84, 0x96, 0x80, 0x6e, 0x21, 0x74, 0xed, 0xc9, 0x6b, 0xbc, 0x1f, 0x9, 0x7e, 0x60, 0x67, 0x57, 0x45, 0x8f, 0x77}
	cots := NewConciseTaStore().
		SetTagIdentity("ab0f44b1-bfdc-4604-ab4a-30f80407ebcc", nil).
		AddEnvironmentGroup(
			*NewEnvironmentGroup().
				SetNamedTaStore("Miscellaneous TA Store"),
		).
		SetKeys(
			TasAndCas{
				Tas: []TrustAnchor{
					*NewTrustAnchor().
						SetData(exampleTaData1).
						SetFormat(0),

					*NewTrustAnchor().
						SetData(exampleTaData2).
						SetFormat(1),

					*NewTrustAnchor().
						SetData(exampleTaData3).
						SetFormat(1),
				},
			},
		)

	cbor, err := cots.ToCBOR()
	if err == nil {
		fmt.Printf("%x\n", cbor)
	}

	json, err := cots.ToJSON()
	if err == nil {
		fmt.Printf("%s\n", string(json))
	}

	// Output:
	// a301a10050ab0f44b1bfdc4604ab4a30f80407ebcc0281a103764d697363656c6c616e656f75732054412053746f726506a1008382005901c1308201bd30820164a003020102021500d09d90bf3d525cc773d522ed77d59e22bba45b88300a06082a8648ce3d040302303e310b300906035504060c0255533110300e060355040a0c074578616d706c65311d301b06035504030c144578616d706c6520547275737420416e63686f72301e170d3232303531393135313330375a170d3332303531363135313330375a303e310b300906035504060c0255533110300e060355040a0c074578616d706c65311d301b06035504030c144578616d706c6520547275737420416e63686f723059301306072a8648ce3d020106082a8648ce3d03010703420004e351aa10392407a4bd037ca0bc1154d0e701f0674f39cb2c4f9210692cebbeec1d27977dc56165751e0e237bfdfb1536e99a884591429661df35cfc0bf2b50cca33f303d301d0603551d0e04160414015c45c9acb0462a715dd710a078c01549f1013f300b0603551d0f040403020284300f0603551d130101ff040530030101ff300a06082a8648ce3d040302034700304402200b06274006c7e9cc6d254cbf4487d6fa0146ea9f317e9281196e0be9a6fbedf5022056813e6e110dd23eb048fcde3e32eb11d0fe3c48328c7279adb035d523eaff5382015902baa28202b6308202b23059301306072a8648ce3d020106082a8648ce3d0301070342000497cf6d70d76a30400c79f1ebab6ad6168871248710d4f4c1207adaef02d19867136791f2c2b29f6f2dcc3c2125c2f2d2fdee4f59094b67af43332f0dfb5ca4400414f6dad1e5128bbf0de9e95343b371c6f7ffe7e26e3082023d3052310b300906035504060c025553311a3018060355040a0c115a657374792048616e64732c20496e632e3127302506035504030c1e5a657374792048616e64732c20496e632e20547275737420416e63686f72a08201e53082018ba00302010202140bdc4aa05179503e58f275d552467347bcafb533300a06082a8648ce3d0403023052310b300906035504060c025553311a3018060355040a0c115a657374792048616e64732c20496e632e3127302506035504030c1e5a657374792048616e64732c20496e632e20547275737420416e63686f72301e170d3232303531393135313330375a170d3332303531363135313330375a3052310b300906035504060c025553311a3018060355040a0c115a657374792048616e64732c20496e632e3127302506035504030c1e5a657374792048616e64732c20496e632e20547275737420416e63686f723059301306072a8648ce3d020106082a8648ce3d0301070342000497cf6d70d76a30400c79f1ebab6ad6168871248710d4f4c1207adaef02d19867136791f2c2b29f6f2dcc3c2125c2f2d2fdee4f59094b67af43332f0dfb5ca440a33f303d301d0603551d0e04160414f6dad1e5128bbf0de9e95343b371c6f7ffe7e26e300b0603551d0f040403020284300f0603551d130101ff040530030101ff300a06082a8648ce3d040302034800304502201da58be7fa442c6cd849ef3567224a9203c225158966b21afd40f3192cf347980221009827b3e0a1aba2514a3994fa6efa9fd6c610b8905fbed93fcb5250754be8aa5882015902d9a28202d5308202d13059301306072a8648ce3d020106082a8648ce3d03010703420004cdd1fe64cf2c04cd93986da576dcedcdf1c92edfd2682cd8e51cfc0409b5c30d6a742e900ed73db8f3f868cba516fd364c4cf3d1ffddd7c07b06d7a992172f1704148a84cff98095a3bc36d6eea518d6978d9bd71f603082025c305c310b300906035504060c025553311f301d060355040a0c16536e6f6262697368204170706172656c2c20496e632e312c302a06035504030c23536e6f6262697368204170706172656c2c20496e632e20547275737420416e63686f72a08201fa3082019fa0030201020214101b934465c01045441e1bb8c5a7c09ea9bea988300a06082a8648ce3d040302305c310b300906035504060c025553311f301d060355040a0c16536e6f6262697368204170706172656c2c20496e632e312c302a06035504030c23536e6f6262697368204170706172656c2c20496e632e20547275737420416e63686f72301e170d3232303531393135313330385a170d3332303531363135313330385a305c310b300906035504060c025553311f301d060355040a0c16536e6f6262697368204170706172656c2c20496e632e312c302a06035504030c23536e6f6262697368204170706172656c2c20496e632e20547275737420416e63686f723059301306072a8648ce3d020106082a8648ce3d03010703420004cdd1fe64cf2c04cd93986da576dcedcdf1c92edfd2682cd8e51cfc0409b5c30d6a742e900ed73db8f3f868cba516fd364c4cf3d1ffddd7c07b06d7a992172f17a33f303d301d0603551d0e041604148a84cff98095a3bc36d6eea518d6978d9bd71f60300b0603551d0f040403020284300f0603551d130101ff040530030101ff300a06082a8648ce3d0403020349003046022100b671fe377f73cf9423bafddc6fe347ed220c714e828717bd94cc436fefecb8ca02210081ad0bfafa48668531a3fbcc3d8496806e2174edc96bbc1f097e606757458f77
	// {"tag-identity":{"id":"ab0f44b1-bfdc-4604-ab4a-30f80407ebcc"},"environments":[{"namedtastore":"Miscellaneous TA Store"}],"keys":{"tas":[{"format":0,"data":"MIIBvTCCAWSgAwIBAgIVANCdkL89UlzHc9Ui7XfVniK7pFuIMAoGCCqGSM49BAMCMD4xCzAJBgNVBAYMAlVTMRAwDgYDVQQKDAdFeGFtcGxlMR0wGwYDVQQDDBRFeGFtcGxlIFRydXN0IEFuY2hvcjAeFw0yMjA1MTkxNTEzMDdaFw0zMjA1MTYxNTEzMDdaMD4xCzAJBgNVBAYMAlVTMRAwDgYDVQQKDAdFeGFtcGxlMR0wGwYDVQQDDBRFeGFtcGxlIFRydXN0IEFuY2hvcjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABONRqhA5JAekvQN8oLwRVNDnAfBnTznLLE+SEGks677sHSeXfcVhZXUeDiN7/fsVNumaiEWRQpZh3zXPwL8rUMyjPzA9MB0GA1UdDgQWBBQBXEXJrLBGKnFd1xCgeMAVSfEBPzALBgNVHQ8EBAMCAoQwDwYDVR0TAQH/BAUwAwEB/zAKBggqhkjOPQQDAgNHADBEAiALBidABsfpzG0lTL9Eh9b6AUbqnzF+koEZbgvppvvt9QIgVoE+bhEN0j6wSPzePjLrEdD+PEgyjHJ5rbA11SPq/1M="},{"format":1,"data":"ooICtjCCArIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASXz21w12owQAx58euratYWiHEkhxDU9MEgetrvAtGYZxNnkfLCsp9vLcw8ISXC8tL97k9ZCUtnr0MzLw37XKRABBT22tHlEou/DenpU0Ozccb3/+fibjCCAj0wUjELMAkGA1UEBgwCVVMxGjAYBgNVBAoMEVplc3R5IEhhbmRzLCBJbmMuMScwJQYDVQQDDB5aZXN0eSBIYW5kcywgSW5jLiBUcnVzdCBBbmNob3KgggHlMIIBi6ADAgECAhQL3EqgUXlQPljyddVSRnNHvK+1MzAKBggqhkjOPQQDAjBSMQswCQYDVQQGDAJVUzEaMBgGA1UECgwRWmVzdHkgSGFuZHMsIEluYy4xJzAlBgNVBAMMHlplc3R5IEhhbmRzLCBJbmMuIFRydXN0IEFuY2hvcjAeFw0yMjA1MTkxNTEzMDdaFw0zMjA1MTYxNTEzMDdaMFIxCzAJBgNVBAYMAlVTMRowGAYDVQQKDBFaZXN0eSBIYW5kcywgSW5jLjEnMCUGA1UEAwweWmVzdHkgSGFuZHMsIEluYy4gVHJ1c3QgQW5jaG9yMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEl89tcNdqMEAMefHrq2rWFohxJIcQ1PTBIHra7wLRmGcTZ5HywrKfby3MPCElwvLS/e5PWQlLZ69DMy8N+1ykQKM/MD0wHQYDVR0OBBYEFPba0eUSi78N6elTQ7Nxxvf/5+JuMAsGA1UdDwQEAwIChDAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIB2li+f6RCxs2EnvNWciSpIDwiUViWayGv1A8xks80eYAiEAmCez4KGrolFKOZT6bvqf1sYQuJBfvtk/y1JQdUvoqlg="},{"format":1,"data":"ooIC1TCCAtEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATN0f5kzywEzZOYbaV23O3N8cku39JoLNjlHPwECbXDDWp0LpAO1z248/hoy6UW/TZMTPPR/93XwHsG16mSFy8XBBSKhM/5gJWjvDbW7qUY1peNm9cfYDCCAlwwXDELMAkGA1UEBgwCVVMxHzAdBgNVBAoMFlNub2JiaXNoIEFwcGFyZWwsIEluYy4xLDAqBgNVBAMMI1Nub2JiaXNoIEFwcGFyZWwsIEluYy4gVHJ1c3QgQW5jaG9yoIIB+jCCAZ+gAwIBAgIUEBuTRGXAEEVEHhu4xafAnqm+qYgwCgYIKoZIzj0EAwIwXDELMAkGA1UEBgwCVVMxHzAdBgNVBAoMFlNub2JiaXNoIEFwcGFyZWwsIEluYy4xLDAqBgNVBAMMI1Nub2JiaXNoIEFwcGFyZWwsIEluYy4gVHJ1c3QgQW5jaG9yMB4XDTIyMDUxOTE1MTMwOFoXDTMyMDUxNjE1MTMwOFowXDELMAkGA1UEBgwCVVMxHzAdBgNVBAoMFlNub2JiaXNoIEFwcGFyZWwsIEluYy4xLDAqBgNVBAMMI1Nub2JiaXNoIEFwcGFyZWwsIEluYy4gVHJ1c3QgQW5jaG9yMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzdH+ZM8sBM2TmG2ldtztzfHJLt/SaCzY5Rz8BAm1ww1qdC6QDtc9uPP4aMulFv02TEzz0f/d18B7BtepkhcvF6M/MD0wHQYDVR0OBBYEFIqEz/mAlaO8NtbupRjWl42b1x9gMAsGA1UdDwQEAwIChDAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0kAMEYCIQC2cf43f3PPlCO6/dxv40ftIgxxToKHF72UzENv7+y4ygIhAIGtC/r6SGaFMaP7zD2EloBuIXTtyWu8Hwl+YGdXRY93"}]}}
}

func Example_encode_environment_SWID_keys_cert() {
	exampleTaData := []byte{0x30, 0x82, 0x1, 0xe5, 0x30, 0x82, 0x1, 0x8b, 0xa0, 0x3, 0x2, 0x1, 0x2, 0x2, 0x14, 0xb, 0xdc, 0x4a, 0xa0, 0x51, 0x79, 0x50, 0x3e, 0x58, 0xf2, 0x75, 0xd5, 0x52, 0x46, 0x73, 0x47, 0xbc, 0xaf, 0xb5, 0x33, 0x30, 0xa, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x4, 0x3, 0x2, 0x30, 0x52, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x1a, 0x30, 0x18, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x11, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x31, 0x27, 0x30, 0x25, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x1e, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x30, 0x1e, 0x17, 0xd, 0x32, 0x32, 0x30, 0x35, 0x31, 0x39, 0x31, 0x35, 0x31, 0x33, 0x30, 0x37, 0x5a, 0x17, 0xd, 0x33, 0x32, 0x30, 0x35, 0x31, 0x36, 0x31, 0x35, 0x31, 0x33, 0x30, 0x37, 0x5a, 0x30, 0x52, 0x31, 0xb, 0x30, 0x9, 0x6, 0x3, 0x55, 0x4, 0x6, 0xc, 0x2, 0x55, 0x53, 0x31, 0x1a, 0x30, 0x18, 0x6, 0x3, 0x55, 0x4, 0xa, 0xc, 0x11, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x31, 0x27, 0x30, 0x25, 0x6, 0x3, 0x55, 0x4, 0x3, 0xc, 0x1e, 0x5a, 0x65, 0x73, 0x74, 0x79, 0x20, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x2c, 0x20, 0x49, 0x6e, 0x63, 0x2e, 0x20, 0x54, 0x72, 0x75, 0x73, 0x74, 0x20, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x30, 0x59, 0x30, 0x13, 0x6, 0x7, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x2, 0x1, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x3, 0x1, 0x7, 0x3, 0x42, 0x0, 0x4, 0x97, 0xcf, 0x6d, 0x70, 0xd7, 0x6a, 0x30, 0x40, 0xc, 0x79, 0xf1, 0xeb, 0xab, 0x6a, 0xd6, 0x16, 0x88, 0x71, 0x24, 0x87, 0x10, 0xd4, 0xf4, 0xc1, 0x20, 0x7a, 0xda, 0xef, 0x2, 0xd1, 0x98, 0x67, 0x13, 0x67, 0x91, 0xf2, 0xc2, 0xb2, 0x9f, 0x6f, 0x2d, 0xcc, 0x3c, 0x21, 0x25, 0xc2, 0xf2, 0xd2, 0xfd, 0xee, 0x4f, 0x59, 0x9, 0x4b, 0x67, 0xaf, 0x43, 0x33, 0x2f, 0xd, 0xfb, 0x5c, 0xa4, 0x40, 0xa3, 0x3f, 0x30, 0x3d, 0x30, 0x1d, 0x6, 0x3, 0x55, 0x1d, 0xe, 0x4, 0x16, 0x4, 0x14, 0xf6, 0xda, 0xd1, 0xe5, 0x12, 0x8b, 0xbf, 0xd, 0xe9, 0xe9, 0x53, 0x43, 0xb3, 0x71, 0xc6, 0xf7, 0xff, 0xe7, 0xe2, 0x6e, 0x30, 0xb, 0x6, 0x3, 0x55, 0x1d, 0xf, 0x4, 0x4, 0x3, 0x2, 0x2, 0x84, 0x30, 0xf, 0x6, 0x3, 0x55, 0x1d, 0x13, 0x1, 0x1, 0xff, 0x4, 0x5, 0x30, 0x3, 0x1, 0x1, 0xff, 0x30, 0xa, 0x6, 0x8, 0x2a, 0x86, 0x48, 0xce, 0x3d, 0x4, 0x3, 0x2, 0x3, 0x48, 0x0, 0x30, 0x45, 0x2, 0x20, 0x1d, 0xa5, 0x8b, 0xe7, 0xfa, 0x44, 0x2c, 0x6c, 0xd8, 0x49, 0xef, 0x35, 0x67, 0x22, 0x4a, 0x92, 0x3, 0xc2, 0x25, 0x15, 0x89, 0x66, 0xb2, 0x1a, 0xfd, 0x40, 0xf3, 0x19, 0x2c, 0xf3, 0x47, 0x98, 0x2, 0x21, 0x0, 0x98, 0x27, 0xb3, 0xe0, 0xa1, 0xab, 0xa2, 0x51, 0x4a, 0x39, 0x94, 0xfa, 0x6e, 0xfa, 0x9f, 0xd6, 0xc6, 0x10, 0xb8, 0x90, 0x5f, 0xbe, 0xd9, 0x3f, 0xcb, 0x52, 0x50, 0x75, 0x4b, 0xe8, 0xaa, 0x58}
	swname := "Bitter Paper"
	cots := NewConciseTaStore().
		AddEnvironmentGroup(
			*NewEnvironmentGroup().
				SetAbbreviatedSwidTag(
					AbbreviatedSwidTag{
						Entities: swid.Entities{
							makeZestyEntityWithRoles(swid.RoleSoftwareCreator),
						},
					},
				),
		).
		AddPermClaims(
			EatCWTClaim{
				SoftwareNameLabel: &swname,
			},
		).
		SetKeys(
			TasAndCas{
				Tas: []TrustAnchor{
					*NewTrustAnchor().
						SetData(exampleTaData).
						SetFormat(0),
				},
			},
		)

	cbor, err := cots.ToCBOR()
	if err == nil {
		fmt.Printf("%x\n", cbor)
	}

	json, err := cots.ToJSON()
	if err == nil {
		fmt.Printf("%s\n", string(json))
	}

	// Output:
	// a30281a102a102a2181f715a657374792048616e64732c20496e632e1821020481a11903e66c42697474657220506170657206a1008182005901e9308201e53082018ba00302010202140bdc4aa05179503e58f275d552467347bcafb533300a06082a8648ce3d0403023052310b300906035504060c025553311a3018060355040a0c115a657374792048616e64732c20496e632e3127302506035504030c1e5a657374792048616e64732c20496e632e20547275737420416e63686f72301e170d3232303531393135313330375a170d3332303531363135313330375a3052310b300906035504060c025553311a3018060355040a0c115a657374792048616e64732c20496e632e3127302506035504030c1e5a657374792048616e64732c20496e632e20547275737420416e63686f723059301306072a8648ce3d020106082a8648ce3d0301070342000497cf6d70d76a30400c79f1ebab6ad6168871248710d4f4c1207adaef02d19867136791f2c2b29f6f2dcc3c2125c2f2d2fdee4f59094b67af43332f0dfb5ca440a33f303d301d0603551d0e04160414f6dad1e5128bbf0de9e95343b371c6f7ffe7e26e300b0603551d0f040403020284300f0603551d130101ff040530030101ff300a06082a8648ce3d040302034800304502201da58be7fa442c6cd849ef3567224a9203c225158966b21afd40f3192cf347980221009827b3e0a1aba2514a3994fa6efa9fd6c610b8905fbed93fcb5250754be8aa58
	// {"environments":[{"swidtag":{"entity":[{"entity-name":"Zesty Hands, Inc.","role":"softwareCreator"}]}}],"permclaims":[{"swname":"Bitter Paper"}],"keys":{"tas":[{"format":0,"data":"MIIB5TCCAYugAwIBAgIUC9xKoFF5UD5Y8nXVUkZzR7yvtTMwCgYIKoZIzj0EAwIwUjELMAkGA1UEBgwCVVMxGjAYBgNVBAoMEVplc3R5IEhhbmRzLCBJbmMuMScwJQYDVQQDDB5aZXN0eSBIYW5kcywgSW5jLiBUcnVzdCBBbmNob3IwHhcNMjIwNTE5MTUxMzA3WhcNMzIwNTE2MTUxMzA3WjBSMQswCQYDVQQGDAJVUzEaMBgGA1UECgwRWmVzdHkgSGFuZHMsIEluYy4xJzAlBgNVBAMMHlplc3R5IEhhbmRzLCBJbmMuIFRydXN0IEFuY2hvcjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJfPbXDXajBADHnx66tq1haIcSSHENT0wSB62u8C0ZhnE2eR8sKyn28tzDwhJcLy0v3uT1kJS2evQzMvDftcpECjPzA9MB0GA1UdDgQWBBT22tHlEou/DenpU0Ozccb3/+fibjALBgNVHQ8EBAMCAoQwDwYDVR0TAQH/BAUwAwEB/zAKBggqhkjOPQQDAgNIADBFAiAdpYvn+kQsbNhJ7zVnIkqSA8IlFYlmshr9QPMZLPNHmAIhAJgns+Chq6JRSjmU+m76n9bGELiQX77ZP8tSUHVL6KpY"}]}}
}

func Example_decode_JSON() {
	cots := ConciseTaStore{}
	err := cots.FromJSON([]byte(ConciseTaStoreTemplateSingleOrg))

	if err != nil {
		fmt.Printf("FAIL: %v", err)
	} else {
		fmt.Println("OK")
	}

	// Output: OK
}

func Example_decode_CBOR() {
	cots := ConciseTaStore{}
	err := cots.FromCBOR(cotsCBOR)

	if err != nil {
		fmt.Printf("FAIL: %v", err)
	} else {
		fmt.Println("OK")
	}

	// Output: OK
}

func Example_list_of_cots_roundtrip() {
	snobTa, _ := os.ReadFile("../cocli/data/cots/Snobbish Apparel_ta.ta")
	sharedTa, _ := os.ReadFile("../cocli/data/cots/shared_ta.ta")

	// cts1
	egSnob := NewEnvironmentGroup()
	egSnob.SetAbbreviatedSwidTag(AbbreviatedSwidTag{})
	egSnob.SwidTag.Entities = swid.Entities{}
	eSnob := swid.Entity{EntityName: "Snobbish Apparel, Inc."}
	_ = eSnob.SetRoles(swid.RoleSoftwareCreator)
	egSnob.SwidTag.Entities = append(egSnob.SwidTag.Entities, eSnob)

	cts1 := ConciseTaStore{}
	cts1.Keys = NewTasAndCas()
	cts1.Keys.AddTaCert(snobTa)
	cts1.Environments = *NewEnvironmentGroups()
	cts1.AddEnvironmentGroup(*egSnob)

	exclName := "Legal Lawyer"
	exclClaims1 := EatCWTClaim{SoftwareNameLabel: &exclName}
	cts1.AddExclClaims(exclClaims1)

	// cts2
	egShared := NewEnvironmentGroup()
	egShared.Environment = &comid.Environment{}
	egShared.Environment.Class = comid.NewClassOID("1.2.3.4.5")

	cts2 := ConciseTaStore{}
	cts2.Keys = NewTasAndCas()
	cts2.Keys.AddTaCert(sharedTa)
	cts2.Environments = *NewEnvironmentGroups()
	cts2.AddEnvironmentGroup(*egShared)

	cts := NewConciseTaStores()
	cts.AddConciseTaStores(cts1)
	cts.AddConciseTaStores(cts2)

	ctsCBOR, _ := cts.ToCBOR()
	var roundtripCBOR ConciseTaStores
	_ = roundtripCBOR.FromCBOR(ctsCBOR)
	err := roundtripCBOR.Valid()

	if err != nil {
		fmt.Printf("CBOR roundtrip FAIL %s\n", err)
	} else {
		fmt.Printf("CBOR roundtrip OK\n")
	}

	ctsJSON, _ := cts.ToJSON()
	var roundtripJSON ConciseTaStores
	_ = roundtripJSON.FromJSON(ctsJSON)
	err = roundtripJSON.Valid()

	if err != nil {
		fmt.Printf("JSON roundtrip FAIL %s\n", err)
	} else {
		fmt.Printf("JSON roundtrip OK\n")
	}

	// Output:
	// CBOR roundtrip OK
	// JSON roundtrip OK
}
