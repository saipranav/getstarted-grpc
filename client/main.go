package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/saipranav/getstarted-grpc-go/store"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8070", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := store.NewStoreClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testSave(ctx, c)
}

func testSave(ctx context.Context, c store.StoreClient) {
	in := &store.Entity{
		Id:     1,
		Name:   "Sai",
		Update: ptypes.TimestampNow(),
	}
	res1, err := c.Save(ctx, in)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", res1.Entity)

	res2, err := c.Restore(ctx, &store.EntityRequest{Id: in.Id})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", res2.Entity)

	fmt.Printf("%t\n", res1.Entity.Id == res2.Entity.Id)
}
