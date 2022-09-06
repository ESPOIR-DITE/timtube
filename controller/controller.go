package controller

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	//"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"html/template"
	"net/http"
	"timtube/config"
	user3 "timtube/controller/user"
	"timtube/controller/util"
	"timtube/domain"
	channel_video "timtube/io/channel-video"
	user4 "timtube/io/channel/channel"
	user2 "timtube/io/video/video"
	user "timtube/io/video/video-data"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	//mux.Use(middleware.RequestID)
	//mux.Use(middleware.RealIP)
	//mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	//mux.Handle("/", homeHandler(env))
	mux.Handle("/", homeHandler(env))
	//mux.Handle("/swagger/*any", swaggerHandler(env))
	mux.Handle("/channel/{channel}", channelVideosHandler(env))
	mux.Handle("/play/{id}", homePlayHandler(env))
	mux.Mount("/user", user3.Home(env))
	mux.Handle("/out", outHandler(env))

	//// documentation for developers
	//opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	//sh := middleware.SwaggerUI(opts, nil)
	//mux.Handle("/docs", sh)

	// documentation for share
	// opts1 := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	// sh1 := middleware.Redoc(opts1, nil)
	// r.Handle("/docs", sh1)

	fileServer := http.FileServer(http.Dir("./view/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux
}

func swaggerHandler(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
func getPicturesOutOfVideo(videoId string) (domain.VideoData, error) {
	videoData, err := user.ReadVideoData(videoId)
	if err != nil {
		fmt.Println(err, " error reading video")
		return domain.VideoData{}, err
	}
	sEnc := base64.StdEncoding.EncodeToString(videoData.Picture)
	videoObject := domain.VideoData{videoData.Id, []byte{}, []byte{}, sEnc, videoData.FileType}
	return videoObject, nil
}

// Get videos from a channel godoc
// @Summary Retrieves Videos based on given ID
// @Produce json
// @Param channel path string true "Channel ID"
// @Success 200 {object} domain.Video
// @Router /users/{id} [get]
func channelVideosHandler(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "channel")
		var videoDatas []domain.VideoVideoData
		if id != "" {
			channelVideos, err := channel_video.ReadChannelVideoByChannelId(id)
			if err != nil {
				fmt.Println(err, "error reading channelVideo")
			}
			for _, channelVideo := range channelVideos {
				video, err := user2.ReadVideo(channelVideo.VideoId)
				videoData, err := getPicturesOutOfVideo(channelVideo.VideoId)
				if err != nil {
					fmt.Println(err, " error reading video")
				} else {
					videoDatas = append(videoDatas, domain.VideoVideoData{video, videoData})
				}
			}
		}
		result, err := json.Marshal(videoDatas)
		if err != nil {
			fmt.Println("couldn't marshal")
			render.Render(w, r, util.ErrInvalidRequest(errors.New("error marshalling")))
			return
		}
		_, err = w.Write([]byte(result))
		if err != nil {
			return
		}
	}
}

func outHandler(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		env.Session.Clear(r.Context())
		http.Redirect(w, r, "/", 301)
		return
	}
}

type ChannelCollection struct {
	Channels []domain.Channel
}

func getChannels() []ChannelCollection {
	var channelList []ChannelCollection
	channels, err := user4.ReadChannels()
	if err != nil {
		fmt.Println(err, " error reading channels")
		return channelList
	}
	var singleChannelList []domain.Channel
	var myIndex = 1

	for index, channel := range channels {

		if len(channel.Image) > 0 {
			sEnc := base64.StdEncoding.EncodeToString(channel.Image)
			channelObject := domain.Channel{channel.Id, channel.Name, channel.ChannelTypeId, channel.UserId, channel.Region, channel.Date, []byte{}, sEnc, channel.Description}
			singleChannelList = append(singleChannelList, channelObject)
			if index > 0 && myIndex%4 == 0 {
				channelList = append(channelList, ChannelCollection{singleChannelList})
				singleChannelList = []domain.Channel{}
			} else if len(channels) < 4 {
				channelList = append(channelList, ChannelCollection{singleChannelList})
				singleChannelList = []domain.Channel{}
			} else if myIndex%4 != 0 && myIndex > 1 {
				channelList = append(channelList, ChannelCollection{singleChannelList})
				singleChannelList = []domain.Channel{}
			}
			myIndex += 1
		}
	}
	return channelList
}
func getChannels2() []ChannelCollection {
	var channelList []ChannelCollection
	channels, err := user4.ReadChannels()
	if err != nil {
		fmt.Println(err, " error reading channels")
		return channelList
	}
	var singleChannelList []domain.Channel
	var myIndex = 1
	listIndex := len(channels)
	fmt.Println(listIndex)
	for index, channel := range channels {
		if len(channel.Image) > 0 {
			sEnc := base64.StdEncoding.EncodeToString(channel.Image)
			channelObject := domain.Channel{channel.Id, channel.Name, channel.ChannelTypeId, channel.UserId, channel.Region, channel.Date, []byte{}, sEnc, channel.Description}
			singleChannelList = append(singleChannelList, channelObject)
			if index > 0 && myIndex%4 == 0 {
				channelList = append(channelList, ChannelCollection{singleChannelList})
				singleChannelList = []domain.Channel{}
			} else if len(channels) < 4 {
				channelList = append(channelList, ChannelCollection{singleChannelList})
				singleChannelList = []domain.Channel{}
			} else if index+1 == len(channels) {
				channelList = append(channelList, ChannelCollection{singleChannelList})
				singleChannelList = []domain.Channel{}
			}
			myIndex += 1
		}
	}
	return channelList
}
func homePlayHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, name, surname, role := util.GetPermenentSession(app, r)
		id := chi.URLParam(r, "id")
		var videoPresentation []domain.VideoVideoData
		var message string
		var VideoDataRaw domain.Video
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
			VideoDataRaw, err = user2.ReadVideo(id)
			//VideoDataRaw, err = user.ReadVideoRawData(id)
			if err != nil {
				fmt.Println(err, "error reading Video")
			}
		}
		if role != "" {
			role, err = util.GetRoleName(role)
			if err != nil {
				fmt.Println("No Role")
			}
		}

		type PageData struct {
			Public   []domain.VideoVideoData
			Message  string
			Video    domain.Video
			Name     string
			Surname  string
			UserRole string
			Email    string
			Channels []ChannelCollection
		}
		date := PageData{videoPresentation, message, VideoDataRaw, name, surname, role, email, getChannels()}
		files := []string{
			app.Path + "index-updated.html",
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
		email, name, surname, role := util.GetPermenentSession(app, r)
		var videoPresentation []domain.VideoVideoData
		var message string
		var VideoDataRaw domain.Video
		publicVideos, err := user.ReadAllPublicVideoData()
		if err != nil {
			message = "Error has occurred"
			fmt.Println(err, "error reading public videos")
		}
		for _, publicVideo := range publicVideos {
			sEnc := base64.StdEncoding.EncodeToString(publicVideo.VideoData.Picture)
			videoPresentation = append(videoPresentation, domain.VideoVideoData{publicVideo.Video, domain.VideoData{publicVideo.VideoData.Id, []byte{}, []byte{}, sEnc, publicVideo.VideoData.Id}})
		}
		if role != "" {
			role, err = util.GetRoleName(role)
			if err != nil {
				fmt.Println("No Role")
			}
		}

		//for _, channels1 := range getChannels() {
		//	for _, channels := range channels1.Channels {
		//		fmt.Println(channels)
		//	}
		//}
		fmt.Println()
		type PageData struct {
			Public   []domain.VideoVideoData
			Message  string
			Video    domain.Video
			Name     string
			Surname  string
			UserRole string
			Email    string
			Channels []ChannelCollection
		}

		date := PageData{videoPresentation, message, VideoDataRaw, name, surname, role, email, getChannels2()}

		files := []string{
			app.Path + "index-updated.html",
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
