// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"log"
	"net"

	pc "github.com/opiproject/opi-api/common/v1/gen/go"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
)

// DoBackend executes the back end code
func DoBackend(ctx context.Context, conn grpc.ClientConnInterface) error {
	nvme := pb.NewNVMfRemoteControllerServiceClient(conn)
	null := pb.NewNullDebugServiceClient(conn)
	aio := pb.NewAioControllerServiceClient(conn)

	err := executeNVMfRemoteController(ctx, nvme)
	if err != nil {
		return err
	}
	err = executeNullDebug(ctx, null)
	if err != nil {
		return err
	}
	err = executeAioController(ctx, aio)
	if err != nil {
		return err
	}
	return nil
}

func executeNVMfRemoteController(ctx context.Context, c4 pb.NVMfRemoteControllerServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewNVMfRemoteControllerServiceClient")
	log.Printf("=======================================")

	addr, err := net.LookupIP("spdk")
	if err != nil {
		return err
	}
	rr0, err := c4.CreateNVMfRemoteController(ctx, &pb.CreateNVMfRemoteControllerRequest{
		NvMfRemoteController: &pb.NVMfRemoteController{
			Id:      &pc.ObjectKey{Value: "OpiNvme8"},
			Trtype:  pb.NvmeTransportType_NVME_TRANSPORT_TCP,
			Adrfam:  pb.NvmeAddressFamily_NVMF_ADRFAM_IPV4,
			Traddr:  addr[0].String(),
			Trsvcid: 4444,
			Subnqn:  "nqn.2016-06.io.spdk:cnode1",
			Hostnqn: "nqn.2014-08.org.nvmexpress:uuid:feb98abe-d51f-40c8-b348-2753f3571d3c"}})
	if err != nil {
		return err
	}
	log.Printf("Connected NVMf: %v", rr0)
	rr2, err := c4.NVMfRemoteControllerReset(ctx, &pb.NVMfRemoteControllerResetRequest{Id: &pc.ObjectKey{Value: "OpiNvme8"}})
	if err != nil {
		return err
	}
	log.Printf("Reset NVMf: %v", rr2)
	rr3, err := c4.ListNVMfRemoteControllers(ctx, &pb.ListNVMfRemoteControllersRequest{})
	if err != nil {
		return err
	}
	log.Printf("List NVMf: %v", rr3)
	rr4, err := c4.GetNVMfRemoteController(ctx, &pb.GetNVMfRemoteControllerRequest{Name: "OpiNvme8"})
	if err != nil {
		return err
	}
	log.Printf("Got NVMf: %v", rr4)
	rr5, err := c4.NVMfRemoteControllerStats(ctx, &pb.NVMfRemoteControllerStatsRequest{Id: &pc.ObjectKey{Value: "OpiNvme8"}})
	if err != nil {
		return err
	}
	log.Printf("Stats NVMf: %v", rr5)
	rr1, err := c4.DeleteNVMfRemoteController(ctx, &pb.DeleteNVMfRemoteControllerRequest{Name: "OpiNvme8"})
	if err != nil {
		return err
	}
	log.Printf("Disconnected NVMf: %v -> %v", rr0, rr1)
	return nil
}

func executeNullDebug(ctx context.Context, c1 pb.NullDebugServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewNullDebugServiceClient")
	log.Printf("=======================================")
	rs1, err := c1.CreateNullDebug(ctx, &pb.CreateNullDebugRequest{NullDebug: &pb.NullDebug{Handle: &pc.ObjectKey{Value: "OpiNull9"}}})
	if err != nil {
		return err
	}
	log.Printf("Added Null: %v", rs1)
	rs3, err := c1.UpdateNullDebug(ctx, &pb.UpdateNullDebugRequest{NullDebug: &pb.NullDebug{Handle: &pc.ObjectKey{Value: "OpiNull9"}}})
	if err != nil {
		return err
	}
	log.Printf("Updated Null: %v", rs3)
	rs4, err := c1.ListNullDebugs(ctx, &pb.ListNullDebugsRequest{})
	if err != nil {
		return err
	}
	log.Printf("Listed Null: %v", rs4)
	rs5, err := c1.GetNullDebug(ctx, &pb.GetNullDebugRequest{Name: "OpiNull9"})
	if err != nil {
		return err
	}
	log.Printf("Got Null: %s", rs5.Handle.Value)
	rs6, err := c1.NullDebugStats(ctx, &pb.NullDebugStatsRequest{Handle: &pc.ObjectKey{Value: "OpiNull9"}})
	if err != nil {
		return err
	}
	log.Printf("Stats Null: %s", rs6.Stats)
	rs2, err := c1.DeleteNullDebug(ctx, &pb.DeleteNullDebugRequest{Name: "OpiNull9"})
	if err != nil {
		return err
	}
	log.Printf("Deleted Null: %v -> %v", rs1, rs2)
	return nil
}

func executeAioController(ctx context.Context, c2 pb.AioControllerServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewAioControllerServiceClient")
	log.Printf("=======================================")
	ra1, err := c2.CreateAioController(ctx, &pb.CreateAioControllerRequest{AioController: &pb.AioController{Handle: &pc.ObjectKey{Value: "OpiAio4"}, Filename: "/tmp/aio_bdev_file"}})
	if err != nil {
		return err
	}
	log.Printf("Added Aio: %v", ra1)
	ra3, err := c2.UpdateAioController(ctx, &pb.UpdateAioControllerRequest{AioController: &pb.AioController{Handle: &pc.ObjectKey{Value: "OpiAio4"}, Filename: "/tmp/aio_bdev_file"}})
	if err != nil {
		return err
	}
	log.Printf("Updated Aio: %v", ra3)
	ra4, err := c2.ListAioControllers(ctx, &pb.ListAioControllersRequest{})
	if err != nil {
		return err
	}
	log.Printf("Listed Aio: %v", ra4)
	ra5, err := c2.GetAioController(ctx, &pb.GetAioControllerRequest{Name: "OpiAio4"})
	if err != nil {
		return err
	}
	log.Printf("Got Aio: %s", ra5.Handle.Value)
	ra6, err := c2.AioControllerStats(ctx, &pb.AioControllerStatsRequest{Handle: &pc.ObjectKey{Value: "OpiAio4"}})
	if err != nil {
		return err
	}
	log.Printf("Stats Aio: %s", ra6.Stats)
	ra2, err := c2.DeleteAioController(ctx, &pb.DeleteAioControllerRequest{Name: "OpiAio4"})
	if err != nil {
		return err
	}
	log.Printf("Deleted Aio: %v -> %v", ra1, ra2)
	return nil
}
