package elastic_index

import (
	"os"
	"path/filepath"
)

type MetaType string

const (
	mappingsMetaType MetaType = "mappings"
	settingsMetaType MetaType = "settings"

	indexNameTime = "2006-01-02--15-04-05"
)

type IndexModel interface {
	GetTypeMapping() string
	GetDocumentId() string
}

type Index interface {
	GetName() string
	GetAlias() string
	GetSettings() string
	GetMappings() string
}

type esIndex struct {
	Name  string
	Alias string
}

func (i *esIndex) GetName() string {
	return i.Name
}

func (i *esIndex) GetAlias() string {
	return i.Alias
}

func (i *esIndex) GetSettings() string {
	return getIndexMetaDataFromJson(settingsMetaType, i.Alias)
}

func (i *esIndex) GetMappings() string {
	return getIndexMetaDataFromJson(mappingsMetaType, i.Alias)
}

func GetRecipe() Index {
	return NewRecipe("recipe")
}

func getIndexMetaDataFromJson(metaType MetaType, fileName string) string {
	file, err := filepath.Abs("app/config/elastic/" + string(metaType) + "/" + fileName + ".json")
	if err != nil {
		panic(err)
	}

	mapping, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(mapping)
}
