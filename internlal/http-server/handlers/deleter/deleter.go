package deleter

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"url-shorter/internlal/lib/api/responce"
	"url-shorter/internlal/lib/logger/sl"
)

type Request struct {
	//URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	responce.Responce
	Alias string `json:"alias,omitempty"`
}

type URLDeleter interface {
	DeletURL(urlToDelet string) error
}

func New(log *slog.Logger, urlDelete URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handle.url.deleter.New"

		log = log.With(slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, responce.Error("failed request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		//if err := validator.New().Struct(req); err != nil {
		//
		//	validateErr := err.(validator.ValidationErrors)
		//
		//	log.Error("invalid request", sl.Err(err))
		//
		//	render.JSON(w, r, responce.ValidatorError(validateErr))
		//
		//	return
		//}

		alias := req.Alias

		err = urlDelete.DeletURL(alias)
		//if errors.Is(err, storage.ErrURLExist) {
		//	log.Info("url already exists", slog.String("url", req.URL))
		//	render.JSON(w, r, responce.Error("url already exists"))
		//	return
		//}
		//TODO: handle generation alias

		if err != nil {
			log.Error("failed to deleter url", sl.Err(err))
			render.JSON(w, r, responce.Error("failed to deleter url"))
			return
		}

		//log.Info("url added", slog.Int64("id", id))

		//TODO: responceOK()
		render.JSON(w, r, Response{
			Responce: responce.OK(),
			Alias:    alias,
		})
	}
}
