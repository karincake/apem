package lang

// THe built-in language item list for error messages
var defaultList = langItem{
	"required":        "value is required",
	"integerPositive": "value must be a positive integer",
	"uuid":            "value must be a valid UUID",
	"parse-fail":      "parsing failed",
	"data-notFound":   "data can not be found",
}
