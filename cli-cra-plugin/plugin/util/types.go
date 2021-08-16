package util

type CraContext struct {
	ToolchainId  string
	ToolchainCrn string
	IamToken     string
	ContextTrace string
	ServiceUrls  map[string]string
}
