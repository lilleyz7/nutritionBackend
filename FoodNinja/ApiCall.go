package foodninja

import (
	"encoding/json"
	models "goTest/Models"
	"net/http"
	"os"
)

func GetNutritionFacts(url string) models.NutritionFacts {
	var res []models.NutritionFacts
	err := GetJSON(url, &res)
	if err != nil {
		panic(err)
	} else {
		return res[0]
	}

}

func GetJSON(url string, target interface{}) error {
	key := os.Getenv("API_KEY")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-RapidAPI-Key", key)
	req.Header.Add("X-RapidAPI-Host", "nutrition-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	return nil
}
