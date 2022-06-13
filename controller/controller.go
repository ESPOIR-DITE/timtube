package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"net/http"
	"timtube/config"
	"timtube/domain"
	user "timtube/io/video/video-data"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	//mux.Handle("/", homeHandler(env))
	mux.Handle("/", homeHandler(env))
	mux.Handle("/play/{id}", homePlayHandler(env))

	fileServer := http.FileServer(http.Dir("./view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
}

func homePlayHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var videoPresentation []domain.VideoVideoData
		var message string
		var VideoDataRaw domain.VideoData
		publicVideos, err := user.ReadAllPublicVideoData()
		if err != nil {
			message = "Error has occurred"
			fmt.Println(err, "error reading public videos")
		}
		for _, publicVideo := range publicVideos {
			sEnc := base64.StdEncoding.EncodeToString(publicVideo.VideoData.Picture)
			videoPresentation = append(videoPresentation, domain.VideoVideoData{publicVideo.Video, domain.VideoData{publicVideo.VideoData.Id, []byte{}, []byte{}, sEnc, publicVideo.VideoData.FileSize}})
		}
		if id != "" {
			VideoDataRaw, err = user.ReadVideoRawData(id)
			if err != nil {
				fmt.Println(err, "error reading Video Data")
			}
		}
		type PageData struct {
			Public    []domain.VideoVideoData
			Message   string
			VideoData domain.VideoData
		}

		date := PageData{videoPresentation, message, VideoDataRaw}

		files := []string{
			app.Path + "index.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, date)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(bookingLink)
		var videoPresentation []domain.VideoVideoData
		var message string
		var VideoDataRaw domain.VideoData
		publicVideos, err := user.ReadAllPublicVideoData()
		if err != nil {
			message = "Error has occurred"
			fmt.Println(err, "error reading public videos")
		}
		for _, publicVideo := range publicVideos {
			sEnc := base64.StdEncoding.EncodeToString(publicVideo.VideoData.Picture)
			videoPresentation = append(videoPresentation, domain.VideoVideoData{publicVideo.Video, domain.VideoData{publicVideo.VideoData.Id, []byte{}, []byte{}, sEnc, publicVideo.VideoData.Id}})
		}
		type PageData struct {
			Public    []domain.VideoVideoData
			Message   string
			VideoData domain.VideoData
		}

		date := PageData{videoPresentation, message, VideoDataRaw}

		files := []string{
			app.Path + "index.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, date)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
