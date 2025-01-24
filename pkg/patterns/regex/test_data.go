// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package regex

import (
	"github.com/andrerrcosta2/gtools/core/data/str/iterables"
)

var simpleRegexes = iterables.Map[string, string]{
	"word":            `^\w+$`,                    // Match a single word
	"digits":          `^\d+$`,                    // Match a string of digits
	"hex-color":       `^#[0-9a-fA-F]{3,6}$`,      // Match a hex color code
	"whitespace":      `^\s+$`,                    // Match a string of whitespace characters
	"uppercase":       `^[A-Z]+$`,                 // Match a string of uppercase letters
	"lowercase":       `^[a-z]+$`,                 // Match a string of lowercase letters
	"alphanumeric":    `^[a-zA-Z0-9]+$`,           // Match alphanumeric characters
	"comma-separated": `^(\w+)(,\s*\w+)*$`,        // Match comma-separated words
	"binary":          `^[01]+$`,                  // Match a binary number
	"float":           `^[+-]?(\d*\.\d+|\d+\.?)$`, // Match a floating point number
}

var complexRegexes = iterables.Map[string, string]{
	"email":           `^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$`,
	"url":             `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`,
	"ip":              `^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`,
	"date-format":     `^\d{4}-\d{2}-\d{2}$`,
	"date":            `^(\d{4})-(\d{2})-(\d{2})$`,
	"strong-password": `^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[\W_]).{8,}$`,
	"phone-number":    `^\+?(\d{1,3})?[-.\s]?\(?\d{1,4}\)?[-.\s]?\d{1,4}[-.\s]?\d{1,9}$`,
	"hex":             `^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`,
	"file-path":       `^(/[\w_\-\.]+)+$|^[a-zA-Z]:\\([\w_\-\.]+\\?)+$`,
}
