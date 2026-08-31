package assets

import "embed"

//go:embed db/ddl.sql
//go:embed swagger/*
//go:embed swagger/node_modules/swagger-ui-dist/*
//go:embed web/*
//go:embed web/_app/*
//go:embed web/_app/**/*
var Files embed.FS
