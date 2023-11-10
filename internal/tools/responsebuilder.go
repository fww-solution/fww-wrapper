package tools

import "fww-wrapper/internal/data/dto"

func ResponseBuilder(data interface{}, meta interface{}) dto.BaseResponse {
	return dto.BaseResponse{
		Meta: meta,
		Data: data,
	}
}