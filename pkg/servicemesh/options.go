package servicemesh

type Options struct {
	IstioPilotHost            string `json:"istioPilotHost,omitempty" yaml:"istioPilotHost"`
	JaegerQueryHost           string `json:"jaegerQueryHost,omitempty" yaml:"jaegerQueryHost"`
	ServicemeshPrometheusHost string `json:"servicemeshPrometheusHost,omitempty" yaml:"servicemeshPrometheusHost"`
}

func NewServiceMeshOptions() *Options {
	return &Options{
		IstioPilotHost:            "",
		JaegerQueryHost:           "",
		ServicemeshPrometheusHost: "",
	}
}
