package httpclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	d, err := Get("https://www.bing.com")
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(d) > 0)
}

func TestGetWithParams(t *testing.T) {
	m := make(map[string]string, 2)
	m["qs"] = "utf-8"
	m["q"] = "test"
	d, err := GetWithParams("https://www.bing.com", m)
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(d) > 0)
	w, err2 := GetWithParams("https://www.bing.com/sasdasdas", nil)
	assert.Nil(t, err2)
	assert.Equal(t, true, len(w) == 0)
}

func TestPostDelete(t *testing.T) {
	r1, err := Delete("https://cn.bing.com/api/dx", "")
	assert.NotNil(t, err)
	assert.Empty(t, r1)
	r2, err := Put("https://cn.bing.com/api/dx", "")
	r3, _ := Post("https://cn.bing.com/api/dx", "")
	assert.NotNil(t, err)
	assert.Empty(t, r2)
	assert.Empty(t, r3)
}
