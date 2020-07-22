package version_helper

type Versioner interface {
	GetVersion() string
}

type versioner struct {
	version string
}

func (v *versioner) GetVersion() string {
	if v.version == "" {
		return "unknown"
	}

	return v.version
}

func NewVersioner(v string) func() Versioner {
	return func() Versioner { return &versioner{version: v} }
}
