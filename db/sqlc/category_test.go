package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wil-ckaew/gofinance-backend/util"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)
	arg := CreateCategoryParams{
		UserID:      user.ID,
		Title:       util.RandomString(12),
		Type:        "debit",
		Description: util.RandomString(20),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.UserID, category.UserID)
	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.Type, category.Type)
	require.Equal(t, arg.Description, category.Description)
	require.NotEmpty(t, category.CreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.Title, category2.Title)
	require.Equal(t, category1.Description, category2.Description)
	require.Equal(t, category1.Type, category2.Type)
	require.NotEmpty(t, category2.CreatedAt)

}

func TestDeleteCategory(t *testing.T) {
	category := createRandomCategory(t)
	err := testQueries.DeleteCategories(context.Background(), category.ID)
	require.NoError(t, err)
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	arg := UpdateCategoriesParams{
		ID:          category1.ID,
		Title:       util.RandomString(12),
		Description: util.RandomString(20),
	}

	category2, err := testQueries.UpdateCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.Title, category2.Title)
	require.Equal(t, arg.Description, category2.Description)
	require.NotEmpty(t, category2.CreatedAt)
}

func TestListCategories(t *testing.T) {
	lasCategory := createRandomCategory(t)

	arg := GetCategoriesParams{
		UserID:      lasCategory.UserID,
		Type:        lasCategory.Type,
		Title:       lasCategory.Title,
		Description: lasCategory.Description,
	}

	categorys, err := testQueries.GetCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, categorys)

	for _, category := range categorys {
		require.Equal(t, lasCategory.ID, category.ID)
		require.Equal(t, lasCategory.UserID, category.UserID)
		require.Equal(t, lasCategory.Title, category.Title)
		require.Equal(t, lasCategory.Description, category.Description)
		require.NotEmpty(t, lasCategory.CreatedAt)
	}

}
