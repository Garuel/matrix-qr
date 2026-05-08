package main

import (
	Bootstrap "matriz-procesador-go-api/bootstrap"
	Config "matriz-procesador-go-api/config"
)

func main() {
    cfg := Config.Load()
    app := Bootstrap.Initialize(cfg)
    app.Listen(":" + cfg.PORT)
}