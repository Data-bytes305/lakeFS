package committed_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/treeverse/lakefs/pkg/graveler"
	"github.com/treeverse/lakefs/pkg/graveler/committed"
	"github.com/treeverse/lakefs/pkg/graveler/committed/mock"
	"github.com/treeverse/lakefs/pkg/graveler/testutil"
)

func makeV(k, id string) *graveler.ValueRecord {
	return &graveler.ValueRecord{Key: graveler.Key(k), Value: &graveler.Value{Identity: []byte(id)}}
}

func makeTombstoneV(k string) *graveler.ValueRecord {
	return &graveler.ValueRecord{Key: graveler.Key(k)}
}

func TestApplyAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	base := testutil.NewFakeIterator()
	base.
		AddRange(&committed.Range{ID: "one", MinKey: committed.Key("a"), MaxKey: committed.Key("cz"), Count: 2}).
		AddValueRecords(makeV("a", "base:a"), makeV("c", "base:c")).
		AddRange(&committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("d"), Count: 1}).
		AddValueRecords(makeV("d", "base:d"))
	changes := testutil.NewFakeIterator()
	changes.
		AddRange(&committed.Range{ID: "b", MinKey: committed.Key("b"), MaxKey: committed.Key("f"), Count: 3}).
		AddValueRecords(makeV("b", "dest:b"), makeV("e", "dest:e"), makeV("f", "dest:f"))
	writer := mock.NewMockMetaRangeWriter(ctrl)
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("a", "base:a")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("b", "dest:b")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("c", "base:c")))
	writer.EXPECT().WriteRange(gomock.Eq(committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("d"), Count: 1}))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("e", "dest:e")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("f", "dest:f")))

	summary, err := committed.Apply(context.Background(), writer, base, changes, &committed.ApplyOptions{})
	assert.NoError(t, err)
	assert.Equal(t, graveler.DiffSummary{
		Count: map[graveler.DiffType]int{
			graveler.DiffTypeAdded: 3,
		},
	}, summary)
}

func TestChangeRangesWithinBaseRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	base := testutil.NewFakeIterator()
	base.
		AddRange(&committed.Range{ID: "outer-range", MinKey: committed.Key("k1"), MaxKey: committed.Key("k6"), Count: 2}).
		AddValueRecords(makeV("k1", "base:k1"), makeV("k6", "base:k6")).
		AddRange(&committed.Range{ID: "last", MinKey: committed.Key("k10"), MaxKey: committed.Key("k13"), Count: 2}).
		AddValueRecords(makeV("k10", "base:k10"), makeV("k13", "base:k13"))
	changes := testutil.NewFakeIterator()
	changes.
		AddRange(&committed.Range{ID: "inner-range-1", MinKey: committed.Key("k2"), MaxKey: committed.Key("k3"), Count: 2}).
		AddValueRecords(makeV("k2", "dest:k2"), makeV("k3", "dest:k3")).
		AddRange(&committed.Range{ID: "inner-range-2", MinKey: committed.Key("k4"), MaxKey: committed.Key("k5"), Count: 2}).
		AddValueRecords(makeV("k4", "dest:k4"), makeV("k5", "dest:k5"))
	writer := mock.NewMockMetaRangeWriter(ctrl)
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("k1", "base:k1")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("k2", "dest:k2")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("k3", "dest:k3")))
	writer.EXPECT().WriteRange(gomock.Eq(committed.Range{ID: "inner-range-2", MinKey: committed.Key("k4"), MaxKey: committed.Key("k5"), Count: 2}))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("k6", "base:k6")))
	writer.EXPECT().WriteRange(gomock.Eq(committed.Range{ID: "last", MinKey: committed.Key("k10"), MaxKey: committed.Key("k13"), Count: 2}))
	summary, err := committed.Apply(context.Background(), writer, base, changes, &committed.ApplyOptions{})
	assert.NoError(t, err)
	assert.Equal(t, graveler.DiffSummary{
		Count: map[graveler.DiffType]int{
			graveler.DiffTypeAdded: 4,
		},
	}, summary)
}

func TestBaseRangesWithinChangeRange(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	base := testutil.NewFakeIterator()
	base.
		AddRange(&committed.Range{ID: "inner-range-1", MinKey: committed.Key("k2"), MaxKey: committed.Key("k3"), Count: 2}).
		AddValueRecords(makeV("k2", "base:k2"), makeV("k3", "base:k3")).
		AddRange(&committed.Range{ID: "inner-range-2", MinKey: committed.Key("k4"), MaxKey: committed.Key("k5"), Count: 2}).
		AddValueRecords(makeV("k4", "base:k4"), makeV("k5", "base:k5"))

	changes := testutil.NewFakeIterator()
	changes.
		AddRange(&committed.Range{ID: "outer-range", MinKey: committed.Key("k1"), MaxKey: committed.Key("k6"), Count: 2}).
		AddValueRecords(makeV("k1", "changes:k1"), makeV("k6", "changes:k6")).
		AddRange(&committed.Range{ID: "last", MinKey: committed.Key("k10"), MaxKey: committed.Key("k13"), Count: 2}).
		AddValueRecords(makeV("k10", "changes:k10"), makeV("k13", "changes:k13"))
	writer := mock.NewMockMetaRangeWriter(ctrl)
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("k1", "changes:k1")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("k2", "base:k2")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("k3", "base:k3")))
	writer.EXPECT().WriteRange(gomock.Eq(committed.Range{ID: "inner-range-2", MinKey: committed.Key("k4"), MaxKey: committed.Key("k5"), Count: 2}))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("k6", "changes:k6")))
	writer.EXPECT().WriteRange(gomock.Eq(committed.Range{ID: "last", MinKey: committed.Key("k10"), MaxKey: committed.Key("k13"), Count: 2}))
	summary, err := committed.Apply(context.Background(), writer, base, changes, &committed.ApplyOptions{})
	assert.NoError(t, err)
	assert.Equal(t, graveler.DiffSummary{
		Count: map[graveler.DiffType]int{
			graveler.DiffTypeAdded: 4,
		},
	}, summary)
}

func TestApplyReplace(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	base := testutil.NewFakeIterator()
	base.
		AddRange(&committed.Range{ID: "base:one", MinKey: committed.Key("a"), MaxKey: committed.Key("cz"), Count: 3}).
		AddValueRecords(makeV("a", "base:a"), makeV("b", "base:b"), makeV("c", "base:c")).
		AddRange(&committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("dz"), Count: 1}).
		AddValueRecords(makeV("d", "base:d")).
		AddRange(&committed.Range{ID: "three", MinKey: committed.Key("e"), MaxKey: committed.Key("ez"), Count: 1}).
		AddValueRecords(makeV("e", "base:e"))
	changes := testutil.NewFakeIterator()
	changes.
		AddRange(&committed.Range{ID: "changes:one", MinKey: committed.Key("b"), MaxKey: committed.Key("e"), Count: 2}).
		AddValueRecords(makeV("b", "changes:b"), makeV("e", "changes:e"))

	writer := mock.NewMockMetaRangeWriter(ctrl)
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("a", "base:a")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("b", "changes:b")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("c", "base:c")))
	writer.EXPECT().WriteRange(gomock.Eq(committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("dz"), Count: 1}))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("e", "changes:e")))

	summary, err := committed.Apply(context.Background(), writer, base, changes, &committed.ApplyOptions{})
	assert.NoError(t, err)
	assert.Equal(t, graveler.DiffSummary{
		Count: map[graveler.DiffType]int{
			graveler.DiffTypeChanged: 2,
		},
	}, summary)
}

func TestApplyDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	base := testutil.NewFakeIterator()
	base.
		AddRange(&committed.Range{ID: "one", MinKey: committed.Key("a"), MaxKey: committed.Key("cz"), Count: 3}).
		AddValueRecords(makeV("a", "base:a"), makeV("b", "base:b"), makeV("c", "base:c")).
		AddRange(&committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("dz"), Count: 1}).
		AddValueRecords(makeV("d", "base:d")).
		AddRange(&committed.Range{ID: "three", MinKey: committed.Key("ez"), MaxKey: committed.Key("ez"), Count: 1}).
		AddValueRecords(makeV("e", "base:e"))
	changes := testutil.NewFakeIterator()
	changes.
		AddRange(&committed.Range{ID: "changes:one", MinKey: committed.Key("b"), MaxKey: committed.Key("e"), Count: 2}).
		AddValueRecords(makeTombstoneV("b"), makeTombstoneV("e"))
	writer := mock.NewMockMetaRangeWriter(ctrl)
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("a", "base:a")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("c", "base:c")))
	writer.EXPECT().WriteRange(gomock.Eq(committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("dz"), Count: 1}))

	summary, err := committed.Apply(context.Background(), writer, base, changes, &committed.ApplyOptions{})
	assert.NoError(t, err)
	assert.Equal(t, graveler.DiffSummary{
		Count: map[graveler.DiffType]int{
			graveler.DiffTypeRemoved: 2,
		},
	}, summary)
}

func TestApplyCopiesLeftoverChanges(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	base := testutil.NewFakeIterator()
	base.
		AddRange(&committed.Range{ID: "one", MinKey: committed.Key("a"), MaxKey: committed.Key("cz"), Count: 3}).
		AddValueRecords(makeV("a", "base:a"), makeV("b", "base:b"), makeV("c", "base:c")).
		AddRange(&committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("dz"), Count: 1}).
		AddValueRecords(makeV("d", "base:d"))
	changes := testutil.NewFakeIterator()
	changes.
		AddRange(&committed.Range{ID: "changes:one", MinKey: committed.Key("b"), MaxKey: committed.Key("f"), Count: 3}).
		AddValueRecords(makeV("b", "changes:b"), makeV("e", "changes:e"), makeV("f", "changes:f"))

	writer := mock.NewMockMetaRangeWriter(ctrl)
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("a", "base:a")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("b", "changes:b")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("c", "base:c")))
	writer.EXPECT().WriteRange(gomock.Eq(committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("dz"), Count: 1}))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("e", "changes:e")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("f", "changes:f")))
	summary, err := committed.Apply(context.Background(), writer, base, changes, &committed.ApplyOptions{})
	assert.NoError(t, err)
	assert.Equal(t, graveler.DiffSummary{
		Count: map[graveler.DiffType]int{
			graveler.DiffTypeChanged: 1,
			graveler.DiffTypeAdded:   2,
		},
	}, summary)
}

func TestApplyTombstoneNoBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	base := testutil.NewFakeIterator()

	changes := testutil.NewFakeIterator()
	changes.
		AddRange(&committed.Range{ID: "changes:one", MinKey: committed.Key("b"), MaxKey: committed.Key("f"), Count: 3, Tombstone: true}).
		AddValueRecords(makeTombstoneV("b"), makeTombstoneV("e"), makeTombstoneV("f"))

	writer := mock.NewMockMetaRangeWriter(ctrl)

	summary, err := committed.Apply(context.Background(), writer, base, changes, &committed.ApplyOptions{})
	assert.Error(t, err, graveler.ErrNoChanges)
	assert.Equal(t, graveler.DiffSummary{
		Count: map[graveler.DiffType]int{},
	}, summary)
}

func TestApplyDeleteNonExistingRecord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	base := testutil.NewFakeIterator().
		AddRange(&committed.Range{ID: "base:one", MinKey: committed.Key("c"), MaxKey: committed.Key("d"), Count: 2}).
		AddValueRecords(makeV("c", "base:c"), makeV("d", "base:d"))
	changes := testutil.NewFakeIterator()
	changes.
		AddRange(&committed.Range{ID: "changes:one", MinKey: committed.Key("a"), MaxKey: committed.Key("f"), Count: 3}).
		AddValueRecords(makeV("a", "changes:a"), makeTombstoneV("e"), makeTombstoneV("f"))

	writer := mock.NewMockMetaRangeWriter(ctrl)
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("a", "changes:a")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("c", "base:c")))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("d", "base:d")))

	summary, err := committed.Apply(context.Background(), writer, base, changes, &committed.ApplyOptions{})

	assert.NoError(t, err)
	assert.Equal(t, graveler.DiffSummary{
		Count: map[graveler.DiffType]int{
			graveler.DiffTypeAdded: 1,
		},
	}, summary)
}
func TestApplyCopiesLeftoverBase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	range1 := &committed.Range{ID: "one", MinKey: committed.Key("a"), MaxKey: committed.Key("cz"), Count: 3}
	range2 := &committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("dz"), Count: 1}
	range4 := &committed.Range{ID: "four", MinKey: committed.Key("g"), MaxKey: committed.Key("hz"), Count: 2}
	base := testutil.NewFakeIterator()
	base.
		AddRange(range1).
		AddValueRecords(makeV("a", "base:a"), makeV("b", "base:b"), makeV("c", "base:c")).
		AddRange(range2).
		AddValueRecords(makeV("d", "base:d")).
		AddRange(&committed.Range{ID: "three", MaxKey: committed.Key("ez"), Count: 1}).
		AddValueRecords(makeV("e", "base:e"), makeV("f", "base:f")).
		AddRange(range4).
		AddValueRecords(makeV("g", "base:g"), makeV("h", "base:h"))

	changes := testutil.NewValueIteratorFake([]graveler.ValueRecord{
		*makeTombstoneV("e"),
	})

	writer := mock.NewMockMetaRangeWriter(ctrl)
	writer.EXPECT().WriteRange(gomock.Eq(*range1))
	writer.EXPECT().WriteRange(gomock.Eq(*range2))
	writer.EXPECT().WriteRecord(gomock.Eq(*makeV("f", "base:f")))
	writer.EXPECT().WriteRange(gomock.Eq(*range4))

	summary, err := committed.Apply(context.Background(), writer, base, committed.NewIteratorWrapper(changes), &committed.ApplyOptions{})
	assert.NoError(t, err)
	assert.Equal(t, graveler.DiffSummary{
		Count: map[graveler.DiffType]int{
			graveler.DiffTypeRemoved: 1,
		},
	}, summary)
}

func TestApplyNoChangesFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	range1 := &committed.Range{ID: "one", MinKey: committed.Key("a"), MaxKey: committed.Key("cz"), Count: 3}
	range2 := &committed.Range{ID: "two", MinKey: committed.Key("d"), MaxKey: committed.Key("dz"), Count: 1}
	base := testutil.NewFakeIterator()
	base.
		AddRange(range1).
		AddValueRecords(makeV("a", "base:a"), makeV("b", "base:b"), makeV("c", "base:c")).
		AddRange(range2).
		AddValueRecords(makeV("d", "base:d"))

	changes := testutil.NewFakeIterator()

	writer := mock.NewMockMetaRangeWriter(ctrl)
	writer.EXPECT().WriteRange(gomock.Any()).AnyTimes()

	_, err := committed.Apply(context.Background(), writer, base, changes, &committed.ApplyOptions{})
	assert.Error(t, err, graveler.ErrNoChanges)
}

func TestApplyCancelContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("base", func(t *testing.T) {
		base := testutil.NewFakeIterator().
			AddRange(&committed.Range{ID: "one", MinKey: committed.Key("a"), MaxKey: committed.Key("cz"), Count: 2}).
			AddValueRecords(makeV("a", "base:a"), makeV("c", "base:c"))
		changes := testutil.NewFakeIterator()
		writer := mock.NewMockMetaRangeWriter(ctrl)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := committed.Apply(ctx, writer, base, changes, &committed.ApplyOptions{})
		assert.True(t, errors.Is(err, context.Canceled), "context canceled error")
	})

	t.Run("change", func(t *testing.T) {
		base := testutil.NewFakeIterator()
		changes := testutil.NewFakeIterator().
			AddRange(&committed.Range{ID: "one", MinKey: committed.Key("b"), MaxKey: committed.Key("b"), Count: 1}).
			AddValueRecords(makeV("b", "dest:b"))
		writer := mock.NewMockMetaRangeWriter(ctrl)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := committed.Apply(ctx, writer, base, changes, &committed.ApplyOptions{})
		assert.True(t, errors.Is(err, context.Canceled), "context canceled error")
	})

	t.Run("base_and_change", func(t *testing.T) {
		base := testutil.NewFakeIterator().
			AddRange(&committed.Range{ID: "one", MinKey: committed.Key("a"), MaxKey: committed.Key("cz"), Count: 3}).
			AddValueRecords(makeV("a", "base:a"), makeV("c", "base:c"))
		changes := testutil.NewFakeIterator().
			AddRange(&committed.Range{ID: "one", MinKey: committed.Key("b"), MaxKey: committed.Key("b"), Count: 1}).
			AddValueRecords(makeV("b", "dest:b"))
		writer := mock.NewMockMetaRangeWriter(ctrl)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := committed.Apply(ctx, writer, base, changes, &committed.ApplyOptions{})
		assert.True(t, errors.Is(err, context.Canceled), "context canceled error")
	})
}
