// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package openapi

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// DirectionStep defines model for DirectionStep.
type DirectionStep struct {
	Id   *string `json:"id,omitempty"`
	Step *string `json:"step,omitempty"`
}

// DirectionStepList defines model for DirectionStepList.
type DirectionStepList = []DirectionStep

// Ingredient defines model for Ingredient.
type Ingredient struct {
	Id       *string `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
}

// IngredientList defines model for IngredientList.
type IngredientList = []Ingredient

// Recipe defines model for Recipe.
type Recipe struct {
	DirectionSteps *[]DirectionStep `json:"directionSteps,omitempty"`
	Id             *string          `json:"id,omitempty"`
	Ingredients    *[]Ingredient    `json:"ingredients,omitempty"`
	Name           *string          `json:"name,omitempty"`
}

// RecipeBook defines model for RecipeBook.
type RecipeBook struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// RecipeBookList defines model for RecipeBookList.
type RecipeBookList = []RecipeBook

// RecipeList defines model for RecipeList.
type RecipeList = []Recipe

// PostRecipeBookJSONRequestBody defines body for PostRecipeBook for application/json ContentType.
type PostRecipeBookJSONRequestBody = RecipeBook

// PutBookJSONRequestBody defines body for PutBook for application/json ContentType.
type PutBookJSONRequestBody = RecipeBook

// PostRecipeJSONRequestBody defines body for PostRecipe for application/json ContentType.
type PostRecipeJSONRequestBody = Recipe

// PutRecipeJSONRequestBody defines body for PutRecipe for application/json ContentType.
type PutRecipeJSONRequestBody = Recipe

// PostDirectionStepJSONRequestBody defines body for PostDirectionStep for application/json ContentType.
type PostDirectionStepJSONRequestBody = DirectionStep

// PutStepJSONRequestBody defines body for PutStep for application/json ContentType.
type PutStepJSONRequestBody = DirectionStep

// PostIngredientJSONRequestBody defines body for PostIngredient for application/json ContentType.
type PostIngredientJSONRequestBody = Ingredient

// PutIngredientJSONRequestBody defines body for PutIngredient for application/json ContentType.
type PutIngredientJSONRequestBody = Ingredient
