package restapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ResultInfo struct {
	StatusCode  int         `json:"status_code"`  // HTTP status code
	StatusText  string      `json:"status_text"`  // HTTP status text
	Message     string      `json:"message"`      // Hata mesajı
	Data        interface{} `json:"data"`         // Veri
	ContentType string      `json:"content_type"` // Content-Type
}

func NewResultInfo(code int, message string, content_type string, data interface{}) *ResultInfo {
	return &ResultInfo{
		StatusCode:  code,
		StatusText:  http.StatusText(code),
		Message:     message,
		Data:        data,
		ContentType: content_type,
	}
}

func (s *ResultInfo) SetStatusText(code int) {
	s.StatusText = http.StatusText(code)
}

func Initialize() {
	fmt.Println("Initializing restapi...")

	r := mux.NewRouter()

	amw := AuthMiddleWare{}
	amw.Init()
	r.Use(amw.Middleware)

	Routes(r)

	// port'u config'e taşı
	http.ListenAndServe(":8080", r)
}

func ApiResult(w http.ResponseWriter, r *http.Request, i ResultInfo) (bool, error) {

	if i.StatusCode >= 400 {
		return ApiResultError(w, r, i)
	}

	return ApiResultSuccess(w, r, i)

}

// / Hatalı response iletmek için kullanılır
func ApiResultError(w http.ResponseWriter, _ *http.Request, i ResultInfo) (bool, error) {

	w.Header().Set("Content-Type", i.ContentType)
	w.WriteHeader(i.StatusCode)

	// TODO: Burada json veya xml olarak dönüş yapılabilir
	w.Write([]byte(i.Data.(string)))

	return true, nil
}

// / Başarılı response iletmek için kullanılır
func ApiResultSuccess(w http.ResponseWriter, r *http.Request, i ResultInfo) (bool, error) {

	return true, nil
}
