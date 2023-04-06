package models

type DailyIntake struct {
	Totals []NutritionFacts
}

type NutritionFacts struct {
	Sugar       float32 `json:"sugar_g" bson:"sugar_g"`
	Fiber       float32 `json:"fiber_g" bson:"fiber_g"`
	Sodium      float32 `json:"sodium_mg" bson:"sodium_mg"`
	Potassium   float32 `json:"potassium_mg" bson:"potassium_mg"`
	TotalFat    float32 `json:"total_fat_g" bson:"total_fat_g"`
	TotalSatFat float32 `json:"fat_saturated_g" bson:"fat_saturated_g"`
	Calories    float32 `json:"calories" bson:"calories"`
	Cholestoral float32 `json:"cholestoral_mg" bson:"cholestoral_mg"`
	Protein     float32 `json:"protein_g" bson:"protein_g"`
	Carbs       float32 `json:"carbohydrates_total_g" bson:"carbohydrates_total_g"`
	Name        string  `json:"name" bson:"name"`
	Username    string  `json:"username" bson:"username"`
}

type UpdatedNutritionFacts struct {
	Sugar       float32 `json:"sugar_g,omitempty" bson:"sugar_g,omitempty"`
	Fiber       float32 `json:"fiber_g,omitempty" bson:"fiber_g,omitempty"`
	Sodium      float32 `json:"sodium_mg,omitempty" bson:"sodium_mg,omitempty"`
	Potassium   float32 `json:"potassium_mg,omitempty" bson:"potassium_mg,omitempty"`
	TotalFat    float32 `json:"total_fat_g,omitempty" bson:"total_fat_g,omitempty"`
	TotalSatFat float32 `json:"fat_saturated_g,omitempty" bson:"fat_saturated_g,omitempty"`
	Calories    float32 `json:"calories,omitempty" bson:"calories,omitempty"`
	Cholestoral float32 `json:"cholestoral_mg,omitempty" bson:"cholestoral_mg,omitempty"`
	Protein     float32 `json:"protein_g,omitempty" bson:"protein_g,omitempty"`
	Carbs       float32 `json:"carbohydrates_total_g,omitempty" bson:"carbohydrates_total_g,omitempty"`
	Name        string  `json:"name,omitempty" bson:"name,omitempty"`
	Username    string  `json:"username,omitempty" bson:"username,omitempty"`
}
