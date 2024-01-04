package uuidutils

import "github.com/google/uuid"

func IsValid(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
