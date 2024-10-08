// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/takazu8108180/go-demo-app/ent/article"
	"github.com/takazu8108180/go-demo-app/ent/comment"
	"github.com/takazu8108180/go-demo-app/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeArticle = "Article"
	TypeComment = "Comment"
)

// ArticleMutation represents an operation that mutates the Article nodes in the graph.
type ArticleMutation struct {
	config
	op             Op
	typ            string
	id             *uuid.UUID
	title          *string
	contents       *string
	username       *string
	nice           *int
	addnice        *int
	created_at     *time.Time
	clearedFields  map[string]struct{}
	comment        map[uuid.UUID]struct{}
	removedcomment map[uuid.UUID]struct{}
	clearedcomment bool
	done           bool
	oldValue       func(context.Context) (*Article, error)
	predicates     []predicate.Article
}

var _ ent.Mutation = (*ArticleMutation)(nil)

// articleOption allows management of the mutation configuration using functional options.
type articleOption func(*ArticleMutation)

// newArticleMutation creates new mutation for the Article entity.
func newArticleMutation(c config, op Op, opts ...articleOption) *ArticleMutation {
	m := &ArticleMutation{
		config:        c,
		op:            op,
		typ:           TypeArticle,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withArticleID sets the ID field of the mutation.
func withArticleID(id uuid.UUID) articleOption {
	return func(m *ArticleMutation) {
		var (
			err   error
			once  sync.Once
			value *Article
		)
		m.oldValue = func(ctx context.Context) (*Article, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Article.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withArticle sets the old Article of the mutation.
func withArticle(node *Article) articleOption {
	return func(m *ArticleMutation) {
		m.oldValue = func(context.Context) (*Article, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ArticleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ArticleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Article entities.
func (m *ArticleMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ArticleMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ArticleMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Article.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *ArticleMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *ArticleMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *ArticleMutation) ResetTitle() {
	m.title = nil
}

// SetContents sets the "contents" field.
func (m *ArticleMutation) SetContents(s string) {
	m.contents = &s
}

// Contents returns the value of the "contents" field in the mutation.
func (m *ArticleMutation) Contents() (r string, exists bool) {
	v := m.contents
	if v == nil {
		return
	}
	return *v, true
}

// OldContents returns the old "contents" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldContents(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContents is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContents requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContents: %w", err)
	}
	return oldValue.Contents, nil
}

// ResetContents resets all changes to the "contents" field.
func (m *ArticleMutation) ResetContents() {
	m.contents = nil
}

// SetUsername sets the "username" field.
func (m *ArticleMutation) SetUsername(s string) {
	m.username = &s
}

// Username returns the value of the "username" field in the mutation.
func (m *ArticleMutation) Username() (r string, exists bool) {
	v := m.username
	if v == nil {
		return
	}
	return *v, true
}

// OldUsername returns the old "username" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldUsername(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUsername is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUsername requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUsername: %w", err)
	}
	return oldValue.Username, nil
}

// ResetUsername resets all changes to the "username" field.
func (m *ArticleMutation) ResetUsername() {
	m.username = nil
}

// SetNice sets the "nice" field.
func (m *ArticleMutation) SetNice(i int) {
	m.nice = &i
	m.addnice = nil
}

// Nice returns the value of the "nice" field in the mutation.
func (m *ArticleMutation) Nice() (r int, exists bool) {
	v := m.nice
	if v == nil {
		return
	}
	return *v, true
}

// OldNice returns the old "nice" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldNice(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNice is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNice requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNice: %w", err)
	}
	return oldValue.Nice, nil
}

// AddNice adds i to the "nice" field.
func (m *ArticleMutation) AddNice(i int) {
	if m.addnice != nil {
		*m.addnice += i
	} else {
		m.addnice = &i
	}
}

// AddedNice returns the value that was added to the "nice" field in this mutation.
func (m *ArticleMutation) AddedNice() (r int, exists bool) {
	v := m.addnice
	if v == nil {
		return
	}
	return *v, true
}

// ResetNice resets all changes to the "nice" field.
func (m *ArticleMutation) ResetNice() {
	m.nice = nil
	m.addnice = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ArticleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ArticleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ArticleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// AddCommentIDs adds the "comment" edge to the Comment entity by ids.
func (m *ArticleMutation) AddCommentIDs(ids ...uuid.UUID) {
	if m.comment == nil {
		m.comment = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.comment[ids[i]] = struct{}{}
	}
}

// ClearComment clears the "comment" edge to the Comment entity.
func (m *ArticleMutation) ClearComment() {
	m.clearedcomment = true
}

// CommentCleared reports if the "comment" edge to the Comment entity was cleared.
func (m *ArticleMutation) CommentCleared() bool {
	return m.clearedcomment
}

// RemoveCommentIDs removes the "comment" edge to the Comment entity by IDs.
func (m *ArticleMutation) RemoveCommentIDs(ids ...uuid.UUID) {
	if m.removedcomment == nil {
		m.removedcomment = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.comment, ids[i])
		m.removedcomment[ids[i]] = struct{}{}
	}
}

// RemovedComment returns the removed IDs of the "comment" edge to the Comment entity.
func (m *ArticleMutation) RemovedCommentIDs() (ids []uuid.UUID) {
	for id := range m.removedcomment {
		ids = append(ids, id)
	}
	return
}

// CommentIDs returns the "comment" edge IDs in the mutation.
func (m *ArticleMutation) CommentIDs() (ids []uuid.UUID) {
	for id := range m.comment {
		ids = append(ids, id)
	}
	return
}

// ResetComment resets all changes to the "comment" edge.
func (m *ArticleMutation) ResetComment() {
	m.comment = nil
	m.clearedcomment = false
	m.removedcomment = nil
}

// Where appends a list predicates to the ArticleMutation builder.
func (m *ArticleMutation) Where(ps ...predicate.Article) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ArticleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ArticleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Article, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ArticleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ArticleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Article).
func (m *ArticleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ArticleMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.title != nil {
		fields = append(fields, article.FieldTitle)
	}
	if m.contents != nil {
		fields = append(fields, article.FieldContents)
	}
	if m.username != nil {
		fields = append(fields, article.FieldUsername)
	}
	if m.nice != nil {
		fields = append(fields, article.FieldNice)
	}
	if m.created_at != nil {
		fields = append(fields, article.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ArticleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case article.FieldTitle:
		return m.Title()
	case article.FieldContents:
		return m.Contents()
	case article.FieldUsername:
		return m.Username()
	case article.FieldNice:
		return m.Nice()
	case article.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ArticleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case article.FieldTitle:
		return m.OldTitle(ctx)
	case article.FieldContents:
		return m.OldContents(ctx)
	case article.FieldUsername:
		return m.OldUsername(ctx)
	case article.FieldNice:
		return m.OldNice(ctx)
	case article.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Article field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case article.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case article.FieldContents:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContents(v)
		return nil
	case article.FieldUsername:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUsername(v)
		return nil
	case article.FieldNice:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNice(v)
		return nil
	case article.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ArticleMutation) AddedFields() []string {
	var fields []string
	if m.addnice != nil {
		fields = append(fields, article.FieldNice)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ArticleMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case article.FieldNice:
		return m.AddedNice()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) AddField(name string, value ent.Value) error {
	switch name {
	case article.FieldNice:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddNice(v)
		return nil
	}
	return fmt.Errorf("unknown Article numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ArticleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ArticleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ArticleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Article nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ArticleMutation) ResetField(name string) error {
	switch name {
	case article.FieldTitle:
		m.ResetTitle()
		return nil
	case article.FieldContents:
		m.ResetContents()
		return nil
	case article.FieldUsername:
		m.ResetUsername()
		return nil
	case article.FieldNice:
		m.ResetNice()
		return nil
	case article.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ArticleMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.comment != nil {
		edges = append(edges, article.EdgeComment)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ArticleMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case article.EdgeComment:
		ids := make([]ent.Value, 0, len(m.comment))
		for id := range m.comment {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ArticleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedcomment != nil {
		edges = append(edges, article.EdgeComment)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ArticleMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case article.EdgeComment:
		ids := make([]ent.Value, 0, len(m.removedcomment))
		for id := range m.removedcomment {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ArticleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedcomment {
		edges = append(edges, article.EdgeComment)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ArticleMutation) EdgeCleared(name string) bool {
	switch name {
	case article.EdgeComment:
		return m.clearedcomment
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ArticleMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Article unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ArticleMutation) ResetEdge(name string) error {
	switch name {
	case article.EdgeComment:
		m.ResetComment()
		return nil
	}
	return fmt.Errorf("unknown Article edge %s", name)
}

// CommentMutation represents an operation that mutates the Comment nodes in the graph.
type CommentMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	message       *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Comment, error)
	predicates    []predicate.Comment
}

var _ ent.Mutation = (*CommentMutation)(nil)

// commentOption allows management of the mutation configuration using functional options.
type commentOption func(*CommentMutation)

// newCommentMutation creates new mutation for the Comment entity.
func newCommentMutation(c config, op Op, opts ...commentOption) *CommentMutation {
	m := &CommentMutation{
		config:        c,
		op:            op,
		typ:           TypeComment,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCommentID sets the ID field of the mutation.
func withCommentID(id uuid.UUID) commentOption {
	return func(m *CommentMutation) {
		var (
			err   error
			once  sync.Once
			value *Comment
		)
		m.oldValue = func(ctx context.Context) (*Comment, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Comment.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withComment sets the old Comment of the mutation.
func withComment(node *Comment) commentOption {
	return func(m *CommentMutation) {
		m.oldValue = func(context.Context) (*Comment, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CommentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CommentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Comment entities.
func (m *CommentMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CommentMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CommentMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Comment.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetMessage sets the "message" field.
func (m *CommentMutation) SetMessage(s string) {
	m.message = &s
}

// Message returns the value of the "message" field in the mutation.
func (m *CommentMutation) Message() (r string, exists bool) {
	v := m.message
	if v == nil {
		return
	}
	return *v, true
}

// OldMessage returns the old "message" field's value of the Comment entity.
// If the Comment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommentMutation) OldMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessage: %w", err)
	}
	return oldValue.Message, nil
}

// ResetMessage resets all changes to the "message" field.
func (m *CommentMutation) ResetMessage() {
	m.message = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *CommentMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *CommentMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Comment entity.
// If the Comment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommentMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *CommentMutation) ResetCreatedAt() {
	m.created_at = nil
}

// Where appends a list predicates to the CommentMutation builder.
func (m *CommentMutation) Where(ps ...predicate.Comment) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the CommentMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *CommentMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Comment, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *CommentMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *CommentMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Comment).
func (m *CommentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CommentMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.message != nil {
		fields = append(fields, comment.FieldMessage)
	}
	if m.created_at != nil {
		fields = append(fields, comment.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CommentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case comment.FieldMessage:
		return m.Message()
	case comment.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CommentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case comment.FieldMessage:
		return m.OldMessage(ctx)
	case comment.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Comment field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CommentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case comment.FieldMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessage(v)
		return nil
	case comment.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Comment field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CommentMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CommentMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CommentMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Comment numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CommentMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CommentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CommentMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Comment nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CommentMutation) ResetField(name string) error {
	switch name {
	case comment.FieldMessage:
		m.ResetMessage()
		return nil
	case comment.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Comment field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CommentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CommentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CommentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CommentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CommentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CommentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CommentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Comment unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CommentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Comment edge %s", name)
}
