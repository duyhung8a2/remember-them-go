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

type PageMod interface {
	Apply(*PageTemplate)
}

type PageModFunc func(*PageTemplate)

func (f PageModFunc) Apply(n *PageTemplate) {
	f(n)
}

type PageModSlice []PageMod

func (mods PageModSlice) Apply(n *PageTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// PageTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type PageTemplate struct {
	ID        func() int32
	Title     func() string
	UserID    func() int32
	ParentID  func() null.Val[int32]
	CreatedAt func() null.Val[time.Time]
	UpdatedAt func() null.Val[time.Time]

	r pageR
	f *Factory
}

type pageR struct {
	Blocks         []*pageRBlocksR
	PageProperties []*pageRPagePropertiesR
	User           *pageRUserR
	Parent         *pageRParentR
	ReverseParents []*pageRReverseParentsR
}

type pageRBlocksR struct {
	number int
	o      *BlockTemplate
}
type pageRPagePropertiesR struct {
	number int
	o      *PagePropertyTemplate
}
type pageRUserR struct {
	o *UserTemplate
}
type pageRParentR struct {
	o *PageTemplate
}
type pageRReverseParentsR struct {
	number int
	o      *PageTemplate
}

// Apply mods to the PageTemplate
func (o *PageTemplate) Apply(mods ...PageMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.Page
// this does nothing with the relationship templates
func (o PageTemplate) toModel() *models.Page {
	m := &models.Page{}

	if o.ID != nil {
		m.ID = o.ID()
	}
	if o.Title != nil {
		m.Title = o.Title()
	}
	if o.UserID != nil {
		m.UserID = o.UserID()
	}
	if o.ParentID != nil {
		m.ParentID = o.ParentID()
	}
	if o.CreatedAt != nil {
		m.CreatedAt = o.CreatedAt()
	}
	if o.UpdatedAt != nil {
		m.UpdatedAt = o.UpdatedAt()
	}

	return m
}

// toModels returns an models.PageSlice
// this does nothing with the relationship templates
func (o PageTemplate) toModels(number int) models.PageSlice {
	m := make(models.PageSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.Page
// according to the relationships in the template. Nothing is inserted into the db
func (t PageTemplate) setModelRels(o *models.Page) {
	if t.r.Blocks != nil {
		rel := models.BlockSlice{}
		for _, r := range t.r.Blocks {
			related := r.o.toModels(r.number)
			for _, rel := range related {
				rel.PageID = o.ID
				rel.R.Page = o
			}
			rel = append(rel, related...)
		}
		o.R.Blocks = rel
	}

	if t.r.PageProperties != nil {
		rel := models.PagePropertySlice{}
		for _, r := range t.r.PageProperties {
			related := r.o.toModels(r.number)
			for _, rel := range related {
				rel.PageID = o.ID
				rel.R.Page = o
			}
			rel = append(rel, related...)
		}
		o.R.PageProperties = rel
	}

	if t.r.User != nil {
		rel := t.r.User.o.toModel()
		rel.R.Pages = append(rel.R.Pages, o)
		o.UserID = rel.ID
		o.R.User = rel
	}

	if t.r.Parent != nil {
		rel := t.r.Parent.o.toModel()
		rel.R.Parent = o
		o.ParentID = null.From(rel.ID)
		o.R.Parent = rel
	}

	if t.r.ReverseParents != nil {
		rel := models.PageSlice{}
		for _, r := range t.r.ReverseParents {
			related := r.o.toModels(r.number)
			for _, rel := range related {
				rel.ParentID = null.From(o.ID)
				rel.R.ReverseParents = append(rel.R.ReverseParents, o)
			}
			rel = append(rel, related...)
		}
		o.R.ReverseParents = rel
	}
}

// BuildSetter returns an *models.PageSetter
// this does nothing with the relationship templates
func (o PageTemplate) BuildSetter() *models.PageSetter {
	m := &models.PageSetter{}

	if o.ID != nil {
		m.ID = omit.From(o.ID())
	}
	if o.Title != nil {
		m.Title = omit.From(o.Title())
	}
	if o.UserID != nil {
		m.UserID = omit.From(o.UserID())
	}
	if o.ParentID != nil {
		m.ParentID = omitnull.FromNull(o.ParentID())
	}
	if o.CreatedAt != nil {
		m.CreatedAt = omitnull.FromNull(o.CreatedAt())
	}
	if o.UpdatedAt != nil {
		m.UpdatedAt = omitnull.FromNull(o.UpdatedAt())
	}

	return m
}

// BuildManySetter returns an []*models.PageSetter
// this does nothing with the relationship templates
func (o PageTemplate) BuildManySetter(number int) []*models.PageSetter {
	m := make([]*models.PageSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.Page
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use PageTemplate.Create
func (o PageTemplate) Build() *models.Page {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.PageSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use PageTemplate.CreateMany
func (o PageTemplate) BuildMany(number int) models.PageSlice {
	m := make(models.PageSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatablePage(m *models.PageSetter) {
	if m.Title.IsUnset() {
		m.Title = omit.From(random[string](nil))
	}
	if m.UserID.IsUnset() {
		m.UserID = omit.From(random[int32](nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.Page
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *PageTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.Page) (context.Context, error) {
	var err error

	if o.r.Blocks != nil {
		for _, r := range o.r.Blocks {
			var rel0 models.BlockSlice
			ctx, rel0, err = r.o.createMany(ctx, exec, r.number)
			if err != nil {
				return ctx, err
			}

			err = m.AttachBlocks(ctx, exec, rel0...)
			if err != nil {
				return ctx, err
			}
		}
	}

	if o.r.PageProperties != nil {
		for _, r := range o.r.PageProperties {
			var rel1 models.PagePropertySlice
			ctx, rel1, err = r.o.createMany(ctx, exec, r.number)
			if err != nil {
				return ctx, err
			}

			err = m.AttachPageProperties(ctx, exec, rel1...)
			if err != nil {
				return ctx, err
			}
		}
	}

	if o.r.Parent != nil {
		var rel3 *models.Page
		ctx, rel3, err = o.r.Parent.o.create(ctx, exec)
		if err != nil {
			return ctx, err
		}
		err = m.AttachParent(ctx, exec, rel3)
		if err != nil {
			return ctx, err
		}
	}

	if o.r.ReverseParents != nil {
		for _, r := range o.r.ReverseParents {
			var rel4 models.PageSlice
			ctx, rel4, err = r.o.createMany(ctx, exec, r.number)
			if err != nil {
				return ctx, err
			}

			err = m.AttachReverseParents(ctx, exec, rel4...)
			if err != nil {
				return ctx, err
			}
		}
	}

	return ctx, err
}

// Create builds a page and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *PageTemplate) Create(ctx context.Context, exec bob.Executor) (*models.Page, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// create builds a page and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *PageTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.Page, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatablePage(opt)

	var rel2 *models.User
	if o.r.User == nil {
		var ok bool
		rel2, ok = userCtx.Value(ctx)
		if !ok {
			PageMods.WithNewUser().Apply(o)
		}
	}
	if o.r.User != nil {
		ctx, rel2, err = o.r.User.o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}
	opt.UserID = omit.From(rel2.ID)

	m, err := models.Pages.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = pageCtx.WithValue(ctx, m)

	m.R.User = rel2

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple pages and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o PageTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.PageSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// createMany builds multiple pages and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o PageTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.PageSlice, error) {
	var err error
	m := make(models.PageSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// Page has methods that act as mods for the PageTemplate
var PageMods pageMods

type pageMods struct{}

func (m pageMods) RandomizeAllColumns(f *faker.Faker) PageMod {
	return PageModSlice{
		PageMods.RandomID(f),
		PageMods.RandomTitle(f),
		PageMods.RandomUserID(f),
		PageMods.RandomParentID(f),
		PageMods.RandomCreatedAt(f),
		PageMods.RandomUpdatedAt(f),
	}
}

// Set the model columns to this value
func (m pageMods) ID(val int32) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.ID = func() int32 { return val }
	})
}

// Set the Column from the function
func (m pageMods) IDFunc(f func() int32) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.ID = f
	})
}

// Clear any values for the column
func (m pageMods) UnsetID() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.ID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pageMods) RandomID(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.ID = func() int32 {
			return random[int32](f)
		}
	})
}

func (m pageMods) ensureID(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		if o.ID != nil {
			return
		}

		o.ID = func() int32 {
			return random[int32](f)
		}
	})
}

// Set the model columns to this value
func (m pageMods) Title(val string) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.Title = func() string { return val }
	})
}

// Set the Column from the function
func (m pageMods) TitleFunc(f func() string) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.Title = f
	})
}

// Clear any values for the column
func (m pageMods) UnsetTitle() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.Title = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pageMods) RandomTitle(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.Title = func() string {
			return random[string](f)
		}
	})
}

func (m pageMods) ensureTitle(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		if o.Title != nil {
			return
		}

		o.Title = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m pageMods) UserID(val int32) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.UserID = func() int32 { return val }
	})
}

// Set the Column from the function
func (m pageMods) UserIDFunc(f func() int32) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.UserID = f
	})
}

// Clear any values for the column
func (m pageMods) UnsetUserID() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.UserID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pageMods) RandomUserID(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.UserID = func() int32 {
			return random[int32](f)
		}
	})
}

func (m pageMods) ensureUserID(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		if o.UserID != nil {
			return
		}

		o.UserID = func() int32 {
			return random[int32](f)
		}
	})
}

// Set the model columns to this value
func (m pageMods) ParentID(val null.Val[int32]) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.ParentID = func() null.Val[int32] { return val }
	})
}

// Set the Column from the function
func (m pageMods) ParentIDFunc(f func() null.Val[int32]) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.ParentID = f
	})
}

// Clear any values for the column
func (m pageMods) UnsetParentID() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.ParentID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pageMods) RandomParentID(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.ParentID = func() null.Val[int32] {
			return randomNull[int32](f)
		}
	})
}

func (m pageMods) ensureParentID(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		if o.ParentID != nil {
			return
		}

		o.ParentID = func() null.Val[int32] {
			return randomNull[int32](f)
		}
	})
}

// Set the model columns to this value
func (m pageMods) CreatedAt(val null.Val[time.Time]) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.CreatedAt = func() null.Val[time.Time] { return val }
	})
}

// Set the Column from the function
func (m pageMods) CreatedAtFunc(f func() null.Val[time.Time]) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.CreatedAt = f
	})
}

// Clear any values for the column
func (m pageMods) UnsetCreatedAt() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.CreatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pageMods) RandomCreatedAt(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.CreatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

func (m pageMods) ensureCreatedAt(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		if o.CreatedAt != nil {
			return
		}

		o.CreatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

// Set the model columns to this value
func (m pageMods) UpdatedAt(val null.Val[time.Time]) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.UpdatedAt = func() null.Val[time.Time] { return val }
	})
}

// Set the Column from the function
func (m pageMods) UpdatedAtFunc(f func() null.Val[time.Time]) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.UpdatedAt = f
	})
}

// Clear any values for the column
func (m pageMods) UnsetUpdatedAt() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.UpdatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m pageMods) RandomUpdatedAt(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.UpdatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

func (m pageMods) ensureUpdatedAt(f *faker.Faker) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		if o.UpdatedAt != nil {
			return
		}

		o.UpdatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

func (m pageMods) WithUser(rel *UserTemplate) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.User = &pageRUserR{
			o: rel,
		}
	})
}

func (m pageMods) WithNewUser(mods ...UserMod) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		related := o.f.NewUser(mods...)

		m.WithUser(related).Apply(o)
	})
}

func (m pageMods) WithoutUser() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.User = nil
	})
}

func (m pageMods) WithParent(rel *PageTemplate) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.Parent = &pageRParentR{
			o: rel,
		}
	})
}

func (m pageMods) WithNewParent(mods ...PageMod) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		related := o.f.NewPage(mods...)

		m.WithParent(related).Apply(o)
	})
}

func (m pageMods) WithoutParent() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.Parent = nil
	})
}

func (m pageMods) WithBlocks(number int, related *BlockTemplate) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.Blocks = []*pageRBlocksR{{
			number: number,
			o:      related,
		}}
	})
}

func (m pageMods) WithNewBlocks(number int, mods ...BlockMod) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		related := o.f.NewBlock(mods...)
		m.WithBlocks(number, related).Apply(o)
	})
}

func (m pageMods) AddBlocks(number int, related *BlockTemplate) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.Blocks = append(o.r.Blocks, &pageRBlocksR{
			number: number,
			o:      related,
		})
	})
}

func (m pageMods) AddNewBlocks(number int, mods ...BlockMod) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		related := o.f.NewBlock(mods...)
		m.AddBlocks(number, related).Apply(o)
	})
}

func (m pageMods) WithoutBlocks() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.Blocks = nil
	})
}

func (m pageMods) WithPageProperties(number int, related *PagePropertyTemplate) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.PageProperties = []*pageRPagePropertiesR{{
			number: number,
			o:      related,
		}}
	})
}

func (m pageMods) WithNewPageProperties(number int, mods ...PagePropertyMod) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		related := o.f.NewPageProperty(mods...)
		m.WithPageProperties(number, related).Apply(o)
	})
}

func (m pageMods) AddPageProperties(number int, related *PagePropertyTemplate) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.PageProperties = append(o.r.PageProperties, &pageRPagePropertiesR{
			number: number,
			o:      related,
		})
	})
}

func (m pageMods) AddNewPageProperties(number int, mods ...PagePropertyMod) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		related := o.f.NewPageProperty(mods...)
		m.AddPageProperties(number, related).Apply(o)
	})
}

func (m pageMods) WithoutPageProperties() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.PageProperties = nil
	})
}

func (m pageMods) WithReverseParents(number int, related *PageTemplate) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.ReverseParents = []*pageRReverseParentsR{{
			number: number,
			o:      related,
		}}
	})
}

func (m pageMods) WithNewReverseParents(number int, mods ...PageMod) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		related := o.f.NewPage(mods...)
		m.WithReverseParents(number, related).Apply(o)
	})
}

func (m pageMods) AddReverseParents(number int, related *PageTemplate) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.ReverseParents = append(o.r.ReverseParents, &pageRReverseParentsR{
			number: number,
			o:      related,
		})
	})
}

func (m pageMods) AddNewReverseParents(number int, mods ...PageMod) PageMod {
	return PageModFunc(func(o *PageTemplate) {
		related := o.f.NewPage(mods...)
		m.AddReverseParents(number, related).Apply(o)
	})
}

func (m pageMods) WithoutReverseParents() PageMod {
	return PageModFunc(func(o *PageTemplate) {
		o.r.ReverseParents = nil
	})
}
