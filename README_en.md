<h1 href="https://github.com/longkui-clown/clown" align="center">clown</h1>
<p align="center">
  <a href="https://github.com/longkui-clown/clown" title="中文文档" rel="nofollow"><img src="https://img.shields.io/badge/Doc-中文-blue.svg?style=flat" alt="中文文档"></a>
  <a href="https://github.com/longkui-clown/clown/blob/main/README_en.md" title="English document" rel="nofollow"><img src="https://img.shields.io/badge/Doc-English-blue.svg?style=flat" alt="English document"></a>
</p>

---

## Why use `clown`

- #### `clown` provides an idea of project initialization, dynamic addition between services, and dependency agreement between businesses

  - ##### Start order：

    - `BeforeStartup`
    - Services (positive sequence `Init` + `OnStart`)
    - `AfterServiceStartup`
    - App (`Init` + `OnStart`)
    - `AfterAppStartup`
    - wait signal

  - ##### Stop order：

    - signal received
    - `BeforeShutdown`
    - App (`OnStop`)
    - `AfterAppShutdown`
    - Services (reverse sequence `OnStop`)
    - `AfterServiceShutdown`

- #### Provides a series of tool functions, easy to develop

---

## The `utils` package

- ### [Jump docs](/docs/utils.md)

## Interfaces define

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

  - `BaseService` implements the default method, by nesting `BaseService` to implement inherited and overridden features.

---

## How to use `clown`

- #### make sure use `go mod` to manage packages.

- #### go get `clown` package

  ```
  go get -u github.com/longkui-clown/clown
  ```

- #### use `clown` in code

  ```go
  package main

  import (
    "github.com/longkui-clown/clown/core"
    "github.com/longkui-clown/clown/pkg/app"
  )

  func main() {
    // Create application manager
    manager := core.NewManager(
      core.Name("BLOG"),
      core.KillWaitTTL(5000),                            // Optional，5000 ms wait when stop services
      core.BeforeStartup(before_startup),                // Optional，callback function before start up
      core.AfterServiceStartup(after_service_startup),   // Optional，callback function after service start up
      core.AfterAppStartup(after_app_startup),           // Optional，callback function after app startup
      core.BeforeShutdown(before_shutdown),              // Optional，callback function before shutdown
      core.AfterAppShutdown(after_app_shutdown),         // Optional，callback function after app shutdown
      core.AfterServiceShutdown(after_service_shutdown), // Optional，callback function after service shutdown
      core.Services(
        ....            // Some services，need to implement core.Service interface
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

- #### `Manage` optional functions desc

  | Functions                              | Optional | Desc                                                            |
  | :------------------------------------- | :--: | :-------------------------------------------------------------- |
  | `Name(name string) Option`             |      | set manager name                                                  |
  | `KillWaitTTL(ttl int64) Option`        |      | Set timeout to wait for stopping, if not set, wait for exit to completely, pass in `milliseconds`         |
  | `Services(services ...Service) Option` |      | Some dependent services，need to implement `core.Service` interface            |
  | `App(app Service) Option`              | `Yes` | Main service set, need to implement `core.Service` interface        |
  | `SetLogger(l Logger) Option`           |      | Optional，set manager logger, need to implement `core.Logger` interface, if not set, use `DefaultManagerLogger` |
  | `BeforeStartup(fn Option)`             |      | Optional，callback hook function before start up                      |
  | `AfterServiceStartup(fn Option)`       |      | Optional，callback hook function after service start up                  |
  | `AfterAppStartup(fn Option)`           |      | Optional，callback hook function after app startup                  |
  | `BeforeShutdown(fn Option)`            |      | Optional，callback hook function before shutdown                  |
  | `AfterAppShutdown(fn Option)`          |      | Optional，callback hook function after app shutdown                  |
  | `AfterServiceShutdown(fn Option)`      |      | Optional，callback hook function after service shutdown                  |
