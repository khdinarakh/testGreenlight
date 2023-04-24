package main

import (
	"greenlight.bcc/internal/assert"
	"net/http"
	"testing"
)

func TestRateLimitMiddleware(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routesTest())
	defer ts.Close()

	for i := 0; i < 4; i++ {

		t.Run("Valid", func(t *testing.T) {
			code, _, _ := ts.get(t, "/v1/movies")

			assert.Equal(t, code, http.StatusOK)
		})
	}

	code, _, _ := ts.get(t, "/v1/movies")
	assert.Equal(t, code, http.StatusTooManyRequests)
}

//func TestRecoverPanic(t *testing.T) {
//	app := &application{}
//
//	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		panic("Something went wrong")
//	})
//
//	rr := httptest.NewRecorder()
//	req, err := http.NewRequest("GET", "/", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	app.recoverPanic(handler).ServeHTTP(rr, req)
//
//	if rr.Code != http.StatusInternalServerError {
//		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, rr.Code)
//	}
//}

//func TestMetricsMiddleware(t *testing.T) {
//	// Create a new expvar map to store response codes and counts
//	totalResponsesSentByStatus := expvar.NewMap("total_responses_sent_by_status")
//
//	// Create a new test HTTP handler that always returns a 200 response
//	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//	})
//
//	// Create a new request and recorder for the test
//	req := httptest.NewRequest(http.MethodGet, "/", nil)
//	rr := httptest.NewRecorder()
//
//	// Create a new application instance and add the metrics middleware to it
//	app := &application{}
//	handler := app.metrics(testHandler)
//
//	// Call the handler and record the response metrics
//	handler.ServeHTTP(rr, req)
//
//	// Verify that the total requests received count has been incremented
//	if count := expvar.Get("total_requests_received").(*expvar.Int).Value(); count != 1 {
//		t.Errorf("expected total requests received count of 1, but got %d", count)
//	}
//
//	// Verify that the total responses sent count has been incremented
//	if count := expvar.Get("total_responses_sent").(*expvar.Int).Value(); count != 1 {
//		t.Errorf("expected total responses sent count of 1, but got %d", count)
//	}
//
//	// Verify that the total processing time has been recorded
//	if count := expvar.Get("total_processing_time_μs").(*expvar.Int).Value(); count <= 0 {
//		t.Errorf("expected total processing time to be greater than 0, but got %d", count)
//	}
//
//	// Verify that the response code has been recorded in the expvar map
//	if count := totalResponsesSentByStatus.Get("200").(*expvar.Int).Value(); count != 1 {
//		t.Errorf("expected response code 200 to have count of 1, but got %d", count)
//	}
//}
//
//func TestEnableCORS(t *testing.T) {
//	// Set up test data
//	trustedOrigins := []string{"http://localhost:3000", "http://localhost:8080"}
//	app := &application{
//		config: config{
//			cors: struct {
//				trustedOrigins []string
//			}{trustedOrigins: trustedOrigins},
//		},
//	}
//
//	// Create a mock handler that will be called by the enableCORS handler
//	handlerCalled := false
//	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		handlerCalled = true
//	})
//
//	// Create a mock request and response
//	req := httptest.NewRequest(http.MethodGet, "/", nil)
//	req.Header.Set("Origin", "http://localhost:8080")
//	w := httptest.NewRecorder()
//
//	// Call the enableCORS handler with the mock handler and the mock request and response
//	app.enableCORS(mockHandler).ServeHTTP(w, req)
//
//	// Check that the Access-Control-Allow-Origin header was set
//	if w.Header().Get("Access-Control-Allow-Origin") != "http://localhost:8080" {
//		t.Errorf("Access-Control-Allow-Origin header was not set correctly")
//	}
//
//	// Check that the Access-Control-Allow-Methods header was set
//	if w.Header().Get("Access-Control-Allow-Methods") != "OPTIONS, PUT, PATCH, DELETE" {
//		t.Errorf("Access-Control-Allow-Methods header was not set correctly")
//	}
//
//	// Check that the Access-Control-Allow-Headers header was set
//	if w.Header().Get("Access-Control-Allow-Headers") != "Authorization, Content-Type" {
//		t.Errorf("Access-Control-Allow-Headers header was not set correctly")
//	}
//
//	// Check that the handler was called
//	if !handlerCalled {
//		t.Errorf("handler was not called")
//	}
//}
//
//func TestRateLimitMiddleware(t *testing.T) {
//	// Initialize a mock HTTP handler for testing.
//	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//	})
//
//	// Initialize a mock application configuration with enabled rate limiting.
//	cfg := config{
//		limiter: struct {
//			rps     float64
//			burst   int
//			enabled bool
//		}{
//			rps:     1,
//			burst:   2,
//			enabled: true,
//		},
//	}
//
//	// Initialize the application with the mock configuration and logger.
//	app := &application{
//		config: cfg,
//		logger: jsonlog.New(os.Stdout, jsonlog.LevelInfo),
//	}
//
//	// Initialize a request with a mocked IP address.
//	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
//	req.RemoteAddr = "127.0.0.1:12345"
//
//	// Initialize a recorder to capture the response.
//	recorder := httptest.NewRecorder()
//
//	// Call the middleware with the mock handler.
//	middleware := app.rateLimit(handler)
//	middleware.ServeHTTP(recorder, req)
//
//	// Assert that the response has a status code of 200.
//	if recorder.Code != http.StatusOK {
//		t.Errorf("expected response status code to be %d, got %d", http.StatusOK, recorder.Code)
//	}
//}
//
//func TestApplication_RateLimit(t *testing.T) {
//	// Set up the test application
//	app := &application{
//		config: config{
//			limiter: struct {
//				rps     float64
//				burst   int
//				enabled bool
//			}{rps: 1, burst: 2, enabled: true},
//		},
//	}
//
//	// Create a new request with a dummy IP address
//	req := httptest.NewRequest(http.MethodGet, "/", nil)
//	req.RemoteAddr = "127.0.0.1:12345"
//
//	// Create a new recorder to capture the response
//	rec := httptest.NewRecorder()
//
//	// Create a dummy handler function to pass to the rateLimit middleware
//	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//	})
//
//	// Call the rateLimit middleware
//	app.rateLimit(handler).ServeHTTP(rec, req)
//
//	// Check that the response is a 200 OK
//	if status := rec.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
//	}
//
//	// Call the rateLimit middleware again with the same IP address
//	app.rateLimit(handler).ServeHTTP(rec, req)
//
//	// Check that the response is still a 200 OK
//	if status := rec.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
//	}
//
//	// Call the rateLimit middleware a third time with the same IP address
//	app.rateLimit(handler).ServeHTTP(rec, req)
//
//	// Check that the response is a 429 Too Many Requests
//	if status := rec.Code; status != http.StatusTooManyRequests {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusTooManyRequests)
//	}
//
//	// Wait for the rate limiter to expire the IP address
//	time.Sleep(2 * time.Second)
//
//	// Call the rateLimit middleware again with the same IP address
//	app.rateLimit(handler).ServeHTTP(rec, req)
//
//	// Check that the response is a 200 OK
//	if status := rec.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
//	}
//}
