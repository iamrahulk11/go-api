package dto

type BaseResponseDto[T any] struct {
	Result ResultResponseDto `json:"result"`
	Data   T                 `json:"data,omitempty"`
}

type ResultResponseDto struct {
	Flag        int    `json:"flag"`
	FlagMessage string `json:"flag_message"`
}
