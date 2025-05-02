package main

import (
	"github.com/aAmer0neee/load-balancer/internal/balancer/domain"
	"github.com/aAmer0neee/load-balancer/pkg/config"
	"github.com/aAmer0neee/load-balancer/pkg/logger"
)

func main(){
	cfg := domain.Cfg{}
	config.LoadConfig(cfg)
	
	logger.New()
}