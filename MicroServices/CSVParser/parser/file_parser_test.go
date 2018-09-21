package parser_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"dbstore/MicroServices/CSVParser/parser"
)

func TestParse_FileNotThere(t *testing.T) {

	err := parser.Parse("blah blah", nil)
	expectedErr := fmt.Errorf("error parsing file: error opening csv file: open blah blah: no such file or directory")
	require.Error(t, err)
	assert.Equal(t, err, expectedErr)
}

//TODO: mock dbClient and then inject the client into file parser and write sunny day and failure cases.
