package edit

import (
	"encoding/json"
	"fmt"
	"log"
	"news/internal/model/news"
	"news/internal/model/newscategories"
	"news/internal/model/validation"
	"news/internal/repository/postgres"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	Id         int32   `json:"id"`
	Title      string  `json:"title,omitempty"`
	Content    *string `json:"content,omitempty"`
	Categories []int32 `json:"categories,omitempty"`
}

func UpdateNews(db *postgres.Postgres) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")
		if err != nil {
			return err
		}

		updatesCounter := 0
		var req Request
		if err := json.Unmarshal(ctx.Body(), &req); err != nil {
			log.Println(err)
			return err
		}

		newsToUpdate, err := db.DB.FindByPrimaryKeyFrom(news.NewsTable, id)
		if err != nil {
			log.Println(err)
			return err
		}

		categoriesRows, err := db.DB.SelectAllFrom(newscategories.NewsCategoriesView, "WHERE news_id = $1", id)
		if err != nil {
			log.Println(err)
			return err
		}

		if isEqual, err := validation.ValidateEqualsIds(id, req.Id); !isEqual && err != nil {
			log.Println(err)
			return err
		}

		if isValid, err := validation.ValidateCategories(req.Categories); isValid && err == nil {
			if len(categoriesRows) == len(req.Categories) {
				tail := "WHERE news_id = $2 AND category_id = $3"
				var columns []string
				columns = append(columns, "category_id")
				for i, category := range categoriesRows {
					categoryVal := category.(*newscategories.NewsCategories).CategoryID
					category.(*newscategories.NewsCategories).CategoryID = req.Categories[i]
					_, err = db.DB.UpdateView(category, columns, tail, id, categoryVal)
					if err != nil {
						log.Println(err)
						return err
					}

					updatesCounter++
				}
			}
		}
		if isValid, err := validation.ValidateTitle(req.Title); isValid && err == nil {
			newsToUpdate.(*news.News).Title = req.Title
			updatesCounter++
		}
		if isValid, err := validation.ValidateContent(req.Content); isValid && err == nil {
			newsToUpdate.(*news.News).Content = req.Content
			updatesCounter++
		}

		err = db.DB.Update(newsToUpdate)
		if err != nil {
			log.Println(err)
			return err
		}

		err = ctx.JSON(fmt.Sprintf("News succesfully updated: %d with %d updated rows", id, updatesCounter))
		if err != nil {
			log.Println(err)
			return err
		}

		return nil
	}
}
