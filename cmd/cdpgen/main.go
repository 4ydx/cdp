// The cdpgen command generates the package cdp from the provided protocol definitions.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/4ydx/cdp/cmd/cdpgen/proto"
)

// Global constants.
const (
	OptionalPropPrefix = ""
	realEnum           = false
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func mkdir(name string) error {
	err := os.Mkdir(name, 0755)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func main() {
	var (
		destPkg          string
		browserProtoJSON string
		jsProtoFileJSON  string
	)
	flag.StringVar(&destPkg, "dest-pkg", "", "Destination for generated cdp package (inside $GOPATH)")
	flag.StringVar(&browserProtoJSON, "browser-proto", "./protodef/browser_protocol.json", "Path to browser protocol")
	flag.StringVar(&jsProtoFileJSON, "js-proto", "./protodef/js_protocol.json", "Path to JS protocol")
	flag.Parse()

	if destPkg == "" {
		fmt.Fprintln(os.Stderr, "error: dest-pkg must be set")
		os.Exit(1)
	}

	var protocol, jsProtocol proto.Protocol
	protocolData, err := ioutil.ReadFile(browserProtoJSON)
	panicErr(err)

	err = json.Unmarshal(protocolData, &protocol)
	panicErr(err)

	jsProtocolData, err := ioutil.ReadFile(jsProtoFileJSON)
	panicErr(err)

	err = json.Unmarshal(jsProtocolData, &jsProtocol)
	panicErr(err)

	protocol.Domains = append(protocol.Domains, jsProtocol.Domains...)
	sort.Slice(protocol.Domains, func(i, j int) bool {
		return protocol.Domains[i].Domain < protocol.Domains[j].Domain
	})

	protoDest := path.Join(destPkg, "protocol")
	for i, d := range protocol.Domains {
		for ii, t := range d.Types {
			// Reference the FrameId in the Frame type.
			if d.Domain == "Page" && t.IDName == "Frame" {
				for iii, p := range t.Properties {
					if p.NameName == "id" || p.NameName == "parentId" {
						p.Ref = "FrameId"
						t.Properties[iii] = p
					}
				}
			}
			d.Types[ii] = t
		}
		protocol.Domains[i] = d
	}

	var g Generator

	// Package cdp/protocol.
	g.pkg = "shared"
	g.dir = protoDest
	err = mkdir(g.path())
	panicErr(err)

	// SharedTypes are types from one domain being used in another domain.
	// They must be moved to a shared file.
	sharedTypes := map[string]bool{
		"FrameID":        true,
		"RemoteObjectID": true,
		"ResourceType":   true,
	}
	g.PackageHeader("")
	for _, d := range protocol.Domains {
		if len(d.Types) > 0 {
			for _, t := range d.Types {
				g.DomainSharedType(d, t, sharedTypes)
			}
		}
	}
	g.writeFile("shared_types.go")
	fmt.Printf("%+v\n", sharedTypes)

	imports := []string{
		"github.com/4ydx/cdp/protocol",
	}
	g.imports = imports

	// Generate the protocol definitions.
	for _, d := range protocol.Domains {
		dLower := strings.ToLower(d.Domain)
		g.dir = path.Join(protoDest, dLower)
		g.pkg = dLower
		err := mkdir(g.path())
		panicErr(err)

		if len(d.Types) > 0 {
			g.PackageHeader("")
			for _, t := range d.Types {
				g.DomainType(d, t, sharedTypes)
			}
			g.writeFile("types.go")
		}

		if len(d.Commands) > 0 {
			g.PackageHeader("")
			for _, c := range d.Commands {
				if c.Redirect != "" {
					continue
				}
				g.DomainCmd(d, c, sharedTypes)
			}
			g.writeFile("command.go")
		}

		if len(d.Events) > 0 {
			g.PackageHeader("")
			for _, e := range d.Events {
				g.DomainEvent(d, e, sharedTypes)
			}
			g.writeFile("event.go")
		}
	}
	g.dir = destPkg

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	goimports := exec.CommandContext(ctx, "goimports", "-w", g.path())
	out, err := goimports.CombinedOutput()
	if err != nil {
		log.Printf("goimports failed: %s", out)
		log.Println(err)
		os.Exit(1)
	}

	goinstall := exec.CommandContext(ctx, "go", "install", path.Join(destPkg, "..."))
	out, err = goinstall.CombinedOutput()
	if err != nil {
		log.Printf("install failed: %s", out)
		log.Println(err)
		os.Exit(1)
	}
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	dir        string
	pkg        string
	imports    []string
	buf        bytes.Buffer // Accumulated output.
	testbuf    bytes.Buffer // Accumulated test output.
	hasContent bool
	hasHeader  bool

	// For managing circular types for go 1.8/1.9.
	writingType bool
	typebuf     bytes.Buffer // Accumulated type output.
}

func (g *Generator) path() string {
	return path.Join(os.Getenv("GOPATH"), "src", g.dir)
}

// Printf prints to the Generator buffer.
func (g *Generator) Printf(format string, args ...interface{}) {
	buf := &g.buf
	if g.writingType {
		buf = &g.typebuf
	}
	fmt.Fprintf(buf, format, args...)
}

func (g *Generator) beginType() {
	g.writingType = true
}

func (g *Generator) commitType() {
	g.writingType = false
	_, err := g.typebuf.WriteTo(&g.buf)
	panicErr(err)
	g.typebuf.Reset()
}

func (g *Generator) writeFile(f string) {
	fp := path.Join(g.path(), f)
	if !g.hasContent {
		log.Printf("No content, skipping %s...", fp)
		g.clear()
		return
	}
	if g.buf.Len() == 0 {
		log.Printf("Empty buffer, skipping %s...", fp)
		return
	}
	log.Printf("Writing %s...", fp)
	err := ioutil.WriteFile(fp, g.format(), 0644)
	panicErr(err)

	if g.testbuf.Len() > 0 {
		g.buf.Reset()
		g.Printf("package %s\n\n", g.pkg)
		_, err = g.testbuf.WriteTo(&g.buf)
		panicErr(err)
		fptest := strings.Replace(fp, ".go", "_test.go", 1)
		log.Printf("Writing %s...", fptest)

		err = ioutil.WriteFile(fptest, g.format(), 0644)
		panicErr(err)
	}
	g.clear()
}

func (g *Generator) clear() {
	g.hasContent = false
	g.hasHeader = false
	g.writingType = false
	g.buf.Reset()
	g.testbuf.Reset()
	g.typebuf.Reset()
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// PackageHeader writes the header for a package.
func (g *Generator) PackageHeader(comment string) {
	if g.hasHeader {
		return
	}
	g.hasHeader = true
	g.Printf(`// Code generated by cdpgen. DO NOT EDIT.

	%s
	package %s

	import (
		"context"
		"encoding/json"
		"fmt"

		%s
	)
	`, comment, g.pkg, quotedImports(g.imports))
}

func (g *Generator) PrintTypeHelper(d proto.Domain, t proto.AnyType, sharedTypes map[string]bool) {
	g.hasContent = true
	g.beginType()

	var comment string
	desc := t.Desc(0, len(t.Name(d))+1)
	if t.Deprecated {
		desc = "\n//\n// Deprecated: " + t.Desc(0, 12)
	}
	if t.Experimental {
		desc = desc + "\n//\n// Note: This type is experimental."
	}

	g.Printf(`
	// %[1]s %[2]s
	%[3]stype %[1]s `, t.Name(d), desc, comment)

	typ := t.GoType(g.pkg, d)
	switch typ {
	case "struct":
		g.domainTypeStruct(d, t, sharedTypes)
	case "enum":
		g.domainTypeEnum(d, t)
	case "time.Time":
		g.domainTypeTime(d, t)
	case "RawMessage":
		g.domainTypeRawMessage(d, t)
	default:
		g.Printf(typ)
	}
	g.Printf("\n\n")
	g.commitType()
}

// DomainSharedType creates the type definition for all non-struct, enu, time, rawmessage types.
func (g *Generator) DomainSharedType(d proto.Domain, t proto.AnyType, sharedTypes map[string]bool) {
	if _, ok := sharedTypes[t.Name(d)]; !ok {
		return
	}
	g.PrintTypeHelper(d, t, sharedTypes)
}

// DomainType creates the type definition.
func (g *Generator) DomainType(d proto.Domain, t proto.AnyType, sharedTypes map[string]bool) {
	if _, ok := sharedTypes[t.Name(d)]; ok {
		return
	}
	g.PrintTypeHelper(d, t, sharedTypes)
}

func (g *Generator) printStructProperties(d proto.Domain, name string, props []proto.AnyType, ptrOptional, renameOptional bool, sharedTypes map[string]bool) {
	for _, prop := range props {
		jsontag := prop.NameName
		ptype := prop.GoType(g.pkg, d)

		// Check for types that were pushed into the shared simple_types.go file.
		suffix := prop.GetSuffix(g.pkg, d)
		if suffix != "" {
			if _, ok := sharedTypes[suffix]; ok {
				ptype = "shared." + suffix
			}
		}
		if _, ok := sharedTypes[ptype]; ok {
			ptype = "shared." + ptype
		}

		// Make all optional properties into pointers, unless they are slices.
		if prop.Optional {
			jsontag += ",omitempty"
		}

		// Avoid recursive type definitions.
		if ptype == name {
			ptype = "*" + ptype
		}

		exportedName := prop.ExportedName(d)
		if renameOptional && prop.Optional {
			exportedName = OptionalPropPrefix + exportedName
		}

		var preDesc, postDesc string

		desc := prop.Desc(8, len(exportedName)+1)
		var deprecated, localEnum, experimental string
		if prop.Deprecated {
			desc = prop.Desc(8, 12)
			if desc == "" {
				desc = "This property should not be used."
			}
			deprecated = "//\n// Deprecated: " + desc + "\n"
			desc = "is deprecated."
		}
		if prop.IsLocalEnum() {
			var enums []string
			for _, e := range prop.Enum {
				enums = append(enums, fmt.Sprintf("%q", e))
			}
			localEnum = "//\n// Values: " + strings.Join(enums, ", ") + ".\n"
		}
		if prop.Experimental {
			experimental = "//\n// Note: This property is experimental.\n"
		}
		if deprecated != "" || localEnum != "" || experimental != "" {
			preDesc = "// " + exportedName + " " + desc + "\n" + deprecated + localEnum + experimental
		} else {
			if desc == "" {
				desc = "No description."
			}
			postDesc = "// " + enforceSingleLine(desc)
		}

		g.Printf("\t%s%s %s `json:\"%s\"` %s\n", preDesc, exportedName, ptype, jsontag, postDesc)
	}
}

func enforceSingleLine(s string) string {
	return strings.Replace(s, "\n//", "", -1)
}

func (g *Generator) domainTypeStruct(d proto.Domain, t proto.AnyType, sharedTypes map[string]bool) {
	g.Printf("struct{\n")
	g.printStructProperties(d, t.Name(d), t.Properties, true, false, sharedTypes)
	g.Printf("}")
}

func (g *Generator) domainTypeTime(d proto.Domain, t proto.AnyType) {
	g.Printf(`float64

	// String calls (time.Time).String().
	func (t %[1]s) String() string {
		return t.Time().String()
	}

	// Time parses the Unix time with millisecond accuracy.
	func (t %[1]s) Time() time.Time {
		secs := int64(t)
		// The Unix time in t only has ms accuracy.
		ms := int64((float64(t)-float64(secs))*1000000)
		return time.Unix(secs, ms*1000)
	}

	// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
	func (t %[1]s) MarshalJSON() ([]byte, error) {
		if t == 0 {
			return []byte("null"), nil
		}
		f := float64(t)
		return json.Marshal(&f)
	}

	// UnmarshalJSON implements json.Unmarshaler.
	func (t *%[1]s) UnmarshalJSON(data []byte) error {
		*t = 0
		if len(data) == 0 {
			return nil
		}
		var f float64
		if err := json.Unmarshal(data, &f); err != nil {
			return errors.New("%[2]s.%[1]s: " + err.Error())
		}
		*t = %[1]s(f)
		return nil
	}

	var _ json.Marshaler = (*%[1]s)(nil)
	var _ json.Unmarshaler = (*%[1]s)(nil)
	`, t.Name(d), g.pkg)

}

func (g *Generator) domainTypeRawMessage(d proto.Domain, t proto.AnyType) {
	g.Printf(`[]byte

	// MarshalJSON copies behavior of json.RawMessage.
	func (%[3]s %[1]s) MarshalJSON() ([]byte, error) {
		if %[3]s == nil {
			return []byte("null"), nil
		}
		return %[3]s, nil
	}

	// UnmarshalJSON copies behavior of json.RawMessage.
	func (%[3]s *%[1]s) UnmarshalJSON(data []byte) error {
		if %[3]s == nil {
			return errors.New("%[2]s.%[1]s: UnmarshalJSON on nil pointer")
		}
		*%[3]s = append((*%[3]s)[0:0], data...)
		return nil
	}

	var _ json.Marshaler = (*%[1]s)(nil)
	var _ json.Unmarshaler = (*%[1]s)(nil)
	`, t.Name(d), g.pkg, t.Recvr(d))
}

func (g *Generator) domainTypeEnum(d proto.Domain, t proto.AnyType) {
	if t.Type != "string" {
		log.Panicf("unknown enum type: %s", t.Type)
	}

	// Verify assumption about enums never being the empty string, this
	// allows us to consider optional enums as type string instead of
	// *string when encoding to JSON (omitempty).
	for _, e := range t.Enum {
		if e == "" {
			panic("enum " + t.Name(d) + " has unexpected empty enum value")
		}
	}

	name := strings.Title(t.Name(d))
	if realEnum {
		g.Printf("int\n\n")

		if name != "PageResourceType" {
			format := `
			// %s as enums.
			const (
				%sNotSet %s = iota`
			g.Printf(format, name, name, name)
			for _, e := range t.Enum {
				g.Printf("\n\t%s%s", name, e.Name())
			}
			g.Printf(`
			)
			`)
		}
		g.Printf(`
		// Valid returns true if enum is set.
		func (e %[1]s) Valid() bool {
			return e >= 1 && e <= %[2]d
		}

		func (e %[1]s) String() string {
			switch e {
			case 0:
				return "%[1]sNotSet"`, name, len(t.Enum))
		for i, e := range t.Enum {
			g.Printf(`
				case %d:
					return "%s"`, i+1, e)
		}
		g.Printf(`
			}
			return fmt.Sprintf("%[1]s(%%d)", e)
		}

		// MarshalJSON encodes enum into a string or null when not set.
		func (e %[1]s) MarshalJSON() ([]byte, error) {
			if e == 0 {
				return []byte("null"), nil
			}
			if !e.Valid() {
				return nil, errors.New("%[2]s.%[1]s: MarshalJSON on bad enum value: " + e.String())
			}
			return json.Marshal(e.String())
		}

		// UnmarshalJSON decodes a string value into a enum.
		func (e *%[1]s) UnmarshalJSON(data []byte) error {
			switch string(data) {
			case "null":
				*e = 0`, name, g.pkg)
		for i, e := range t.Enum {
			g.Printf(`
				case "\"%s\"":
					*e = %d`, e, i+1)
		}
		g.Printf(`
			default:
				return fmt.Errorf("%s.%s: UnmarshalJSON on bad input: %%s", data)
			}
			return nil
		}`, g.pkg, name)
	} else {
		g.Printf("string\n\n")

		if name != "PageResourceType" {
			g.Printf(`
			// %s as enums.
			const (`, name)
			format := "\n\t%s%s %s = %q"
			g.Printf(format, name, "NotSet", name, "")

			for _, e := range t.Enum {
				g.Printf(format, name, e.Name(), name, e)
			}
			g.Printf(`
		)
		`)
		}

		var enumValues []string
		for _, e := range t.Enum {
			enumValues = append(enumValues, fmt.Sprintf("%q", e))
		}

		g.Printf(`
	func (e %[1]s) Valid() bool {
		switch e {
		case %[2]s:
			return true
		default:
			return false
		}
	}

	func (e %[1]s) String() string {
		return string(e)
	}
	`, t.Name(d), strings.Join(enumValues, ", "))
	}
}

// CmdType generates the type for CDP methods names.
func (g *Generator) CmdType(doms []proto.Domain) {
	g.hasContent = true
	g.Printf(`
	// CmdType is the type for CDP methods names.
	type CmdType string

	func (c CmdType) String() string {
		return string(c)
	}

	// Cmd methods.
	const (`)
	for _, d := range doms {
		for _, c := range d.Commands {
			g.Printf("\n\t%s CmdType = %q", c.CmdName(d, true), d.Domain+"."+c.NameName)
		}
	}
	g.Printf("\n)\n")
}

// DomainCmd defines the command args and reply.
func (g *Generator) DomainCmd(d proto.Domain, c proto.Command, sharedTypes map[string]bool) {
	g.beginType()
	g.hasContent = true
	g.domainCmdArgs(d, c, sharedTypes)
	g.commitType()

	g.beginType()
	g.hasContent = true
	g.domainCmdReply(d, c, sharedTypes)
	g.commitType()
}

func (g *Generator) domainCmdArgs(d proto.Domain, c proto.Command, sharedTypes map[string]bool) {
	g.Printf("const %[1]s = %[2]q\n", "Command"+d.Domain+strings.Title(c.NameName), d.Domain+"."+c.NameName)

	g.Printf(`
	// %[1]s represents the arguments for %[2]s in the %[3]s domain.
	type %[1]s struct {
		`, c.ArgsName(d), c.Name(), d.Name())
	g.printStructProperties(d, c.ArgsName(d), c.Parameters, true, true, sharedTypes)
	g.Printf("}\n\n")

	g.Printf(`
		// Unmarshal the byte array into a return value for %[2]s in the %[3]s domain.
		func (a * %[1]s) UnmarshalJSON(b []byte) error {
			type Copy %[1]s
			c := &Copy{}
			err := json.Unmarshal(b, c)
			if err != nil {
				return err
			}
			*a = %[1]s(*c)
			return nil
		}
		`, c.ArgsName(d), c.Name(), d.Name())

	g.Printf(`
		// Marshall the byte array into a return value for %[2]s in the %[3]s domain.
		func (a * %[1]s) MarshalJSON() ([]byte, error) {
			type Copy %[1]s
			c := &Copy{}
			*c = Copy(*a)
			return json.Marshal(&c)
		}
		`, c.ArgsName(d), c.Name(), d.Name())
}

func (g *Generator) domainCmdReply(d proto.Domain, c proto.Command, sharedTypes map[string]bool) {
	g.Printf(`
		// %[1]s represents the return values for %[2]s in the %[3]s domain.
		type %[1]s struct {
			`, c.ReplyName(d), c.Name(), d.Name())
	g.printStructProperties(d, c.ReplyName(d), c.Returns, true, false, sharedTypes)
	g.Printf("}\n\n")

	hasFrameId, hasNode, hasArrayNodes := false, false, false
	for _, a := range c.Returns {
		if a.NameName == "frameId" {
			hasFrameId = true
		}
		if a.NameName == "node" {
			hasNode = true
		}
		if a.NameName == "nodes" {
			if d.Name() == "DOMSnapshot" {
				// TODO: It does have a FrameID value but I have to look at a live example first.
			}
			if d.Name() != "Accessibility" && d.Name() != "DOMSnapshot" && d.Name() != "Memory" {
				hasArrayNodes = true
			}
		}
	}
	if hasFrameId {
		g.Printf(`
				// %[1]s returns whether or not the FrameID matches the reply value for %[2]s in the %[3]s domain.
				func (a * %[1]s) MatchFrameID(frameID string, m []byte) bool {
					`, c.ReplyName(d), c.Name(), d.Name())

		g.Printf(`err := a.UnmarshalJSON(m)
					if err != nil {
						log.Fatalf("unmarshal error: %s", err)
					}
					return a.FrameID == shared.FrameID(frameID)
				}
				`, c.ReplyName(d))
	} else if hasNode {
		g.Printf(`
				// %[1]s returns whether or not the FrameID matches the reply value for %[2]s in the %[3]s domain.
				func (a * %[1]s) MatchFrameID(frameID string, m []byte) bool {
					`, c.ReplyName(d), c.Name(), d.Name())

		g.Printf(`err := a.UnmarshalJSON(m)
					if err != nil {
						log.Fatalf("unmarshal error: %s", err)
					}
					return a.Node.FrameID == shared.FrameID(frameID)
				}
				`, c.ReplyName(d))
	} else if hasArrayNodes {
		g.Printf(`
				// %[1]s returns whether or not the FrameID matches the reply value for %[2]s in the %[3]s domain.
				func (a * %[1]s) MatchFrameID(frameID string, m []byte) bool {
					`, c.ReplyName(d), c.Name(), d.Name())

		g.Printf(`err := a.UnmarshalJSON(m)
					if err != nil {
						log.Fatalf("unmarshal error: %s", err)
					}
					fid := ""
					for _, n := range a.Nodes {
						if n.FrameID != "" {
							fid = string(n.FrameID)
						}
					}
					return fid == frameID
				}
				`, c.ReplyName(d))
	} else {
		g.Printf(`
				// %[1]s returns whether or not the FrameID matches the reply value for %[2]s in the %[3]s domain.
				func (a * %[1]s) MatchFrameID(frameID string, m []byte) bool {
					`, c.ReplyName(d), c.Name(), d.Name())

		g.Printf(`err := a.UnmarshalJSON(m)
					if err != nil {
						log.Fatalf("unmarshal error: %s", err)
					}
					return true
				}
				`, c.ReplyName(d))
	}

	g.Printf(`
			// Unmarshal the byte array into a return value for %[2]s in the %[3]s domain.
			func (a * %[1]s) UnmarshalJSON(b []byte) error {
				type Copy %[1]s
				c := &Copy{}
				err := json.Unmarshal(b, c)
				if err != nil {
					return err
				}
				*a = %[1]s(*c)
				return nil
			}
			`, c.ReplyName(d), c.Name(), d.Name())
}

// EventType generates the type for CDP event names.
func (g *Generator) EventType(doms []proto.Domain) {
	g.hasContent = true
	g.Printf(`
			// EventType is the type for CDP event names.
			type EventType string

			func (e EventType) String() string {
				return string(e)
			}

			// Event methods.
			const (`)
	for _, d := range doms {
		for _, e := range d.Events {
			g.Printf("\n\t%s EventType = %q", e.EventName(d), d.Domain+"."+e.NameName)
		}
	}
	g.Printf("\n)\n")
}

// DomainEvent defines the event client and reply.
func (g *Generator) DomainEvent(d proto.Domain, e proto.Event, sharedTypes map[string]bool) {
	g.hasContent = true

	g.beginType()
	g.domainEventReply(d, e, sharedTypes)
	g.commitType()
}

func (g *Generator) domainEventReply(d proto.Domain, e proto.Event, sharedTypes map[string]bool) {
	g.Printf("const %[1]s = %[2]q\n", "Event"+d.Domain+strings.Title(e.NameName), d.Domain+"."+e.NameName)

	g.Printf(`
			// %[1]s is the reply for %[2]s events.
			type %[1]s struct {
				`, e.ReplyName(d), e.Name())
	g.printStructProperties(d, e.ReplyName(d), e.Parameters, true, false, sharedTypes)
	g.Printf("}\n")

	g.Printf(`
				// Unmarshal the byte array into a return value for %[2]s in the %[3]s domain.
				func (a * %[1]s) UnmarshalJSON(b []byte) error {
					type Copy %[1]s
					c := &Copy{}
					err := json.Unmarshal(b, c)
					if err != nil {
						return err
					}
					*a = %[1]s(*c)
					return nil
				}
				`, e.ReplyName(d), e.Name(), d.Name())
}

func quotedImports(imports []string) string {
	if len(imports) == 0 {
		return ""
	}
	return "\"" + strings.Join(imports, "\"\n\"") + "\""
}

func isNonPointer(pkg string, d proto.Domain, t proto.AnyType) bool {
	typ := t.GoType(pkg, d)
	switch {
	case t.IsEnum():
	case strings.HasPrefix(typ, "[]"):
	case strings.HasPrefix(typ, "map["):
	case typ == "time.Time":
	case typ == "json.RawMessage":
	case typ == "RawMessage":
	case typ == "interface{}":
	default:
		return false
	}
	return true
}
