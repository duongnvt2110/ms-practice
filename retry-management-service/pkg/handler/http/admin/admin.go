package admin

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"ms-practice/retry-management-service/pkg/model"
	"ms-practice/retry-management-service/pkg/usecase"

	"github.com/gorilla/mux"
)

type AdminHandler struct {
	dlqUC     usecase.DLQUsecase
	templates *template.Template
}

type listViewData struct {
	Records    []model.DLQRecord
	Page       int
	PageSize   int
	Total      int64
	TotalPages int
	PrevPage   int
	NextPage   int
}

type detailViewData struct {
	Record        *model.DLQRecord
	HeadersJSON   string
	PayloadText   string
	PayloadBase64 string
}

func NewAdminHandler(dlqUC usecase.DLQUsecase) (*AdminHandler, error) {
	tpls, err := parseTemplates()
	if err != nil {
		return nil, err
	}
	return &AdminHandler{
		dlqUC:     dlqUC,
		templates: tpls,
	}, nil
}

func (h *AdminHandler) List(w http.ResponseWriter, r *http.Request) {
	page := parseInt(r.URL.Query().Get("page"), 1)
	pageSize := parseInt(r.URL.Query().Get("page_size"), 20)

	records, total, err := h.dlqUC.List(r.Context(), page, pageSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("list dlq records failed: %v", err), http.StatusInternalServerError)
		return
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	prevPage := 0
	nextPage := 0
	if page > 1 {
		prevPage = page - 1
	}
	if page < totalPages {
		nextPage = page + 1
	}

	data := listViewData{
		Records:    records,
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
		PrevPage:   prevPage,
		NextPage:   nextPage,
	}

	if err := h.templates.ExecuteTemplate(w, "list.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *AdminHandler) Detail(w http.ResponseWriter, r *http.Request) {
	idRaw := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idRaw, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	record, err := h.dlqUC.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("get record failed: %v", err), http.StatusInternalServerError)
		return
	}
	if record == nil {
		http.Error(w, "record not found", http.StatusNotFound)
		return
	}

	payloadText := string(record.Payload)
	if len(payloadText) > 5000 {
		payloadText = payloadText[:5000] + "..."
	}

	data := detailViewData{
		Record:        record,
		HeadersJSON:   record.Headers,
		PayloadText:   payloadText,
		PayloadBase64: base64.StdEncoding.EncodeToString(record.Payload),
	}

	if err := h.templates.ExecuteTemplate(w, "detail.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func parseTemplates() (*template.Template, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	base := filepath.Join(cwd, "pkg", "handler", "http", "admin", "templates")
	listPath := filepath.Join(base, "list.html")
	detailPath := filepath.Join(base, "detail.html")
	return template.ParseFiles(listPath, detailPath)
}

func parseInt(raw string, fallback int) int {
	if raw == "" {
		return fallback
	}
	val, err := strconv.Atoi(raw)
	if err != nil || val <= 0 {
		return fallback
	}
	return val
}
