package service

import (
	"backend/internal/app/middleware/database"
	// "backend/internal/app/payload/request"
	// "backend/internal/app/payload/response"
)

type Article struct {
	db database.Article	
}

func NewArticle(db database.Article) Article {
	return Article{db}
}

// func (a *Article) View(req request.Artcile) response.Article {
// 	return a.db.View(req.Id)
// }

