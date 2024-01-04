# gin-application-template

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=LeslieLeung_gin-application-template&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=LeslieLeung_gin-application-template)

## Description

This is a template for a gin application. 

The following features have already been implemented to simplify the initialization of a new project:

- [x] Project Structure
- [x] Dockerfile
- [x] Makefile
- [x] [Gin](https://github.com/gin-gonic/gin) Server
- [x] Config Management ([viper](https://github.com/spf13/viper))
- [x] [GoReleaser](https://goreleaser.com/) (GitHub Actions)
- [x] API Documentation ([swag](https://github.com/swaggo/swag))
- [x] Access Logs, Request ID
- [x] Health Check ([tavsec/gin-healthcheck](https://github.com/tavsec/gin-healthcheck))
- [x] Database Connection
- [x] Logging ([uber-go/zap](https://github.com/uber-go/zap))
- [ ] i18n Support
- [ ] Run as Script
- [ ] Observability
- [x] Graceful Shutdown ([gin-contrib/graceful](https://github.com/gin-contrib/graceful))

## Usage

### GitHub Template

1. Click on the green "Use this template" button
2. Click on "Create a new repository"
3. Input a name for your new repository, then click on "Create repository"

### gonew

1. Install gonew: `go install golang.org/x/tools/cmd/gonew@latest`
2. Open a terminal and run `gonew github.com/LeslieLeung/gin-application-template github.com/<your-username>/<your-repo-name>`

## Useful Make Commands

- `make run`: Run the application
- `make swag`: Generate API documentation
- `make dry-build`: Compile the application