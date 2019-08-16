package pages

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func Test(t *testing.T){
	tests := []struct {
		name string
		in *http.Request
		out *httptest.ResponseRecorder
		expectedStatus int
		expectedBody string
	}{
		{
			name: "good",
			in : httptest.NewRequest("GET","/",nil),
			out: httptest.NewRecorder(),
			expectedStatus : http.StatusOK,
			expectedBody : message,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T){
			h := NewHandler(nil)
			h.Home(test.out,test.in)
			if test.out.Code != test.expectedStatus {
				t.Logf("expected: %d\ngot: %d\n",test.expectedStatus,test.out.Code)
				t.Fail()
			}
			body := test.out.Body.String()
			if body != test.expectedBody {
				t.Logf("expected: %s\ngot: %s\n",test.expectedBody,body)
				t.Fail()
			}
		})
	}
}