package beater

import (
    "fmt"
    "time"

    "github.com/elastic/beats/libbeat/beat"
    "github.com/elastic/beats/libbeat/common"
    "github.com/elastic/beats/libbeat/logp"
    "github.com/elastic/beats/libbeat/publisher"

    "github.com/regiocom/jolokiaperfbeat/config"
    "strings"
)

type Jolokiaperfbeat struct {
    done   chan struct{}
    config config.Config
    client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
    config := config.DefaultConfig
    if err := cfg.Unpack(&config); err != nil {
        return nil, fmt.Errorf("Error reading config file: %v", err)
    }

    bt := &Jolokiaperfbeat{
        done:   make(chan struct{}),
        config: config,
    }

    return bt, ValidateConfig(bt.config)
}
func ValidateConfig(cfg config.Config) error {
    var validProviders = []string {"cxfCounter", "cxfMetrics"}
    if !stringInArray(cfg.Provider, validProviders) {
        return fmt.Errorf(`Provider "%s" invalid. Valid options are [%s].`, cfg.Provider, strings.Join(validProviders, "|"))
    }

    return nil
}

func stringInArray(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func (bt *Jolokiaperfbeat) Run(b *beat.Beat) error {
    logp.Info("jolokiaperfbeat is running! Hit CTRL-C to stop it.")

    bt.client = b.Publisher.Connect()
    ticker := time.NewTicker(bt.config.Period)
    counter := 1
    for {
        select {
        case <-bt.done:
            return nil
        case <-ticker.C:
        }

        serverResponse, errHttp := GetStatsFromServer(bt.config)
        if errHttp != nil {
            logp.Err("GetStatsFromServer", errHttp)
        } else {
            perfCounters, errJson := Convert(serverResponse)
            if errJson != nil {
                logp.Err("Convert", errHttp)
            } else {
                for key, value := range perfCounters.Counters {
                    sd := ServiceDataExtract(key)
                    fmt.Println("Key:", key, "Value:", value)
                    event := CreateEvent(counter, b.Name, sd, value)
                    bt.client.PublishEvent(event)
                    logp.Info("Event sent")
                }

                counter++
            }
        }
    }
}

func (bt *Jolokiaperfbeat) Stop() {
    bt.client.Close()
    close(bt.done)
}
