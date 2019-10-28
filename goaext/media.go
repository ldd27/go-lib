package goaext

//
//import (
//	"fmt"
//
//	. "github.com/goadesign/goaext/design"
//	. "github.com/goadesign/goaext/design/apidsl"
//)
//
//type ConcatOption struct {
//	Name  string
//	Email string
//	URL   string
//}
//
//type TraitOption struct {
//	Name    string
//	DSLFunc func()
//}
//
//type Option struct {
//	Name        string
//	Title       string
//	Description string
//	Concat      *ConcatOption
//	Host        string
//	Scheme      []string
//	BasePath    string
//	Version     string
//	Consumes    []interface{}
//	Produces    []interface{}
//}
//
//const (
//	ApplicationJSON               = "application/json"
//	ApplicationXWWWFormUrlEncoded = "application/x-www-form-urlencoded"
//)
//
//var defaultOption = Option{
//	Name:     "goaext",
//	Scheme:   []string{"https"},
//	BasePath: "v1",
//	Version:  "1.0",
//	Consumes: []interface{}{ApplicationJSON},
//	Produces: []interface{}{ApplicationJSON},
//}
//
//func NewAPI(opts ...func(*Option)) *APIDefinition {
//	conf := defaultOption
//	for _, o := range opts {
//		o(&conf)
//	}
//
//	return API(conf.Name, func() {
//		Title(conf.Title)
//		Description(conf.Description)
//
//		// 联系人
//		if conf.Concat != nil {
//			Contact(func() {
//				Name(conf.Concat.Name)
//				Email(conf.Concat.Email)
//				URL(conf.Concat.URL)
//			})
//		}
//
//		Host(conf.Host)
//		Scheme(conf.Scheme...)
//		BasePath(conf.BasePath)
//		Version(conf.Version)
//		Consumes(conf.Consumes)
//		Produces(conf.Produces)
//	})
//}
//
//func DefinedMedia(name string, dsl func()) *MediaTypeDefinition {
//	return MediaType(fmt.Sprintf("application/vnd.%s+json", name), func() {
//		UseTrait("JsonAPI")
//		dsl()
//	})
//}
//
//func DefinedType(name string, dsl func()) *UserTypeDefinition {
//	return Type(fmt.Sprintf("application/vnd.%s+json", name), func() {
//		UseTrait("JsonAPI")
//		dsl()
//	})
//}
