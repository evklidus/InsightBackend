package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Category struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type CoursePreview struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Name     string `json:"name"`
	Tag      string `json:"tag"`
}

type CoursePage struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Lessons []Lesson `json:"lessons"`
}

type Lesson struct {
	Name     string `json:"name"`
	VideoUrl string `json:"videoUrl"`
}

var categories = []Category{
	{Name: "Спорт", Tag: "sport"},
	{Name: "Программирование", Tag: "programming"},
}

var coursePreviews = []CoursePreview{
	{Id: 1, ImageUrl: "https://www.apple.com/v/apple-fitness-plus/l/images/meta/apple-fitness-plus__eafl9rq9woom_og.png", Name: "Пилатес от Веты", Tag: "sport"},
	{Id: 2, ImageUrl: "https://www.apple.com/v/apple-fitness-plus/l/images/meta/apple-fitness-plus__eafl9rq9woom_og.png", Name: "Силовая", Tag: "sport"},
	{Id: 3, ImageUrl: "https://miro.medium.com/max/1200/1*yvz6FsBEh-JGN_miQIMEXA.jpeg", Name: "Flutter", Tag: "programming"},
}

var coursePages = []CoursePage{
	{Id: 1, Lessons: []Lesson{
		{Name: "Первый день", VideoUrl: "https://player.vimeo.com/external/438451071.hd.mp4?s=863dcc7f2bd294d7968b25a2a867bd0ca1b6522e&profile_id=175&oauth2_token_id=57447761"},
	}},
}

func getCoursePreviewsByCategory(context *gin.Context) {
	category := context.Param("category")

	var coursePreviewsByCategory []CoursePreview

	for _, coursePreview := range coursePreviews {
		if coursePreview.Tag == category {
			coursePreviewsByCategory = append(coursePreviewsByCategory, coursePreview)
		}
	}
	if len(coursePreviewsByCategory) != 0 {
		context.IndentedJSON(http.StatusOK, coursePreviewsByCategory)
	} else {
		context.IndentedJSON(http.StatusNoContent, gin.H{"message": "Course previews not found"})
	}
}

func getCoursePageById(context *gin.Context) {
	stringId := context.Param("id")

	id, error := strconv.Atoi(stringId)
	if error == nil {
		for _, coursePage := range coursePages {
			if coursePage.Id == id {
				context.IndentedJSON(http.StatusOK, coursePage)
				return
			}
		}
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Please course is still building..."})
	} else {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please enter id as int"})
	}
}

func main() {
	router := gin.Default()
	router.GET("/categories", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, categories)
	})
	router.GET("/course_previews", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, coursePreviews)
	})
	router.GET("/course_previews/:category", getCoursePreviewsByCategory)
	router.GET("/course_pages", func(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, coursePages)
	})
	router.GET("/course_pages/:id", getCoursePageById)
	router.Run() // listen and serve on 0.0.0.0:8080
}
