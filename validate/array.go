package validate

import "fmt"

// Array validates array length.
type Array struct {
	MinLength    int
	MinLengthSet bool
	MaxLength    int
	MaxLengthSet bool
}

// SetMaxLength sets MaxLength validation.
func (t *Array) SetMaxLength(v int) {
	t.MaxLengthSet = true
	t.MaxLength = v
}

// SetMinLength sets MinLength validation.
func (t *Array) SetMinLength(v int) {
	t.MinLengthSet = true
	t.MinLength = v
}

// ValidateLength returns error if array length v is invalid.
func (t Array) ValidateLength(v int) error {
	if t.MaxLengthSet && v > t.MaxLength {
		return fmt.Errorf("len %d greater than maximum %d", v, t.MaxLength)
	}
	if t.MinLengthSet && v < t.MinLength {
		return fmt.Errorf("len %d less than minimum %d", v, t.MinLength)
	}

	return nil
}

// Set reports whether any validations are set.
func (t Array) Set() bool {
	return t.MaxLengthSet || t.MinLengthSet
}