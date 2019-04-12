package kinds

import (
	"context"

	"github.com/creativesoftwarefdn/weaviate/database"
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/creativesoftwarefdn/weaviate/validation"
)

// AddAction Class Instance to the connected DB. If the class contains a network
// ref, it has a side-effect on the schema: The schema will be updated to
// include this particular network ref class.
func (m *Manager) AddAction(ctx context.Context, class *models.Action) (*models.Action, error) {
	schemaLock, err := m.db.SchemaLock()
	if err != nil {
		return nil, newErrInternal("could not aquire lock: %v", err)
	}
	defer unlock(schemaLock)
	dbConnector := schemaLock.Connector()

	class.ID = generateUUID()

	err = m.validateAction(ctx, schemaLock, class)
	if err != nil {
		return nil, newErrInvalidUserInput("invalid action: %v", err)
	}

	err = m.addNetworkDataTypesForAction(ctx, schemaLock.SchemaManager(), class)
	if err != nil {
		return nil, newErrInternal("could not update schema for network refs: %v", err)
	}

	dbConnector.AddAction(ctx, class, class.ID)
	if err != nil {
		return nil, newErrInternal("could not store action: %v", err)
	}

	return class, nil
}

// AddThing Class Instance to the connected DB. If the class contains a network
// ref, it has a side-effect on the schema: The schema will be updated to
// include this particular network ref class.
func (m *Manager) AddThing(ctx context.Context, class *models.Thing) (*models.Thing, error) {
	schemaLock, err := m.db.SchemaLock()
	if err != nil {
		return nil, newErrInternal("could not aquire lock: %v", err)
	}
	defer unlock(schemaLock)
	dbConnector := schemaLock.Connector()

	class.ID = generateUUID()

	err = m.validateThing(ctx, schemaLock, class)
	if err != nil {
		return nil, newErrInvalidUserInput("invalid thing: %v", err)
	}

	err = m.addNetworkDataTypesForThing(ctx, schemaLock.SchemaManager(), class)
	if err != nil {
		return nil, newErrInternal("could not update schema for network refs: %v", err)
	}

	dbConnector.AddThing(ctx, class, class.ID)
	if err != nil {
		return nil, newErrInternal("could not store thing: %v", err)
	}

	return class, nil
}

func (m *Manager) validateThing(ctx context.Context, schemaLock database.SchemaLock, class *models.Thing) error {
	// Validate schema given in body with the weaviate schema
	databaseSchema := schema.HackFromDatabaseSchema(schemaLock.GetSchema())
	return validation.ValidateThingBody(
		ctx, class, databaseSchema, schemaLock.Connector(), m.network, m.config)
}

func (m *Manager) validateAction(ctx context.Context, schemaLock database.SchemaLock, class *models.Action) error {
	// Validate schema given in body with the weaviate schema
	databaseSchema := schema.HackFromDatabaseSchema(schemaLock.GetSchema())
	return validation.ValidateActionBody(
		ctx, class, databaseSchema, schemaLock.Connector(), m.network, m.config)
}

func (m *Manager) addNetworkDataTypesForThing(ctx context.Context, sm database.SchemaManager, class *models.Thing) error {
	refSchemaUpdater := newReferenceSchemaUpdater(ctx, sm, m.network, class.Class, kind.THING_KIND)
	return refSchemaUpdater.addNetworkDataTypes(class.Schema)
}

func (m *Manager) addNetworkDataTypesForAction(ctx context.Context, sm database.SchemaManager, class *models.Action) error {
	refSchemaUpdater := newReferenceSchemaUpdater(ctx, sm, m.network, class.Class, kind.ACTION_KIND)
	return refSchemaUpdater.addNetworkDataTypes(class.Schema)
}