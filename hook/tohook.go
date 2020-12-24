package hook

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

//Init ---
func Init(router *gin.RouterGroup, dbs *mgo.Session) {
	repoS := newRepoService(dbs)
	hookS := NewHookService(repoS)
	MakeHTTPHandlers(router, hookS)
}
