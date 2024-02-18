package service

import (
	"backend/internal/app/database"
	"backend/internal/app/payload"
	"github.com/sirupsen/logrus"
)

type Article struct {
	db  database.Article
	log *logrus.Logger
}

func NewArticle(db database.Article, log *logrus.Logger) Article {
	return Article{db, log}
}

func (a *Article) Get(id int64) (article payload.Article, err error) {
	//res, err := a.db.Get(id)
	//if err != nil {
	//	a.log.Error("Ошибка получения статьи: ", err.Error())
	//}

	return
}

func (a *Article) List() (articles payload.ArticleList, err error) {
	//categories, err := a.db.GetAllCategories()
	//if err != nil {
	//	a.log.Errorln("Ошибка получения категорий ", err.Error())
	//	return payload.ArticleList{}, err
	//}
	//
	//var list = make([]payload.ArticleCategory, len(categories))
	//
	//for _, category := range categories {
	//	articles, err := a.db.GetLimitArticlesByCategoryId(category.Id, 10)
	//	if err != nil {
	//		a.log.Errorln("Ошибка получения статей ", err.Error())
	//		continue
	//	}
	//
	//	data := payload.ArticleCategory{
	//		Name:     category.Name,
	//		Articles: articles,
	//	}
	//
	//	list = append(list, data)
	//}

	return
}
