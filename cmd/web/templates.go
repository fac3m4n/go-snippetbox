package main

import "snippetbox.kevo.dev/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
