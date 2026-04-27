package api

import (
	"net/http"

	"github.com/mattrowley10/the_faywood_adapter/internal/core/domain/types"
)

func (s *Server) ListCollections(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	s.logger.Info("listing collections")

	mockCollections := []types.Collection{
		{
			CollectionID: "col_1",
			Title:        "Summer Collection",
			Slug:         "summer-collection",
		},
		{
			CollectionID: "col_2",
			Title:        "Winter Collection",
			Slug:         "winter-collection",
		},
	}

	s.jsonResponse(w, http.StatusOK, mockCollections)
}

func (s *Server) GetCollection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	collectionID := r.PathValue("collectionID")
	if collectionID == "" {
		s.errorResponse(w, http.StatusBadRequest, "invalid collection ID")
		return
	}

	s.logger.Info("fetching collection", "collection_id", collectionID)

	mockCollection := &types.Collection{
		CollectionID: collectionID,
		Title:        "Mock Collection",
		Slug:         "mock-collection",
	}

	s.jsonResponse(w, http.StatusOK, mockCollection)
}
