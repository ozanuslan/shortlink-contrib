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
	"github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mysql/ent/link"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mysql/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeLink = "Link"
)

// LinkMutation represents an operation that mutates the Link nodes in the graph.
type LinkMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	url           *string
	hash          *string
	describe      *string
	json          *v1.Link
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Link, error)
	predicates    []predicate.Link
}

var _ ent.Mutation = (*LinkMutation)(nil)

// linkOption allows management of the mutation configuration using functional options.
type linkOption func(*LinkMutation)

// newLinkMutation creates new mutation for the Link entity.
func newLinkMutation(c config, op Op, opts ...linkOption) *LinkMutation {
	m := &LinkMutation{
		config:        c,
		op:            op,
		typ:           TypeLink,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLinkID sets the ID field of the mutation.
func withLinkID(id uuid.UUID) linkOption {
	return func(m *LinkMutation) {
		var (
			err   error
			once  sync.Once
			value *Link
		)
		m.oldValue = func(ctx context.Context) (*Link, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Link.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLink sets the old Link of the mutation.
func withLink(node *Link) linkOption {
	return func(m *LinkMutation) {
		m.oldValue = func(context.Context) (*Link, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LinkMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LinkMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Link entities.
func (m *LinkMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LinkMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *LinkMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Link.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetURL sets the "url" field.
func (m *LinkMutation) SetURL(s string) {
	m.url = &s
}

// URL returns the value of the "url" field in the mutation.
func (m *LinkMutation) URL() (r string, exists bool) {
	v := m.url
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old "url" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL resets all changes to the "url" field.
func (m *LinkMutation) ResetURL() {
	m.url = nil
}

// SetHash sets the "hash" field.
func (m *LinkMutation) SetHash(s string) {
	m.hash = &s
}

// Hash returns the value of the "hash" field in the mutation.
func (m *LinkMutation) Hash() (r string, exists bool) {
	v := m.hash
	if v == nil {
		return
	}
	return *v, true
}

// OldHash returns the old "hash" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHash: %w", err)
	}
	return oldValue.Hash, nil
}

// ResetHash resets all changes to the "hash" field.
func (m *LinkMutation) ResetHash() {
	m.hash = nil
}

// SetDescribe sets the "describe" field.
func (m *LinkMutation) SetDescribe(s string) {
	m.describe = &s
}

// Describe returns the value of the "describe" field in the mutation.
func (m *LinkMutation) Describe() (r string, exists bool) {
	v := m.describe
	if v == nil {
		return
	}
	return *v, true
}

// OldDescribe returns the old "describe" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldDescribe(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescribe is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescribe requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescribe: %w", err)
	}
	return oldValue.Describe, nil
}

// ClearDescribe clears the value of the "describe" field.
func (m *LinkMutation) ClearDescribe() {
	m.describe = nil
	m.clearedFields[link.FieldDescribe] = struct{}{}
}

// DescribeCleared returns if the "describe" field was cleared in this mutation.
func (m *LinkMutation) DescribeCleared() bool {
	_, ok := m.clearedFields[link.FieldDescribe]
	return ok
}

// ResetDescribe resets all changes to the "describe" field.
func (m *LinkMutation) ResetDescribe() {
	m.describe = nil
	delete(m.clearedFields, link.FieldDescribe)
}

// SetJSON sets the "json" field.
func (m *LinkMutation) SetJSON(v v1.Link) {
	m.json = &v
}

// JSON returns the value of the "json" field in the mutation.
func (m *LinkMutation) JSON() (r v1.Link, exists bool) {
	v := m.json
	if v == nil {
		return
	}
	return *v, true
}

// OldJSON returns the old "json" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldJSON(ctx context.Context) (v v1.Link, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldJSON is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldJSON requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldJSON: %w", err)
	}
	return oldValue.JSON, nil
}

// ResetJSON resets all changes to the "json" field.
func (m *LinkMutation) ResetJSON() {
	m.json = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *LinkMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *LinkMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
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
func (m *LinkMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *LinkMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *LinkMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Link entity.
// If the Link object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LinkMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *LinkMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the LinkMutation builder.
func (m *LinkMutation) Where(ps ...predicate.Link) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the LinkMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *LinkMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Link, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *LinkMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *LinkMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Link).
func (m *LinkMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LinkMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.url != nil {
		fields = append(fields, link.FieldURL)
	}
	if m.hash != nil {
		fields = append(fields, link.FieldHash)
	}
	if m.describe != nil {
		fields = append(fields, link.FieldDescribe)
	}
	if m.json != nil {
		fields = append(fields, link.FieldJSON)
	}
	if m.created_at != nil {
		fields = append(fields, link.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, link.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LinkMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case link.FieldURL:
		return m.URL()
	case link.FieldHash:
		return m.Hash()
	case link.FieldDescribe:
		return m.Describe()
	case link.FieldJSON:
		return m.JSON()
	case link.FieldCreatedAt:
		return m.CreatedAt()
	case link.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LinkMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case link.FieldURL:
		return m.OldURL(ctx)
	case link.FieldHash:
		return m.OldHash(ctx)
	case link.FieldDescribe:
		return m.OldDescribe(ctx)
	case link.FieldJSON:
		return m.OldJSON(ctx)
	case link.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case link.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Link field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LinkMutation) SetField(name string, value ent.Value) error {
	switch name {
	case link.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case link.FieldHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHash(v)
		return nil
	case link.FieldDescribe:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescribe(v)
		return nil
	case link.FieldJSON:
		v, ok := value.(v1.Link)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetJSON(v)
		return nil
	case link.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case link.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Link field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LinkMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LinkMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LinkMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Link numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LinkMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(link.FieldDescribe) {
		fields = append(fields, link.FieldDescribe)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LinkMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LinkMutation) ClearField(name string) error {
	switch name {
	case link.FieldDescribe:
		m.ClearDescribe()
		return nil
	}
	return fmt.Errorf("unknown Link nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LinkMutation) ResetField(name string) error {
	switch name {
	case link.FieldURL:
		m.ResetURL()
		return nil
	case link.FieldHash:
		m.ResetHash()
		return nil
	case link.FieldDescribe:
		m.ResetDescribe()
		return nil
	case link.FieldJSON:
		m.ResetJSON()
		return nil
	case link.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case link.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Link field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LinkMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LinkMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LinkMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LinkMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LinkMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LinkMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LinkMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Link unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LinkMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Link edge %s", name)
}
