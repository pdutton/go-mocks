package mock_sync

import (
	"testing"

	"github.com/pdutton/go-mocks/internal/testutil"
	"go.uber.org/mock/gomock"
)

// TestMockLocker_Lock tests basic lock operation.
func TestMockLocker_Lock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLocker := NewMockLocker(ctrl)

	mockLocker.EXPECT().Lock()

	mockLocker.Lock()
}

// TestMockLocker_Unlock tests basic unlock operation.
func TestMockLocker_Unlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLocker := NewMockLocker(ctrl)

	mockLocker.EXPECT().Unlock()

	mockLocker.Unlock()
}

// TestMockLocker_LockUnlock tests Lock/Unlock sequence with InOrder.
func TestMockLocker_LockUnlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLocker := NewMockLocker(ctrl)

	gomock.InOrder(
		mockLocker.EXPECT().Lock(),
		mockLocker.EXPECT().Unlock(),
	)

	mockLocker.Lock()
	mockLocker.Unlock()
}

// TestMockLocker_MultipleLockUnlock tests multiple lock/unlock cycles.
func TestMockLocker_MultipleLockUnlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLocker := NewMockLocker(ctrl)

	gomock.InOrder(
		mockLocker.EXPECT().Lock(),
		mockLocker.EXPECT().Unlock(),
		mockLocker.EXPECT().Lock(),
		mockLocker.EXPECT().Unlock(),
	)

	mockLocker.Lock()
	mockLocker.Unlock()
	mockLocker.Lock()
	mockLocker.Unlock()
}

// TestMockMutex_Lock tests mutex lock.
func TestMockMutex_Lock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMutex := NewMockMutex(ctrl)

	mockMutex.EXPECT().Lock()

	mockMutex.Lock()
}

// TestMockMutex_Unlock tests mutex unlock.
func TestMockMutex_Unlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMutex := NewMockMutex(ctrl)

	mockMutex.EXPECT().Unlock()

	mockMutex.Unlock()
}

// TestMockMutex_LockUnlockSequence tests mutex lock/unlock sequence.
func TestMockMutex_LockUnlockSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMutex := NewMockMutex(ctrl)

	gomock.InOrder(
		mockMutex.EXPECT().Lock(),
		mockMutex.EXPECT().Unlock(),
	)

	mockMutex.Lock()
	mockMutex.Unlock()
}

// TestMockRWMutex_RLock tests read lock.
func TestMockRWMutex_RLock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRWMutex := NewMockRWMutex(ctrl)

	mockRWMutex.EXPECT().RLock()

	mockRWMutex.RLock()
}

// TestMockRWMutex_RUnlock tests read unlock.
func TestMockRWMutex_RUnlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRWMutex := NewMockRWMutex(ctrl)

	mockRWMutex.EXPECT().RUnlock()

	mockRWMutex.RUnlock()
}

// TestMockRWMutex_Lock tests write lock.
func TestMockRWMutex_Lock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRWMutex := NewMockRWMutex(ctrl)

	mockRWMutex.EXPECT().Lock()

	mockRWMutex.Lock()
}

// TestMockRWMutex_Unlock tests write unlock.
func TestMockRWMutex_Unlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRWMutex := NewMockRWMutex(ctrl)

	mockRWMutex.EXPECT().Unlock()

	mockRWMutex.Unlock()
}

// TestMockRWMutex_ReadLockSequence tests multiple read locks.
func TestMockRWMutex_ReadLockSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRWMutex := NewMockRWMutex(ctrl)

	gomock.InOrder(
		mockRWMutex.EXPECT().RLock(),
		mockRWMutex.EXPECT().RLock(),
		mockRWMutex.EXPECT().RUnlock(),
		mockRWMutex.EXPECT().RUnlock(),
	)

	mockRWMutex.RLock()
	mockRWMutex.RLock()
	mockRWMutex.RUnlock()
	mockRWMutex.RUnlock()
}

// TestMockRWMutex_WriteLockSequence tests write lock/unlock sequence.
func TestMockRWMutex_WriteLockSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRWMutex := NewMockRWMutex(ctrl)

	gomock.InOrder(
		mockRWMutex.EXPECT().Lock(),
		mockRWMutex.EXPECT().Unlock(),
	)

	mockRWMutex.Lock()
	mockRWMutex.Unlock()
}

// TestMockWaitGroup_Add tests counter increment.
func TestMockWaitGroup_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWG := NewMockWaitGroup(ctrl)

	mockWG.EXPECT().Add(1)

	mockWG.Add(1)
}

// TestMockWaitGroup_Done tests counter decrement.
func TestMockWaitGroup_Done(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWG := NewMockWaitGroup(ctrl)

	mockWG.EXPECT().Done()

	mockWG.Done()
}

// TestMockWaitGroup_Wait tests waiting for zero.
func TestMockWaitGroup_Wait(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWG := NewMockWaitGroup(ctrl)

	mockWG.EXPECT().Wait()

	mockWG.Wait()
}

// TestMockWaitGroup_Lifecycle tests Add, Done, Wait sequence.
func TestMockWaitGroup_Lifecycle(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWG := NewMockWaitGroup(ctrl)

	gomock.InOrder(
		mockWG.EXPECT().Add(2),
		mockWG.EXPECT().Done(),
		mockWG.EXPECT().Done(),
		mockWG.EXPECT().Wait(),
	)

	mockWG.Add(2)
	mockWG.Done()
	mockWG.Done()
	mockWG.Wait()
}

// TestMockWaitGroup_MultipleAdds tests multiple Add operations.
func TestMockWaitGroup_MultipleAdds(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockWG := NewMockWaitGroup(ctrl)

	gomock.InOrder(
		mockWG.EXPECT().Add(1),
		mockWG.EXPECT().Add(1),
		mockWG.EXPECT().Add(1),
	)

	mockWG.Add(1)
	mockWG.Add(1)
	mockWG.Add(1)
}

// TestMockOnce_Do tests single execution.
func TestMockOnce_Do(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOnce := NewMockOnce(ctrl)

	f := func() {}
	mockOnce.EXPECT().Do(gomock.Any())

	mockOnce.Do(f)
}

// TestMockCond_Wait tests condition wait.
func TestMockCond_Wait(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCond := NewMockCond(ctrl)

	mockCond.EXPECT().Wait()

	mockCond.Wait()
}

// TestMockCond_Signal tests signaling one waiter.
func TestMockCond_Signal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCond := NewMockCond(ctrl)

	mockCond.EXPECT().Signal()

	mockCond.Signal()
}

// TestMockCond_Broadcast tests signaling all waiters.
func TestMockCond_Broadcast(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCond := NewMockCond(ctrl)

	mockCond.EXPECT().Broadcast()

	mockCond.Broadcast()
}

// TestMockCond_WaitSignalSequence tests wait/signal pattern.
func TestMockCond_WaitSignalSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCond := NewMockCond(ctrl)

	gomock.InOrder(
		mockCond.EXPECT().Wait(),
		mockCond.EXPECT().Signal(),
	)

	mockCond.Wait()
	mockCond.Signal()
}

// TestMockPool_Get tests getting from pool.
func TestMockPool_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPool := NewMockPool(ctrl)

	mockPool.EXPECT().Get().Return("value")

	result := mockPool.Get()
	testutil.AssertEqual(t, "value", result)
}

// TestMockPool_Put tests putting to pool.
func TestMockPool_Put(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPool := NewMockPool(ctrl)

	mockPool.EXPECT().Put("value")

	mockPool.Put("value")
}

// TestMockPool_GetPutSequence tests get/put cycle.
func TestMockPool_GetPutSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPool := NewMockPool(ctrl)

	gomock.InOrder(
		mockPool.EXPECT().Get().Return("item"),
		mockPool.EXPECT().Put("item"),
	)

	item := mockPool.Get()
	testutil.AssertEqual(t, "item", item)
	mockPool.Put(item)
}

// TestMockPool_MultipleGets tests multiple gets.
func TestMockPool_MultipleGets(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPool := NewMockPool(ctrl)

	gomock.InOrder(
		mockPool.EXPECT().Get().Return("item1"),
		mockPool.EXPECT().Get().Return("item2"),
		mockPool.EXPECT().Get().Return("item3"),
	)

	item1 := mockPool.Get()
	testutil.AssertEqual(t, "item1", item1)

	item2 := mockPool.Get()
	testutil.AssertEqual(t, "item2", item2)

	item3 := mockPool.Get()
	testutil.AssertEqual(t, "item3", item3)
}

// TestMockMap_Load tests loading from map.
func TestMockMap_Load(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMap := NewMockMap(ctrl)

	mockMap.EXPECT().Load("key").Return("value", true)

	value, ok := mockMap.Load("key")
	testutil.AssertEqual(t, "value", value)
	testutil.AssertEqual(t, true, ok)
}

// TestMockMap_LoadNotFound tests loading missing key.
func TestMockMap_LoadNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMap := NewMockMap(ctrl)

	mockMap.EXPECT().Load("missing").Return(nil, false)

	value, ok := mockMap.Load("missing")
	testutil.AssertNil(t, value)
	testutil.AssertEqual(t, false, ok)
}

// TestMockMap_Store tests storing to map.
func TestMockMap_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMap := NewMockMap(ctrl)

	mockMap.EXPECT().Store("key", "value")

	mockMap.Store("key", "value")
}

// TestMockMap_Delete tests deleting from map.
func TestMockMap_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMap := NewMockMap(ctrl)

	mockMap.EXPECT().Delete("key")

	mockMap.Delete("key")
}

// TestMockMap_LoadOrStore tests load-or-store operation.
func TestMockMap_LoadOrStore(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMap := NewMockMap(ctrl)

	mockMap.EXPECT().LoadOrStore("key", "value").Return("value", false)

	actual, loaded := mockMap.LoadOrStore("key", "value")
	testutil.AssertEqual(t, "value", actual)
	testutil.AssertEqual(t, false, loaded)
}

// TestMockMap_LoadAndDelete tests load-and-delete operation.
func TestMockMap_LoadAndDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMap := NewMockMap(ctrl)

	mockMap.EXPECT().LoadAndDelete("key").Return("value", true)

	value, loaded := mockMap.LoadAndDelete("key")
	testutil.AssertEqual(t, "value", value)
	testutil.AssertEqual(t, true, loaded)
}

// TestMockMap_Range tests iterating over map.
func TestMockMap_Range(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMap := NewMockMap(ctrl)

	f := func(key, value any) bool { return true }
	mockMap.EXPECT().Range(gomock.Any())

	mockMap.Range(f)
}

// TestMockMap_StoreLoadSequence tests store then load.
func TestMockMap_StoreLoadSequence(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMap := NewMockMap(ctrl)

	gomock.InOrder(
		mockMap.EXPECT().Store("key", "value"),
		mockMap.EXPECT().Load("key").Return("value", true),
	)

	mockMap.Store("key", "value")
	value, ok := mockMap.Load("key")
	testutil.AssertEqual(t, "value", value)
	testutil.AssertEqual(t, true, ok)
}
