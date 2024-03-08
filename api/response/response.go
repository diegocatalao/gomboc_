package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
)

type MetadataResponse struct {
	StatusCode int    `json:"status_code"` // The HTTP status code received
	Message    string `json:"message"`     // The response message for the user
	Trace      string `json:"trace"`       // The trace hash for this request
}

type PaginationResponse struct {
	Page int `json:"page"` // The current page of data
	Size int `json:"size"` // The size of data array
}

type OutboundResponse struct {
	Metadata   *MetadataResponse   `json:"metadata"`   // The info about this request
	Data       any                 `json:"data"`       // The JSON data for this response
	Pagination *PaginationResponse `json:"pagination"` // The pagination section
}

type InboundResponse struct {
	Writer  http.ResponseWriter
	Message string
	Status  int
	Data    any
	Page    *int
	Size    *int
}

// MakeResponse constructs and sends a JSON response to the client.
func (inbound InboundResponse) MakeResponse() {
	var pagination *PaginationResponse = nil

	inbound.Writer.WriteHeader(inbound.Status)         // Write the HTTP status code into the response
	trace := inbound.Writer.Header().Get("x-trace-id") // Get the trace ID for outgoing responses

	if inbound.Status == http.StatusOK && (inbound.Page != nil || inbound.Size != nil) {
		pagination = &PaginationResponse{
			Page: *inbound.Page,
			Size: *inbound.Size,
		}
	}

	json.NewEncoder(inbound.Writer).Encode(OutboundResponse{
		Metadata: &MetadataResponse{
			Trace:      trace,
			Message:    inbound.Message,
			StatusCode: inbound.Status,
		},
		Pagination: pagination,
		Data:       inbound.Data,
	})

	log.Trace().Msgf("Request %s: returns response with status code %d", trace, inbound.Status)
}

func SuccessResponse(w http.ResponseWriter, data any, status int) {
	inbound := InboundResponse{w, http.StatusText(status), status, data, nil, nil}
	inbound.MakeResponse()
}

func SuccessPaginatedResponse(w http.ResponseWriter, data any, status int, page int, size int) {
	inbound := InboundResponse{w, http.StatusText(status), status, data, &page, &size}
	inbound.MakeResponse()
}

func BadResponse(w http.ResponseWriter, status int, messages ...string) {
	message := http.StatusText(status) + ". " + strings.Join(messages, " ")

	inbound := InboundResponse{w, message, status, nil, nil, nil}
	inbound.MakeResponse()
}
