package tdx

var (
	TestRegID = "https://intel.com"
	TestOID   = "2.16.840.1.113741.1.2.3.4.5"

	TDXPCERefValTemplate = `{
  "lang": "en-GB",
  "tag-identity": {
    "id": "43BBE37F-2E61-4B33-AED3-53CFF1428B17",
    "version": 0
  },
  "entities": [
    {
      "name": "INTEL",
      "regid": "https://intel.com",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
    }
  ],
  "triples": {
    "reference-values": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "oid",
              "value": "2.16.840.1.113741.1.2.3.4.4"
            },
            "vendor": "Intel Corporation",
            "model": "0123456789ABCDEF"
          }
        },
        "measurements": [
          {
            "value": {
              "attributes": "AwM=",
              "tcbevalnum": 5,
              "pceid": "0000"
            },
            "authorized-by": {
              "type": "pkix-base64-key",
              "value": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFn0taoAwR3PmrKkYLtAsD9o05KSM6mbgfNCgpuL0g6VpTHkZl73wk5BDxoV7n+Oeee0iIqkW3HMZT3ETiniJdg==\n-----END PUBLIC KEY-----"
            }
          }
        ]
      }
    ]
  }
}
`
	TDXQERefValTemplate = `{
  "lang": "en-GB",
  "tag-identity": {
    "id": "43BBE37F-2E61-4B33-AED3-53CFF1428B16",
    "version": 0
  },
  "entities": [
    {
      "name": "INTEL",
      "regid": "https://intel.com",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
    }
  ],
  "triples": {
    "reference-values": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "oid",
              "value": "2.16.840.1.113741.1.2.3.4.1"
            },
            "vendor": "Intel Corporation",
            "model": "TDX QE TCB"
          }
        },
        "measurements": [
          {
            "value": {
              "miscselect": "wAAAAPv/AAA=",
              "tcbevalnum": 11, 
              "mrsigner": [
                "sha-256:h0KPxSKAPTEGXnvOPPA/5HUJZjHl4Hu9eg/eYMTPJcc=",
                "sha-512:oxT8LcZjrnpra8Z4dZQFc5bms/VpzVD9XdtNG7r9K2qjFPwtxmOuemtrxnh1lAVzluaz9WnNUP1d200buv0rag=="
              ],
              "isvprodid": "AwM="
            },
            "authorized-by": {
              "type": "pkix-base64-key",
              "value": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFn0taoAwR3PmrKkYLtAsD9o05KSM6mbgfNCgpuL0g6VpTHkZl73wk5BDxoV7n+Oeee0iIqkW3HMZT3ETiniJdg==\n-----END PUBLIC KEY-----"
            }
          }
        ]
      }
    ]
  }
}	
`
	TDXSeamRefValJSONTemplate = ` {
  "lang": "en-GB",
  "tag-identity": {
    "id": "43BBE37F-2E61-4B33-AED3-53CFF1428B20",
    "version": 0
  },
  "entities": [
    {
      "name": "INTEL",
      "regid": "https://intel.com",
      "roles": [
        "tagCreator",
        "creator",
        "maintainer"
      ]
    }
  ],
  "triples": {
    "reference-values": [
      {
        "environment": {
          "class": {
            "id": {
              "type": "oid",
              "value": "2.16.840.1.113741.1.2.3.4.5"
            },
            "vendor": "Intel Corporation",
            "model": "TDX SEAM"
          }
        },
        "measurements": [
          {
            "value": {
              "isvprodid": "AwM=",
              "isvsvn": 10,
              "attributes": "8AoL",
              "tcbevalnum": 11,
              "mrsigner": [
                "sha-256:h0KPxSKAPTEGXnvOPPA/5HUJZjHl4Hu9eg/eYMTPJcc=",
                "sha-512:oxT8LcZjrnpra8Z4dZQFc5bms/VpzVD9XdtNG7r9K2qjFPwtxmOuemtrxnh1lAVzluaz9WnNUP1d200buv0rag=="
              ]
            },
            "authorized-by": {
              "type": "pkix-base64-key",
              "value": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFn0taoAwR3PmrKkYLtAsD9o05KSM6mbgfNCgpuL0g6VpTHkZl73wk5BDxoV7n+Oeee0iIqkW3HMZT3ETiniJdg==\n-----END PUBLIC KEY-----"
            }
          }
        ]
      }
    ]
  }
}
`
)
