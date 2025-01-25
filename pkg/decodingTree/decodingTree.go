package decodingTree

import (
	"errors"
	"fmt"
	"strings"
)

type DecodingTree struct {
	Value *rune
	Zero  *DecodingTree
	One   *DecodingTree
}

func BuildDecodingTree(table map[rune]string) (*DecodingTree, error) {
	root := &DecodingTree{}

	for char, code := range table {
		current := root
		for i, bit := range code {
			switch bit {
			case '0':
				if current.Zero == nil {
					current.Zero = &DecodingTree{}
				}
				current = current.Zero
			case '1':
				if current.One == nil {
					current.One = &DecodingTree{}
				}
				current = current.One
			default:
				return nil, fmt.Errorf("invalid bit '%c' in code for '%c'", bit, char)
			}

			if i < len(code)-1 && current.Value != nil {
				return nil, fmt.Errorf("code conflict: '%s' is prefix of '%c'", code[:i+1], *current.Value)
			}
		}

		if current.Zero != nil || current.One != nil {
			return nil, fmt.Errorf("code '%s' is prefix of another code", code)
		}

		if current.Value != nil {
			return nil, fmt.Errorf("duplicate code '%s'", code)
		}

		current.Value = &char
	}

	return root, nil
}

func (dt *DecodingTree) Decode(encoded string) (string, error) {
	var builder strings.Builder
	current := dt

	for pos, bit := range encoded {
		switch bit {
		case '0':
			if current.Zero == nil {
				return "", fmt.Errorf("unexpected 0 at position %d", pos)
			}
			current = current.Zero
		case '1':
			if current.One == nil {
				return "", fmt.Errorf("unexpected 1 at position %d", pos)
			}
			current = current.One
		default:
			return "", fmt.Errorf("invalid bit '%c' at position %d", bit, pos)
		}

		if current.Value != nil {
			builder.WriteRune(*current.Value)
			current = dt
		}
	}

	if current != dt {
		return "", errors.New("incomplete encoding")
	}

	return builder.String(), nil
}
