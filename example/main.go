package main

import "github.com/metooweb/goservice"

func main() {

	cfg := &goservice.Config{}
	cfg.Name = "test_service"
	cfg.DisplayName = "测试"
	cfg.Run = func() error {

		return nil
	}

	goservice.Init(cfg)

}
