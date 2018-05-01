package stringslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHas(t *testing.T) {
	assert.True(t, Has([]string{"foo", "bar"}, "foo"))
	assert.True(t, Has([]string{"foo", "bar"}, "bar"))
	assert.False(t, Has([]string{"foo", "bar"}, "baz"))
	assert.False(t, Has([]string{"foo", "bar"}, "baR"))
}

func TestHasI(t *testing.T) {
	assert.True(t, HasI([]string{"foO", "bAr"}, "foo"))
	assert.True(t, HasI([]string{"foo", "baR"}, "bar"))
	assert.False(t, HasI([]string{"foo", "bar"}, "baz"))
}
