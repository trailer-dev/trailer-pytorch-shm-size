package recommendation

type AuthenticationType string

const (
	InternalAuthentication AuthenticationType = "internal"
	BasicAuthentication    AuthenticationType = "basic"
)

type BasicAuth struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Hash     string `json:"hash,omitempty"`
}

type Authentication struct {
	Type      AuthenticationType `json:"type"`
	BasicAuth BasicAuth          `json:"basicAuth,omitempty"`
}

type URL struct {
	Authentication Authentication `json:"authentication"`
	Subdomain      string         `json:"subdomain"`
	Domain         string         `json:"domain"`
	Port           int64          `json:"port"`
}

type ResourceRequest struct {
	Memory int64 `json:"memory"`
	CPU    int64 `json:"cpu"`
}

type ResourceLimit struct {
	Memory int64 `json:"memory"`
	CPU    int64 `json:"cpu"`
}

type Resources struct {
	Requests ResourceRequest `json:"requests"`
	Limits   ResourceLimit   `json:"limits"`
	ShmSize  int64           `json:"shmSize"`
}

// PortBinding represents a binding between a Host IP address and a Host Port
type PortBinding struct {
	// HostIP is the host IP Address
	HostIP string `json:"HostIp"`
	// HostPort is the host port number
	HostPort string
}

// PortMap is a collection of PortBinding indexed by Port
type PortMap map[Port][]PortBinding

// PortSet is a collection of structs indexed by Port
type PortSet map[Port]struct{}

// Port is a string containing port number and protocol in the format "80/tcp"
type Port string

type StartupConfiguration struct {
	Command   []string `json:"command,omitempty"`
	Arguments []string `json:"arguments,omitempty"`
}

type VersionSpecifier string

const (
	ExactVersion         VersionSpecifier = "exactly"
	MinimumVersion       VersionSpecifier = "minimum"
	MaximumVersion       VersionSpecifier = "maximum"
	ExcludeVersion       VersionSpecifier = "exclude"
	CompatibleVersion    VersionSpecifier = "compatible"
	UnconstrainedVersion VersionSpecifier = "unconstrained"
	CustomVersion        VersionSpecifier = "custom"
)

type Version struct {
	Specifier  VersionSpecifier `json:"specifier,omitempty"`
	Constraint string           `json:"constraint"`
}

type Package struct {
	Name        string  `json:"name"`
	Version     Version `json:"version,omitempty"`
	Description string  `json:"description,omitempty"`
	Channel     string  `json:"channel"`
}

type Environment struct {
	Channels []string  `json:"channels"`
	Packages []Package `json:"packages"`
}

type ImageConfiguration struct {
	Environments map[string]Environment `json:"environments"`
}

type Image struct {
	Configuration ImageConfiguration `json:"configuration"`
}

type Host struct {
	TotalMemory int64 `json:"total_memory"`
}

type Workspace struct {
	Resources Resources `json:"resources,omitempty"`
	Image     Image     `json:"image,omitempty"`
	Host      Host      `json:"host"`
}
