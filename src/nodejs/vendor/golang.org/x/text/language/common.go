// Code generated by running "go generate" in golang.google.cn/x/text. DO NOT EDIT.

package language

// This file contains code common to the maketables.go and the package code.

// langAliasType is the type of an alias in langAliasMap.
type langAliasType int8

const (
	langDeprecated langAliasType = iota
	langMacro
	langLegacy

	langAliasTypeUnknown langAliasType = -1
)
