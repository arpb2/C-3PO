package http

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"golang.org/x/xerrors"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/gin-gonic/gin"
)

type worker struct {
	recorder *httptest.ResponseRecorder
	aborted  bool
	values   map[string]interface{}
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

func (w *worker) NextHandler() {
	// Nothing
}

func (w *worker) IsAborted() bool {
	return w.aborted
}

func (w *worker) AbortTransaction() {
	defer func() {
		w.aborted = true
	}()
	if !w.aborted {
		w.recorder.WriteHeader(500)
	}
}

func (w *worker) AbortTransactionWithStatus(code int, jsonObj interface{}) {
	defer func() {
		w.aborted = true
	}()
	if !w.aborted {
		w.recorder.WriteHeader(code)

		j, err := json.Marshal(jsonObj)
		if err != nil {
			panic(err)
		}
		_, err = w.recorder.Write(j)
		if err != nil {
			panic(err)
		}
	}
}

func (w *worker) SetValue(key string, value interface{}) {
	w.values[key] = value
}

func (w *worker) GetValue(key string) interface{} {
	return w.values[key]
}

func (w *worker) AbortTransactionWithError(err error) {
	defer func() {
		w.aborted = true
	}()
	if !w.aborted {
		var httpError http.Error
		if xerrors.As(err, &httpError) {
			if httpError.Code >= 200 && httpError.Code < 300 {
				fmt.Printf(
					"Request halted with code '%d' and message '%s' when its a successful response",
					httpError.Code,
					httpError.Reason,
				)
			} else {
				w.recorder.WriteHeader(httpError.Code)

				j, err := json.Marshal(map[string]string{
					"error": httpError.Reason,
				})
				if err != nil {
					panic(err)
				}
				_, err = w.recorder.Write(j)
				if err != nil {
					panic(err)
				}
			}
		} else {
			w.recorder.WriteHeader(500)

			j, err := json.Marshal(map[string]string{
				"error": "internal error",
			})
			if err != nil {
				panic(err)
			}
			_, err = w.recorder.Write(j)
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
		values:   make(map[string]interface{}),
	}
	ctx := &http.Context{
		Writer:     worker,
		Middleware: worker,
	}

	return ctx, worker.recorder
}
