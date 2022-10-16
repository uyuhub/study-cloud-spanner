// Code generated by yo. DO NOT EDIT.
// Package models contains the types.
package models

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
)

// User represents a row from 'users'.
type User struct {
	UserID string            `spanner:"user_id" json:"user_id"` // user_id
	Age    spanner.NullInt64 `spanner:"age" json:"age"`         // age
	Name   string            `spanner:"name" json:"name"`       // name
}

func UserPrimaryKeys() []string {
	return []string{
		"user_id",
	}
}

func UserColumns() []string {
	return []string{
		"user_id",
		"age",
		"name",
	}
}

func (u *User) columnsToPtrs(cols []string, customPtrs map[string]interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		if val, ok := customPtrs[col]; ok {
			ret = append(ret, val)
			continue
		}

		switch col {
		case "user_id":
			ret = append(ret, &u.UserID)
		case "age":
			ret = append(ret, &u.Age)
		case "name":
			ret = append(ret, &u.Name)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (u *User) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "user_id":
			ret = append(ret, u.UserID)
		case "age":
			ret = append(ret, u.Age)
		case "name":
			ret = append(ret, u.Name)
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newUser_Decoder returns a decoder which reads a row from *spanner.Row
// into User. The decoder is not goroutine-safe. Don't use it concurrently.
func newUser_Decoder(cols []string) func(*spanner.Row) (*User, error) {
	customPtrs := map[string]interface{}{}

	return func(row *spanner.Row) (*User, error) {
		var u User
		ptrs, err := u.columnsToPtrs(cols, customPtrs)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &u, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (u *User) Insert(ctx context.Context) *spanner.Mutation {
	return spanner.Insert("users", UserColumns(), []interface{}{
		u.UserID, u.Age, u.Name,
	})
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (u *User) Update(ctx context.Context) *spanner.Mutation {
	return spanner.Update("users", UserColumns(), []interface{}{
		u.UserID, u.Age, u.Name,
	})
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (u *User) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	return spanner.InsertOrUpdate("users", UserColumns(), []interface{}{
		u.UserID, u.Age, u.Name,
	})
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (u *User) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, UserPrimaryKeys()...)

	values, err := u.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "User.UpdateColumns", "users", err)
	}

	return spanner.Update("users", colsWithPKeys, values), nil
}

// FindUser gets a User by primary key
func FindUser(ctx context.Context, db YORODB, userID string) (*User, error) {
	key := spanner.Key{userID}
	row, err := db.ReadRow(ctx, "users", key, UserColumns())
	if err != nil {
		return nil, newError("FindUser", "users", err)
	}

	decoder := newUser_Decoder(UserColumns())
	u, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindUser", "users", err)
	}

	return u, nil
}

// ReadUser retrieves multiples rows from User by KeySet as a slice.
func ReadUser(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*User, error) {
	var res []*User

	decoder := newUser_Decoder(UserColumns())

	rows := db.Read(ctx, "users", keys, UserColumns())
	err := rows.Do(func(row *spanner.Row) error {
		u, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, u)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadUser", "users", err)
	}

	return res, nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := u.columnsToValues(UserPrimaryKeys())
	return spanner.Delete("users", spanner.Key(values))
}

// FindUsersByName retrieves multiple rows from 'users' as a slice of User.
//
// Generated from index 'idx_users_name'.
func FindUsersByName(ctx context.Context, db YORODB, name string) ([]*User, error) {
	const sqlstr = "SELECT " +
		"user_id, age, name " +
		"FROM users@{FORCE_INDEX=idx_users_name} " +
		"WHERE name = @param0"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = name

	decoder := newUser_Decoder(UserColumns())

	// run query
	YOLog(ctx, sqlstr, name)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*User{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindUsersByName", "users", err)
		}

		u, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindUsersByName", "users", err)
		}

		res = append(res, u)
	}

	return res, nil
}

// ReadUsersByName retrieves multiples rows from 'users' by KeySet as a slice.
//
// This does not retrives all columns of 'users' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'idx_users_name'.
func ReadUsersByName(ctx context.Context, db YORODB, keys spanner.KeySet) ([]*User, error) {
	var res []*User
	columns := []string{
		"user_id",
		"name",
	}

	decoder := newUser_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "users", "idx_users_name", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		u, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, u)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadUsersByName", "users", err)
	}

	return res, nil
}
