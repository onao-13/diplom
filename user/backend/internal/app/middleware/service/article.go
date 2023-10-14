package service

import (
	"backend/internal/app/middleware/database"
	"backend/internal/app/payload/response"

	"github.com/sirupsen/logrus"
)

type Article struct {
	db database.Article	
	log logrus.Logger
}

func NewArticle(db database.Article, log logrus.Logger) Article {
	return Article{db, log}
}

func (a *Article) Get(id int64) (response.Article, error) {
	res, err := a.db.Get(id)
	if err != nil {
		a.log.Error("Ошибка получения статьи: ", err.Error())		
	}

	return res, err
}

func (a *Article) List() (response.ArticleList, error) {
	categories, err := a.db.GetAllCategories()
	if err != nil {
		a.log.Errorln("Ошибка получения категорий ", err.Error())
		return response.ArticleList{}, err
	}

	var list = make([]response.ArticleCategory, len(categories))

	for _, category := range categories {
		articles, err := a.db.GetLimitArticlesByCategoryId(category.Id, 10)
		if err != nil {
			a.log.Errorln("Ошибка получения статей ", err.Error())
			continue
		}

		data := response.ArticleCategory{
			Name: category.Name,
			Articles: articles,
		}

		list = append(list, data)
	}

	return response.ArticleList{List: list}, nil
}

