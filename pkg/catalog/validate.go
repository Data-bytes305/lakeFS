package catalog

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/treeverse/lakefs/pkg/graveler"
	"github.com/treeverse/lakefs/pkg/ident"
)

const (
	MaxPathLength = 1024
)

var (
	reValidRef          = regexp.MustCompile(`^[^\s]+$`)
	reValidBranchID     = regexp.MustCompile(`^\w[-\w]*$`)
	reValidRepositoryID = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{2,62}$`)
)

var (
	ErrInvalidType       = fmt.Errorf("invalid type: %w", ErrInvalid)
	ErrRequiredValue     = fmt.Errorf("required value: %w", ErrInvalid)
	ErrPathRequiredValue = fmt.Errorf("missing path: %w", ErrRequiredValue)
)

type ValidateFunc func(v interface{}) error

type ValidateArg struct {
	Name  string
	Value interface{}
	Fn    ValidateFunc
}

func Validate(args []ValidateArg) error {
	for _, arg := range args {
		err := arg.Fn(arg.Value)
		if err != nil {
			return fmt.Errorf("argument %s: %w", arg.Name, err)
		}
	}
	return nil
}

func MakeValidateOptional(fn ValidateFunc) ValidateFunc {
	return func(v interface{}) error {
		switch s := v.(type) {
		case string:
			if len(s) == 0 {
				return nil
			}
		case fmt.Stringer:
			if len(s.String()) == 0 {
				return nil
			}
		case nil:
			return nil
		}
		return fn(v)
	}
}

func ValidateStorageNamespace(v interface{}) error {
	s, ok := v.(graveler.StorageNamespace)
	if !ok {
		panic(ErrInvalidType)
	}
	if len(s) == 0 {
		return ErrRequiredValue
	}
	return nil
}

func ValidateRef(v interface{}) error {
	s, ok := v.(graveler.Ref)
	if !ok {
		panic(ErrInvalidType)
	}
	if len(s) == 0 {
		return ErrRequiredValue
	}
	if !reValidRef.MatchString(s.String()) {
		return ErrInvalidValue
	}
	return nil
}

func ValidateBranchID(v interface{}) error {
	s, ok := v.(graveler.BranchID)
	if !ok {
		panic(ErrInvalidType)
	}
	if len(s) == 0 {
		return ErrRequiredValue
	}
	branchName := s.String()
	if !reValidBranchID.MatchString(branchName) {
		return ErrInvalidValue
	}
	return nil
}

func ValidateTagID(v interface{}) error {
	s, ok := v.(graveler.TagID)
	if !ok {
		panic(ErrInvalidType)
	}
	// https://git-scm.com/docs/git-check-ref-format
	tag := string(s)
	if len(tag) == 0 {
		return ErrRequiredValue
	}
	if tag == "@" {
		return ErrInvalidValue
	}
	if strings.HasSuffix(tag, ".") || strings.HasSuffix(tag, ".lock") || strings.HasSuffix(tag, "/") {
		return ErrInvalidValue
	}
	if strings.Contains(tag, "..") || strings.Contains(tag, "//") || strings.Contains(tag, "@{") {
		return ErrInvalidValue
	}
	// Unlike git, we do allow '~'.  That supports migration from our previous ref format where commits started with a tilde.
	if strings.ContainsAny(tag, "^:?*[\\") {
		return ErrInvalidValue
	}
	for _, r := range tag {
		if isControlCodeOrSpace(r) {
			return ErrInvalidValue
		}
	}
	return nil
}

func isControlCodeOrSpace(r rune) bool {
	const space = 0x20
	return r <= space
}

func ValidateCommitID(v interface{}) error {
	s, ok := v.(graveler.CommitID)
	if !ok {
		panic(ErrInvalidType)
	}
	if len(s) == 0 {
		return ErrRequiredValue
	}
	if !ident.IsContentAddress(s.String()) {
		return ErrInvalidValue
	}
	return nil
}

func ValidateRepositoryID(v interface{}) error {
	s, ok := v.(graveler.RepositoryID)
	if !ok {
		panic(ErrInvalidType)
	}
	if len(s) == 0 {
		return ErrRequiredValue
	}
	if !reValidRepositoryID.MatchString(s.String()) {
		return ErrInvalidValue
	}
	return nil
}

func ValidatePath(v interface{}) error {
	s, ok := v.(Path)
	if !ok {
		panic(ErrInvalidType)
	}
	l := len(s.String())
	if l == 0 {
		return ErrPathRequiredValue
	}
	if l > MaxPathLength {
		return fmt.Errorf("%w: %d is above maximum length (%d)", ErrInvalidValue, l, MaxPathLength)
	}
	return nil
}

func ValidateRequiredString(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		panic(ErrInvalidType)
	}
	if len(s) == 0 {
		return ErrRequiredValue
	}
	return nil
}

func ValidateNonNegativeInt(v interface{}) error {
	i, ok := v.(int)
	if !ok {
		panic(ErrInvalidType)
	}
	if i < 0 {
		return ErrInvalidValue
	}
	return nil
}

var ValidatePathOptional = MakeValidateOptional(ValidatePath)
var ValidateTagIDOptional = MakeValidateOptional(ValidateTagID)
