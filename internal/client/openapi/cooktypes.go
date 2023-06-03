// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
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

// PostRecipeBooksJSONRequestBody defines body for PostRecipeBooks for application/json ContentType.
type PostRecipeBooksJSONRequestBody = RecipeBook

// PutRecipeBooksBookIdJSONRequestBody defines body for PutRecipeBooksBookId for application/json ContentType.
type PutRecipeBooksBookIdJSONRequestBody = RecipeBook

// PostRecipeBooksBookIdRecipesJSONRequestBody defines body for PostRecipeBooksBookIdRecipes for application/json ContentType.
type PostRecipeBooksBookIdRecipesJSONRequestBody = Recipe

// PutRecipeBooksBookIdRecipesRecipeIdJSONRequestBody defines body for PutRecipeBooksBookIdRecipesRecipeId for application/json ContentType.
type PutRecipeBooksBookIdRecipesRecipeIdJSONRequestBody = Recipe

// PostRecipeBooksBookIdRecipesRecipeIdDirectionStepsJSONRequestBody defines body for PostRecipeBooksBookIdRecipesRecipeIdDirectionSteps for application/json ContentType.
type PostRecipeBooksBookIdRecipesRecipeIdDirectionStepsJSONRequestBody = DirectionStep

// PutRecipeBooksBookIdRecipesRecipeIdDirectionStepsStepIdJSONRequestBody defines body for PutRecipeBooksBookIdRecipesRecipeIdDirectionStepsStepId for application/json ContentType.
type PutRecipeBooksBookIdRecipesRecipeIdDirectionStepsStepIdJSONRequestBody = DirectionStep

// PostRecipeBooksBookIdRecipesRecipeIdIngredientsJSONRequestBody defines body for PostRecipeBooksBookIdRecipesRecipeIdIngredients for application/json ContentType.
type PostRecipeBooksBookIdRecipesRecipeIdIngredientsJSONRequestBody = Ingredient

// PutRecipeBooksBookIdRecipesRecipeIdIngredientsIngredientIdJSONRequestBody defines body for PutRecipeBooksBookIdRecipesRecipeIdIngredientsIngredientId for application/json ContentType.
type PutRecipeBooksBookIdRecipesRecipeIdIngredientsIngredientIdJSONRequestBody = Ingredient
