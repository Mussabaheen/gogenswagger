package swagger

// JSON represents the json generated by the swagger docs
type JSON struct {
	Swagger string          `json:"swagger"` // represents the swagger version
	Host    string          `json:"host"`    // url of the host
	Paths   map[string]Path `json:"paths"`   // represents the endpoints in swagger
}

// RestAPI represents the API
type RestAPI struct {
	Description string              `json:"description"` // description of the API
	Tags        []string            `json:"tags"`        // tags used to group the API
	OperationID string              `json:"operationId"` // unique id to represent the API
	Responses   map[string]Response `json:"responses"`   // expected responses for the API
}

// Path represents the different type of endpoints
type Path struct {
	Get    *RestAPI `json:"get"`    // GET endpoint
	Post   *RestAPI `json:"post"`   // POST endpoint
	Put    *RestAPI `json:"put"`    // PUT endpoint
	Delete *RestAPI `json:"delete"` // DELETE endpoint
	Update *RestAPI `json:"update"` // UPDATE endpoint
}

// Response represents response for the API
type Response struct {
	Description string `json:"description"` // description for the specific response
}
