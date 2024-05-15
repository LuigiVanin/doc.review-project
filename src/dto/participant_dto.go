package dto

type AddParticipantDto struct {
	UserId string `json:"user_id" validate:"required,uuid4"`
}
