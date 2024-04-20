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

type UserMod interface {
	Apply(*UserTemplate)
}

type UserModFunc func(*UserTemplate)

func (f UserModFunc) Apply(n *UserTemplate) {
	f(n)
}

type UserModSlice []UserMod

func (mods UserModSlice) Apply(n *UserTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// UserTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type UserTemplate struct {
	ID        func() int32
	Username  func() string
	Email     func() string
	Password  func() string
	CreatedAt func() null.Val[time.Time]
	UpdatedAt func() null.Val[time.Time]

	r userR
	f *Factory
}

type userR struct {
	Pages []*userRPagesR
}

type userRPagesR struct {
	number int
	o      *PageTemplate
}

// Apply mods to the UserTemplate
func (o *UserTemplate) Apply(mods ...UserMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.User
// this does nothing with the relationship templates
func (o UserTemplate) toModel() *models.User {
	m := &models.User{}

	if o.ID != nil {
		m.ID = o.ID()
	}
	if o.Username != nil {
		m.Username = o.Username()
	}
	if o.Email != nil {
		m.Email = o.Email()
	}
	if o.Password != nil {
		m.Password = o.Password()
	}
	if o.CreatedAt != nil {
		m.CreatedAt = o.CreatedAt()
	}
	if o.UpdatedAt != nil {
		m.UpdatedAt = o.UpdatedAt()
	}

	return m
}

// toModels returns an models.UserSlice
// this does nothing with the relationship templates
func (o UserTemplate) toModels(number int) models.UserSlice {
	m := make(models.UserSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.User
// according to the relationships in the template. Nothing is inserted into the db
func (t UserTemplate) setModelRels(o *models.User) {
	if t.r.Pages != nil {
		rel := models.PageSlice{}
		for _, r := range t.r.Pages {
			related := r.o.toModels(r.number)
			for _, rel := range related {
				rel.UserID = o.ID
				rel.R.User = o
			}
			rel = append(rel, related...)
		}
		o.R.Pages = rel
	}
}

// BuildSetter returns an *models.UserSetter
// this does nothing with the relationship templates
func (o UserTemplate) BuildSetter() *models.UserSetter {
	m := &models.UserSetter{}

	if o.ID != nil {
		m.ID = omit.From(o.ID())
	}
	if o.Username != nil {
		m.Username = omit.From(o.Username())
	}
	if o.Email != nil {
		m.Email = omit.From(o.Email())
	}
	if o.Password != nil {
		m.Password = omit.From(o.Password())
	}
	if o.CreatedAt != nil {
		m.CreatedAt = omitnull.FromNull(o.CreatedAt())
	}
	if o.UpdatedAt != nil {
		m.UpdatedAt = omitnull.FromNull(o.UpdatedAt())
	}

	return m
}

// BuildManySetter returns an []*models.UserSetter
// this does nothing with the relationship templates
func (o UserTemplate) BuildManySetter(number int) []*models.UserSetter {
	m := make([]*models.UserSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.User
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use UserTemplate.Create
func (o UserTemplate) Build() *models.User {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.UserSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use UserTemplate.CreateMany
func (o UserTemplate) BuildMany(number int) models.UserSlice {
	m := make(models.UserSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatableUser(m *models.UserSetter) {
	if m.Username.IsUnset() {
		m.Username = omit.From(random[string](nil))
	}
	if m.Email.IsUnset() {
		m.Email = omit.From(random[string](nil))
	}
	if m.Password.IsUnset() {
		m.Password = omit.From(random[string](nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.User
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *UserTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.User) (context.Context, error) {
	var err error

	if o.r.Pages != nil {
		for _, r := range o.r.Pages {
			var rel0 models.PageSlice
			ctx, rel0, err = r.o.createMany(ctx, exec, r.number)
			if err != nil {
				return ctx, err
			}

			err = m.AttachPages(ctx, exec, rel0...)
			if err != nil {
				return ctx, err
			}
		}
	}

	return ctx, err
}

// Create builds a user and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *UserTemplate) Create(ctx context.Context, exec bob.Executor) (*models.User, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// create builds a user and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *UserTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.User, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatableUser(opt)

	m, err := models.Users.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = userCtx.WithValue(ctx, m)

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple users and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o UserTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.UserSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// createMany builds multiple users and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o UserTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.UserSlice, error) {
	var err error
	m := make(models.UserSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// User has methods that act as mods for the UserTemplate
var UserMods userMods

type userMods struct{}

func (m userMods) RandomizeAllColumns(f *faker.Faker) UserMod {
	return UserModSlice{
		UserMods.RandomID(f),
		UserMods.RandomUsername(f),
		UserMods.RandomEmail(f),
		UserMods.RandomPassword(f),
		UserMods.RandomCreatedAt(f),
		UserMods.RandomUpdatedAt(f),
	}
}

// Set the model columns to this value
func (m userMods) ID(val int32) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.ID = func() int32 { return val }
	})
}

// Set the Column from the function
func (m userMods) IDFunc(f func() int32) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.ID = f
	})
}

// Clear any values for the column
func (m userMods) UnsetID() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.ID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m userMods) RandomID(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.ID = func() int32 {
			return random[int32](f)
		}
	})
}

func (m userMods) ensureID(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		if o.ID != nil {
			return
		}

		o.ID = func() int32 {
			return random[int32](f)
		}
	})
}

// Set the model columns to this value
func (m userMods) Username(val string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Username = func() string { return val }
	})
}

// Set the Column from the function
func (m userMods) UsernameFunc(f func() string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Username = f
	})
}

// Clear any values for the column
func (m userMods) UnsetUsername() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Username = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m userMods) RandomUsername(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Username = func() string {
			return random[string](f)
		}
	})
}

func (m userMods) ensureUsername(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		if o.Username != nil {
			return
		}

		o.Username = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m userMods) Email(val string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Email = func() string { return val }
	})
}

// Set the Column from the function
func (m userMods) EmailFunc(f func() string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Email = f
	})
}

// Clear any values for the column
func (m userMods) UnsetEmail() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Email = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m userMods) RandomEmail(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Email = func() string {
			return random[string](f)
		}
	})
}

func (m userMods) ensureEmail(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		if o.Email != nil {
			return
		}

		o.Email = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m userMods) Password(val string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Password = func() string { return val }
	})
}

// Set the Column from the function
func (m userMods) PasswordFunc(f func() string) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Password = f
	})
}

// Clear any values for the column
func (m userMods) UnsetPassword() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Password = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m userMods) RandomPassword(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.Password = func() string {
			return random[string](f)
		}
	})
}

func (m userMods) ensurePassword(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		if o.Password != nil {
			return
		}

		o.Password = func() string {
			return random[string](f)
		}
	})
}

// Set the model columns to this value
func (m userMods) CreatedAt(val null.Val[time.Time]) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.CreatedAt = func() null.Val[time.Time] { return val }
	})
}

// Set the Column from the function
func (m userMods) CreatedAtFunc(f func() null.Val[time.Time]) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.CreatedAt = f
	})
}

// Clear any values for the column
func (m userMods) UnsetCreatedAt() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.CreatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m userMods) RandomCreatedAt(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.CreatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

func (m userMods) ensureCreatedAt(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		if o.CreatedAt != nil {
			return
		}

		o.CreatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

// Set the model columns to this value
func (m userMods) UpdatedAt(val null.Val[time.Time]) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.UpdatedAt = func() null.Val[time.Time] { return val }
	})
}

// Set the Column from the function
func (m userMods) UpdatedAtFunc(f func() null.Val[time.Time]) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.UpdatedAt = f
	})
}

// Clear any values for the column
func (m userMods) UnsetUpdatedAt() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.UpdatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m userMods) RandomUpdatedAt(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.UpdatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

func (m userMods) ensureUpdatedAt(f *faker.Faker) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		if o.UpdatedAt != nil {
			return
		}

		o.UpdatedAt = func() null.Val[time.Time] {
			return randomNull[time.Time](f)
		}
	})
}

func (m userMods) WithPages(number int, related *PageTemplate) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.r.Pages = []*userRPagesR{{
			number: number,
			o:      related,
		}}
	})
}

func (m userMods) WithNewPages(number int, mods ...PageMod) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		related := o.f.NewPage(mods...)
		m.WithPages(number, related).Apply(o)
	})
}

func (m userMods) AddPages(number int, related *PageTemplate) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.r.Pages = append(o.r.Pages, &userRPagesR{
			number: number,
			o:      related,
		})
	})
}

func (m userMods) AddNewPages(number int, mods ...PageMod) UserMod {
	return UserModFunc(func(o *UserTemplate) {
		related := o.f.NewPage(mods...)
		m.AddPages(number, related).Apply(o)
	})
}

func (m userMods) WithoutPages() UserMod {
	return UserModFunc(func(o *UserTemplate) {
		o.r.Pages = nil
	})
}
