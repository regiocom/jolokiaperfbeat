// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
    Period  time.Duration `config:"period"`
    Baseurl string `config:"baseurl"`
    Provider string `config:"provider"`
}

var DefaultConfig = Config{
    Period: 60 * time.Second,
}
