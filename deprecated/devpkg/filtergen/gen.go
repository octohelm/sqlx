package filtergen

import (
	"cmp"
	"context"
	"fmt"
	"go/types"
	"reflect"
	"strings"

	"github.com/octohelm/gengo/pkg/camelcase"
	"github.com/octohelm/gengo/pkg/gengo"
	"github.com/octohelm/gengo/pkg/gengo/snippet"
	"github.com/octohelm/storage/deprecated/pkg/dal"
	dalcompose "github.com/octohelm/storage/deprecated/pkg/dal/compose"
	tablegenutil "github.com/octohelm/storage/devpkg/tablegen/util"
	"github.com/octohelm/storage/pkg/filter"
	"github.com/octohelm/storage/pkg/sqlbuilder"
)

func init() {
	gengo.Register(&filterGen{})
}

type filterGen struct{}

func (*filterGen) Name() string {
	return "filter"
}

func (*filterGen) New(ctx gengo.Context) gengo.Generator {
	return &filterGen{}
}

func (g *filterGen) GenerateType(c gengo.Context, srcNamed *types.Named) error {
	if typStruct, ok := srcNamed.Underlying().(*types.Struct); ok {
		tables := make(map[*types.Named]sqlbuilder.Table)

		for i := range typStruct.NumFields() {
			f := typStruct.Field(i)
			tag := reflect.StructTag(typStruct.Tag(i))

			if named, ok := f.Type().(*types.Named); ok {
				t, err := tablegenutil.ScanTable(c, named)
				if err != nil {
					return err
				}

				tables[named] = t

				if s, ok := tag.Lookup("select"); ok {
					g.generateSubFilter(c, tables, cmp.Or(tag.Get("as"), tag.Get("by")), named, s)
				} else {
					g.generateIndexedFilter(c, t, named, tag.Get("domain"))
				}
			}
		}
	}

	return gengo.ErrSkip
}

func (g *filterGen) generateSubFilter(c gengo.Context, tables map[*types.Named]sqlbuilder.Table, as string, fromModeType *types.Named, fromModelFieldName string) {
	values := strings.Split(as, ".")

	modelTypeName := values[0]
	modelTypeFieldName := values[1]

	var modelType *types.Named

	for named := range tables {
		if named.Obj().Name() == modelTypeName {
			modelType = named
		}
	}

	if modelType == nil {
		return
	}

	c.RenderT(`
func @ModelTypeName'By@ModelFieldName'From@FromModelTypeName'(ctx @contextContext, patchers ...@patcherTyped[@FromModelType]) @patcherTyped[@ModelType] {
	return @patcherWhere[@ModelType](
		@ModelType'T.@ModelFieldName.V(
			@patcherInSelectIfExists(ctx, @FromModelType'T.@FromModelFieldName, patchers...),
		),
	)
}
`, snippet.Args{
		"ModelTypeName":  snippet.ID(modelType.Obj().Name()),
		"ModelType":      snippet.ID(modelType.String()),
		"ModelFieldName": snippet.ID(modelTypeFieldName),

		"FromModelTypeName":  snippet.ID(fromModeType.Obj().Name()),
		"FromModelType":      snippet.ID(fromModeType.String()),
		"FromModelFieldName": snippet.ID(fromModelFieldName),

		"contextContext":          snippet.ID(reflect.TypeFor[context.Context]()),
		"patcherTyped":            snippet.ID("github.com/octohelm/storage/deprecated/pkg/dal/compose/querierpatcher.Typed"),
		"patcherWhere":            snippet.ID("github.com/octohelm/storage/deprecated/pkg/dal/compose/querierpatcher.Where"),
		"patcherInSelectIfExists": snippet.ID("github.com/octohelm/storage/deprecated/pkg/dal/compose/querierpatcher.InSelectIfExists"),
	})
}

func (g *filterGen) generateIndexedFilter(c gengo.Context, t sqlbuilder.Table, named *types.Named, domainName string) {
	indexedFields := make([]string, 0)

	cols := map[string]bool{}

	for k := range t.Keys() {
		for col := range k.Cols() {
			cols[col.FieldName()] = true
		}
	}

	for col := range t.Cols() {
		if cols[col.FieldName()] {
			indexedFields = append(indexedFields, col.FieldName())
		}
	}

	if domainName == "" {
		domainName = strings.TrimPrefix(t.TableName(), "t_")
	}

	for _, fieldName := range indexedFields {
		f := t.F(fieldName)
		fieldType := sqlbuilder.GetColumnDef(f).Type

		fieldComment := fmt.Sprintf("%s", func() string {
			if comment := sqlbuilder.GetColumnDef(f).Comment; comment != "" {
				return comment
			}
			return ""
		}())

		c.RenderT(`
type @ModelTypeName'By@FieldName struct {
	@composeFrom[@Type] 

	@fieldComment
	@FieldName *@filterFilter[@FieldType] `+"`"+`name:"@domainName~@fieldName,omitempty" in:"query"`+"`"+`
}

func (f *@ModelTypeName'By@FieldName) ApplyQuerier(q @dalQuerier) @dalQuerier {
	return @composeApplyQuerierFromFilter(q, @Type'T.@FieldName, f.@FieldName)
}
`, snippet.Args{
			"ModelTypeName": snippet.ID(named.Obj().Name()),
			"Type":          snippet.ID(named.Obj()),

			"FieldName":    snippet.ID(fieldName),
			"FieldType":    snippet.ID(fieldType.String()),
			"fieldComment": snippet.Comment(fieldComment),
			"domainName":   snippet.ID(camelcase.LowerKebabCase(domainName)),
			"fieldName":    snippet.ID(camelcase.LowerCamelCase(fieldName)),

			"dalQuerier": snippet.ID(reflect.TypeFor[dal.Querier]()),

			"composeApplyQuerierFromFilter": snippet.PkgExposeFor[dalcompose.P]("ApplyQuerierFromFilter"),
			"composeFrom":                   snippet.PkgExposeFor[dalcompose.P]("From"),
			"filterFilter":                  snippet.PkgExposeFor[filter.P]("Filter"),
		})
	}
}

type defaultModel struct{}

func (defaultModel) TableName() string {
	return "t"
}
