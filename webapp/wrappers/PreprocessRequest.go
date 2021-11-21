package wrappers

import(
	"net/http"
	"strings"

	"go-webapp/webapp/sessions"
)


const urlAuthRedirect = "http://127.0.0.1:8080/login"

type Preprocessor struct {
	handler http.Handler
}

func (p *Preprocessor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    uri := r.URL.RequestURI()

    if strings.Contains(uri, "login") || strings.Contains(uri, "register") {
    	p.handler.ServeHTTP(w, r)
    	return
    }

	var token string
    var session *sessions.Session
    cookie, err := r.Cookie("gsession")
    if err != nil {
        token = sessions.SetSession()
        http.SetCookie(w, &http.Cookie{Name: "gsession", Value: token, HttpOnly: true, 
                                       SameSite: 3})
        http.Redirect(w, r, urlAuthRedirect, http.StatusSeeOther)
        return
    }
    
    token = cookie.Value
    session, err = sessions.GetSession(token)
    if err != nil {
        token = sessions.SetSession()
        http.SetCookie(w, &http.Cookie{Name: "gsession", Value: token, HttpOnly: true, 
                                       SameSite: 3})
        http.Redirect(w, r, urlAuthRedirect, http.StatusSeeOther)
        return
        }
    if session.Is_Authenticated != true {
        http.Redirect(w, r, urlAuthRedirect, http.StatusSeeOther)
        return
    }
    
    p.handler.ServeHTTP(w, r)
}


func PreprocessRequest(handlerToWrap http.Handler) *Preprocessor {
	return &Preprocessor{handlerToWrap}
}