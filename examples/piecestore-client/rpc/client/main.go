// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli"
	"github.com/zeebo/errs"
	"google.golang.org/grpc"

	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/piecestore/rpc/client"
	"storj.io/storj/pkg/provider"
)

var ctx = context.Background()
var argError = errs.Class("argError")

func main() {

	app := cli.NewApp()

	ca, err := provider.NewCA(ctx, 12, 4)
	if err != nil {
		log.Fatal(err)
	}
	identity, err := ca.NewIdentity()
	if err != nil {
		log.Fatal(err)
	}
	identOpt, err := identity.DialOption()
	if err != nil {
		log.Fatal(err)
	}

	// Set up connection with rpc server
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(":7777", identOpt)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer printError(conn.Close)

	psClient, err := client.NewPSClient(conn, 1024*32, identity.Key.(*ecdsa.PrivateKey))
	if err != nil {
		log.Fatalf("could not initialize PSClient: %s", err)
	}

	app.Commands = []cli.Command{
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "upload data",
			Action: func(c *cli.Context) error {

				if c.Args().Get(0) == "" {
					return argError.New("No input file specified")
				}

				file, err := os.Open(c.Args().Get(0))
				if err != nil {
					return err
				}
				// Close the file when we are done
				defer printError(file.Close)

				fileInfo, err := file.Stat()
				if err != nil {
					return err
				}

				if fileInfo.IsDir() {
					return argError.New(fmt.Sprintf("Path (%s) is a directory, not a file", c.Args().Get(0)))
				}

				var length = fileInfo.Size()
				var ttl = time.Now().Add(24 * time.Hour)

				// Created a section reader so that we can concurrently retrieve the same file.
				dataSection := io.NewSectionReader(file, 0, length)

				id := client.NewPieceID()

				if err := psClient.Put(context.Background(), id, dataSection, ttl, &pb.PayerBandwidthAllocation{}); err != nil {
					fmt.Printf("Failed to Store data of id: %s\n", id)
					return err
				}

				fmt.Printf("Successfully stored file of id: %s\n", id)

				return nil
			},
		},
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "download data",
			Action: func(c *cli.Context) error {
				const (
					id int = iota
					outputDir
				)

				if c.Args().Get(id) == "" {
					return argError.New("No id specified")
				}

				if c.Args().Get(outputDir) == "" {
					return argError.New("No output file specified")
				}

				_, err := os.Stat(c.Args().Get(outputDir))
				if err != nil && !os.IsNotExist(err) {
					return err
				}

				if err == nil {
					return argError.New("File already exists")
				}

				if err = os.MkdirAll(filepath.Dir(c.Args().Get(outputDir)), 0700); err != nil {
					return err
				}

				// Create File on file system
				dataFile, err := os.OpenFile(c.Args().Get(outputDir), os.O_RDWR|os.O_CREATE, 0755)
				if err != nil {
					return err
				}
				defer printError(dataFile.Close)

				pieceInfo, err := psClient.Meta(context.Background(), client.PieceID(c.Args().Get(id)))
				if err != nil {
					errRemove := os.Remove(c.Args().Get(outputDir))
					if errRemove != nil {
						log.Println(errRemove)
					}
					return err
				}

				rr, err := psClient.Get(ctx, client.PieceID(c.Args().Get(id)), pieceInfo.Size, &pb.PayerBandwidthAllocation{})
				if err != nil {
					fmt.Printf("Failed to retrieve file of id: %s\n", c.Args().Get(id))
					errRemove := os.Remove(c.Args().Get(outputDir))
					if errRemove != nil {
						log.Println(errRemove)
					}
					return err
				}

				reader, err := rr.Range(ctx, 0, pieceInfo.Size)
				if err != nil {
					fmt.Printf("Failed to retrieve file of id: %s\n", c.Args().Get(id))
					errRemove := os.Remove(c.Args().Get(outputDir))
					if errRemove != nil {
						log.Println(errRemove)
					}
					return err
				}

				_, err = io.Copy(dataFile, reader)
				if err != nil {
					fmt.Printf("Failed to retrieve file of id: %s\n", c.Args().Get(id))
					errRemove := os.Remove(c.Args().Get(outputDir))
					if errRemove != nil {
						log.Println(errRemove)
					}
				} else {
					fmt.Printf("Successfully retrieved file of id: %s\n", c.Args().Get(id))
				}

				return reader.Close()
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"x"},
			Usage:   "delete data",
			Action: func(c *cli.Context) error {
				if c.Args().Get(0) == "" {
					return argError.New("Missing data Id")
				}
				err = psClient.Delete(context.Background(), client.PieceID(c.Args().Get(0)))

				return err
			},
		},
		{
			Name:    "stat",
			Aliases: []string{"s"},
			Usage:   "retrieve stats",
			Action: func(c *cli.Context) error {
				var summary *pb.StatSummary
				summary, err := psClient.Stats(context.Background())
				if err != nil {
					return err
				}

				log.Printf("Space Used: %v, Space Available: %v\n", summary.UsedSpace, summary.AvailableSpace)

				return nil
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func printError(fn func() error) {
	err := fn()
	if err != nil {
		fmt.Println(err)
	}
}
