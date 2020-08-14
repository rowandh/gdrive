// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package asset_test

import (
	"context"

	asset "cloud.google.com/go/asset/apiv1"
	"google.golang.org/api/iterator"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_ExportAssets() {
	// import assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"

	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.ExportAssetsRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ExportAssets(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_BatchGetAssetsHistory() {
	// import assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"

	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.BatchGetAssetsHistoryRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.BatchGetAssetsHistory(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateFeed() {
	// import assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"

	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.CreateFeedRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateFeed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetFeed() {
	// import assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"

	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.GetFeedRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetFeed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListFeeds() {
	// import assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"

	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.ListFeedsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListFeeds(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateFeed() {
	// import assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"

	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.UpdateFeedRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateFeed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteFeed() {
	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.DeleteFeedRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteFeed(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_SearchAllResources() {
	// import assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.SearchAllResourcesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.SearchAllResources(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_SearchAllIamPolicies() {
	// import assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := asset.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assetpb.SearchAllIamPoliciesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.SearchAllIamPolicies(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
