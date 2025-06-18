package service

import (
	"ecommerceGO/config"
	"ecommerceGO/internal/helper"
	"ecommerceGO/internal/repository"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config config.Config
}
