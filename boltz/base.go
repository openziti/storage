/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package boltz

import (
	"github.com/openziti/foundation/v2/errorz"
	"github.com/openziti/storage/ast"
	"go.etcd.io/bbolt"
	"reflect"
	"time"
)

const (
	IndexesBucket = "indexes"
)

const (
	FieldId             = "id"
	FieldCreatedAt      = "createdAt"
	FieldUpdatedAt      = "updatedAt"
	FieldTags           = "tags"
	FieldIsSystemEntity = "isSystem"
)

type EntityEventType byte

const (
	EntityCreated EntityEventType = 1
	EntityUpdated EntityEventType = 2
	EntityDeleted EntityEventType = 3

	EntityCreatedAsync EntityEventType = 4
	EntityUpdatedAsync EntityEventType = 5
	EntityDeletedAsync EntityEventType = 6
)

func (self EntityEventType) IsCreate() bool {
	return self == EntityCreated || self == EntityCreatedAsync
}

func (self EntityEventType) IsUpdate() bool {
	return self == EntityUpdated || self == EntityUpdatedAsync
}

func (self EntityEventType) IsDelete() bool {
	return self == EntityDeleted || self == EntityDeletedAsync
}

func (self EntityEventType) IsAsync() bool {
	return self == EntityCreatedAsync || self == EntityUpdatedAsync || self == EntityDeletedAsync
}

// A Checkable can be checked for consistency. This could be an index, an FK or something which contains
// multiple Checkables and can delegate to them
type Checkable interface {
	CheckIntegrity(ctx MutateContext, fix bool, errorSink func(err error, fixed bool)) error
}

// storeInternal contains all the store methods which are only used internally. These are generally called
// from parent to child or vice-versa, and are thus happening through the interface, rather than directly
// from a store to itself.
type storeInternal interface {
	getOrCreateEntitiesBucket(tx *bbolt.Tx) *TypedBucket
	getOrCreateEntityBucket(tx *bbolt.Tx, id []byte) *TypedBucket

	newRowComparator(sort []ast.SortField) (RowComparator, error)
	addSymbol(name string, public bool, symbol EntitySymbol) EntitySymbol
	newEntitySymbol(name string, nodeType ast.NodeType, key string, linkedType Store, prefix ...string) *entitySymbol

	getLinks() map[string]LinkCollection
	inheritMapSymbol(symbol *entityMapSymbol)
	processDeleteConstraints(ctx MutateContext, id string) (entityChangeFlow, error)
	newIndexingContext(isCreate bool, ctx MutateContext, id string, holder errorz.ErrorHolder) *IndexingContext
	newEntityChangeFlow() entityChangeFlow
}

// UntypedEntityChangeState instances are passed to entity event listeners that don't need the concrete entity types
type UntypedEntityChangeState interface {
	GetEventId() string
	GetEntityId() string
	GetCtx() MutateContext
	GetChangeType() EntityEventType
	GetInitialState() Entity
	GetFinalState() Entity
	GetInitialParentEntity() Entity
	GetFinalParentEntity() Entity
	GetStore() Store
	IsParentEvent() bool
}

// Store contains all the methods for interacting with an entity store that don't require knowedge of the concrete
// entity type.
type Store interface {
	storeInternal

	ast.SymbolTypes
	Checkable
	Constrained

	GetEntityType() string
	GetSingularEntityType() string
	GetRootPath() []string
	GetEntitiesBucket(tx *bbolt.Tx) *TypedBucket
	GetEntityBucket(tx *bbolt.Tx, id []byte) *TypedBucket
	IsChildStore() bool
	IsEntityPresent(tx *bbolt.Tx, id string) bool
	IsExtended() bool

	GetSymbol(name string) EntitySymbol

	GetPublicSymbols() []string
	IsPublicSymbol(symbol string) bool

	FindMatching(tx *bbolt.Tx, readIndex SetReadIndex, values []string) []string

	GetLinkCollection(name string) LinkCollection
	GetRefCountedLinkCollection(name string) RefCountedLinkCollection

	GetRelatedEntitiesIdList(tx *bbolt.Tx, id string, field string) []string
	GetRelatedEntitiesCursor(tx *bbolt.Tx, id string, field string, forward bool) ast.SetCursor
	IsEntityRelated(tx *bbolt.Tx, id string, field string, relatedEntityId string) bool

	// QueryIds compiles the query and runs it against the store
	QueryIds(tx *bbolt.Tx, query string) ([]string, int64, error)

	// QueryIdsC executes a compile query against the store
	QueryIdsC(tx *bbolt.Tx, query ast.Query) ([]string, int64, error)

	QueryWithCursorC(tx *bbolt.Tx, cursorProvider ast.SetCursorProvider, query ast.Query) ([]string, int64, error)

	IterateIds(tx *bbolt.Tx, filter ast.BoolNode) ast.SeekableSetCursor

	// IterateValidIds skips non-present entities in extended stores
	IterateValidIds(tx *bbolt.Tx, filter ast.BoolNode) ast.SeekableSetCursor

	GetParentStore() Store
	GrantSymbols(child ConfigurableStore)

	DeleteById(ctx MutateContext, id string) error
	DeleteWhere(ctx MutateContext, query string) error

	AddListener(listener func(Entity), changeType EntityEventType, changeTypes ...EntityEventType)
	AddEntityIdListener(listener func(string), changeType EntityEventType, changeTypes ...EntityEventType)
	AddUntypedEntityConstraint(constraint UntypedEntityConstraint)
	GetEntityReflectType() reflect.Type
}

// ConfigurableStore has all the APIs that store implementations need to configure themselves
type ConfigurableStore interface {
	Store
	MapSymbol(name string, wrapper SymbolMapper)
	AddIdSymbol(name string, nodeType ast.NodeType) EntitySymbol
	AddSymbol(name string, nodeType ast.NodeType, path ...string) EntitySymbol
	AddFkSymbol(name string, linkedType Store, path ...string) EntitySymbol
	AddSymbolWithKey(name string, nodeType ast.NodeType, key string, path ...string) EntitySymbol
	AddFkSymbolWithKey(name string, key string, linkedType Store, path ...string) EntitySymbol
	AddMapSymbol(name string, nodeType ast.NodeType, key string, path ...string)
	AddSetSymbol(name string, nodeType ast.NodeType) EntitySetSymbol
	AddPublicSetSymbol(name string, nodeType ast.NodeType) EntitySetSymbol
	AddFkSetSymbol(name string, linkedType Store) EntitySetSymbol
	NewEntitySymbol(name string, nodeType ast.NodeType) EntitySymbol
	AddEntitySymbol(symbol EntitySymbol)

	AddExtEntitySymbols()
	MakeSymbolPublic(name string)

	AddLinkCollection(local EntitySymbol, remove EntitySymbol) LinkCollection
	AddRefCountedLinkCollection(local EntitySymbol, remove EntitySymbol) RefCountedLinkCollection
}

// ChildStoreStrategy instances are used to allow a parent store to properly delegate to a child store where needed
type ChildStoreStrategy[E Entity] interface {
	HandleUpdate(ctx MutateContext, entity E, checker FieldChecker) (bool, error)
	HandleDelete(ctx MutateContext, entity E) error
	GetStore() Store
}

// EntityConstraint implementations allow reacting to entity changes, both pre and post commit
type EntityConstraint[E Entity] interface {
	ProcessPreCommit(state *EntityChangeState[E]) error
	ProcessPostCommit(state *EntityChangeState[E])
}

// UntypedEntityConstraint instances can react to entity changes in cases where you don't care about the entity type
// or need to react to changes of multiple entity types
type UntypedEntityConstraint interface {
	ProcessPreCommit(state UntypedEntityChangeState) error
	ProcessPostCommit(state UntypedEntityChangeState)
}

// EntityEventListener instances will be notified after an entity change has been committed
type EntityEventListener[E Entity] interface {
	HandleEntityEvent(entity E)
}

// EntityStore extends Store with the methods that need concrete implementation types
type EntityStore[E Entity] interface {
	Store

	RegisterChildStoreStrategy(childStoreStrategy ChildStoreStrategy[E])

	Create(ctx MutateContext, entity E) error
	Update(ctx MutateContext, entity E, checker FieldChecker) error

	// FindById returns the entity for the given Id and true if the eneity exists. If the entity
	// doesn't exist it returns the default value (usually nil) and a false
	FindById(tx *bbolt.Tx, id string) (E, bool, error)

	// LoadById return the entity for the given Id if it exists, otherwise it returns nil and a
	// RecordNotFoundError
	LoadById(tx *bbolt.Tx, id string) (E, error)
	LoadEntity(tx *bbolt.Tx, id string, entity E) (bool, error)

	GetEntityStrategy() EntityStrategy[E]
	AddEntityConstraint(constraint EntityConstraint[E])
	AddEntityEventListener(listener EntityEventListener[E], changeType EntityEventType, changeTypes ...EntityEventType)
	AddEntityEventListenerF(listener func(E), changeType EntityEventType, changeTypes ...EntityEventType)
}

type EntityStrategy[E Entity] interface {
	NewEntity() E
	FillEntity(entity E, bucket *TypedBucket)
	PersistEntity(entity E, ctx *PersistContext)
}

// A PersistContext wraps all the state needed when persisting an entity to a store
type PersistContext struct {
	MutateContext
	Id           string
	Store        Store
	Bucket       *TypedBucket
	FieldChecker FieldChecker
	IsCreate     bool
}

func (ctx *PersistContext) GetParentContext() *PersistContext {
	result := &PersistContext{
		MutateContext: ctx.MutateContext,
		Id:            ctx.Id,
		Store:         ctx.Store.GetParentStore(),
		Bucket:        ctx.Store.GetParentStore().GetEntityBucket(ctx.Bucket.Tx(), []byte(ctx.Id)),
		FieldChecker:  ctx.FieldChecker,
		IsCreate:      ctx.IsCreate,
	}
	// inherit error context
	result.Bucket.ErrorHolderImpl = ctx.Bucket.ErrorHolderImpl
	return result
}

func (ctx *PersistContext) WithFieldOverrides(overrides map[string]string) {
	if ctx.FieldChecker != nil {
		ctx.FieldChecker = NewMappedFieldChecker(ctx.FieldChecker, overrides)
	}
}

func (ctx *PersistContext) GetAndSetString(field string, value string) (*string, bool) {
	return ctx.Bucket.GetAndSetString(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetRequiredString(field string, value string) {
	if ctx.ProceedWithSet(field) {
		if value == "" {
			ctx.Bucket.SetError(errorz.NewFieldError(field+" is required", field, value))
			return
		}
		ctx.Bucket.setTyped(TypeString, field, []byte(value))
	}
}

func (ctx *PersistContext) SetString(field string, value string) {
	ctx.Bucket.SetString(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetStringP(field string, value *string) {
	ctx.Bucket.SetStringP(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetTimeP(field string, value *time.Time) {
	ctx.Bucket.SetTimeP(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetBool(field string, value bool) {
	ctx.Bucket.SetBool(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetInt32(field string, value int32) {
	ctx.Bucket.SetInt32(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetInt64(field string, value int64) {
	ctx.Bucket.SetInt64(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetMap(field string, value map[string]interface{}) {
	ctx.Bucket.PutMap(field, value, ctx.FieldChecker, true)
}

func (ctx *PersistContext) SetStringList(field string, value []string) {
	ctx.Bucket.SetStringList(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) GetAndSetStringList(field string, value []string) ([]string, bool) {
	return ctx.Bucket.GetAndSetStringList(field, value, ctx.FieldChecker)
}

func (ctx *PersistContext) SetLinkedIds(field string, value []string) {
	if ctx.ProceedWithSet(field) {
		collection := ctx.Store.GetLinkCollection(field)
		ctx.Bucket.SetError(collection.SetLinks(ctx.Bucket.Tx(), ctx.Id, value))
	}
}

func (ctx *PersistContext) ProceedWithSet(field string) bool {
	return ctx.Bucket.ProceedWithSet(field, ctx.FieldChecker)
}

// Entity represents the minimal methods needed for an entity
type Entity interface {
	GetId() string
	SetId(id string)
	GetEntityType() string
}

// ExtEntity extends Entity with common additional attributes
type ExtEntity interface {
	Entity
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetTags() map[string]interface{}
	IsSystemEntity() bool

	SetCreatedAt(createdAt time.Time)
	SetUpdatedAt(updatedAt time.Time)
	SetTags(tags map[string]interface{})
}

// NamedExtEntity extends ExtEntity with a Name attribute
type NamedExtEntity interface {
	ExtEntity
	GetName() string
}

func NewExtEntity(id string, tags map[string]interface{}) *BaseExtEntity {
	return &BaseExtEntity{
		Id:   id,
		Tags: tags,
	}
}

type BaseExtEntity struct {
	Id        string                 `json:"id"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
	Tags      map[string]interface{} `json:"tags"`
	IsSystem  bool                   `json:"isSystem"`
	Migrate   bool                   `json:"-"`
}

func (entity *BaseExtEntity) GetId() string {
	return entity.Id
}

func (entity *BaseExtEntity) SetId(id string) {
	entity.Id = id
}

func (entity *BaseExtEntity) GetCreatedAt() time.Time {
	return entity.CreatedAt
}

func (entity *BaseExtEntity) GetUpdatedAt() time.Time {
	return entity.UpdatedAt
}

func (entity *BaseExtEntity) GetTags() map[string]interface{} {
	return entity.Tags
}

func (entity *BaseExtEntity) IsSystemEntity() bool {
	return entity.IsSystem
}

func (entity *BaseExtEntity) SetCreatedAt(createdAt time.Time) {
	entity.CreatedAt = createdAt
}

func (entity *BaseExtEntity) SetUpdatedAt(updatedAt time.Time) {
	entity.UpdatedAt = updatedAt
}

func (entity *BaseExtEntity) SetTags(tags map[string]interface{}) {
	entity.Tags = tags
}

func (entity *BaseExtEntity) LoadBaseValues(bucket *TypedBucket) {
	entity.CreatedAt = bucket.GetTimeOrError(FieldCreatedAt)
	entity.UpdatedAt = bucket.GetTimeOrError(FieldUpdatedAt)
	entity.Tags = bucket.GetMap(FieldTags)
	entity.IsSystem = bucket.GetBoolWithDefault(FieldIsSystemEntity, false)
}

func (entity *BaseExtEntity) SetBaseValues(ctx *PersistContext) {
	if ctx.IsCreate {
		entity.CreateBaseValues(ctx)
	} else {
		entity.UpdateBaseValues(ctx)
	}
}

func (entity *BaseExtEntity) CreateBaseValues(ctx *PersistContext) {
	now := time.Now()
	if entity.Migrate {
		ctx.Bucket.SetTimeP(FieldCreatedAt, &entity.CreatedAt, nil)
		ctx.Bucket.SetTimeP(FieldUpdatedAt, &entity.UpdatedAt, nil)
	} else {
		ctx.Bucket.SetTimeP(FieldCreatedAt, &now, nil)
		ctx.Bucket.SetTimeP(FieldUpdatedAt, &now, nil)
	}
	ctx.Bucket.PutMap(FieldTags, entity.Tags, nil, false)
	if entity.IsSystem {
		ctx.Bucket.SetBool(FieldIsSystemEntity, true, nil)
	}
}

func (entity *BaseExtEntity) UpdateBaseValues(ctx *PersistContext) {
	now := time.Now()
	ctx.Bucket.SetTimeP(FieldUpdatedAt, &now, nil)
	ctx.Bucket.PutMap(FieldTags, entity.Tags, ctx.FieldChecker, false)
}
