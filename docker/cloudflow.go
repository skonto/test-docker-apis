package docker

import "encoding/json"

type Pair struct {
	a, b interface{}
}

type ImageReference struct {
	Registry   string
	Repository string
	Image      string
	Tag        string
	FullURI    string
}

// Connection is one inlet/outlet connection
type Connection struct {
	InletName           string `json:"inlet_name"`
	InletStreamletName  string `json:"inlet_streamlet_name"`
	OutletName          string `json:"outlet_name"`
	OutletStreamletName string `json:"outlet_streamlet_name"`
}

// PortMapping maps outlets
type PortMapping struct {
	AppID     string `json:"app_id"`
	Outlet    string `json:"outlet"`
	Streamlet string `json:"streamlet"`
}

// Endpoint contains deployment endpoint information
type Endpoint struct {
	AppID         string `json:"app_id,omitempty"`
	Streamlet     string `json:"streamlet,omitempty"`
	ContainerPort int    `json:"container_port,omitempty"`
}

// Deployment contains a streamlet deployment
type Deployment struct {
	ClassName     string                  `json:"class_name"`
	Config        json.RawMessage         `json:"config"`
	Image         string                  `json:"image"`
	Name          string                  `json:"name"`
	PortMappings  map[string]PortMapping  `json:"port_mappings"`
	VolumeMounts  []VolumeMountDescriptor `json:"volume_mounts"`
	Runtime       string                  `json:"runtime"`
	StreamletName string                  `json:"streamlet_name"`
	SecretName    string                  `json:"secret_name"`
	Endpoint      *Endpoint               `json:"endpoint,omitempty"`
	Replicas      int                     `json:"replicas,omitempty"`
}

// InOutletSchema contains the schema of a in/out-let
type InOutletSchema struct {
	Fingerprint string `json:"fingerprint"`
	Schema      string `json:"schema"`
	Name        string `json:"name"`
	Format      string `json:"format"`
}

// InOutlet defines a in/out-let and its schema
type InOutlet struct {
	Name   string         `json:"name"`
	Schema InOutletSchema `json:"schema"`
}

// Attribute TBD
type Attribute struct {
	AttributeName string `json:"attribute_name"`
	ConfigPath    string `json:"config_path"`
}

// ConfigParameterDescriptor TBD
type ConfigParameterDescriptor struct {
	Key          string `json:"key"`
	Description  string `json:"description"`
	Type         string `json:"validation_type"`
	Pattern      *string `json:"validation_pattern,omitempty"`
	DefaultValue *string `json:"default_value,omitempty"`
}

// ReadWriteMany is a name of a VolumeMount access mode
const ReadWriteMany = "ReadWriteMany"

// ReadOnlyMany is a name of a VolumeMount access mode
const ReadOnlyMany = "ReadOnlyMany"

// VolumeMountDescriptor TBD
type VolumeMountDescriptor struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	AccessMode string `json:"access_mode"`
	PVCName    string `json:"pvc_name,omitempty"`
}


// Descriptor TBD
type Descriptor struct {
	Attributes       []Attribute                 `json:"attributes"`
	ClassName        string                      `json:"class_name"`
	ConfigParameters []ConfigParameterDescriptor `json:"config_parameters"`
	VolumeMounts     []VolumeMountDescriptor     `json:"volume_mounts"`
	Inlets           []InOutlet                  `json:"inlets"`
	Labels           []string                    `json:"labels"`
	Outlets          []InOutlet                  `json:"outlets"`
	Runtime          string                      `json:"runtime"`
	Description      string                      `json:"description"`
}


type Descriptors struct {
	StreamletDescriptors []Descriptor `json:"streamlet-descriptors"`
	APIVersion           string       `json:"api-version"`
}
