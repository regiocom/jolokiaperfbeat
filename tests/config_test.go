// +build !integration

package tests

import (
    "testing"
    "github.com/regiocom/jolokiaperfbeat/config"
    "github.com/regiocom/jolokiaperfbeat/beater"
)

func TestValidateConfigEmptyProvider(t *testing.T) {
    c := config.Config{}

    c.Provider = ""
    err := beater.ValidateConfig(c)

    if err == nil {
        t.Fatalf("Empty provider must return error")
    }
}

func TestValidateConfigInvalidProvider(t *testing.T) {
    c := config.Config{}

    c.Provider = "KLMNJ"
    err := beater.ValidateConfig(c)

    if err == nil {
        t.Fatalf("Invalid provider must return error")
    }
}
