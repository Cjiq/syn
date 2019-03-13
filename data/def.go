package data

type Result struct {
	ResultSet ResultSet `json:"data"`
}

type ResultSet struct {
	DefinitionData DefinitionData `json:"definitionData"`
}

type DefinitionData struct {
	Definitions []Definition `json:"definitions"`
	Entry       string       `json:"entry"`
}

type Definition struct {
	Synonyms []Synonym `json:"synonyms"`
	Type     string    `json:"pos"`
	Text     string    `json:"definition"`
}

type Synonym struct {
	Similarity int    `json:"similarity,string"`
	Term       string `json:"term"`
}
