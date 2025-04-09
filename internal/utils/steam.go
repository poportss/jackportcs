package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchSteamProfile(apiKey, steamID string) (string, string, error) {
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=%s&steamids=%s", apiKey, steamID)

	resp, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("erro ao chamar API da Steam: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Response struct {
			Players []struct {
				PersonaName string `json:"personaname"`
				AvatarFull  string `json:"avatarfull"`
			} `json:"players"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", fmt.Errorf("erro ao decodificar resposta da Steam: %w", err)
	}

	if len(result.Response.Players) == 0 {
		return "", "", fmt.Errorf("usuário Steam não encontrado")
	}

	return result.Response.Players[0].PersonaName, result.Response.Players[0].AvatarFull, nil
}
