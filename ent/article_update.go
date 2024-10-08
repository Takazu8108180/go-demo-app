// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/takazu8108180/go-demo-app/ent/article"
	"github.com/takazu8108180/go-demo-app/ent/comment"
	"github.com/takazu8108180/go-demo-app/ent/predicate"
)

// ArticleUpdate is the builder for updating Article entities.
type ArticleUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleMutation
}

// Where appends a list predicates to the ArticleUpdate builder.
func (au *ArticleUpdate) Where(ps ...predicate.Article) *ArticleUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetTitle sets the "title" field.
func (au *ArticleUpdate) SetTitle(s string) *ArticleUpdate {
	au.mutation.SetTitle(s)
	return au
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableTitle(s *string) *ArticleUpdate {
	if s != nil {
		au.SetTitle(*s)
	}
	return au
}

// SetContents sets the "contents" field.
func (au *ArticleUpdate) SetContents(s string) *ArticleUpdate {
	au.mutation.SetContents(s)
	return au
}

// SetNillableContents sets the "contents" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableContents(s *string) *ArticleUpdate {
	if s != nil {
		au.SetContents(*s)
	}
	return au
}

// SetUsername sets the "username" field.
func (au *ArticleUpdate) SetUsername(s string) *ArticleUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableUsername(s *string) *ArticleUpdate {
	if s != nil {
		au.SetUsername(*s)
	}
	return au
}

// SetNice sets the "nice" field.
func (au *ArticleUpdate) SetNice(i int) *ArticleUpdate {
	au.mutation.ResetNice()
	au.mutation.SetNice(i)
	return au
}

// SetNillableNice sets the "nice" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableNice(i *int) *ArticleUpdate {
	if i != nil {
		au.SetNice(*i)
	}
	return au
}

// AddNice adds i to the "nice" field.
func (au *ArticleUpdate) AddNice(i int) *ArticleUpdate {
	au.mutation.AddNice(i)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *ArticleUpdate) SetCreatedAt(t time.Time) *ArticleUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *ArticleUpdate) SetNillableCreatedAt(t *time.Time) *ArticleUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// AddCommentIDs adds the "comment" edge to the Comment entity by IDs.
func (au *ArticleUpdate) AddCommentIDs(ids ...uuid.UUID) *ArticleUpdate {
	au.mutation.AddCommentIDs(ids...)
	return au
}

// AddComment adds the "comment" edges to the Comment entity.
func (au *ArticleUpdate) AddComment(c ...*Comment) *ArticleUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.AddCommentIDs(ids...)
}

// Mutation returns the ArticleMutation object of the builder.
func (au *ArticleUpdate) Mutation() *ArticleMutation {
	return au.mutation
}

// ClearComment clears all "comment" edges to the Comment entity.
func (au *ArticleUpdate) ClearComment() *ArticleUpdate {
	au.mutation.ClearComment()
	return au
}

// RemoveCommentIDs removes the "comment" edge to Comment entities by IDs.
func (au *ArticleUpdate) RemoveCommentIDs(ids ...uuid.UUID) *ArticleUpdate {
	au.mutation.RemoveCommentIDs(ids...)
	return au
}

// RemoveComment removes "comment" edges to Comment entities.
func (au *ArticleUpdate) RemoveComment(c ...*Comment) *ArticleUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return au.RemoveCommentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArticleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArticleUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArticleUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ArticleUpdate) check() error {
	if v, ok := au.mutation.Title(); ok {
		if err := article.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Article.title": %w`, err)}
		}
	}
	if v, ok := au.mutation.Username(); ok {
		if err := article.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Article.username": %w`, err)}
		}
	}
	if v, ok := au.mutation.Nice(); ok {
		if err := article.NiceValidator(v); err != nil {
			return &ValidationError{Name: "nice", err: fmt.Errorf(`ent: validator failed for field "Article.nice": %w`, err)}
		}
	}
	return nil
}

func (au *ArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
	}
	if value, ok := au.mutation.Contents(); ok {
		_spec.SetField(article.FieldContents, field.TypeString, value)
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.SetField(article.FieldUsername, field.TypeString, value)
	}
	if value, ok := au.mutation.Nice(); ok {
		_spec.SetField(article.FieldNice, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedNice(); ok {
		_spec.AddField(article.FieldNice, field.TypeInt, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.CommentTable,
			Columns: []string{article.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedCommentIDs(); len(nodes) > 0 && !au.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.CommentTable,
			Columns: []string{article.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.CommentTable,
			Columns: []string{article.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArticleUpdateOne is the builder for updating a single Article entity.
type ArticleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleMutation
}

// SetTitle sets the "title" field.
func (auo *ArticleUpdateOne) SetTitle(s string) *ArticleUpdateOne {
	auo.mutation.SetTitle(s)
	return auo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableTitle(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetTitle(*s)
	}
	return auo
}

// SetContents sets the "contents" field.
func (auo *ArticleUpdateOne) SetContents(s string) *ArticleUpdateOne {
	auo.mutation.SetContents(s)
	return auo
}

// SetNillableContents sets the "contents" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableContents(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetContents(*s)
	}
	return auo
}

// SetUsername sets the "username" field.
func (auo *ArticleUpdateOne) SetUsername(s string) *ArticleUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableUsername(s *string) *ArticleUpdateOne {
	if s != nil {
		auo.SetUsername(*s)
	}
	return auo
}

// SetNice sets the "nice" field.
func (auo *ArticleUpdateOne) SetNice(i int) *ArticleUpdateOne {
	auo.mutation.ResetNice()
	auo.mutation.SetNice(i)
	return auo
}

// SetNillableNice sets the "nice" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableNice(i *int) *ArticleUpdateOne {
	if i != nil {
		auo.SetNice(*i)
	}
	return auo
}

// AddNice adds i to the "nice" field.
func (auo *ArticleUpdateOne) AddNice(i int) *ArticleUpdateOne {
	auo.mutation.AddNice(i)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *ArticleUpdateOne) SetCreatedAt(t time.Time) *ArticleUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *ArticleUpdateOne) SetNillableCreatedAt(t *time.Time) *ArticleUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// AddCommentIDs adds the "comment" edge to the Comment entity by IDs.
func (auo *ArticleUpdateOne) AddCommentIDs(ids ...uuid.UUID) *ArticleUpdateOne {
	auo.mutation.AddCommentIDs(ids...)
	return auo
}

// AddComment adds the "comment" edges to the Comment entity.
func (auo *ArticleUpdateOne) AddComment(c ...*Comment) *ArticleUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.AddCommentIDs(ids...)
}

// Mutation returns the ArticleMutation object of the builder.
func (auo *ArticleUpdateOne) Mutation() *ArticleMutation {
	return auo.mutation
}

// ClearComment clears all "comment" edges to the Comment entity.
func (auo *ArticleUpdateOne) ClearComment() *ArticleUpdateOne {
	auo.mutation.ClearComment()
	return auo
}

// RemoveCommentIDs removes the "comment" edge to Comment entities by IDs.
func (auo *ArticleUpdateOne) RemoveCommentIDs(ids ...uuid.UUID) *ArticleUpdateOne {
	auo.mutation.RemoveCommentIDs(ids...)
	return auo
}

// RemoveComment removes "comment" edges to Comment entities.
func (auo *ArticleUpdateOne) RemoveComment(c ...*Comment) *ArticleUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return auo.RemoveCommentIDs(ids...)
}

// Where appends a list predicates to the ArticleUpdate builder.
func (auo *ArticleUpdateOne) Where(ps ...predicate.Article) *ArticleUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArticleUpdateOne) Select(field string, fields ...string) *ArticleUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Article entity.
func (auo *ArticleUpdateOne) Save(ctx context.Context) (*Article, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArticleUpdateOne) SaveX(ctx context.Context) *Article {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArticleUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ArticleUpdateOne) check() error {
	if v, ok := auo.mutation.Title(); ok {
		if err := article.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Article.title": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Username(); ok {
		if err := article.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Article.username": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Nice(); ok {
		if err := article.NiceValidator(v); err != nil {
			return &ValidationError{Name: "nice", err: fmt.Errorf(`ent: validator failed for field "Article.nice": %w`, err)}
		}
	}
	return nil
}

func (auo *ArticleUpdateOne) sqlSave(ctx context.Context) (_node *Article, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(article.Table, article.Columns, sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Article.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for _, f := range fields {
			if !article.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != article.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
	}
	if value, ok := auo.mutation.Contents(); ok {
		_spec.SetField(article.FieldContents, field.TypeString, value)
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.SetField(article.FieldUsername, field.TypeString, value)
	}
	if value, ok := auo.mutation.Nice(); ok {
		_spec.SetField(article.FieldNice, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedNice(); ok {
		_spec.AddField(article.FieldNice, field.TypeInt, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.CommentTable,
			Columns: []string{article.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedCommentIDs(); len(nodes) > 0 && !auo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.CommentTable,
			Columns: []string{article.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   article.CommentTable,
			Columns: []string{article.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Article{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{article.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
