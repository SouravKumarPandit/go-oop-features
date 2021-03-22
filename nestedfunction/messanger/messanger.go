package messanger
//https://www.youtube.com/watch?v=HOlklLaFgfM&ab_channel=DavyWybiral
import "net/http"


func AuthRequired(handler http.HandlerFunc) http.HandlerFunc{

	return func(writer http.ResponseWriter, request *http.Request) {
		//here write logic before calling handler
		//session := writer.Header()//todo check for session expired
		handler.ServeHTTP(writer, request)
	}
}

/*func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")
		_, ok := session.Values["user_id"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler.ServeHTTP(w, r)
	}
}*/