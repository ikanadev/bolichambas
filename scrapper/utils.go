package main

import "strings"

const (
	LaPaz      = "La Paz"
	Oruro      = "Oruro"
	Potosi     = "Potosí"
	Cochabamba = "Cochabamba"
	Chuquisaca = "Chuquisaca"
	Tarija     = "Tarija"
	SantaCruz  = "Santa Cruz"
	Pando      = "Pando"
	Beni       = "Beni"
)

func normalizeString(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ToLower(s)
	return s
}

func parseDepto(depto string) string {
	depto = normalizeString(depto)
	if strings.Contains(depto, "lapaz") {
		return LaPaz
	}
	if strings.Contains(depto, "oruro") {
		return Oruro
	}
	if strings.Contains(depto, "potos") {
		return Potosi
	}
	if strings.Contains(depto, "cochabamba") {
		return Cochabamba
	}
	if strings.Contains(depto, "chuquisaca") {
		return Chuquisaca
	}
	if strings.Contains(depto, "tarija") {
		return Tarija
	}
	if strings.Contains(depto, "santacruz") {
		return SantaCruz
	}
	if strings.Contains(depto, "pando") {
		return Pando
	}
	if strings.Contains(depto, "beni") {
		return Beni
	}
	return ""
}
