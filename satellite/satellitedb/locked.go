// Code generated by lockedgen using 'go generate'. DO NOT EDIT.

// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"sync"
	"time"

	"storj.io/storj/pkg/accounting"
	"storj.io/storj/pkg/bwagreement"
	"storj.io/storj/pkg/datarepair/irreparable"
	"storj.io/storj/pkg/datarepair/queue"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/statdb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite"
	"storj.io/storj/storage"
)

// locked implements a locking wrapper around satellite.DB.
type locked struct {
	sync.Locker
	db satellite.DB
}

// newLocked returns database wrapped with locker.
func newLocked(db satellite.DB) satellite.DB {
	return &locked{&sync.Mutex{}, db}
}

// Accounting returns database for storing information about data use
func (m *locked) Accounting() accounting.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedAccounting{m.Locker, m.db.Accounting()}
}

// BandwidthAgreement returns database for storing bandwidth agreements
func (m *locked) BandwidthAgreement() bwagreement.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedBandwidthAgreement{m.Locker, m.db.BandwidthAgreement()}
}

// Close closes the database
func (m *locked) Close() error {
	m.Lock()
	defer m.Unlock()
	return m.db.Close()
}

// CreateTables initializes the database
func (m *locked) CreateTables() error {
	m.Lock()
	defer m.Unlock()
	return m.db.CreateTables()
}

// Irreparable returns database for failed repairs
func (m *locked) Irreparable() irreparable.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedIrreparable{m.Locker, m.db.Irreparable()}
}

// OverlayCache returns database for caching overlay information
func (m *locked) OverlayCache() storage.KeyValueStore {
	m.Lock()
	defer m.Unlock()
	return &lockedOverlayCache{m.Locker, m.db.OverlayCache()}
}

// RepairQueue returns queue for segments that need repairing
func (m *locked) RepairQueue() queue.RepairQueue {
	m.Lock()
	defer m.Unlock()
	return &lockedRepairQueue{m.Locker, m.db.RepairQueue()}
}

// StatDB returns database for storing node statistics
func (m *locked) StatDB() statdb.DB {
	m.Lock()
	defer m.Unlock()
	return &lockedStatDB{m.Locker, m.db.StatDB()}
}

// lockedAccounting implements locking wrapper for accounting.DB
type lockedAccounting struct {
	sync.Locker
	db accounting.DB
}

// LastRawTime records the latest last tallied time.
func (m *lockedAccounting) LastRawTime(ctx context.Context, timestampType string) (time.Time, bool, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.LastRawTime(ctx, timestampType)
}

// SaveAtRestRaw records raw tallies of at-rest-data.
func (m *lockedAccounting) SaveAtRestRaw(ctx context.Context, latestTally time.Time, nodeData map[storj.NodeID]int64) error {
	m.Lock()
	defer m.Unlock()
	return m.db.SaveAtRestRaw(ctx, latestTally, nodeData)
}

// SaveBWRaw records raw sums of agreement values to the database and updates the LastRawTime.
func (m *lockedAccounting) SaveBWRaw(ctx context.Context, latestBwa time.Time, bwTotals accounting.BWTally) error {
	m.Lock()
	defer m.Unlock()
	return m.db.SaveBWRaw(ctx, latestBwa, bwTotals)
}

// lockedBandwidthAgreement implements locking wrapper for bwagreement.DB
type lockedBandwidthAgreement struct {
	sync.Locker
	db bwagreement.DB
}

// CreateAgreement adds a new bandwidth agreement.
func (m *lockedBandwidthAgreement) CreateAgreement(ctx context.Context, a1 string, a2 bwagreement.Agreement) error {
	m.Lock()
	defer m.Unlock()
	return m.db.CreateAgreement(ctx, a1, a2)
}

// GetAgreements gets all bandwidth agreements.
func (m *lockedBandwidthAgreement) GetAgreements(ctx context.Context) ([]bwagreement.Agreement, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetAgreements(ctx)
}

// GetAgreementsSince gets all bandwidth agreements since specific time.
func (m *lockedBandwidthAgreement) GetAgreementsSince(ctx context.Context, a1 time.Time) ([]bwagreement.Agreement, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetAgreementsSince(ctx, a1)
}

// lockedIrreparable implements locking wrapper for irreparable.DB
type lockedIrreparable struct {
	sync.Locker
	db irreparable.DB
}

// Delete removes irreparable segment info based on segmentPath.
func (m *lockedIrreparable) Delete(ctx context.Context, segmentPath []byte) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Delete(ctx, segmentPath)
}

// Get returns irreparable segment info based on segmentPath.
func (m *lockedIrreparable) Get(ctx context.Context, segmentPath []byte) (*irreparable.RemoteSegmentInfo, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Get(ctx, segmentPath)
}

// IncrementRepairAttempts increments the repair attempts.
func (m *lockedIrreparable) IncrementRepairAttempts(ctx context.Context, segmentInfo *irreparable.RemoteSegmentInfo) error {
	m.Lock()
	defer m.Unlock()
	return m.db.IncrementRepairAttempts(ctx, segmentInfo)
}

// lockedOverlayCache implements locking wrapper for storage.KeyValueStore
type lockedOverlayCache struct {
	sync.Locker
	db storage.KeyValueStore
}

// Close closes the store
func (m *lockedOverlayCache) Close() error {
	m.Lock()
	defer m.Unlock()
	return m.db.Close()
}

// Delete deletes key and the value
func (m *lockedOverlayCache) Delete(a0 storage.Key) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Delete(a0)
}

// Get gets a value to store
func (m *lockedOverlayCache) Get(a0 storage.Key) (storage.Value, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Get(a0)
}

// GetAll gets all values from the store
func (m *lockedOverlayCache) GetAll(a0 storage.Keys) (storage.Values, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.GetAll(a0)
}

// Iterate iterates over items based on opts
func (m *lockedOverlayCache) Iterate(opts storage.IterateOptions, fn func(storage.Iterator) error) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Iterate(opts, fn)
}

// List lists all keys starting from start and upto limit items
func (m *lockedOverlayCache) List(start storage.Key, limit int) (storage.Keys, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.List(start, limit)
}

// Put adds a value to store
func (m *lockedOverlayCache) Put(a0 storage.Key, a1 storage.Value) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Put(a0, a1)
}

// ReverseList lists all keys in revers order
func (m *lockedOverlayCache) ReverseList(a0 storage.Key, a1 int) (storage.Keys, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.ReverseList(a0, a1)
}

// lockedRepairQueue implements locking wrapper for queue.RepairQueue
type lockedRepairQueue struct {
	sync.Locker
	db queue.RepairQueue
}

// Dequeue removes an injured segment.
func (m *lockedRepairQueue) Dequeue(ctx context.Context) (pb.InjuredSegment, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Dequeue(ctx)
}

// Enqueue adds an injured segment.
func (m *lockedRepairQueue) Enqueue(ctx context.Context, qi *pb.InjuredSegment) error {
	m.Lock()
	defer m.Unlock()
	return m.db.Enqueue(ctx, qi)
}

// Peekqueue lists limit amount of injured segments.
func (m *lockedRepairQueue) Peekqueue(ctx context.Context, limit int) ([]pb.InjuredSegment, error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Peekqueue(ctx, limit)
}

// lockedStatDB implements locking wrapper for statdb.DB
type lockedStatDB struct {
	sync.Locker
	db statdb.DB
}

// Create adds a new stats entry for node.
func (m *lockedStatDB) Create(ctx context.Context, nodeID storj.NodeID, initial *statdb.NodeStats) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Create(ctx, nodeID, initial)
}

// CreateEntryIfNotExists creates a node stats entry if it didn't already exist.
func (m *lockedStatDB) CreateEntryIfNotExists(ctx context.Context, nodeID storj.NodeID) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.CreateEntryIfNotExists(ctx, nodeID)
}

// FindInvalidNodes finds a subset of storagenodes that have stats below provided reputation requirements.
func (m *lockedStatDB) FindInvalidNodes(ctx context.Context, nodeIDs storj.NodeIDList, maxStats *statdb.NodeStats) (invalid storj.NodeIDList, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.FindInvalidNodes(ctx, nodeIDs, maxStats)
}

// Get returns node stats.
func (m *lockedStatDB) Get(ctx context.Context, nodeID storj.NodeID) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Get(ctx, nodeID)
}

// Update all parts of single storagenode's stats.
func (m *lockedStatDB) Update(ctx context.Context, request *statdb.UpdateRequest) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.Update(ctx, request)
}

// UpdateAuditSuccess updates a single storagenode's audit stats.
func (m *lockedStatDB) UpdateAuditSuccess(ctx context.Context, nodeID storj.NodeID, auditSuccess bool) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.UpdateAuditSuccess(ctx, nodeID, auditSuccess)
}

// UpdateBatch for updating multiple storage nodes' stats.
func (m *lockedStatDB) UpdateBatch(ctx context.Context, requests []*statdb.UpdateRequest) (statslist []*statdb.NodeStats, failed []*statdb.UpdateRequest, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.UpdateBatch(ctx, requests)
}

// UpdateUptime updates a single storagenode's uptime stats.
func (m *lockedStatDB) UpdateUptime(ctx context.Context, nodeID storj.NodeID, isUp bool) (stats *statdb.NodeStats, err error) {
	m.Lock()
	defer m.Unlock()
	return m.db.UpdateUptime(ctx, nodeID, isUp)
}
