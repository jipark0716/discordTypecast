// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jipark0716/discordTypecast/ent/migrate"

	"github.com/jipark0716/discordTypecast/ent/usertypecastsetting"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// UserTypecastSetting is the client for interacting with the UserTypecastSetting builders.
	UserTypecastSetting *UserTypecastSettingClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.UserTypecastSetting = NewUserTypecastSettingClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		UserTypecastSetting: NewUserTypecastSettingClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		UserTypecastSetting: NewUserTypecastSettingClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		UserTypecastSetting.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.UserTypecastSetting.Use(hooks...)
}

// UserTypecastSettingClient is a client for the UserTypecastSetting schema.
type UserTypecastSettingClient struct {
	config
}

// NewUserTypecastSettingClient returns a client for the UserTypecastSetting from the given config.
func NewUserTypecastSettingClient(c config) *UserTypecastSettingClient {
	return &UserTypecastSettingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usertypecastsetting.Hooks(f(g(h())))`.
func (c *UserTypecastSettingClient) Use(hooks ...Hook) {
	c.hooks.UserTypecastSetting = append(c.hooks.UserTypecastSetting, hooks...)
}

// Create returns a builder for creating a UserTypecastSetting entity.
func (c *UserTypecastSettingClient) Create() *UserTypecastSettingCreate {
	mutation := newUserTypecastSettingMutation(c.config, OpCreate)
	return &UserTypecastSettingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserTypecastSetting entities.
func (c *UserTypecastSettingClient) CreateBulk(builders ...*UserTypecastSettingCreate) *UserTypecastSettingCreateBulk {
	return &UserTypecastSettingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserTypecastSetting.
func (c *UserTypecastSettingClient) Update() *UserTypecastSettingUpdate {
	mutation := newUserTypecastSettingMutation(c.config, OpUpdate)
	return &UserTypecastSettingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserTypecastSettingClient) UpdateOne(uts *UserTypecastSetting) *UserTypecastSettingUpdateOne {
	mutation := newUserTypecastSettingMutation(c.config, OpUpdateOne, withUserTypecastSetting(uts))
	return &UserTypecastSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserTypecastSettingClient) UpdateOneID(id int) *UserTypecastSettingUpdateOne {
	mutation := newUserTypecastSettingMutation(c.config, OpUpdateOne, withUserTypecastSettingID(id))
	return &UserTypecastSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserTypecastSetting.
func (c *UserTypecastSettingClient) Delete() *UserTypecastSettingDelete {
	mutation := newUserTypecastSettingMutation(c.config, OpDelete)
	return &UserTypecastSettingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserTypecastSettingClient) DeleteOne(uts *UserTypecastSetting) *UserTypecastSettingDeleteOne {
	return c.DeleteOneID(uts.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserTypecastSettingClient) DeleteOneID(id int) *UserTypecastSettingDeleteOne {
	builder := c.Delete().Where(usertypecastsetting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserTypecastSettingDeleteOne{builder}
}

// Query returns a query builder for UserTypecastSetting.
func (c *UserTypecastSettingClient) Query() *UserTypecastSettingQuery {
	return &UserTypecastSettingQuery{
		config: c.config,
	}
}

// Get returns a UserTypecastSetting entity by its id.
func (c *UserTypecastSettingClient) Get(ctx context.Context, id int) (*UserTypecastSetting, error) {
	return c.Query().Where(usertypecastsetting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserTypecastSettingClient) GetX(ctx context.Context, id int) *UserTypecastSetting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserTypecastSettingClient) Hooks() []Hook {
	return c.hooks.UserTypecastSetting
}
