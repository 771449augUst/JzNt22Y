// 代码生成时间: 2025-10-09 01:41:25
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "net/http"
    "revel"
    "strings"
)

// Node represents a blockchain node
type Node struct {
    Id   string `json:"id"`
    URL  string `json:"url"`
    Peers []string `json:"peers"`
}

// NodeService manages blockchain nodes
type NodeService struct {
    Nodes map[string]Node
}

// NewNodeService creates a new NodeService
func NewNodeService() *NodeService {
    return &NodeService{
        Nodes: make(map[string]Node),
    }
}

// AddNode adds a new node to the blockchain network
func (service *NodeService) AddNode(node Node) error {
    if _, exists := service.Nodes[node.Id]; exists {
        return errors.New("node with this ID already exists")
    }
    service.Nodes[node.Id] = node
    return nil
}

// RemoveNode removes a node from the blockchain network
func (service *NodeService) RemoveNode(nodeId string) error {
    if _, exists := service.Nodes[nodeId]; !exists {
        return errors.New("node not found")
    }
    delete(service.Nodes, nodeId)
    return nil
}

// GetNodes returns all nodes in the blockchain network
func (service *NodeService) GetNodes() []Node {
    nodes := make([]Node, 0, len(service.Nodes))
    for _, node := range service.Nodes {
        nodes = append(nodes, node)
    }
    return nodes
}

// AppController is the controller for our blockchain node management app
type AppController struct {
    revel.Controller
    NodeService *NodeService
}

// Index is the action that renders the index page
func (c AppController) Index() revel.Result {
    return c.Render(c.NodeService.GetNodes())
}

// AddNodeAction is the action to add a new node
func (c AppController) AddNodeAction(node Node) revel.Result {
    err := c.NodeService.AddNode(node)
    if err != nil {
        return c.RenderError(err)
    }
    return c.Redirect(Index)
}

// RemoveNodeAction is the action to remove a node
func (c AppController) RemoveNodeAction(nodeId string) revel.Result {
    err := c.NodeService.RemoveNode(nodeId)
    if err != nil {
        return c.RenderError(err)
    }
    return c.Redirect(Index)
}

func init() {
    revel.OnAppStart(func() {
        // Initialize the node service with some nodes
        nodeService := NewNodeService()
        revel.DEFAULT_CONTROLLER = &AppController{NodeService: nodeService}
    })
}

// Start the Revel server
func main() {
    revel.Run()
}
