package server

import "testing"

func TestGetToken(t *testing.T) {

	for i := 0; i < 10; i++ {
		_, err := GetToken("0", "ljkljlkj", "jadsljdsjldskjldsa", "http://localhost:8080/testing")
		if err != nil {
			t.Fail()
		}
	}

}
