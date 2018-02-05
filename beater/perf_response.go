package beater

import (
    "encoding/json"
    "regexp"
    "github.com/elastic/beats/libbeat/common"
    "time"
)

type PerformanceCounter struct {
    NumInvocations    int
    MaxResponseTime   int
    MinResponseTime   int
    AvgResponseTime   int
    NumRuntimeFaults  int
    Percentile75th    float64 `json:"75thPercentile"`
    StandardDeviation float64 `json:"StdDev"`
    Mean              float64 `json:"Mean"`
    Percentile98th    float64 `json:"98thPercentile"`
    Percentile95th    float64 `json:"95thPercentile"`
    Percentile99th    float64 `json:"99thPercentile"`
    Max               float64 `json:"Max"`
    Count             int     `json:"Count"`
    FiveMinuteRate    float64 `json:"FiveMinuteRate"`
    Percentile50th    float64 `json:"50thPercentile"`
    MeanRate          float64 `json:"MeanRate"`
    Min               float64 `json:"Min"`
    OneMinuteRate     float64 `json:"OneMinuteRate"`
    DurationUnit      string  `json:"DurationUnit"`
    Percentile999th   float64 `json:"999thPercentile"`
    FifteenMinuteRate float64 `json:"FifteenMinuteRate"`
}

type PerformanceCounterResponse struct {
    Counters map[string]PerformanceCounter `json:"value"`
}

type ServiceData struct {
    Source      string
    Attribute   string
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

var performanceCounterRegex = regexp.MustCompile("bus\\.id=(?P<service>\\S+)-(?P<version>\\d+.\\d+.\\d+).*,operation=\"(?P<operation>\\S+)\",port")
var performanceCounterNames = performanceCounterRegex.SubexpNames()

var metricsServerRegex = regexp.MustCompile(`Attribute=(?P<attribute>[\S ]*?),Operation=(?P<operation>\S+),bus\.id=(?P<service>\S*?)-(?P<version>\d+.\d+.\d+)`)
var metricsServerNames = metricsServerRegex.SubexpNames();

func ServiceDataExtract(key string) ServiceData {
    var r ServiceData
    if (performanceCounterRegex.MatchString(key)) {
        matches := performanceCounterRegex.FindStringSubmatch(key)
        paramsMap := make(map[string]string)
        for i, name := range performanceCounterNames {
            if i > 0 && i <= len(matches) {
                paramsMap[name] = matches[i]
            }
        }

        r.Source = "Performance.Counter.Server"
        r.Attribute = ""
        r.ServiceName = paramsMap["service"]
        r.Version = paramsMap["version"]
        r.Operation = paramsMap["operation"]
    } else if (metricsServerRegex.MatchString(key)) {
        matches := metricsServerRegex.FindStringSubmatch(key);
        paramsMap := make(map[string]string)
        for i, name := range metricsServerNames {
            if i > 0 && i <= len(matches) {
                paramsMap[name] = matches[i]
            }
        }

        r.Source = "Metrics.Server"
        r.Attribute = paramsMap["attribute"]
        r.ServiceName = paramsMap["service"]
        r.Version = paramsMap["version"]
        r.Operation = paramsMap["operation"]
    }

    return r
}

func CreateEvent(counter int, name string, sd ServiceData, value PerformanceCounter) common.MapStr {
    event := common.MapStr{
        "@timestamp":        common.Time(time.Now()),
        "type":              name,
        "counter":           counter,
        "service.name":      sd.ServiceName,
        "service.version":   sd.Version,
        "service.operation": sd.Operation,
    }

    switch sd.Source {
    case "Performance.Counter.Server":
        event.Update(
            common.MapStr{
                "avgresponsetime": value.AvgResponseTime,
                "minresponsetime": value.MinResponseTime,
                "maxresponsetime": value.MaxResponseTime,
                "numinvocations":  value.NumInvocations,
                "numfaults":       value.NumRuntimeFaults,
            })
    case "Metrics.Server":
        // Minimal data set: attribute and count
        event.Update(common.MapStr{
            "service.attribute": sd.Attribute,
            "numinvocations":    value.Count,
        })
        // Everything except "In Flight" has more keys
        if sd.Attribute != "In Flight" {
            event.Update(
                common.MapStr{
                    "rate.mean":  value.MeanRate,
                    "rate.avg1":  value.OneMinuteRate,
                    "rate.avg5":  value.FiveMinuteRate,
                    "rate.avg15": value.FifteenMinuteRate,
                })
            // Exclude "Data Written" and "Data Read" from remaining keys
            if sd.Attribute != "Data Written" && sd.Attribute != "Data Read" {
                event.Update(
                    common.MapStr{
                        "avgresponsetime":     value.Mean,
                        "minresponsetime":     value.Min,
                        "maxresponsetime":     value.Max,
                        "numinvocations":      value.Count,
                        "responsetime.stddev": value.StandardDeviation,
                        "responsetime.p50":    value.Percentile50th,
                        "responsetime.p75":    value.Percentile75th,
                        "responsetime.p98":    value.Percentile98th,
                        "responsetime.p95":    value.Percentile95th,
                        "responsetime.p99":    value.Percentile99th,
                        "responsetime.p999":   value.Percentile999th,
                        "responsetime.unit":   value.DurationUnit,
                    })
            }
        }
    }

    return event
}
