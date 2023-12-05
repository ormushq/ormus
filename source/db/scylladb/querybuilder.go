package scylladb

import (
	"context"
	"fmt"
	"log"
)

type queryBuilder[T any] struct {
	model   TableInterface
	session SessionxInterface
}

func (queryBuilder *queryBuilder[T]) Insert(ctx context.Context, insertData *T) error {
	insertStatement, insertNames := queryBuilder.model.Insert()
	insertQuery := queryBuilder.session.Query(insertStatement, insertNames)
	err := insertQuery.BindStruct(insertData).ExecRelease()
	if err != nil {
		log.Println("Insert error:", err.Error())
		return err
	}

	return nil
}

func (queryBuilder *queryBuilder[T]) Select(ctx context.Context, dataToGet *T) ([]T, error) {
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

func (queryBuilder *queryBuilder[T]) Get(ctx context.Context, dataToGet *T) (*T, error) {
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

func (queryBuilder *queryBuilder[T]) Delete(ctx context.Context, dataToBeDeleted *T) error {
	deleteStatment, deleteName := queryBuilder.model.Delete()
	fmt.Print(deleteName, deleteStatment)
	deleteQuery := queryBuilder.session.Query(deleteStatment, deleteName)

	if err := deleteQuery.BindStruct(dataToBeDeleted).WithContext(ctx).ExecRelease(); err != nil {
		log.Println("Delete error", err.Error())
		return err

	}

	return nil
}

func NewQueryBuilder[T any](model TableInterface, session SessionxInterface) *queryBuilder[T] {
	return &queryBuilder[T]{
		model:   model,
		session: session,
	}
}
