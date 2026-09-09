package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	ctx := context.Background()

	endpoint := "localhost:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal("Error creating MinIO client:", err)
		return
	}
	fmt.Println("=== MinIO Client Initialized ===")

	bucketName := "matin"

	fmt.Println("\n[1] Press ENTER to create bucket...")
	fmt.Scanln()

	err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		log.Println("MakeBucket status/error:", err)
	} else {
		fmt.Printf("Bucket '%s' created successfully.\n", bucketName)
	}

	////////////////////////////////////////
	fmt.Println("\n[2] Press ENTER to run FPutObject (Uploading local './a.cpp' as 'a.cpp')...")
	fmt.Scanln()

	finfo, err := client.FPutObject(ctx, bucketName, "a.cpp", "./a.cpp", minio.PutObjectOptions{
		ContentType: "text/x-c",
	})
	if err != nil {
		log.Println("FPutObject error:", err)
	} else {
		fmt.Printf("FPutObject success: Key=%s, Size=%d bytes, ETag=%s\n", finfo.Key, finfo.Size, finfo.ETag)
	}
	////////////////////////////////////////
	fmt.Println("\n[3] Press ENTER to run PutObject (Uploading in-memory stream as 'docs/hello.txt')...")
	fmt.Scanln()

	content := []byte("Hello MinIO from Go Stream!")
	reader := bytes.NewReader(content)
	pinfo, err := client.PutObject(ctx, bucketName, "docs/hello.txt", reader, int64(len(content)), minio.PutObjectOptions{
		ContentType: "text/plain",
	})
	if err != nil {
		log.Println("PutObject error:", err)
	} else {
		fmt.Printf("PutObject success: Key=%s, Size=%d bytes\n", pinfo.Key, pinfo.Size)
	}

	////////////////////////////////////////
	fmt.Println("\n[4] Press ENTER to list all objects in bucket...")
	fmt.Scanln()

	fmt.Println("Objects in bucket:")
	objectCh := client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectCh {
		if object.Err != nil {
			log.Println("ListObjects item error:", object.Err)

			continue
		}
		fmt.Printf(" - Key: %-20s | Size: %d bytes | LastModified: %s\n", object.Key, object.Size, object.LastModified.Format("15:04:05"))
	}

	////////////////////////////////////////
	fmt.Println("\n[5] Press ENTER to read 'docs/hello.txt' stream directly...")
	fmt.Scanln()

	obj, err := client.GetObject(ctx, bucketName, "docs/hello.txt", minio.GetObjectOptions{})
	if err != nil {
		log.Println("GetObject error:", err)
	} else {
		data, err := io.ReadAll(obj)
		if err != nil {
			log.Println("Error reading object body:", err)
		} else {
			fmt.Printf("Content of 'docs/hello.txt': %s\n", string(data))
		}
		obj.Close()
	}

	////////////////////////////////////////
	fmt.Println("\n[6] Press ENTER to download 'a.cpp' to './downloaded_a.cpp'...")
	fmt.Scanln()

	err = client.FGetObject(ctx, bucketName, "a.cpp", "./downloaded_a.cpp", minio.GetObjectOptions{})
	if err != nil {
		log.Println("FGetObject error:", err)
	} else {
		fmt.Println("FGetObject success: Saved to ./downloaded_a.cpp")
	}

	////////////////////////////////////////
	fmt.Println("\n[7] Press ENTER to generate Presigned Download URL (valid for 15 mins)...")
	fmt.Scanln()

	presignedURL, err := client.PresignedGetObject(ctx, bucketName, "a.cpp", 15*time.Minute, nil)
	if err != nil {
		log.Println("PresignedGetObject error:", err)
	} else {
		fmt.Println("Presigned URL:")
		fmt.Println(presignedURL.String())
	}

	////////////////////////////////////////
	fmt.Println("\n[8] Press ENTER to delete uploaded objects...")
	fmt.Scanln()

	err = client.RemoveObject(ctx, bucketName, "a.cpp", minio.RemoveObjectOptions{})
	if err != nil {
		log.Println("RemoveObject (a.cpp) error:", err)
	} else {
		fmt.Println("Deleted 'a.cpp'")
	}

	err = client.RemoveObject(ctx, bucketName, "docs/hello.txt", minio.RemoveObjectOptions{})
	if err != nil {
		log.Println("RemoveObject (docs/hello.txt) error:", err)
	} else {
		fmt.Println("Deleted 'docs/hello.txt'")
	}

	////////////////////////////////////////
	fmt.Println("\n[9] Press ENTER to remove bucket 'matin'...")
	fmt.Scanln()

	err = client.RemoveBucket(ctx, bucketName)
	if err != nil {
		log.Println("RemoveBucket error:", err)
	} else {
		fmt.Printf("Bucket '%s' removed successfully.\n", bucketName)
	}

	fmt.Println("\n=== Test Suite Completed ===")
}
