package apis

import (
	"testing"
)

func TestSubscribe(t *testing.T) {
	output := Subscribe("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	if output != nil {
		t.Errorf("Expected status code %v, got %v", nil, output)
	}
}

func TestSubscribeError(t *testing.T) {
	output := Subscribe("test")
	if output != nil {
		t.Errorf("Expected status code %v, got %v", nil, output)
	}
	output2 := Subscribe("test")
	if output2 == nil {
		t.Errorf("Expected status code %v, got %v", nil, output2)
	}
}

func TestUnsubscribe(t *testing.T) {
	output := Unsubscribe("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	if output != nil {
		t.Errorf("Expected status code %v, got %v", nil, output)
	}
}

func TestUnsubscribeError(t *testing.T) {
	Unsubscribe("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	output := Unsubscribe("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	if output == nil {
		t.Errorf("Expected status code %v, got %v", nil, output)
	}

}
