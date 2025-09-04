package lists_test

import (
	"context"
	"testing"

	"top10/internal/lists"
)

func TestNewService(t *testing.T) {
	t.Run("should create service successfully with embedded data", func(t *testing.T) {
		t.Helper()

		service, err := lists.NewService(context.Background())
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if service == nil {
			t.Fatal("expected service to be non-nil")
		}

		count := service.GetListCount()
		if count == 0 {
			t.Fatal("expected list count to be greater than 0")
		}

		if count != 1857 {
			t.Errorf("expected 1857 lists, got %d", count)
		}
	})
}

func TestService_GetRandomList(t *testing.T) {
	t.Run("should return a valid top ten list", func(t *testing.T) {
		t.Helper()

		service, err := lists.NewService(context.Background())
		if err != nil {
			t.Fatalf("failed to create service: %v", err)
		}

		list, err := service.GetRandomList()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if list.Title == "" {
			t.Error("expected list title to be non-empty")
		}

		if len(list.Items) != 10 {
			t.Errorf("expected 10 items, got %d", len(list.Items))
		}

		if list.Year != 0 && (list.Year < 1980 || list.Year > 2020) {
			t.Errorf("expected year to be 0 or between 1980 and 2020, got %d", list.Year)
		}

		if list.Show == "" {
			t.Error("expected show to be non-empty")
		}

		if list.Date == "" {
			t.Error("expected date to be non-empty")
		}
	})

	t.Run("should return different lists on multiple calls", func(t *testing.T) {
		t.Helper()

		service, err := lists.NewService(context.Background())
		if err != nil {
			t.Fatalf("failed to create service: %v", err)
		}

		firstList, err := service.GetRandomList()
		if err != nil {
			t.Fatalf("failed to get first list: %v", err)
		}

		differentFound := false
		for i := 0; i < 10; i++ {
			secondList, err := service.GetRandomList()
			if err != nil {
				t.Fatalf("failed to get second list: %v", err)
			}

			if firstList.Title != secondList.Title {
				differentFound = true
				break
			}
		}

		if !differentFound {
			t.Error("expected to find different lists after multiple calls, but all were the same")
		}
	})
}

func TestService_GetListCount(t *testing.T) {
	t.Run("should return correct count of lists", func(t *testing.T) {
		t.Helper()

		service, err := lists.NewService(context.Background())
		if err != nil {
			t.Fatalf("failed to create service: %v", err)
		}

		count := service.GetListCount()
		expectedCount := 1857

		if count != expectedCount {
			t.Errorf("expected count %d, got %d", expectedCount, count)
		}
	})
}
