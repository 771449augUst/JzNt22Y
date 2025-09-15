// 代码生成时间: 2025-09-16 01:18:44
package controllers

import (
    "math/rand"
    "time"
    "encoding/json"
    "revel"
)

// RandomNumberGeneratorController is the controller for generating random numbers.
type RandomNumberGeneratorController struct {
    revel.Controller
}

// Generate action generates a random number within a specified range.
func (c RandomNumberGeneratorController) Generate(min, max int) revel.Result {
    // Initialize the random seed with the current time to ensure randomness.
    rand.Seed(time.Now().UnixNano())

    // Generate a random number within the range [min, max].
    randomNumber := rand.Intn(max - min + 1) + min

    // Return the result as a JSON object.
    return c.RenderJSON(RandomNumberResult{
        Minimum:   min,
        Maximum:   max,
        Generated: randomNumber,
    })
}

// RandomNumberResult represents the result of the random number generation.
type RandomNumberResult struct {
    Minimum   int `json:"minimum"`
    Maximum  int `json:"maximum"`
    Generated int `json:"generated"`
}
