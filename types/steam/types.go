package steam

// CommonResponse
// for some interface which need publisher api key are hard to test
// and no way to know their response, so use this type to replace
type CommonResponse = map[string]any
