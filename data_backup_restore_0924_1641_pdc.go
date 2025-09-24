// 代码生成时间: 2025-09-24 16:41:52
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "archive/zip"
    "io"
    "log"
    "revel"
)

// BackupService 定义数据备份和恢复的服务
type BackupService struct {
    *revel.Controller
}

// Backup 处理数据备份的请求
func (c BackupService) Backup() revel.Result {
    // 获取备份文件的目录
    backupDir := revel.Config.StringDefault("backup.dir", "./backup")
    // 获取需要备份的文件或目录
    source := revel.Config.StringDefault("backup.source", "./")
    
    // 确保备份目录存在
    if _, err := os.Stat(backupDir); os.IsNotExist(err) {
        if err := os.MkdirAll(backupDir, 0755); err != nil {
            return c.RenderError(err)
        }
    }

    // 构建备份文件名和路径
    backupFile, err := filepath.Abs(filepath.Join(backupDir, fmt.Sprintf("backup_%s.zip", time.Now().Format("20060102_150405"))))
    if err != nil {
        return c.RenderError(err)
    }

    // 创建ZIP文件
    file, err := os.Create(backupFile)
    if err != nil {
        return c.RenderError(err)
    }
    defer file.Close()

    // 创建一个ZIP写入器
    writer := zip.NewWriter(file)
    defer writer.Close()

    // 添加文件到ZIP
    if err := addFilesToZip(source, writer, ""); err != nil {
        return c.RenderError(err)
    }

    // 返回备份文件的路径
    return c.RenderText(backupFile)
}

// Restore 处理数据恢复的请求
func (c BackupService) Restore() revel.Result {
    // 获取备份文件的目录
    backupDir := revel.Config.StringDefault("backup.dir", "./backup")
    // 获取备份文件的名称
    backupFile := c.Params.Files["backupFile"]

    // 检查备份文件是否存在
    if _, err := os.Stat(backupFile); os.IsNotExist(err) {
        return c.RenderError(err)
    }

    // 解压备份文件
    dest := revel.Config.StringDefault("restore.dest", "./")
    if err := unzip(backupFile, dest); err != nil {
        return c.RenderError(err)
    }

    // 返回恢复成功消息
    return c.RenderText("Restore successful")
}

// addFilesToZip 递归添加文件到ZIP
func addFilesToZip(path string, zw *zip.Writer, basePath string) error {
    files, _ := ioutil.ReadDir(path)
    for _, file := range files {
        filePath := filepath.Join(path, file.Name())
        if file.IsDir() {
            // 创建目录
            if _, err := zw.Create(filepath.Join(basePath, file.Name())); err != nil {
                return err
    }
            // 递归添加文件
            if err := addFilesToZip(filePath, zw, filepath.Join(basePath, file.Name())); err != nil {
                return err
            }
        } else {
            // 添加文件
            if _, err := zw.Create(filepath.Join(basePath, file.Name())); err != nil {
                return err
            }
            srcFile, err := os.Open(filePath)
            if err != nil {
                return err
            }
            defer srcFile.Close()
            dstFile, err := zw.Create(filepath.Join(basePath, file.Name()))
            if err != nil {
                return err
            }
            _, err = io.Copy(dstFile, srcFile)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

// unzip 解压ZIP文件
func unzip(src, dest string) error {
    srcFile, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    for _, f := range srcFile.File {
        filePath := filepath.Join(dest, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(filePath, os.ModePerm)
        } else {
            if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
                return err
            }
            file, err := srcFile.File.Open()
            if err != nil {
                return err
            }
            defer file.Close()
            
            destFile, err := os.Create(filePath)
            if err != nil {
                return err
            }
            defer destFile.Close()
            _, err = io.Copy(destFile, file)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

// RenderError 错误处理
func (c BackupService) RenderError(err error) revel.Result {
    return c.RenderJSON(map[string]string{
        "error": err.Error(),
    })
}
