package main

import (
	"testing"

	"github.com/sarulabs/dingo/dingo/tests/app/pkg"

	"github.com/sarulabs/dingo/dingo/tests/app/generated_services/dic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildStruct(t *testing.T) {
	container, err := dic.NewContainer()
	require.Nil(t, err)

	expected3 := &pkg.BuildStructTestC{P1: "C"}
	expected2 := &pkg.BuildStructTestB{P1: "B", P2: expected3}
	expected1 := &pkg.BuildStructTestA{P1: "A", P2: expected2, P3: expected3}
	expected4 := &pkg.BuildStructTestA{P1: "value1", P2: expected2, P3: &pkg.BuildStructTestC{P1: "value2"}}

	assert.Equal(t, expected1, container.GetTestBuildStruct1())
	assert.Equal(t, expected2, container.GetTestBuildStruct2())
	assert.Equal(t, expected3, container.GetTestBuildStruct3())
	assert.Equal(t, expected4, container.GetTestBuildStruct4())

	res1, err := container.SafeGetTestBuildStruct1()
	assert.Nil(t, err)
	assert.Equal(t, expected1, res1)

	res2, err := container.SafeGetTestBuildStruct2()
	assert.Nil(t, err)
	assert.Equal(t, expected2, res2)

	res3, err := container.SafeGetTestBuildStruct3()
	assert.Nil(t, err)
	assert.Equal(t, expected3, res3)

	res4, err := container.SafeGetTestBuildStruct4()
	assert.Nil(t, err)
	assert.Equal(t, expected4, res4)
}
