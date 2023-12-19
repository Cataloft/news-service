package edit

import (
	"encoding/json"
	"fmt"
	"news/internal/model/news"
	"news/internal/model/newscategories"
	"news/internal/repository/postgres"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Id         int32   `json:"id"`
	Title      string  `json:"title"`
	Content    *string `json:"content"`
	Categories []int32 `json:"categories"`
}

func UpdateNews(db *postgres.Postgres) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			return err
		}

		var req Request
		if err := json.Unmarshal(ctx.Body(), &req); err != nil {
			return err
		}

		newsToUpdate, err := db.DB.FindByPrimaryKeyFrom(news.NewsTable, id)
		if err != nil {
			return err
		}

		categoriesRows, err := db.DB.SelectAllFrom(newscategories.NewsCategoriesView, "WHERE news_id = $1", id)
		if err != nil {
			return err
		}

		if len(req.Categories) == len(categoriesRows) {
			tail := "WHERE news_id = $2 AND category_id = $3"
			var columns []string
			columns = append(columns, "category_id")
			for i, category := range categoriesRows {
				categoryVal := category.(*newscategories.NewsCategories).CategoryID
				category.(*newscategories.NewsCategories).CategoryID = req.Categories[i]
				_, err = db.DB.UpdateView(category, columns, tail, id, categoryVal)
				if err != nil {
					return err
				}
			}
		}

		newsToUpdate.(*news.News).Title = req.Title
		newsToUpdate.(*news.News).Content = req.Content

		err = db.DB.Update(newsToUpdate)
		if err != nil {
			return err
		}

		err = ctx.JSON(fmt.Sprintf("News succesfully updated: %d", id))
		if err != nil {
			return err
		}

		return nil
	}
}
