// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package audit

import (
	"context"
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"storj.io/storj/internal/testidentity"
	"storj.io/storj/internal/teststorj"
	"storj.io/storj/pkg/auth"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/pointerdb"
	"storj.io/storj/pkg/storage/meta"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storage/teststore"
)

var (
	ctx       = context.Background()
	ErrNoList = errors.New("list error: failed to get list")
	ErrNoNum  = errors.New("num error: failed to get num")
)

func TestAuditSegment(t *testing.T) {
	type pathCount struct {
		path  storj.Path
		count int
	}

	ca, err := testidentity.NewTestCA(ctx)
	assert.NoError(t, err)
	identity, err := ca.NewIdentity()
	assert.NoError(t, err)

	// note: to simulate better,
	// change limit in library to 5 in
	// list api call, default is  0 == 1000 listing
	tests := []struct {
		bm   string
		path storj.Path
	}{
		{
			bm:   "success-1",
			path: "folder1/file1",
		},
		{
			bm:   "success-2",
			path: "foodFolder1/file1/file2",
		},
		{
			bm:   "success-3",
			path: "foodFolder1/file1/file2/foodFolder2/file3",
		},
		{
			bm:   "success-4",
			path: "projectFolder/project1.txt/",
		},
		{
			bm:   "success-5",
			path: "newProjectFolder/project2.txt",
		},
		{
			bm:   "success-6",
			path: "Pictures/image1.png",
		},
		{
			bm:   "success-7",
			path: "Pictures/Nature/mountains.png",
		},
		{
			bm:   "success-8",
			path: "Pictures/City/streets.png",
		},
		{
			bm:   "success-9",
			path: "Pictures/Animals/Dogs/dogs.png",
		},
		{
			bm:   "success-10",
			path: "Nada/ビデオ/😶",
		},
	}

	ctx = auth.WithAPIKey(ctx, nil)

	// PointerDB instantiation
	db := teststore.New()
	c := pointerdb.Config{MaxInlineSegmentSize: 8000}

	cache := overlay.NewCache(teststore.New(), nil)

	pointers := pointerdb.NewServer(db, cache, zap.NewNop(), c, identity)

	// create a pdb client and instance of audit
	cursor := NewCursor(pointers)

	// put 10 paths in db
	t.Run("putToDB", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.bm, func(t *testing.T) {
				assert1 := assert.New(t)

				// create a pointer and put in db
				putRequest := makePutRequest(tt.path)

				// create putreq. object
				req := &pb.PutRequest{Path: tt.path, Pointer: putRequest.Pointer}

				// put pointer into db
				_, err := pointers.Put(ctx, req)
				if err != nil {
					t.Fatalf("failed to put %v: error: %v", req.Pointer, err)
					assert1.NotNil(err)
				}
				if err != nil {
					t.Error("cant instantiate the piece store client")
				}
			})
		}
	})

	t.Run("NextStripe", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.bm, func(t *testing.T) {
				assert1 := assert.New(t)
				stripe, err := cursor.NextStripe(ctx)
				if err != nil {
					assert1.Error(err)
					assert1.Nil(stripe)
				}
				if stripe != nil {
					assert1.Nil(err)
				}
			})
		}
	})

	// test to see how random paths are
	t.Run("probabilisticTest", func(t *testing.T) {
		listRes, err := pointers.List(ctx, &pb.ListRequest{
			Prefix:     "",
			StartAfter: "",
			EndBefore:  "",
			Recursive:  true,
			Limit:      10,
			MetaFlags:  meta.None,
		})
		if err != nil {
			t.Error(ErrNoList)
		}

		list := listRes.GetItems()

		// get count of items picked at random
		uniquePathCounted := []pathCount{}
		pathCounter := []pathCount{}

		// get a list of 100 paths generated from random
		for i := 0; i < 100; i++ {
			randomNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
			if err != nil {
				t.Error(ErrNoNum)
			}
			pointerItem := list[randomNum.Int64()]
			path := pointerItem.Path
			val := pathCount{path: path, count: 1}
			pathCounter = append(pathCounter, val)
		}

		// get a count for paths in list
		for _, pc := range pathCounter {
			skip := false
			for i, up := range uniquePathCounted {
				if reflect.DeepEqual(pc.path, up.path) {
					up.count = up.count + 1
					uniquePathCounted[i] = up
					skip = true
					break
				}
			}
			if !skip {
				uniquePathCounted = append(uniquePathCounted, pc)
			}
		}

		// Section: binomial test for randomness
		n := float64(100) // events
		p := float64(.10) // theoretical probability of getting  1/10 paths
		m := n * p
		s := math.Sqrt(m * (1 - p)) // binomial distribution

		// if values fall outside of the critical values of test statistics (ie Z value)
		// in a 2-tail test
		// we can assume, 95% confidence, it's not sampling according to a 10% probability
		for _, v := range uniquePathCounted {
			z := (float64(v.count) - m) / s
			if z <= -1.96 || z >= 1.96 {
				t.Log(false)
			} else {
				t.Log(true)
			}
		}
	})
}

func makePutRequest(path storj.Path) pb.PutRequest {
	var rps []*pb.RemotePiece
	rps = append(rps, &pb.RemotePiece{
		PieceNum: 1,
		NodeId:   teststorj.NodeIDFromString("testId"),
	})
	pr := pb.PutRequest{
		Path: path,
		Pointer: &pb.Pointer{
			Type: pb.Pointer_REMOTE,
			Remote: &pb.RemoteSegment{
				Redundancy: &pb.RedundancyScheme{
					Type:             pb.RedundancyScheme_RS,
					MinReq:           1,
					Total:            3,
					RepairThreshold:  2,
					SuccessThreshold: 3,
					ErasureShareSize: 2,
				},
				PieceId:      "testId",
				RemotePieces: rps,
			},
			SegmentSize: int64(10),
		},
	}
	return pr
}
