package probes

/*type http_methods struct {
	GET  string
	POST string
	HEAD string
}*/

type request struct {
	url                      string
	method                   string
	data                     string
	validation_method        string
	validation_type          string
	validation_message       string
	validation_response_code uint8
	flags                    []string
}

type response struct {
	status  uint8
	headers string
	data    string
}

type FullResponse struct {
	request  request
	respnose response
}

var reqdata request

func Probe(reqdata_ request) {
	reqdata = reqdata_
}
