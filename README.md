# registry-father

CLI 和 CUI 结合的 DN11 注册信息管理器

## 项目结构说明

services: 为 CLI 和 CUI 提供处理方法
pages: CUI 的界面
cmd: CLI 命令

## 特性

- [x] AS 信息管理
  - [x] 信息编辑
  - [ ] 信息实时校验
  - [ ] 优雅的提示
- [x] 信息校验
- [ ] Git Commit 结构优化

## 使用方式

### 日常编辑AS信息

直接启动程序

### 检查AS是否冲突

```shell
./register-father check
```


## 贡献

非常感谢您参与到这个项目中来，我们欢迎任何类型的贡献。

如果您想为该项目贡献代码，请您遵守[约定式提交](https://www.conventionalcommits.org/zh-hans/v1.0.0/)规范。