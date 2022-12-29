package webserver

type Request struct {
	Alerts   []Alert `json:"alerts"`
}

type Alert struct {
	Labels struct {
        	GepardecCluster        string `json:"gepardec_cluster"`
	} `json:"labels"`
}
