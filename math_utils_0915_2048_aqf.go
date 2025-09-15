// 代码生成时间: 2025-09-15 20:48:21
package controllers

import (
    "encoding/json"
    "github.com/revel/revel"
)

// MathUtilsController handles requests related to mathematical operations.
type MathUtilsController struct {
    *revel.Controller
}

// Add provides an endpoint to add two numbers.
// @Title Add
// @Description Add two integers and return the result.
// @Param x query int true "First integer"
// @Param y query int true "Second integer"
// @Success 200 {object} controllers.Result{value=int} "The result of the addition."
// @Failure 400 {string} string "Invalid input."
// @Failure 500 {string} string "An internal error occurred."
// @Router /math/add [get]
func (c *MathUtilsController) Add() revel.Result {
    x := c.Params.Query.Get("x")
    y := c.Params.Query.Get("y")

    // Validate and convert input to integers
    intX, errX := strconv.Atoi(x)
    intY, errY := strconv.Atoi(y)

    if errX != nil || errY != nil {
        return c.RenderError(revel.NewError("Invalid input."))
    }

    // Perform addition
    result := intX + intY

    // Return the result
    return c.RenderJSON(Result{Value: result})
}

// Result is a simple struct to wrap the result of a calculation.
type Result struct {
    Value int `json:"value"`
}
