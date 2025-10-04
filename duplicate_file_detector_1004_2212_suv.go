// 代码生成时间: 2025-10-04 22:12:00
package main

import (
    "errors"
    "fmt"
    "io/fs"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "github.com/revel/revel"
)

// FileHash represents a file and its hash.
type FileHash struct {
    Path string
    Hash string
}

// detectDuplicates scans the directory for duplicate files based on content.
func detectDuplicates(path string) ([]FileHash, error) {
    var duplicates []FileHash
    fileMap := make(map[string][]string)

    err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }

        data, err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }

        // Calculate hash of file content.
        hash := string(data) // Simplified for demonstration; in production, use a proper hash function.
        fileMap[hash] = append(fileMap[hash], path)

        return nil
    })

    if err != nil {
        return nil, err
    }

    // Find and collect duplicates.
    for _, paths := range fileMap {
        if len(paths) > 1 {
            for _, path := range paths {
                fileInfo, err := os.Stat(path)
                if err != nil {
                    return nil, err
                }
                duplicates = append(duplicates, FileHash{Path: path, Hash: fmt.Sprintf("%x", hash)})
            }
        }
    }

    return duplicates, nil
}

func main() {
    revel.RunProd()
}

// The following methods are required by the Revel framework for handling requests.

type App struct {
    *revel.Controller
}

func (c App) DetectDuplicates() revel.Result {
    // Get the directory path from the request parameter.
    directory := c.Params.Route.Get("directory")

    if directory == "" {
        return c.RenderError(revel.NewError("Directory path is required"))
    }

    duplicates, err := detectDuplicates(directory)
    if err != nil {
        return c.RenderError(err)
    }

    // Render the result as JSON.
    return c.RenderJSON(duplicates)
}
