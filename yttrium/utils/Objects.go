package yttrium_utils

type About struct {
	Name     string
	Version  string
	Codename string
}

type LoadedColumn struct {
	Exists bool
	Column Column
}

type Column struct {
	Rows []Row
}

type Row struct {
	Key   string
	Value string
}
