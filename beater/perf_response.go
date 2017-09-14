package beater

import (
    "encoding/json"
    "regexp"
)

type PerformanceCounter struct {
    NumInvocations   int
    MaxResponseTime  int
    MinResponseTime  int
    AvgResponseTime  int
    NumRuntimeFaults int
}

type PerformanceCounterResponse struct {
    Counters map[string]PerformanceCounter `json:"value"`
}

type ServiceData struct {
    ServiceName string
    Version     string
    Operation   string
}

func Convert(responseText []byte) (PerformanceCounterResponse, error) {
    var responseCounters PerformanceCounterResponse
    err := json.Unmarshal(responseText, &responseCounters)
    if err != nil {
        return responseCounters, err
    }

    return responseCounters, nil
}

var reg = regexp.MustCompile("bus.id=(?P<service>\\S+)-(?P<version>\\d+.\\d+.\\d+).*,operation=\"(?P<operation>\\S+)\",port")
var names = reg.SubexpNames()

func ServiceDataExtract(key string) ServiceData {
    var r ServiceData
    if (reg.MatchString(key)) {
        matches := reg.FindStringSubmatch(key)
        paramsMap := make(map[string]string)
        for i, name := range names {
            if i > 0 && i <= len(matches) {
                paramsMap[name] = matches[i]
            }
        }

        r.ServiceName = paramsMap["service"]
        r.Version = paramsMap["version"]
        r.Operation = paramsMap["operation"]
    }

    return r
}
