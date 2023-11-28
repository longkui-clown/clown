<h1 href="https://github.com/longkui-clown/clown" align="center">clown</h1>
<p align="center">
  <a href="https://github.com/longkui-clown/clown" title="中文文档" rel="nofollow"><img src="https://img.shields.io/badge/Doc-中文-blue.svg?style=flat" alt="中文文档"></a>
  <a href="https://github.com/longkui-clown/clown/blob/main/README_en.md" title="English document" rel="nofollow"><img src="https://img.shields.io/badge/Doc-English-blue.svg?style=flat" alt="English document"></a>
</p>

---

## 为什么使用 `clown`

- #### `clown` 提供一种项目初始化的思路，服务间动态添加，业务间的依赖关系约定

  - ##### 启动顺序：

    - `BeforeStartup`
    - Services (正序 `Init` + `OnStart`)
    - `AfterServiceStartup`
    - App (`Init` + `OnStart`)
    - `AfterAppStartup`
    - 等待关闭信号

  - ##### 关闭顺序：

    - 收到关闭信号
    - `BeforeShutdown`
    - App (`OnStop`)
    - `AfterAppShutdown`
    - Services (反序 `OnStop`)
    - `AfterServiceShutdown`

- #### 集成了系列工具函数，方便业务逻辑的开发

---

## 工具类文档

- ### [跳转文档](/docs/utils.md)

## 接口定义

- #### `Option`

  ```go
  type Option func(*Manager)
  ```

- #### `Logger`

  ```go
  type Logger interface {
    DEBUG(format string, args ...any)
    INFO(format string, args ...any)
    WARNING(format string, args ...any)
    ERROR(format string, args ...any)
  }
  ```

- #### `Service`

  ```go
  type Service interface {
    Init(*Manager) error
    GetName() string
    SetName(name string) Service
    OnStart() error
    OnStop() error
  }
  ```

  - `BaseService` 实现了默认的方法，通过嵌套 `BaseService` 来实现继承和重新的特性

---

## 如何使用 `clown`

- #### 确保使用 `go mod` 管理项目包

- #### 拉取 `clown` 包

  ```
  go get -u github.com/longkui-clown/clown
  ```

- #### 代码中使用

  ```go
  package main

  import (
    "github.com/longkui-clown/clown/core"
    "github.com/longkui-clown/clown/pkg/app"
  )

  func main() {
    // 配置初始化（显示调用）
    manager := core.NewManager(
      core.Name("BLOG"),
      core.KillWaitTTL(5000),                            // 可选，5000毫秒服务关闭等待
      core.BeforeStartup(before_startup),                // 可选，启动前回调
      core.AfterServiceStartup(after_service_startup),   // 可选，服务启动后回调
      core.AfterAppStartup(after_app_startup),           // 可选，app启动后回调
      core.BeforeShutdown(before_shutdown),              // 可选，关闭前回调
      core.AfterAppShutdown(after_app_shutdown),         // 可选，app关闭后回调
      core.AfterServiceShutdown(after_service_shutdown), // 可选，服务关闭后回调
      core.Services(
        ....            // 系列服务，需要实现 core.Service 接口
      ),
      core.App(app.NewTestApp("TestApp")),
    )

    manager.Run()
  }

  func before_startup(m *core.Manager) {

  }

  func after_service_startup(m *core.Manager) {

  }

  func after_app_startup(m *core.Manager) {

  }

  func before_shutdown(m *core.Manager) {

  }

  func after_app_shutdown(m *core.Manager) {

  }

  func after_service_shutdown(m *core.Manager) {

  }

  ```

- #### 管理器可选初始化选项

  | 函数                                   | 必选 | 描述                                                            |
  | :------------------------------------- | :--: | :-------------------------------------------------------------- |
  | `Name(name string) Option`             |      | 设置管理器名字                                                  |
  | `KillWaitTTL(ttl int64) Option`        |      | 设置停服超时等待设置，不设置则等待退出完成，单位 `毫秒`         |
  | `Services(services ...Service) Option` |      | 依赖服务添加, 系列服务，需要实现 `core.Service` 接口            |
  | `App(app Service) Option`              | `是` | 主体服务设置, 主体服务设置，需要实现 `core.Service` 接口        |
  | `SetLogger(l Logger) Option`           |      | 可选，管理器日志设置, 需要实现 `core.Logger` 接口，不设置使用 `DefaultManagerLogger` |
  | `BeforeStartup(fn Option)`             |      | 可选，生命周期 `钩子函数`, 启动前回调                           |
  | `AfterServiceStartup(fn Option)`       |      | 可选，生命周期 `钩子函数`, 服务启动后回调                       |
  | `AfterAppStartup(fn Option)`           |      | 可选，生命周期 `钩子函数`, app 启动后回调                       |
  | `BeforeShutdown(fn Option)`            |      | 可选，生命周期 `钩子函数`, 关闭前回调                           |
  | `AfterAppShutdown(fn Option)`          |      | 可选，生命周期 `钩子函数`, app 关闭后回调                       |
  | `AfterServiceShutdown(fn Option)`      |      | 可选，生命周期 `钩子函数`, 服务关闭后回调                       |
