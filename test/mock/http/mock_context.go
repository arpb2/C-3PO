package http

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"golang.org/x/xerrors"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/gin-gonic/gin"
)

type worker struct {
	recorder *httptest.ResponseRecorder
	aborted  bool
}

func (w *worker) WriteJson(code int, obj interface{}) {
	if w.aborted {
		return
	}

	w.recorder.WriteHeader(code)

	j, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	_, err = w.recorder.Write(j)
	if err != nil {
		panic(err)
	}
}

func (w *worker) WriteString(code int, format string, values ...interface{}) {
	if w.aborted {
		return
	}

	w.recorder.WriteHeader(code)

	_, err := w.recorder.WriteString(fmt.Sprintf(format, values...))
	if err != nil {
		panic(err)
	}
}

func (w *worker) WriteStatus(code int) {
	if w.aborted {
		return
	}

	w.recorder.WriteHeader(code)
}

func (m *worker) NextHandler() {
	// Nothing
}

func (m *worker) IsAborted() bool {
	return m.aborted
}

func (m *worker) AbortTransaction() {
	defer func() {
		m.aborted = true
	}()
	if !m.aborted {
		m.recorder.WriteHeader(500)
	}
}

func (m *worker) AbortTransactionWithStatus(code int, jsonObj interface{}) {
	defer func() {
		m.aborted = true
	}()
	if !m.aborted {
		m.recorder.WriteHeader(code)

		j, err := json.Marshal(jsonObj)
		if err != nil {
			panic(err)
		}
		_, err = m.recorder.Write(j)
		if err != nil {
			panic(err)
		}
	}
}

func (m *worker) AbortTransactionWithError(err error) {
	defer func() {
		m.aborted = true
	}()
	if !m.aborted {
		var httpError http.Error
		if xerrors.As(err, &httpError) {
			if httpError.Code >= 200 && httpError.Code < 300 {
				fmt.Printf(
					"Request halted with code '%d' and message '%s' when its a successful response",
					httpError.Code,
					httpError.Reason,
				)
			} else {
				m.recorder.WriteHeader(httpError.Code)

				j, err := json.Marshal(map[string]string{
					"error": httpError.Reason,
				})
				if err != nil {
					panic(err)
				}
				_, err = m.recorder.Write(j)
				if err != nil {
					panic(err)
				}
			}
		} else {
			m.recorder.WriteHeader(500)

			j, err := json.Marshal(map[string]string{
				"error": "internal error",
			})
			if err != nil {
				panic(err)
			}
			_, err = m.recorder.Write(j)
			if err != nil {
				panic(err)
			}
		}
	}
}

func CreateTestContext() (*http.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	worker := &worker{
		recorder: httptest.NewRecorder(),
	}
	ctx := &http.Context{
		Writer:     worker,
		Middleware: worker,
	}

	return ctx, worker.recorder
}
