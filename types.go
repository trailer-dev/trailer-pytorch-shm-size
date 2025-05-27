package recommendation

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

type Version struct {
	Constraint string `json:"constraint"`
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
	TotalMemory int64 `json:"totalMemory"`
}

type Workspace struct {
	Resources Resources `json:"resources,omitempty"`
	Image     Image     `json:"image,omitempty"`
	Host      Host      `json:"host"`
}
