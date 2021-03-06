# Go JSON Parsers: Errors

Go Version: go1.10


## Test: Unknown object property.

Body:
```json
{"cat": "Sammy", "dog": "Loki"}
```
Target:
```go
struct { Cat string "json:\"cat\"" }
```


* **pkg [encoding/json](https://godoc.org/encoding/json)**  (stdlib)


	* **json.Unmarshal**

	  Err: `<nil>`

		<details>
		
		```go
		(interface {}) <nil>
		```
		
		</details>


	* **json.Decoder.DisallowUnknownFields**

	  Err: `json: unknown field "dog"`

		<details>
		
		```go
		(*errors.errorString)(0xc4200a4c30)({
		  s: (string) (len=25) "json: unknown field \"dog\""
		})
		```
		
		</details>



* **pkg [github.com/json-iterator/go](https://godoc.org/github.com/json-iterator/go)** 


	* **jsoniter.ConfigCompatibleWithStandardLibrary**

	  Err: `<nil>`

		<details>
		
		```go
		(interface {}) <nil>
		```
		
		</details>



* **pkg [github.com/koki/json](https://godoc.org/github.com/koki/json)** 


	* **json.Unmarshal**

	  Err: `<nil>`

		<details>
		
		```go
		(interface {}) <nil>
		```
		
		</details>


	* **json.UnmarshalMap**

	  Err: `<nil>`

		<details>
		
		```go
		(interface {}) <nil>
		```
		
		</details>


	* **json.UnmarshalMap + jsonutil.ExtraneousFieldPaths**

	  Err: `extraneous fields (typos?) at paths: $.dog`

		<details>
		
		```go
		(*jsonutil.ExtraneousFieldsError)(0xc4200bae80)({
		  Paths: ([][]string) (len=1 cap=1) {
		    ([]string) (len=1 cap=1) {
		      (string) (len=3) "dog"
		    }
		  }
		})
		```
		
		</details>




## Test: Using a string instead of an object.

Body:
```json
{"cat": "sammy"}
```
Target:
```go
struct { Cat struct {} "json:\"cat\"" }
```


* **pkg [encoding/json](https://godoc.org/encoding/json)**  (stdlib)


	* **json.Unmarshal**

	  Err: `json: cannot unmarshal string into Go struct field .cat of type struct {}`

		<details>
		
		```go
		(*json.UnmarshalTypeError)(0xc4200be4b0)({
		  Value: (string) (len=6) "string",
		  Type: (*reflect.rtype)(0x120fb20)({
		    size: (uintptr) <nil>,
		    ptrdata: (uintptr) <nil>,
		    hash: (uint32) 670477339,
		    tflag: (reflect.tflag) 2,
		    align: (uint8) 1,
		    fieldAlign: (uint8) 1,
		    kind: (uint8) 153,
		    alg: (*reflect.typeAlg)(0x13bdd30)({
		      hash: (func(unsafe.Pointer, uintptr) uintptr) 0x1001b60,
		      equal: (func(unsafe.Pointer, unsafe.Pointer) bool) 0x1002360
		    }),
		    gcdata: (*uint8)(0x127a228)(1),
		    str: (reflect.nameOff) 21559,
		    ptrToThis: (reflect.typeOff) 118240
		  }),
		  Offset: (int64) 15,
		  Struct: (string) "",
		  Field: (string) (len=3) "cat"
		})
		```
		
		</details>


	* **json.Decoder.DisallowUnknownFields**

	  Err: `json: cannot unmarshal string into Go struct field .cat of type struct {}`

		<details>
		
		```go
		(*json.UnmarshalTypeError)(0xc4200be500)({
		  Value: (string) (len=6) "string",
		  Type: (*reflect.rtype)(0x120fb20)({
		    size: (uintptr) <nil>,
		    ptrdata: (uintptr) <nil>,
		    hash: (uint32) 670477339,
		    tflag: (reflect.tflag) 2,
		    align: (uint8) 1,
		    fieldAlign: (uint8) 1,
		    kind: (uint8) 153,
		    alg: (*reflect.typeAlg)(0x13bdd30)({
		      hash: (func(unsafe.Pointer, uintptr) uintptr) 0x1001b60,
		      equal: (func(unsafe.Pointer, unsafe.Pointer) bool) 0x1002360
		    }),
		    gcdata: (*uint8)(0x127a228)(1),
		    str: (reflect.nameOff) 21559,
		    ptrToThis: (reflect.typeOff) 118240
		  }),
		  Offset: (int64) 15,
		  Struct: (string) "",
		  Field: (string) (len=3) "cat"
		})
		```
		
		</details>



* **pkg [github.com/json-iterator/go](https://godoc.org/github.com/json-iterator/go)** 


	* **jsoniter.ConfigCompatibleWithStandardLibrary**

	  Err: `struct { Cat struct {} "json:\"cat\"" }.Cat: skipObjectDecoder: expect object or null, error found in #8 byte of ...|{"cat": "sammy"}|..., bigger context ...|{"cat": "sammy"}|...`

		<details>
		
		```go
		(*errors.errorString)(0xc4200a5180)({
		  s: (string) (len=178) "struct { Cat struct {} \"json:\\\"cat\\\"\" }.Cat: skipObjectDecoder: expect object or null, error found in #8 byte of ...|{\"cat\": \"sammy\"}|..., bigger context ...|{\"cat\": \"sammy\"}|..."
		})
		```
		
		</details>



* **pkg [github.com/koki/json](https://godoc.org/github.com/koki/json)** 


	* **json.Unmarshal**

	  Err: `$.cat: json: cannot unmarshal string into Go struct field .cat of type struct {}`

		<details>
		
		```go
		(*structurederrors.ErrorWithContext)(0xc4201381e0)({
		  BaseError: (*json.UnmarshalTypeError)(0xc4200be550)({
		    Value: (string) (len=6) "string",
		    Type: (*reflect.rtype)(0x120fb20)({
		      size: (uintptr) <nil>,
		      ptrdata: (uintptr) <nil>,
		      hash: (uint32) 670477339,
		      tflag: (reflect.tflag) 2,
		      align: (uint8) 1,
		      fieldAlign: (uint8) 1,
		      kind: (uint8) 153,
		      alg: (*reflect.typeAlg)(0x13bdd30)({
		        hash: (func(unsafe.Pointer, uintptr) uintptr) 0x1001b60,
		        equal: (func(unsafe.Pointer, unsafe.Pointer) bool) 0x1002360
		      }),
		      gcdata: (*uint8)(0x127a228)(1),
		      str: (reflect.nameOff) 21559,
		      ptrToThis: (reflect.typeOff) 118240
		    }),
		    Offset: (int64) 15,
		    Struct: (string) "",
		    Field: (string) (len=3) "cat"
		  }),
		  Context: ([]string) (len=1 cap=1) {
		    (string) (len=5) "$.cat"
		  }
		})
		```
		
		</details>


	* **json.UnmarshalMap**

	  Err: `$.cat: json: cannot unmarshal string into Go struct field .cat of type struct {}`

		<details>
		
		```go
		(*structurederrors.ErrorWithContext)(0xc4201382d0)({
		  BaseError: (*json.UnmarshalTypeError)(0xc4200be5a0)({
		    Value: (string) (len=6) "string",
		    Type: (*reflect.rtype)(0x120fb20)({
		      size: (uintptr) <nil>,
		      ptrdata: (uintptr) <nil>,
		      hash: (uint32) 670477339,
		      tflag: (reflect.tflag) 2,
		      align: (uint8) 1,
		      fieldAlign: (uint8) 1,
		      kind: (uint8) 153,
		      alg: (*reflect.typeAlg)(0x13bdd30)({
		        hash: (func(unsafe.Pointer, uintptr) uintptr) 0x1001b60,
		        equal: (func(unsafe.Pointer, unsafe.Pointer) bool) 0x1002360
		      }),
		      gcdata: (*uint8)(0x127a228)(1),
		      str: (reflect.nameOff) 21559,
		      ptrToThis: (reflect.typeOff) 118240
		    }),
		    Offset: (int64) 14,
		    Struct: (string) "",
		    Field: (string) (len=3) "cat"
		  }),
		  Context: ([]string) (len=1 cap=1) {
		    (string) (len=5) "$.cat"
		  }
		})
		```
		
		</details>


	* **json.UnmarshalMap + jsonutil.ExtraneousFieldPaths**

	  Err: `$.cat: json: cannot unmarshal string into Go struct field .cat of type struct {}`

		<details>
		
		```go
		(*structurederrors.ErrorWithContext)(0xc4201383c0)({
		  BaseError: (*json.UnmarshalTypeError)(0xc4200be5f0)({
		    Value: (string) (len=6) "string",
		    Type: (*reflect.rtype)(0x120fb20)({
		      size: (uintptr) <nil>,
		      ptrdata: (uintptr) <nil>,
		      hash: (uint32) 670477339,
		      tflag: (reflect.tflag) 2,
		      align: (uint8) 1,
		      fieldAlign: (uint8) 1,
		      kind: (uint8) 153,
		      alg: (*reflect.typeAlg)(0x13bdd30)({
		        hash: (func(unsafe.Pointer, uintptr) uintptr) 0x1001b60,
		        equal: (func(unsafe.Pointer, unsafe.Pointer) bool) 0x1002360
		      }),
		      gcdata: (*uint8)(0x127a228)(1),
		      str: (reflect.nameOff) 21559,
		      ptrToThis: (reflect.typeOff) 118240
		    }),
		    Offset: (int64) 14,
		    Struct: (string) "",
		    Field: (string) (len=3) "cat"
		  }),
		  Context: ([]string) (len=1 cap=1) {
		    (string) (len=5) "$.cat"
		  }
		})
		```
		
		</details>




