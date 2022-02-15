text="package api

import (
    \"github.com/xm-chentl/go-mvc\"
	\"github.com/xm-chentl/go-mvc/metadata\"
)

func Register() {
	metadata.RegisterMap(map[string]mvc.IApi{})
}"
echo "$text" > api/metadata.go
go run . --api