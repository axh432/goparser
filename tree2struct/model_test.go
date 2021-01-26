package tree2struct

import (
	"fmt"
	. "github.com/axh432/gogex"
	gogex_model "github.com/axh432/goparser/tree2struct/gogex-model"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func Test_GogexModel(t *testing.T) {
	t.Run("parse a gogex file and convert to struct", func(t *testing.T) {
		fileAsBytes, err := ioutil.ReadFile("./import_example.go")
		require.NoError(t, err)
		tree := Match(string(fileAsBytes), GogexFile)
		require.True(t, tree.IsValid)
		justLabels := tree.PruneToLabels()
		gogexFile, err := gogex_model.NewGogexFile(&justLabels.Children[0])
		getTypeDefinitionsFromGogexFile(gogexFile)
		fmt.Println(writeDefinitions())
		require.NoError(t, err)
	})

}
