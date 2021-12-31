package handle

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	Name     string
	TeamName string
	Info     string
)

func EditPlayers(w http.ResponseWriter, r *http.Request) {
	Name = r.FormValue("name")
	TeamName = r.FormValue("teamname")
	if TeamName == "" {
		TeamName = "None"
	}
	Info = r.FormValue("Information")
	w.Header().Set("Content-Type", "text/html")
	r.ParseForm()
	if r.Method != "POST" {
		w.Write([]byte(html))
	} else {
		uploadFile, handle, err := r.FormFile("image")
		errorHandle(err, w)
		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext != ".jpg" && ext != ".png" {
			errorHandle(errors.New("只支持jpg/png图片上传"), w)
			return
		}
		os.Mkdir("./uploaded", 0777)
		saveFile, err := os.OpenFile("./uploaded"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		errorHandle(err, w)
		io.Copy(saveFile, uploadFile)

		defer uploadFile.Close()
		defer saveFile.Close()
		w.Write([]byte("查看上传图片:<a target='_blank' href='/uploaded/" + handle.Filename + "'>" + handle.Filename + "</a>"))
	}

}
func showPicHandle(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("." + req.URL.Path)
	errorHandle(err, w)

	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorHandle(err, w)
	w.Write(buff)
}

func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

const html = `<html>
    <head></head>
    <body>
        <form method="post" enctype="multipart/form-data">
            <input type="file" name="image" />
            <input type="submit" />
        </form>
    </body>
</html>`
