package controller

import (
	"encoding/json"
	"net/http"
	"stakeholders-service/dto"
	"stakeholders-service/service"
)


type JWTController struct {
	JwtService *service.JWTService
}

func (controller *JWTController) GenerateAccsessToken(writer http.ResponseWriter, req *http.Request) {
	var tokenInfo dto.TokenInfo
	err := json.NewDecoder(req.Body).Decode(&tokenInfo)
	
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}


	token,err := controller.JwtService.GenerateAccsessToken(tokenInfo.UserId,tokenInfo.Username,tokenInfo.PresonId,tokenInfo.Role)
	if err != nil {
		println("Error while generating token")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	
	
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	tourJSON, err := json.Marshal(token)
	if err != nil {
    println("Error while encoding tour to JSON")
	}
    writer.WriteHeader(http.StatusOK)
	writer.Write(tourJSON)
    

}