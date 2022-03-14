package model

type Log struct {
	Request             Request             `json:"request"`
	UpStreamUri         string              `json:"upstream_uri"`
	Response            Response            `json:"response"`
	AuthenticatedEntity AuthenticatedEntity `json:"authenticated_entity"`
	Route               Route               `json:"route"`
	Service             Service             `json:"service"`
	Latencies           Latencies           `json:"latencies"`
	ClientIp            string              `json:"client_ip"`
	StartedAt           int64               `json:"started_at"`
}

type Request struct {
	Method      string        `json:"method"`
	Uri         string        `json:"uri"`
	Url         string        `json:"url"`
	Size        int64         `json:"size"`
	QueryString []string      `json:"querystring"`
	Headers     HeaderRequest `json:"headers"`
}

type HeaderRequest struct {
	Accept    string `json:"accept"`
	Host      string `json:"host"`
	UserAgent string `json:"user-agent"`
}

type Response struct {
	Status  int64          `json:"status"`
	Size    int64          `json:"size"`
	Headers HeaderResponse `json:"headers"`
}

type HeaderResponse struct {
	ContentLength                 string `json:"Content-Length"`
	Via                           string `json:"via"`
	Connection                    string `json:"Connection"`
	AccessControlAllowCredentials string `json:"access-control-allow-credentials"`
	ContentType                   string `json:"Content-Type"`
	Server                        string `json:"server"`
	AccessControlAllowOrigin      string `json:"access-control-allow-origin"`
}

type AuthenticatedEntity struct {
	ConsumerID ConsumerID `json:"consumer_id"`
}

type ConsumerID struct {
	Uuid string `json:"uuid"`
}

type Route struct {
	CreatedAt     int64        `json:"created_at"`
	Hosts         string       `json:"hosts"`
	ID            string       `json:"id"`
	Methods       []string     `json:"methods"`
	Paths         []string     `json:"paths"`
	PreserveHost  bool         `json:"preserve_host"`
	Protocols     []string     `json:"protocols"`
	RegexPriority int64        `json:"regex_priority"`
	Service       RouteService `json:"service"`
	StripPath     bool         `json:"strip_path"`
	UpdatedAt     int64        `json:"updated_at"`
}

type RouteService struct {
	ID string `json:"id"`
}

type Service struct {
	ConnectTimeout int64  `json:"connect_timeout"`
	CreatedAt      int64  `json:"created_at"`
	Host           string `json:"host"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	Port           int64  `json:"port"`
	Protocol       string `json:"protocol"`
	ReadTimeout    int64  `json:"read_timeout"`
	Retries        int64  `json:"retries"`
	UpdatedAt      int64  `json:"updated_at"`
	WriteTimeout   int64  `json:"write_timeout"`
}

type Latencies struct {
	Proxy   int64 `json:"proxy"`
	Kong    int64 `json:"kong"`
	Request int64 `json:"request"`
}

type AverageLatencies struct {
	Proxy   float64
	Kong    float64
	Request float64
}
