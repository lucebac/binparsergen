package binparsergen

import "fmt"

func GenerateCode(
	spec *ConversionSpec,
	profile map[string]*StructDefinition) string {
	result := fmt.Sprintf(`
package %s

// Autogenerated code from %s. Do not edit.

import (
    "encoding/binary"
    "fmt"
    "bytes"
    "io"
    "unicode/utf16"
    "unicode/utf8"
)

var (
   // Depending on autogenerated code we may use this. Add a reference
   // to shut the compiler up.
   _ = bytes.MinRead
   _ = fmt.Sprintf
   _ = utf16.Decode
   _ = binary.LittleEndian
   _ = utf8.RuneError
)

`, spec.Module, spec.Filename)
	profile_name := spec.Profile

	result += GenerateProfileCode(profile_name, profile)
	for struct_name, struct_def := range profile {
		struct_name := NormalizeName(struct_name)
		result += GenerateStructCode(struct_name, profile_name, struct_def)
		if spec.GenerateDebugString {
			result += GenerateDebugString(
				struct_name, profile_name, struct_def)
		}
	}

	result += GeneratePrototypes()

	return result
}
