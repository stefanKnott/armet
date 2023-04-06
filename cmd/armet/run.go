package armet

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/stefanKnott/armet/pkg/handlers"
	h "github.com/stefanKnott/armet/pkg/helm"
)

var kubernetes bool
var helm bool
var pathToConfigs string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run armet to check cluster resources",
	Run: func(cmd *cobra.Command, args []string) {
		if kubernetes {
			fmt.Println("polling for cluster resources")
		}

		if helm {
			fmt.Println("polling for helm resources")
			h.BeginLoop(pathToConfigs)
		}

		r := mux.NewRouter()

		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowCredentials: true,
		})
		handlers.RegisterHandlerRoutes(r)
		handler := c.Handler(r)
		http.ListenAndServe(":8080", handler)
	},
}

// access control and  CORS middleware
func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func init() {
	runCmd.Flags().BoolVarP(&kubernetes, "kubernetes", "k", false, "poll cluster for kubernetes objects")
	runCmd.Flags().BoolVarP(&helm, "helm", "z", false, "poll cluster for helm objects")
	runCmd.Flags().StringVarP(&pathToConfigs, "configs", "c", "", "absolute path to kubeconfig files, directory or file")

	rootCmd.AddCommand(runCmd)
}
