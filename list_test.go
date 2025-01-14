package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListLists(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": [
				{
					"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
					"name": "list1",
					"description": "This is a note.",
					"kind": "ip",
					"num_items": 10,
					"num_referencing_filters": 2,
					"created_on": "2020-01-01T08:00:00Z",
					"modified_on": "2020-01-10T14:00:00Z"
				}
			],
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists", handler)

	createdOn, _ := time.Parse(time.RFC3339, "2020-01-01T08:00:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2020-01-10T14:00:00Z")

	want := []List{
		{
			ID:                    "2c0fc9fa937b11eaa1b71c4d701ab86e",
			Name:                  "list1",
			Description:           "This is a note.",
			Kind:                  "ip",
			NumItems:              10,
			NumReferencingFilters: 2,
			CreatedOn:             &createdOn,
			ModifiedOn:            &modifiedOn,
		},
	}

	actual, err := client.ListLists(context.Background(), ListListsParams{
		AccountID: testAccountID,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateList(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
				"name": "list1",
				"description": "This is a note.",
				"kind": "ip",
				"num_items": 10,
				"num_referencing_filters": 2,
				"created_on": "2020-01-01T08:00:00Z",
				"modified_on": "2020-01-10T14:00:00Z"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists", handler)

	createdOn, _ := time.Parse(time.RFC3339, "2020-01-01T08:00:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2020-01-10T14:00:00Z")

	want := List{
		ID:                    "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Name:                  "list1",
		Description:           "This is a note.",
		Kind:                  "ip",
		NumItems:              10,
		NumReferencingFilters: 2,
		CreatedOn:             &createdOn,
		ModifiedOn:            &modifiedOn,
	}

	actual, err := client.CreateList(context.Background(), ListCreateParams{
		AccountID: testAccountID, Name: "list1", Description: "This is a note.", Kind: "ip",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetList(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
				"name": "list1",
				"description": "This is a note.",
				"kind": "ip",
				"num_items": 10,
				"num_referencing_filters": 2,
				"created_on": "2020-01-01T08:00:00Z",
				"modified_on": "2020-01-10T14:00:00Z"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)

	createdOn, _ := time.Parse(time.RFC3339, "2020-01-01T08:00:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2020-01-10T14:00:00Z")

	want := List{
		ID:                    "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Name:                  "list1",
		Description:           "This is a note.",
		Kind:                  "ip",
		NumItems:              10,
		NumReferencingFilters: 2,
		CreatedOn:             &createdOn,
		ModifiedOn:            &modifiedOn,
	}

	actual, err := client.GetList(context.Background(), ListGetParams{
		AccountID: testAccountID, ID: "2c0fc9fa937b11eaa1b71c4d701ab86e",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestUpdateList(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
				"name": "list1",
				"description": "This note was updated.",
				"kind": "ip",
				"num_items": 10,
				"num_referencing_filters": 2,
				"created_on": "2020-01-01T08:00:00Z",
				"modified_on": "2020-01-10T14:00:00Z"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)

	createdOn, _ := time.Parse(time.RFC3339, "2020-01-01T08:00:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2020-01-10T14:00:00Z")

	want := List{
		ID:                    "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Name:                  "list1",
		Description:           "This note was updated.",
		Kind:                  "ip",
		NumItems:              10,
		NumReferencingFilters: 2,
		CreatedOn:             &createdOn,
		ModifiedOn:            &modifiedOn,
	}

	actual, err := client.UpdateList(context.Background(),
		ListUpdateParams{
			AccountID: testAccountID, ID: "2c0fc9fa937b11eaa1b71c4d701ab86e", Description: "This note was updated.",
		},
	)
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteList(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"id": "34b12448945f11eaa1b71c4d701ab86e"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists/2c0fc9fa937b11eaa1b71c4d701ab86e", handler)

	want := ListDeleteResponse{}
	want.Success = true
	want.Errors = []ResponseInfo{}
	want.Messages = []ResponseInfo{}
	want.Result.ID = "34b12448945f11eaa1b71c4d701ab86e"

	actual, err := client.DeleteList(context.Background(), ListDeleteParams{
		AccountID: testAccountID, ID: "2c0fc9fa937b11eaa1b71c4d701ab86e",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestListsItemsIP(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")

		if len(r.URL.Query().Get("cursor")) > 0 && r.URL.Query().Get("cursor") == "yyy" {
			fmt.Fprint(w, `{
			"result": [
    			{
      				"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
      				"ip": "192.0.2.2",
      				"comment": "Another Private IP address",
      				"created_on": "2020-01-01T08:00:00Z",
      				"modified_on": "2020-01-10T14:00:00Z"
    			}
  			],
  			"result_info": {
    			"cursors": {
					"before": "xxx"
				}
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
		} else {
			fmt.Fprint(w, `{
			"result": [
    			{
      				"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
      				"ip": "192.0.2.1",
      				"comment": "Private IP address",
      				"created_on": "2020-01-01T08:00:00Z",
      				"modified_on": "2020-01-10T14:00:00Z"
    			}
  			],
  			"result_info": {
    			"cursors": {
					"before": "xxx",
					"after": "yyy"
				}
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
		}
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists/2c0fc9fa937b11eaa1b71c4d701ab86e/items", handler)

	createdOn, _ := time.Parse(time.RFC3339, "2020-01-01T08:00:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2020-01-10T14:00:00Z")

	want := []ListItem{
		{
			ID:         "2c0fc9fa937b11eaa1b71c4d701ab86e",
			IP:         "192.0.2.1",
			Comment:    "Private IP address",
			CreatedOn:  &createdOn,
			ModifiedOn: &modifiedOn,
		},
		{
			ID:         "2c0fc9fa937b11eaa1b71c4d701ab86e",
			IP:         "192.0.2.2",
			Comment:    "Another Private IP address",
			CreatedOn:  &createdOn,
			ModifiedOn: &modifiedOn,
		},
	}

	actual, err := client.ListListItems(context.Background(), ListListItemsParams{
		AccountID: testAccountID, ID: "2c0fc9fa937b11eaa1b71c4d701ab86e",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestCreateListItemsIP(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'POST', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"operation_id": "4da8780eeb215e6cb7f48dd981c4ea02"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists/2c0fc9fa937b11eaa1b71c4d701ab86e/items", handler)

	want := ListItemCreateResponse{}
	want.Success = true
	want.Errors = []ResponseInfo{}
	want.Messages = []ResponseInfo{}
	want.Result.OperationID = "4da8780eeb215e6cb7f48dd981c4ea02"

	actual, err := client.CreateListItemsAsync(context.Background(), ListCreateItemsParams{
		AccountID: testAccountID,
		ID:        "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Items: []ListItemCreateRequest{{
			IP:      "192.0.2.1",
			Comment: "Private IP",
		}, {
			IP:      "192.0.2.2",
			Comment: "Another Private IP",
		}}})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestReplaceListItemsIP(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"operation_id": "4da8780eeb215e6cb7f48dd981c4ea02"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists/2c0fc9fa937b11eaa1b71c4d701ab86e/items", handler)

	want := ListItemCreateResponse{}
	want.Success = true
	want.Errors = []ResponseInfo{}
	want.Messages = []ResponseInfo{}
	want.Result.OperationID = "4da8780eeb215e6cb7f48dd981c4ea02"

	actual, err := client.ReplaceListItemsAsync(context.Background(), ListReplaceItemsParams{
		AccountID: testAccountID,
		ID:        "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Items: []ListItemCreateRequest{{
			IP:      "192.0.2.1",
			Comment: "Private IP",
		}, {
			IP:      "192.0.2.2",
			Comment: "Another Private IP",
		}}})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestDeleteListItems(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected method 'DELETE', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"operation_id": "4da8780eeb215e6cb7f48dd981c4ea02"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists/2c0fc9fa937b11eaa1b71c4d701ab86e/items", handler)

	want := ListItemDeleteResponse{}
	want.Success = true
	want.Errors = []ResponseInfo{}
	want.Messages = []ResponseInfo{}
	want.Result.OperationID = "4da8780eeb215e6cb7f48dd981c4ea02"

	actual, err := client.DeleteListItemsAsync(context.Background(), ListDeleteItemsParams{
		AccountID: testAccountID,
		ID:        "2c0fc9fa937b11eaa1b71c4d701ab86e",
		Items: ListItemDeleteRequest{[]ListItemDeleteItemRequest{{
			ID: "34b12448945f11eaa1b71c4d701ab86e",
		}}}})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestGetListItemIP(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, `{
			"result": {
				"id": "2c0fc9fa937b11eaa1b71c4d701ab86e",
				"ip": "192.0.2.1",
				"comment": "Private IP address",
				"created_on": "2020-01-01T08:00:00Z",
				"modified_on": "2020-01-10T14:00:00Z"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/accounts/"+testAccountID+"/rules/lists/2c0fc9fa937b11eaa1b71c4d701ab86e/items/"+
		"34b12448945f11eaa1b71c4d701ab86e", handler)

	createdOn, _ := time.Parse(time.RFC3339, "2020-01-01T08:00:00Z")
	modifiedOn, _ := time.Parse(time.RFC3339, "2020-01-10T14:00:00Z")

	want := ListItem{
		ID:         "2c0fc9fa937b11eaa1b71c4d701ab86e",
		IP:         "192.0.2.1",
		Comment:    "Private IP address",
		CreatedOn:  &createdOn,
		ModifiedOn: &modifiedOn,
	}

	actual, err := client.GetListItem(context.Background(), ListGetItemParams{
		AccountID: testAccountID,
		ListID:    "2c0fc9fa937b11eaa1b71c4d701ab86e",
		ID:        "34b12448945f11eaa1b71c4d701ab86e",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestPollListTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	start := time.Now()
	err := client.pollListBulkOperation(ctx, testAccountID, "list1")
	assert.ErrorIs(t, err, context.DeadlineExceeded)
	assert.WithinDuration(t, start, time.Now(), time.Second,
		"pollListBulkOperation took too much time with an expiring context")
}
