// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rustic-beans/spotify-viewer/ent/album"
	"github.com/rustic-beans/spotify-viewer/ent/artist"
	"github.com/rustic-beans/spotify-viewer/ent/schema"
	"github.com/rustic-beans/spotify-viewer/ent/track"
)

// ArtistCreate is the builder for creating a Artist entity.
type ArtistCreate struct {
	config
	mutation *ArtistMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetExternalUrls sets the "external_urls" field.
func (ac *ArtistCreate) SetExternalUrls(sm *schema.StringMap) *ArtistCreate {
	ac.mutation.SetExternalUrls(sm)
	return ac
}

// SetHref sets the "href" field.
func (ac *ArtistCreate) SetHref(s string) *ArtistCreate {
	ac.mutation.SetHref(s)
	return ac
}

// SetName sets the "name" field.
func (ac *ArtistCreate) SetName(s string) *ArtistCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetURI sets the "uri" field.
func (ac *ArtistCreate) SetURI(s string) *ArtistCreate {
	ac.mutation.SetURI(s)
	return ac
}

// SetID sets the "id" field.
func (ac *ArtistCreate) SetID(s string) *ArtistCreate {
	ac.mutation.SetID(s)
	return ac
}

// AddAlbumIDs adds the "albums" edge to the Album entity by IDs.
func (ac *ArtistCreate) AddAlbumIDs(ids ...string) *ArtistCreate {
	ac.mutation.AddAlbumIDs(ids...)
	return ac
}

// AddAlbums adds the "albums" edges to the Album entity.
func (ac *ArtistCreate) AddAlbums(a ...*Album) *ArtistCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAlbumIDs(ids...)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (ac *ArtistCreate) AddTrackIDs(ids ...string) *ArtistCreate {
	ac.mutation.AddTrackIDs(ids...)
	return ac
}

// AddTracks adds the "tracks" edges to the Track entity.
func (ac *ArtistCreate) AddTracks(t ...*Track) *ArtistCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTrackIDs(ids...)
}

// Mutation returns the ArtistMutation object of the builder.
func (ac *ArtistCreate) Mutation() *ArtistMutation {
	return ac.mutation
}

// Save creates the Artist in the database.
func (ac *ArtistCreate) Save(ctx context.Context) (*Artist, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArtistCreate) SaveX(ctx context.Context) *Artist {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArtistCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArtistCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArtistCreate) check() error {
	if _, ok := ac.mutation.ExternalUrls(); !ok {
		return &ValidationError{Name: "external_urls", err: errors.New(`ent: missing required field "Artist.external_urls"`)}
	}
	if _, ok := ac.mutation.Href(); !ok {
		return &ValidationError{Name: "href", err: errors.New(`ent: missing required field "Artist.href"`)}
	}
	if v, ok := ac.mutation.Href(); ok {
		if err := artist.HrefValidator(v); err != nil {
			return &ValidationError{Name: "href", err: fmt.Errorf(`ent: validator failed for field "Artist.href": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Artist.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := artist.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Artist.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.URI(); !ok {
		return &ValidationError{Name: "uri", err: errors.New(`ent: missing required field "Artist.uri"`)}
	}
	if v, ok := ac.mutation.URI(); ok {
		if err := artist.URIValidator(v); err != nil {
			return &ValidationError{Name: "uri", err: fmt.Errorf(`ent: validator failed for field "Artist.uri": %w`, err)}
		}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := artist.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Artist.id": %w`, err)}
		}
	}
	return nil
}

func (ac *ArtistCreate) sqlSave(ctx context.Context) (*Artist, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Artist.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArtistCreate) createSpec() (*Artist, *sqlgraph.CreateSpec) {
	var (
		_node = &Artist{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(artist.Table, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeString))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.ExternalUrls(); ok {
		_spec.SetField(artist.FieldExternalUrls, field.TypeJSON, value)
		_node.ExternalUrls = value
	}
	if value, ok := ac.mutation.Href(); ok {
		_spec.SetField(artist.FieldHref, field.TypeString, value)
		_node.Href = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(artist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.URI(); ok {
		_spec.SetField(artist.FieldURI, field.TypeString, value)
		_node.URI = value
	}
	if nodes := ac.mutation.AlbumsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AlbumsTable,
			Columns: artist.AlbumsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(album.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.TracksTable,
			Columns: artist.TracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Artist.Create().
//		SetExternalUrls(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArtistUpsert) {
//			SetExternalUrls(v+v).
//		}).
//		Exec(ctx)
func (ac *ArtistCreate) OnConflict(opts ...sql.ConflictOption) *ArtistUpsertOne {
	ac.conflict = opts
	return &ArtistUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *ArtistCreate) OnConflictColumns(columns ...string) *ArtistUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &ArtistUpsertOne{
		create: ac,
	}
}

type (
	// ArtistUpsertOne is the builder for "upsert"-ing
	//  one Artist node.
	ArtistUpsertOne struct {
		create *ArtistCreate
	}

	// ArtistUpsert is the "OnConflict" setter.
	ArtistUpsert struct {
		*sql.UpdateSet
	}
)

// SetExternalUrls sets the "external_urls" field.
func (u *ArtistUpsert) SetExternalUrls(v *schema.StringMap) *ArtistUpsert {
	u.Set(artist.FieldExternalUrls, v)
	return u
}

// UpdateExternalUrls sets the "external_urls" field to the value that was provided on create.
func (u *ArtistUpsert) UpdateExternalUrls() *ArtistUpsert {
	u.SetExcluded(artist.FieldExternalUrls)
	return u
}

// SetHref sets the "href" field.
func (u *ArtistUpsert) SetHref(v string) *ArtistUpsert {
	u.Set(artist.FieldHref, v)
	return u
}

// UpdateHref sets the "href" field to the value that was provided on create.
func (u *ArtistUpsert) UpdateHref() *ArtistUpsert {
	u.SetExcluded(artist.FieldHref)
	return u
}

// SetName sets the "name" field.
func (u *ArtistUpsert) SetName(v string) *ArtistUpsert {
	u.Set(artist.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ArtistUpsert) UpdateName() *ArtistUpsert {
	u.SetExcluded(artist.FieldName)
	return u
}

// SetURI sets the "uri" field.
func (u *ArtistUpsert) SetURI(v string) *ArtistUpsert {
	u.Set(artist.FieldURI, v)
	return u
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *ArtistUpsert) UpdateURI() *ArtistUpsert {
	u.SetExcluded(artist.FieldURI)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(artist.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ArtistUpsertOne) UpdateNewValues() *ArtistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(artist.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Artist.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ArtistUpsertOne) Ignore() *ArtistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArtistUpsertOne) DoNothing() *ArtistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArtistCreate.OnConflict
// documentation for more info.
func (u *ArtistUpsertOne) Update(set func(*ArtistUpsert)) *ArtistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArtistUpsert{UpdateSet: update})
	}))
	return u
}

// SetExternalUrls sets the "external_urls" field.
func (u *ArtistUpsertOne) SetExternalUrls(v *schema.StringMap) *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.SetExternalUrls(v)
	})
}

// UpdateExternalUrls sets the "external_urls" field to the value that was provided on create.
func (u *ArtistUpsertOne) UpdateExternalUrls() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateExternalUrls()
	})
}

// SetHref sets the "href" field.
func (u *ArtistUpsertOne) SetHref(v string) *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.SetHref(v)
	})
}

// UpdateHref sets the "href" field to the value that was provided on create.
func (u *ArtistUpsertOne) UpdateHref() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateHref()
	})
}

// SetName sets the "name" field.
func (u *ArtistUpsertOne) SetName(v string) *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ArtistUpsertOne) UpdateName() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateName()
	})
}

// SetURI sets the "uri" field.
func (u *ArtistUpsertOne) SetURI(v string) *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *ArtistUpsertOne) UpdateURI() *ArtistUpsertOne {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateURI()
	})
}

// Exec executes the query.
func (u *ArtistUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArtistCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArtistUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ArtistUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ArtistUpsertOne.ID is not supported by MySQL driver. Use ArtistUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ArtistUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ArtistCreateBulk is the builder for creating many Artist entities in bulk.
type ArtistCreateBulk struct {
	config
	err      error
	builders []*ArtistCreate
	conflict []sql.ConflictOption
}

// Save creates the Artist entities in the database.
func (acb *ArtistCreateBulk) Save(ctx context.Context) ([]*Artist, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Artist, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArtistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArtistCreateBulk) SaveX(ctx context.Context) []*Artist {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArtistCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArtistCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Artist.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ArtistUpsert) {
//			SetExternalUrls(v+v).
//		}).
//		Exec(ctx)
func (acb *ArtistCreateBulk) OnConflict(opts ...sql.ConflictOption) *ArtistUpsertBulk {
	acb.conflict = opts
	return &ArtistUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *ArtistCreateBulk) OnConflictColumns(columns ...string) *ArtistUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &ArtistUpsertBulk{
		create: acb,
	}
}

// ArtistUpsertBulk is the builder for "upsert"-ing
// a bulk of Artist nodes.
type ArtistUpsertBulk struct {
	create *ArtistCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(artist.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ArtistUpsertBulk) UpdateNewValues() *ArtistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(artist.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Artist.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ArtistUpsertBulk) Ignore() *ArtistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ArtistUpsertBulk) DoNothing() *ArtistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ArtistCreateBulk.OnConflict
// documentation for more info.
func (u *ArtistUpsertBulk) Update(set func(*ArtistUpsert)) *ArtistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ArtistUpsert{UpdateSet: update})
	}))
	return u
}

// SetExternalUrls sets the "external_urls" field.
func (u *ArtistUpsertBulk) SetExternalUrls(v *schema.StringMap) *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.SetExternalUrls(v)
	})
}

// UpdateExternalUrls sets the "external_urls" field to the value that was provided on create.
func (u *ArtistUpsertBulk) UpdateExternalUrls() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateExternalUrls()
	})
}

// SetHref sets the "href" field.
func (u *ArtistUpsertBulk) SetHref(v string) *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.SetHref(v)
	})
}

// UpdateHref sets the "href" field to the value that was provided on create.
func (u *ArtistUpsertBulk) UpdateHref() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateHref()
	})
}

// SetName sets the "name" field.
func (u *ArtistUpsertBulk) SetName(v string) *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ArtistUpsertBulk) UpdateName() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateName()
	})
}

// SetURI sets the "uri" field.
func (u *ArtistUpsertBulk) SetURI(v string) *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.SetURI(v)
	})
}

// UpdateURI sets the "uri" field to the value that was provided on create.
func (u *ArtistUpsertBulk) UpdateURI() *ArtistUpsertBulk {
	return u.Update(func(s *ArtistUpsert) {
		s.UpdateURI()
	})
}

// Exec executes the query.
func (u *ArtistUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ArtistCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ArtistCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ArtistUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
