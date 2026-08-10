package assets

import "embed"

//go:embed db/ddl.sql
//go:embed swagger/*
//go:embed swagger/node_modules/swagger-ui-dist/*
var Files embed.FS
