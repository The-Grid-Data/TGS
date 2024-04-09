package tgs_004

type EnumOption struct {
	ID          string
	Caption     string
	Description string
}

type ParamDefinition struct {
	ID                   string
	Description          string
	Validation           string
	IsLinkToAnotherTable bool
	EnumOptions          []EnumOption
}
