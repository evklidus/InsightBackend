package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Category struct {
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
	Tag      string `json:"tag"`
}

type CoursePreview struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Name     string `json:"name"`
	Tag      string `json:"tag"`
}

type CoursePage struct {
	Id       int      `json:"id"`
	ImageUrl string   `json:"imageUrl"`
	Name     string   `json:"name"`
	Lessons  []Lesson `json:"lessons"`
}

type Lesson struct {
	Name     string `json:"name"`
	VideoUrl string `json:"videoUrl"`
}

var categories = []Category{
	{Name: "Спорт", ImageUrl: "https://images.unsplash.com/photo-1518611012118-696072aa579a?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=3540&q=80", Tag: "sport"},
	{Name: "Программирование", ImageUrl: "https://images.unsplash.com/photo-1605379399642-870262d3d051?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=3893&q=80", Tag: "programming"},
}

var coursePreviews = []CoursePreview{
	{Id: 1, ImageUrl: "https://www.apple.com/v/apple-fitness-plus/l/images/meta/apple-fitness-plus__eafl9rq9woom_og.png", Name: "Пилатес от Веты", Tag: "sport"},
	{Id: 2, ImageUrl: "https://avatars.yandex.net/get-music-content/118603/e9de54a9.p.4465783/m1000x1000", Name: "Mr. Miyagi", Tag: "sport"},
	{Id: 3, ImageUrl: "https://logowik.com/content/uploads/images/flutter5786.jpg", Name: "Flutter", Tag: "programming"},
}

var coursePages = []CoursePage{
	{Id: 1, ImageUrl: "https://media1.popsugar-assets.com/files/thumbor/v8KPX6IRPz1wie9wQvhB4iYTdcw/fit-in/2048xorig/filters:format_auto-!!-:strip_icc-!!-/2018/07/30/638/n/1922564/8d1b02c5595eaf1f_GettyImages-1007595384/i/Kim-Kardashian-Blue-PVC-Heels-From-Yeezy.jpg", Name: "Пилатес от Веты", Lessons: []Lesson{
		{Name: "Первый день", VideoUrl: "http://0.0.0.0:8080/videos/0"},
	}},
	{Id: 2, ImageUrl: "https://news.store.rambler.ru/img/dc3b4493acdf7bb2208582027eb15ebd?img-format=auto&img-1-resize=height:355,fit:max&img-2-filter=sharpen", Name: "Mr. Miyagi", Lessons: []Lesson{
		{Name: "Minor", VideoUrl: "http://0.0.0.0:8080/videos/1"},
	}},
}

var videos = [2]string{
	"The Kardashians _ Season 3 Returns May 25 _ Hulu.mp4",
	"Miyagi & Andy Panda - Minor (Mood Video).mp4",
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

func getVideoById(context *gin.Context) {
	var num, err = strconv.Atoi(context.Param("num"))
	if err == nil {
		context.File(videos[num])
	} else {
		context.File(videos[0])
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
	router.GET("/videos/:num", getVideoById)
	router.Run() // listen and serve on 0.0.0.0:8080
}
