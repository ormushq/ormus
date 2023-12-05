/*
Package scylladb provides a QueryBuilder to simplify the interaction with a ScyllaDB database using the gocql library.

Usage:

	func main() {
	    // Assume 'model' and 'session' instances are available
	    // Create a new QueryBuilder
	    queryBuilder := scylladb.NewQueryBuilder[model](model, session)

	    // Example of inserting data into the database
	    insertData := &YourStruct{...}
	    err := queryBuilder.Insert(ctx, insertData)
	    if err != nil {
	        log.Fatal("Insert operation failed:", err)
	    }

	    // Example of selecting data from the database
	    dataToGet := &YourStruct{...}
	    results, err := queryBuilder.Select(ctx, dataToGet)
	    if err != nil {
	        log.Fatal("Select operation failed:", err)
	    }

	    // Example of getting a single record from the database
	    singleResult, err := queryBuilder.Get(ctx, dataToGet)
	    if err != nil {
	        log.Fatal("Get operation failed:", err)
	    }

	    // Example of deleting data from the database
	    dataToBeDeleted := &YourStruct{...}
	    err = queryBuilder.Delete(ctx, dataToBeDeleted)
	    if err != nil {
	        log.Fatal("Delete operation failed:", err)
	    }
	}

This package includes a QueryBuilder struct, which provides methods for common
CRUD operations (Insert, Select, Get, Delete) on a ScyllaDB database using the gocql library.
The QueryBuilder simplifies the process of constructing and executing CQL queries for data manipulation.

Structs:
  - QueryBuilder[T]: A generic struct that takes a model (implementing TableInterface) and
    a session (implementing SessionxInterface) to perform database operations.
*/
package scylladb

import (
	"context"
	"fmt"
	"log"
)

type QueryBuilder[T any] struct {
	model   TableInterface
	session SessionxInterface
}

func (queryBuilder *QueryBuilder[T]) Insert(ctx context.Context, insertData *T) error {
	insertStatement, insertNames := queryBuilder.model.Insert()
	insertQuery := queryBuilder.session.Query(insertStatement, insertNames)
	err := insertQuery.BindStruct(insertData).ExecRelease()
	if err != nil {
		log.Println("Insert error:", err.Error())
		return err
	}

	return nil
}

func (queryBuilder *QueryBuilder[T]) Select(ctx context.Context, dataToGet *T) ([]T, error) {
	selectStatement, selectNames := queryBuilder.model.Select()
	selectQuery := queryBuilder.session.Query(selectStatement, selectNames)

	var results []T
	err := selectQuery.BindStruct(dataToGet).SelectRelease(&results)
	if err != nil {
		log.Println("Select error:", err.Error())
		return nil, err
	}

	return results, nil
}

func (queryBuilder *QueryBuilder[T]) Get(ctx context.Context, dataToGet *T) (*T, error) {
	selectStatement, selectNames := queryBuilder.model.Get()
	selectQuery := queryBuilder.session.Query(selectStatement, selectNames)

	var result []T
	err := selectQuery.BindStruct(dataToGet).WithContext(ctx).SelectRelease(&result)
	if err != nil {
		log.Println("Get error", err.Error())
		return nil, err
	}

	if len(result) > 0 {
		return &result[0], nil
	}

	return nil, nil
}

func (queryBuilder *QueryBuilder[T]) Delete(ctx context.Context, dataToBeDeleted *T) error {
	deleteStatment, deleteName := queryBuilder.model.Delete()
	fmt.Print(deleteName, deleteStatment)
	deleteQuery := queryBuilder.session.Query(deleteStatment, deleteName)

	if err := deleteQuery.BindStruct(dataToBeDeleted).WithContext(ctx).ExecRelease(); err != nil {
		log.Println("Delete error", err.Error())
		return err

	}

	return nil
}

func NewQueryBuilder[T any](model TableInterface, session SessionxInterface) *QueryBuilder[T] {
	return &QueryBuilder[T]{
		model:   model,
		session: session,
	}
}
