package test

import (
	"fmt"
	"strings"

	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

// Paths for packages used by code generated in this file,
// relative to the import_prefix of the generator.Generator.
const (
	contextPkgPath = "context"
)

func init() {
	generator.RegisterPlugin(new(test))
}

// test is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for test support.
type test struct {
	gen *generator.Generator
}

// Name returns the name of this plugin, "test".
func (g *test) Name() string {
	return "test"
}

// The names for packages imported in the generated code.
// They may vary from the final path component of the import path
// if the name is used by other packages.
var (
	contextPkg string
)

// Init initializes the plugin.
func (g *test) Init(gen *generator.Generator) {
	g.gen = gen
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *test) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *test) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *test) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *test) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}

	contextPkg = string(g.gen.AddImport(contextPkgPath))

	for i, service := range file.FileDescriptorProto.Service {
		g.generateService(file, service, i)
	}
}

// GenerateImports generates the import declaration for this file.
func (g *test) GenerateImports(file *generator.FileDescriptor) {
}

func unexport(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// generateService generates all the code for the named service.
func (g *test) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	path := fmt.Sprintf("6,%d", index) // 6 means service.

	origServName := service.GetName()
	servName := generator.CamelCase(origServName)
	servAlias := servName + "Service"

	// strip suffix
	if strings.HasSuffix(servAlias, "ServiceService") {
		servAlias = strings.TrimSuffix(servAlias, "Service")
	}

	g.P()
	g.P("// For example")
	g.P()

	tmp := "aa"
	for _, v := range file.GetExtension() {
		if tmp != "" {
			tmp += "\n"
		}
		tmp += fmt.Sprintf("%v", *v)
	}

	g.P("// ", tmp)

	// Client interface.
	g.P("type ", servAlias, " interface {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
		g.P(g.generateClientSignature(servName, method))
	}
	g.P("}")
}

// generateClientSignature returns the client-side signature for a method.
func (g *test) generateClientSignature(servName string, method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)

	// tmp := "aa"
	// for _, v := range method.GetOptions().GetUninterpretedOption() {
	// 	if tmp != "" {
	// 		tmp += "\n"
	// 	}
	// 	tmp += fmt.Sprintf("%v", *v)
	// }
	return fmt.Sprintf("%s(ctx %s.Context) error", methName, contextPkg)
}

// AddPluginToParams Simplify the protoc call statement by adding 'plugins=test' directly to the command line arguments.
func AddPluginToParams(p string) string {
	params := p
	if strings.Contains(params, "plugins=") {
		params = strings.Replace(params, "plugins=", "plugins=test+", -1)
	} else {
		if len(params) > 0 {
			params += ","
		}
		params += "plugins=test"
	}
	return params
}
