// Bitfan API.
//
// the purpose of this api....
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS
//
//
// Host: 127.0.0.1:5123
// BasePath: /api/v1
// Version: 0.0.1
// License: Apache 2.0 http://www.apache.org/licenses/LICENSE-2.0.html
// Contact: Valere JEANTET<valere.jeantet@gmail.com> http://vjeantet.fr
// Schemes: http
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
//
// swagger:meta
package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vjeantet/bitfan/core"
	"github.com/vjeantet/bitfan/lib"
)

var plugins map[string]map[string]core.ProcessorFactory

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func ServeREST(hostport string, plugs map[string]map[string]core.ProcessorFactory) {
	plugins = plugs

	r := gin.New()
	r.Use(gin.Recovery(), cors())
	v1 := r.Group("api/v1")
	{

		// swagger:operation GET /pipelines pipeline listPipelines
		//
		// Lists pipelines.
		//
		// This will show all running pipelines.
		//
		// ---
		//
		// produces:
		// - application/json
		//
		//
		// responses:
		//   200:
		//     description: pipelines response
		//     schema:
		//       type: array
		//       items:
		//         "$ref": "#/definitions/Pipeline"
		//   default:
		//     description: unexpected error
		//     schema:
		//       "$ref": "#/definitions/Error"
		v1.GET("/pipelines", getPipelines)

		// swagger:operation DELETE /pipelines/{id} pipeline stopPipeline
		//
		// Stop a running pipeline.
		//
		// This will stop pipeline by ID.
		//
		// ---
		//
		//
		// produces:
		// - application/json
		//
		// parameters:
		//   - name: "id"
		//     in: "path"
		//     description: "Pipeline ID"
		//     required: true
		//     type: integer
		//
		//
		// responses:
		//   200:
		//     description: pipelines response
		//   default:
		//     description: unexpected error
		//     schema:
		//       "$ref": "#/definitions/Error"
		v1.DELETE("/pipelines/:id", deletePipeline)

		// swagger:operation POST /pipelines pipeline addPipeline
		//
		// Start a pipeline.
		//
		// This will start pipeline.
		//
		// ---
		// consumes:
		// - application/json
		//
		// produces:
		// - application/json
		//
		// parameters:
		// - in: "body"
		//   name: "body"
		//   description: "Pipeline object that needs to be started"
		//   required: true
		//   schema:
		//     $ref: "#/definitions/Pipeline"
		//
		// responses:
		//   200:
		//     description: pipeline response
		//     schema:
		//       "$ref": "#/definitions/Pipeline"
		//   default:
		//     description: unexpected error
		//     schema:
		//       "$ref": "#/definitions/Error"
		v1.POST("/pipelines", addPipeline)

		// swagger:operation GET /pipelines/{id} pipeline getPipeline
		//
		// Get a pipeline.
		//
		// This will show a running pipeline.
		//
		// ---
		//
		// produces:
		// - application/json
		//
		//
		// parameters:
		//   - name: "id"
		//     in: "path"
		//     description: "Pipeline ID"
		//     required: true
		//     type: integer
		//
		// responses:
		//   200:
		//     description: pipeline response
		//     schema:
		//       "$ref": "#/definitions/Pipeline"
		//   default:
		//     description: unexpected error
		//     schema:
		//       "$ref": "#/definitions/Error"
		v1.GET("/pipelines/:id", getPipeline)

		// swagger:operation GET /docs doc listDocs
		//
		// Lists plugins.
		//
		// This will show all avaialable plugins.
		//
		// ---
		// produces:
		// - application/json
		//
		// responses:
		//   200:
		//     description: processor doc response
		//     schema:
		//       type: array
		//       items:
		//         "$ref": "#/definitions/processorDoc"
		//   default:
		//     description: unexpected error
		//     schema:
		//       "$ref": "#/definitions/Error"
		v1.GET("/docs", getDocs)

		v1.GET("/docs/inputs", getDocsInputs)
		v1.GET("/docs/inputs/:name", getDocsInputsByName)
		v1.GET("/docs/filters", getDocsFilters)
		v1.GET("/docs/filters/:name", getDocsFiltersByName)
		v1.GET("/docs/outputs", getDocsOutputs)
		v1.GET("/docs/outputs/:name", getDocsOutputsByName)
	}
	if hostport == "" {
		hostport = "127.0.0.1:8080"
	}
	go r.Run(hostport)
}

func getPipelines(c *gin.Context) {

	var pipelines []Pipeline
	var err error

	pipelines = []Pipeline{}
	ppls := core.Pipelines()
	for _, p := range ppls {
		pipelines = append(pipelines, Pipeline{
			ID:                 p.ID,
			Label:              p.Label,
			ConfigLocation:     p.ConfigLocation,
			ConfigHostLocation: p.ConfigHostLocation,
		})
	}

	if err == nil {
		c.JSON(200, pipelines)
	} else {
		c.JSON(404, gin.H{"error": "no pipelines(s) running"})
	}
	// curl -i http://localhost:8080/api/v1/pipelines
}

func getPipeline(c *gin.Context) {
	var pipeline Pipeline
	var err error
	var id int

	id, err = strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	pipeline = Pipeline{}
	ppls := core.Pipelines()
	for _, p := range ppls {
		if p.ID == id {
			pipeline = Pipeline{
				ID:                 p.ID,
				Label:              p.Label,
				ConfigLocation:     p.ConfigLocation,
				ConfigHostLocation: p.ConfigHostLocation,
			}
			c.JSON(200, pipeline)
			return
		}
	}

	c.JSON(404, gin.H{"error": "no pipelines(s) running"})

}

func addPipeline(c *gin.Context) {

	// ID, err := core.StartPipeline(&starter.Pipeline, starter.Agents)

	// b, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pp.Println("c-->", string(b))

	//Bind request data
	var pipeline Pipeline
	err := c.BindJSON(&pipeline)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var loc *lib.Location
	cwd, _ := os.Getwd()

	if pipeline.Content != "" {
		loc, err = lib.NewLocationContent(pipeline.Content, cwd)
	} else {
		loc, err = lib.NewLocation(pipeline.ConfigLocation, cwd)
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ppl := loc.ConfigPipeline()
	if pipeline.Label != "" {
		ppl.Name = pipeline.Label
	}

	agt, err := loc.ConfigAgents()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ID, err := core.StartPipeline(&ppl, agt)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("/api/v1/pipelines/%d", ID))
	return

	// c.JSON(200, gin.H{"statut": "started"})

	// curl -i -X DELETE http://localhost:8080/api/v1/pipelines/1
}

func deletePipeline(c *gin.Context) {

	var err error
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("malformed id : %s", c.Params.ByName("id"))})
		return
	}

	err = core.StopPipeline(id)
	if err == nil {
		c.JSON(200, gin.H{"id": c.Params.ByName("id"), "status": "deleted"})
	} else {
		c.JSON(404, gin.H{"error": err.Error()})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/pipelines/1
}

func getDocs(c *gin.Context) {

	data := docsByKind("input")
	data = append(data, docsByKind("filter")...)
	data = append(data, docsByKind("output")...)

	c.JSON(200, data)
	// curl -i http://localhost:8080/api/v1/docs
}

func getDocsInputs(c *gin.Context) {
	c.JSON(200, docsByKind("input"))
	// curl -i http://localhost:8080/api/v1/docs/inputs
}

func getDocsFilters(c *gin.Context) {
	c.JSON(200, docsByKind("filter"))
	// curl -i http://localhost:8080/api/v1/docs/filters
}

func getDocsOutputs(c *gin.Context) {
	c.JSON(200, docsByKind("output"))
	// curl -i http://localhost:8080/api/v1/docs/outputs
}

func getDocsInputsByName(c *gin.Context) {
	data, err := docsByKindByName("input", c.Params.ByName("name"))
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, data)
	}
}

func getDocsFiltersByName(c *gin.Context) {
	data, err := docsByKindByName("filter", c.Params.ByName("name"))
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, data)
	}
}

func getDocsOutputsByName(c *gin.Context) {
	data, err := docsByKindByName("output", c.Params.ByName("name"))
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, data)
	}
}
