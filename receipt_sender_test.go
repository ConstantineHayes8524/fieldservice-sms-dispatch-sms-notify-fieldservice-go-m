package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendSMSRequestFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v1/sms/send" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		var payload map[string]string
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Errorf("decode request: %v", err)
		}
		if len(payload) != 3 || payload["to"] != "+15551234567" || payload["body"] != "follow-up required" || payload["idempotency_key"] != "WO-42" {
			t.Errorf("unexpected SMS payload: %#v", payload)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true,"data":{"message_id":"test-id"}}`))
	}))
	defer server.Close()

	client := &Client{BaseURL: server.URL, Key: "test-key", HTTP: server.Client()}
	result, err := client.SendSMS("+15551234567", "follow-up required", "WO-42")
	if err != nil || result.MessageID != "test-id" {
		t.Fatalf("SendSMS result = %+v, %v", result, err)
	}
}
