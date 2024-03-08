package handlers

import (
	"encoding/json"
	"fmt"
	interfaces "gomboc/api/interfaces"
	"gomboc/api/models"
	response "gomboc/api/response"
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func GetNodeHandler(w http.ResponseWriter, r *http.Request) {
	response.BadResponse(w, http.StatusNotImplemented)
}

func CreateNodeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var iNode, oNode models.NodeModel

	if err = json.NewDecoder(r.Body).Decode(&iNode); err != nil {
		response.BadResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if oNode, err = interfaces.CreateNode(iNode); err != nil {
		response.BadResponse(w, http.StatusBadRequest, err.Error())
	}

	response.SuccessResponse(w, oNode, http.StatusCreated)
}

func GetNodesHandler(w http.ResponseWriter, r *http.Request) {
	var page, limit int
	var isActive, isPending bool

	if sPage := r.URL.Query().Get("page"); sPage == "" {
		page = 0
	} else {
		page, _ = strconv.Atoi(sPage)

		if page < 0 {
			emsg := fmt.Sprintf("Invalid page selection: try page > 0, not %d.", page)
			response.BadResponse(w, http.StatusUnprocessableEntity, emsg)
			return
		}
	}

	if sLimit := r.URL.Query().Get("limit"); sLimit == "" {
		limit = 100
	} else {
		limit, _ = strconv.Atoi(sLimit)

		if limit < 0 || limit > 100 {
			emsg := fmt.Sprintf("Invalid limit range: try 0 <= limit <= 100, not %d.", limit)
			response.BadResponse(w, http.StatusUnprocessableEntity, emsg)
			return
		}
	}

	if sIsActive := r.URL.Query().Get("page"); sIsActive == "" {
		isActive = true
	} else {
		isActive, _ = strconv.ParseBool(sIsActive)
	}

	if sIsPending := r.URL.Query().Get("page"); sIsPending == "" {
		isPending = false
	} else {
		isPending, _ = strconv.ParseBool(sIsPending)
	}

	if oNodes, err := interfaces.GetAllNodes(page, limit, isActive, isPending); err != nil {
		response.BadResponse(w, http.StatusNotFound)
	} else {
		response.SuccessPaginatedResponse(w, oNodes, http.StatusOK, page, len(oNodes))
	}

}

func GrantNodeAccessHandler(w http.ResponseWriter, r *http.Request) {
	response.BadResponse(w, http.StatusNotImplemented)
}

func RevokeNodeAccessHandler(w http.ResponseWriter, r *http.Request) {
	response.BadResponse(w, http.StatusNotImplemented)
}

func StartNodeSessionHandler(w http.ResponseWriter, r *http.Request) {
	response.BadResponse(w, http.StatusNotImplemented)
}

func RequestNodeRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var iNode, oNode models.NodeModel
	var rNode models.NodeRequestRegistrationModel

	if err = json.NewDecoder(r.Body).Decode(&rNode); err != nil {
		response.BadResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if rNode.PublicId == "" {
		log.Warn().Msg("Received a empty or invalid public id")

		response.BadResponse(w, http.StatusBadRequest, "Invalid public id")
		return
	}

	if iNode, err = interfaces.GetNodeByPublicId(rNode.PublicId); err != nil {
		log.Warn().Msgf("Device with public id '%s' not found. Start spoke mode.", rNode.PublicId)

		iNode = models.NodeModel{
			PublicId:   rNode.PublicId,
			IsActive:   false,
			IsPending:  true,
			LastIOTime: time.Now(),
		}

		if oNode, err = interfaces.CreateNode(iNode); err != nil {
			log.Err(err).Msgf("Unable to create node in spoke mode: %s", err.Error())

			response.BadResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		log.Warn().Msgf("Device with public id '%s' already registered.", iNode.PublicId)

		iNode.LastIOTime = time.Now()

		if oNode, err = interfaces.UpdateNode(iNode); err != nil {
			log.Err(err).Msgf("Unable to update device with public %s: ", err.Error())

			response.BadResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	response.SuccessResponse(w, oNode, http.StatusAccepted)
}
