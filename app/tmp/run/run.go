// GENERATED CODE - DO NOT EDIT
// This file is the run file for Revel.
// It registers all the controllers and provides details for the Revel server engine to
// properly inject parameters directly into the action endpoints.
package run

import (
	"reflect"
	"github.com/revel/revel"
	controllers "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers0 "github.com/revel/modules/testrunner/app/controllers"
	_ "nlpf/app"
	controllers1 "nlpf/app/controllers"
	tests "nlpf/tests"
	time "time"
	"github.com/revel/revel/testing"
)

var (
	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

// Register and run the application
func Run(port int) {
	Register()
	revel.Run(port)
}

// Register all the controllers
func Register() {
	revel.AppLog.Info("Running revel server")
	
	revel.RegisterController((*controllers.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModuleDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Admin)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Administration",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "begin_date_input", Type: reflect.TypeOf((*time.Time)(nil)) },
					&revel.MethodArg{Name: "end_date_input", Type: reflect.TypeOf((*time.Time)(nil)) },
					&revel.MethodArg{Name: "motifrejet", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "currentoffer", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "isaccepted", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					61: []string{ 
						"tags",
						"rep",
					},
					66: []string{ 
						"tags",
						"rep",
					},
					68: []string{ 
						"tags",
					},
				},
			},
			&revel.MethodType{
				Name: "AcceptOffer",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					73: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "RefuseOffer",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "tag", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					78: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "LogOut",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "message", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					71: []string{ 
						"message",
					},
				},
			},
			&revel.MethodType{
				Name: "User",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "uid", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"uid",
					},
				},
			},
			&revel.MethodType{
				Name: "Auth",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Inscription",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					117: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "SignIn",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "nom", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prenom", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "phone", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "HTTP403",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					141: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Client)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					47: []string{ 
						"tags",
						"total",
					},
				},
			},
			&revel.MethodType{
				Name: "Facture",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					71: []string{ 
						"tags",
						"total",
					},
				},
			},
			&revel.MethodType{
				Name: "Modify",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					91: []string{ 
						"tag",
					},
				},
			},
			&revel.MethodType{
				Name: "ModifyDemande",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "address", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "motif", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "phone", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ProcessDemande",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "address", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "motif", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "phone", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteDemande",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Demande",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					171: []string{ 
						"y",
						"m",
						"d",
					},
				},
			},
			&revel.MethodType{
				Name: "Tag",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					190: []string{ 
						"tag",
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}
}
