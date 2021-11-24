package version

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"net/http"
)

// VersionEnv ..
type VersionEnv struct {
	ReleaseVersion    	string `envconfig:"VERSION_RELEASE" default:"V.0.1"`
	GitBranchName    	string `envconfig:"VERSION_GIT_BRANCH_NAME" default:"sfa_3.0.0"`
	GitCommit    		string `envconfig:"VERSION_GIT_COMMIT" default:"51ee8631f320b63bdad1e635c9eb882441cd4086"`
	DateTime  			string `envconfig:"VERSION_DATE_TIME" default:"13 Oct 2021"`
	TeamCreated    		string `envconfig:"VERSION_TEAM_CREATED" default:"Wayang Squad - Srikandi Team"`
	NameServices   		string `envconfig:"VERSION_NAME_SERVICES" default:"ottosfa-api-apk"`
	VersionType    		string `envconfig:"VERSION_TYPE" default:"SIT"`
}

var (
	versionEnv VersionEnv
)

// init ..
func init() {
	err := envconfig.Process("Version", &versionEnv)
	if err != nil {
		fmt.Println("Failed to get Version env:", err)
	}
}

// Version ..
func Version(ctx *gin.Context) {
	fmt.Println(">>> Version <<<")

	ctx.String(http.StatusOK,
		"Version Release : [%s]\n"+
		"Branch Name : [%s]\n"+
		"GitCommit : [%s]\n"+
		"DateTime : [%s] \n"+
		"Team By: [%s] \n"+
		"Services :[%s] \n"+
		"VersionType : [%s]\n",
		versionEnv.ReleaseVersion,
		versionEnv.GitBranchName,
		versionEnv.GitCommit,
		versionEnv.DateTime,
		versionEnv.TeamCreated,
		versionEnv.NameServices,
		versionEnv.VersionType)
}

