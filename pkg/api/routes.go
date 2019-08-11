package api
//
//import (
//	"fmt"
//	"github.com/julienschmidt/httprouter"
//	"net/http"
//)
//
//func (s *Server) routes() {
//	s.Router.GET("/", s.Index)
//	s.Router.GET("/v1/images", s.GetImages)
//}
//
//// homepage
//func (s *Server) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Fprint(w, "Welcome to dokuweb-api\n")
//}
//
//// GetImages is used to retrieve all the images from dropbox
//func (s *Server) GetImages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Fprintln(w, "Not Implemented!")
//	// TODO: call dropbox api to return all images
//}
//
//// GetImagesByDateRange is used to retrieve only the images on a specified date range
//func (s *Server) GetImagesByDateRange(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Println(w, "Implement me!")
//}
//
//// get one image by id
