package assets

import "embed"

//go:embed db/ddl.sql
//go:embed swagger/*
var Files embed.FS
