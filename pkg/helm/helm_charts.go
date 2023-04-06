package helm

type HelmRelease struct {
	Name      string     `json:"name"`
	Info      *ChartInfo `json:"info"`
	Chart     *HelmChart `json:"chart"`
	Version   uint       `json:"version"`
	Namespace string     `json:"namespace"`
}

type HelmChart struct {
	Metadata *Metadata              `json:"metadata"`
	Values   map[string]interface{} `json:"values"`
}

type Metadata struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	ApiVersion  string `json:"apiVersion"`
	AppVersion  string `json:"appVersion"`
	Type        string `json:"type"`
}

type ChartInfo struct {
	FirstDeployed string `json:"first_deployed"`
	LastDeployed  string `json:"last_deployed"`
	Deleted       string `json:"deleted"`
	Description   string `json:"description"`
	Status        string `json:"status"`
}
