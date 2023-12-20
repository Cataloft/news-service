package list

import (
	"log"
	"news/internal/model/news"
	"news/internal/model/newscategories"
	"news/internal/repository/postgres"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool
	News    []NewsWithCategories
}

type NewsWithCategories struct {
	Id         int32
	Title      string
	Content    *string
	Categories []int32
}

func GetListNews(db *postgres.Postgres) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var resp Response
		var List NewsWithCategories
		var ListNews []NewsWithCategories

		newsRows, err := db.DB.SelectAllFrom(news.NewsTable, "")
		if err != nil {
			log.Println(err)
			return err
		}

		for _, newsRow := range newsRows {
			List.Id = newsRow.(*news.News).ID
			List.Title = newsRow.(*news.News).Title
			List.Content = newsRow.(*news.News).Content
			categoryRows, _ := db.DB.SelectAllFrom(newscategories.NewsCategoriesView, "WHERE news_id = $1", List.Id)
			for _, categoryRow := range categoryRows {
				List.Categories = append(List.Categories, categoryRow.(*newscategories.NewsCategories).CategoryID)
			}
			ListNews = append(ListNews, List)
			List.Categories = nil
		}

		resp.Success = true
		resp.News = ListNews
		err = ctx.JSON(resp)
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	}
}
