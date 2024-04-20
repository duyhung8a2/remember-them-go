// Code generated by BobGen sqlite v0.25.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"
	"time"

	"github.com/aarondl/opt/null"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/jaswdr/faker/v2"
	"github.com/stephenafamo/bob"
	models "remember_them/models"
)

type PagePropertyMod interface {
	Apply(*PagePropertyTemplate)
}

type PagePropertyModFunc func(*PagePropertyTemplate)

func (f PagePropertyModFunc) Apply(n *PagePropertyTemplate) {
	f(n)
}

type PagePropertyModSlice []PagePropertyMod

func (mods PagePropertyModSlice) Apply(n *PagePropertyTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// PagePropertyTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type PagePropertyTemplate struct {
	ID        func() int32
	PageID    func() int32
	Name      func() string
	Value     func() null.Val[string]
	CreatedAt func() null.Val[time.Time]
	UpdatedAt func() null.Val[time.Time]

	r pagePropertyR
	f *Factory
}

type pagePropertyR struct {
	Page *pagePropertyRPageR
}

type pagePropertyRPageR struct {
	o *PageTemplate
}

// Apply mods to the PagePropertyTemplate
func (o *PagePropertyTemplate) Apply(mods ...PagePropertyMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.PageProperty
// this does nothing with the relationship templates
func (o PagePropertyTemplate) toModel() *models.PageProperty {
	m := &models.PageProperty{}

	if o.ID != nil {
		m.ID = o.ID()
	}
	if o.PageID != nil {
		m.PageID = o.PageID()
	}
	if o.Name != nil {
		m.Name = o.Name()
	}
	if o.Value != nil {
		m.Value = o.Value()
	}
	if o.CreatedAt != nil {
		m.CreatedAt = o.CreatedAt()
	}
	if o.UpdatedAt != nil {
		m.UpdatedAt = o.UpdatedAt()
	}

	return m
}

// toModels returns an models.PagePropertySlice
// this does nothing with the relationship templates
func (o PagePropertyTemplate) toModels(number int) models.PagePropertySlice {
	m := make(models.PagePropertySlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.PageProperty
// according to the relationships in the template. Nothing is inserted into the db
func (t PagePropertyTemplate) setModelRels(o *models.PageProperty) {
	if t.r.Page != nil {
		rel := t.r.Page.o.toModel()
		rel.R.PageProperties = append(rel.R.PageProperties, o)
		o.PageID = rel.ID
		o.R.Page = rel
	}
}

// BuildSetter returns an *models.PagePropertySetter
// this does nothing with the relationship templates
func (o PagePropertyTemplate) BuildSetter() *models.PagePropertySetter {
	m := &models.PagePropertySetter{}

	if o.ID != nil {
		m.ID = omit.From(o.ID())
	}
	if o.PageID != nil {
		m.PageID = omit.From(o.PageID())
	}
	if o.Name != nil {
		m.Name = omit.From(o.Name())
	}
	if o.Value != nil {
		m.Value = omitnull.FromNull(o.Value())
	}
	if o.CreatedAt != nil {
		m.CreatedAt = omitnull.FromNull(o.CreatedAt())
	}
	if o.UpdatedAt != nil {
		m.UpdatedAt = omitnull.FromNull(o.UpdatedAt())
	}

	return m
}

// BuildManySetter returns an []*models.PagePropertySetter
// this does nothing with the relationship templates
func (o PagePropertyTemplate) BuildManySetter(number int) []*models.PagePropertySetter {
	m := make([]*models.PagePropertySetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.PageProperty
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use PagePropertyTemplate.Create
func (o PagePropertyTemplate) Build() *models.PageProperty {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.PagePropertySlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use PagePropertyTemplate.CreateMany
func (o PagePropertyTemplate) BuildMany(number int) models.PagePropertySlice {
	m := make(models.PagePropertySlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatablePageProperty(m *models.PagePropertySetter) {
	if m.PageID.IsUnset() {
		m.PageID = omit.From(random[int32](nil))
	}
	if m.Name.IsUnset() {
		m.Name = omit.From(random[string](nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.PageProperty
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *PagePropertyTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.PageProperty) (context.Context, error) {
	var err error

	return ctx, err
}

// Create builds a pageProperty and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *PagePropertyTemplate) Create(ctx context.Context, exec bob.Executor) (*models.PageProperty, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// create builds a pageProperty and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *PagePropertyTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.PageProperty, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatablePageProperty(opt)

	var rel0 *models.Page
	if o.r.Page == nil {
		var ok bool
		rel0, ok = pageCtx.Value(ctx)
		if !ok {
			PagePropertyMods.WithNewPage().Apply(o)
		}
	}
	if o.r.Page != nil {
		ctx, rel0, err = o.r.Page.o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}
	opt.PageID = omit.From(rel0.ID)

	m, err := models.PageProperties.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = pagePropertyCtx.WithValue(ctx, m)

	m.R.Page = rel0

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple pageProperties and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o PagePropertyTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.PagePropertySlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// createMany builds multiple pageProperties and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o PagePropertyTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.PagePropertySlice, error) {
	var err error
	m := make(models.PagePropertySlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// PageProperty has methods that act as mods for the PagePropertyTemplate
var PagePropertyMods pagePropertyMods

type pagePropertyMods struct{}

func (m pagePropertyMods) RandomizeAllColumns(f *faker.Faker) PagePropertyMod {
	return PagePropertyModSlice{
		PagePropertyMods.RandomID(f),
		PagePropertyMods.RandomPageID(f),
		PagePropertyMods.RandomName(f),
		PagePropertyMods.RandomValue(f),
		PagePropertyMods.RandomCreatedAt(f),
		PagePropertyMods.RandomUpdatedAt(f),
	}
}

// Set the model columns to this value
func (m pagePropertyMods) ID(val int32) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.ID = func() int32 { return val }
	})
}

// Set the Column from the function
func (m pagePropertyMods) IDFunc(f func() int32) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.ID = f
	})
}

// Clear any values for the column
func (m pagePropertyMods) UnsetID() PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.ID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pagePropertyMods) RandomID(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.ID = func() int32 {
			return random[int32](f)
		}
	})
}

func (m pagePropertyMods) ensureID(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		if o.ID != nil {
			return
		}

		o.ID = func() int32 {
			return random[int32](f)
		}
	})
}

// Set the model columns to this value
func (m pagePropertyMods) PageID(val int32) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.PageID = func() int32 { return val }
	})
}

// Set the Column from the function
func (m pagePropertyMods) PageIDFunc(f func() int32) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.PageID = f
	})
}

// Clear any values for the column
func (m pagePropertyMods) UnsetPageID() PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.PageID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pagePropertyMods) RandomPageID(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.PageID = func() int32 {
			return random[int32](f)
		}
	})
}

func (m pagePropertyMods) ensurePageID(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		if o.PageID != nil {
			return
		}

		o.PageID = func() int32 {
			return random[int32](f)
		}
	})
}

// Set the model columns to this value
func (m pagePropertyMods) Name(val string) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.Name = func() string { return val }
	})
}

// Set the Column from the function
func (m pagePropertyMods) NameFunc(f func() string) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.Name = f
	})
}

// Clear any values for the column
func (m pagePropertyMods) UnsetName() PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.Name = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pagePropertyMods) RandomName(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.Name = func() string {
			return random[string](f)
		}
	})
}

func (m pagePropertyMods) ensureName(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		if o.Name != nil {
			return
		}

		o.Name = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m pagePropertyMods) Value(val null.Val[string]) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.Value = func() null.Val[string] { return val }
	})
}

// Set the Column from the function
func (m pagePropertyMods) ValueFunc(f func() null.Val[string]) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.Value = f
	})
}

// Clear any values for the column
func (m pagePropertyMods) UnsetValue() PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.Value = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pagePropertyMods) RandomValue(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.Value = func() null.Val[string] {
			return randomNull[string](f)
		}
	})
}

func (m pagePropertyMods) ensureValue(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		if o.Value != nil {
			return
		}

		o.Value = func() null.Val[string] {
			return randomNull[string](f)
		}
	})
}

// Set the model columns to this value
func (m pagePropertyMods) CreatedAt(val null.Val[time.Time]) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.CreatedAt = func() null.Val[time.Time] { return val }
	})
}

// Set the Column from the function
func (m pagePropertyMods) CreatedAtFunc(f func() null.Val[time.Time]) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.CreatedAt = f
	})
}

// Clear any values for the column
func (m pagePropertyMods) UnsetCreatedAt() PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.CreatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pagePropertyMods) RandomCreatedAt(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.CreatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

func (m pagePropertyMods) ensureCreatedAt(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		if o.CreatedAt != nil {
			return
		}

		o.CreatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

// Set the model columns to this value
func (m pagePropertyMods) UpdatedAt(val null.Val[time.Time]) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.UpdatedAt = func() null.Val[time.Time] { return val }
	})
}

// Set the Column from the function
func (m pagePropertyMods) UpdatedAtFunc(f func() null.Val[time.Time]) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.UpdatedAt = f
	})
}

// Clear any values for the column
func (m pagePropertyMods) UnsetUpdatedAt() PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.UpdatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pagePropertyMods) RandomUpdatedAt(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.UpdatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

func (m pagePropertyMods) ensureUpdatedAt(f *faker.Faker) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		if o.UpdatedAt != nil {
			return
		}

		o.UpdatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

func (m pagePropertyMods) WithPage(rel *PageTemplate) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.r.Page = &pagePropertyRPageR{
			o: rel,
		}
	})
}

func (m pagePropertyMods) WithNewPage(mods ...PageMod) PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		related := o.f.NewPage(mods...)

		m.WithPage(related).Apply(o)
	})
}

func (m pagePropertyMods) WithoutPage() PagePropertyMod {
	return PagePropertyModFunc(func(o *PagePropertyTemplate) {
		o.r.Page = nil
	})
}