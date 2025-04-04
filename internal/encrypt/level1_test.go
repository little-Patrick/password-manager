package encrypt

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	fmt.Println("hellow")
	
    assert.Equal(t, 1, 1) // Simple test to ensure the setup is correct
}

// func TestNewLevel1(t *testing.T) {
// 	l1 := NewLevel1("mypassword", "12345")
// 	assert.NotNil(t, l1)
// 	assert.Equal(t, "mypassword", l1.password)
// 	assert.Equal(t, "12345", l1.key1)
//
// 	l2 := NewLevel1("mypassword", "")
// 	assert.NotNil(t, l2)
// 	assert.Equal(t, "mypassword", l2.password)
// 	assert.NotEmpty(t, l2.key1)
// }
