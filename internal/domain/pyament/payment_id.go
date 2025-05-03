package pyament

import "github.com/google/uuid"

type ID uuid.UUID

func NewPaymentID() ID {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return ID(id)
}

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func ParsePaymentID(id string) (ID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return ID{}, err
	}

	return ID(parsedID), nil
}
