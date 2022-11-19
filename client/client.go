package main

import (
	"context"
	"log"
	pb "rntmailinglist/proto"
	"time"

	"github.com/alexflint/go-arg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func logResponse(res *pb.EmailResponse, err error) {
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	if res.EmailEntry == nil {
		log.Fatalf("Email not found")
	} else {
		log.Fatalf("Response: %v", res.EmailEntry)
	}
}

func createEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry {
	log.Printf("Create email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateEmail(ctx, &pb.CreateEmailRequest{EmailAddr: addr})
	logResponse(res, err)

	return res.EmailEntry
}

func getEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry {
	log.Printf("Get email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetEmail(ctx, &pb.GetEmailRequest{EmailAddr: addr})
	logResponse(res, err)

	return res.EmailEntry
}

func getEmailBatch(client pb.MailingListServiceClient, count int, page int) {
	log.Printf("Get email batch")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetEmailBatch(ctx, &pb.GetEmailBatchRequest{Count: int32(count), Page: int32(page)})
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Println("response:")
	for i := 0; i < len(res.EmailEntries); i++ {
		log.Printf("Item [%v of %v]: %s", i+1, len(res.EmailEntries), res.EmailEntries[i])
	}
}

func updateEmail(client pb.MailingListServiceClient, entry pb.EmailEntry) *pb.EmailEntry {
	log.Printf("Update email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UpdateEmail(ctx, &pb.UpdateEmailRequest{EmailEntry: &entry})
	logResponse(res, err)

	return res.EmailEntry
}

func deleteEmail(client pb.MailingListServiceClient, addr string) *pb.EmailEntry {
	log.Printf("Delete email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.DeleteEmail(ctx, &pb.DeleteEmailRequest{EmailAddr: addr})
	logResponse(res, err)

	return res.EmailEntry
}

var args struct {
	GrpcAddr string `args"env:MAILINGLIST_GRPC_ADDR`
}

func main() {
	arg.MustParse(&args)

	if args.GrpcAddr == "" {
		args.GrpcAddr = ":8081"
	}

	conn, err := grpc.Dial(args.GrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMailingListServiceClient(conn)

	// newEmail := createEmail(client, "rnr@rrr.rrr")
	// newEmail.ConfirmedAt = 10000
	// updateEmail(client, *newEmail)
	// deleteEmail(client, newEmail.Email)
	// getEmailBatch(client, 5, 5)

	getEmailBatch(client, 3, 1)
	getEmailBatch(client, 3, 2)
	getEmailBatch(client, 3, 3)

}
