package test_test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"strategy-pattern/handler"

	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("calculatePrice Handler", func() {
	e := echo.New()
	var (
		defaultReqJSON  = `{"customerID":"Default", "products":["classic","standout", "premium"]}`
		unileverReqJSON = `{"customerID":"unilever","products":["classic", "classic", "classic","premium"]}`
		appleReqJSON    = `{"customerID":"apple","products":["standout", "standout", "standout","premium"]}`
		nikeReqJSON     = `{"customerID":"nike","products":["premium", "premium", "premium","premium"]}`
	)

	Context("If the client is default", func() {
		It("should return 200 OK and total of  987.97", func() {
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(defaultReqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := handler.CalculatePrice(c)
			if err != nil {
				Fail(err.Error())
			}
			Expect(rec.Code).To(Equal(http.StatusOK))
			// log.Println(rec.Body.String())
			Expect(rec.Body.String()).To(Equal(`{"total":"987.97"}`))
		})
	})

	Context("If the client is Unilever", func() {
		It("should return 200 OK and total of  934.97", func() {
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(unileverReqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := handler.CalculatePrice(c)
			if err != nil {
				Fail(err.Error())
			}
			Expect(rec.Code).To(Equal(http.StatusOK))
			// log.Println(rec.Body.String())
			Expect(rec.Body.String()).To(Equal(`{"total":"934.97"}`))
		})
	})

	Context("If the client is Apple", func() {
		It("should return 200 OK and total of  1294.96", func() {
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(appleReqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := handler.CalculatePrice(c)
			if err != nil {
				Fail(err.Error())
			}
			Expect(rec.Code).To(Equal(http.StatusOK))
			// log.Println(rec.Body.String())
			Expect(rec.Body.String()).To(Equal(`{"total":"1294.96"}`))
		})
	})

	Context("If the client is Nike", func() {
		It("should return 200 OK and total of  1519.96", func() {
			req := httptest.NewRequest(echo.POST, "/", strings.NewReader(nikeReqJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			err := handler.CalculatePrice(c)
			if err != nil {
				Fail(err.Error())
			}
			Expect(rec.Code).To(Equal(http.StatusOK))
			// log.Println(rec.Body.String())
			Expect(rec.Body.String()).To(Equal(`{"total":"1519.96"}`))
		})
	})

})
