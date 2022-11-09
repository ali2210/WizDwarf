package storj_bucket

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/ali2210/wizdwarf/other/bucket/storj_bucket/bucket"
	"storj.io/uplink"
)

type BucketObject struct {
	Ctx             context.Context
	Key             string
	Bucket          string
	NodeCredentials *uplink.Access
}

// StoreJCredentials implements Storj_Proteins_Bucket
func (boj *BucketObject) StoreJCredentials(user_phrase string, passphrase ...string) (*bucket.Bucket_Error, *uplink.Access) {

	const api_key string = "1dfJiTZDDGnTZ9gCmZDAFDp1gLCBSsijs838VhzgskptVYifaWwwDwE51cAqTwATc1L7wwfaQ8Jz723JMLjEvFntG6ZccHwgm7GpvaqCABtFMh4rtU2K"
	const address string = "12L9ZFwhzVpuEKMUNUqkaTLGzwY9G24tbiigLiXpmZWKwmcNDDs@eu1.storj.io:7777"

	app_cred, err := uplink.RequestAccessWithPassphrase(boj.Ctx, address, api_key, passphrase[0])
	if err != nil {
		log.Fatalln(err.Error())
		return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}, &uplink.Access{}
	}

	perm := uplink.FullPermission()
	perm.NotBefore = time.Now().Add(-2 * time.Minute)
	perm.NotAfter = time.Now().Add(12 * time.Hour)

	user_cred, err := app_cred.Share(perm, uplink.SharePrefix{Bucket: boj.Bucket, Prefix: boj.Key})
	if err != nil {
		log.Fatalln(err.Error())
		return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}, &uplink.Access{}
	}

	serialized, err := user_cred.Serialize()
	if err != nil && strings.Contains(serialized, " ") {
		log.Fatalln(err.Error())
		return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}, &uplink.Access{}
	}

	access, err := uplink.ParseAccess(serialized)
	if err != nil {
		log.Fatalln(err.Error())
		return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}, &uplink.Access{}
	}

	boj.NodeCredentials = access

	log.Println("Application credentials generate process will almost complete: ...")
	return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Ok}, boj.NodeCredentials
}

// DownloadObject implements Storj_Proteins_Bucket
func (boj *BucketObject) DownloadObject(project *uplink.Project, filetype string, filename ...string) *bucket.Bucket_Error {

	download, err := project.DownloadObject(boj.Ctx, boj.Bucket, boj.Key, nil)
	if err != nil {
		return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	log.Println("Content key: ", boj.Key, " bucket:...", boj.Bucket)

	defer download.Close()

	read, err := os.ReadFile("app_data/" + filename[0] + filetype)
	if err != nil {
		return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	content, err := io.ReadAll(download)
	if err != nil {
		return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	if !bytes.Equal(read, content) {
		return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	log.Println("Downloading process almost completed: ...", download)
	return &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Ok}
}

// ListObject implements Storj_Proteins_Bucket
func (boj *BucketObject) ListObject(project *uplink.Project) (*uplink.Object, *bucket.Bucket_Error) {

	var object *uplink.Object

	for project.ListObjects(boj.Ctx, boj.Bucket, nil).Next() {

		if err := project.ListObjects(boj.Ctx, boj.Bucket, nil).Err(); err != nil {

			return nil, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
		}

		object = project.ListObjects(boj.Ctx, boj.Bucket, nil).Item()
	}

	log.Println("List of contents in bucket: ....", object)

	return object, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Ok}
}

// StoreObject implements Storj_Proteins_Bucket
func (bo *BucketObject) StoreObject(user_phrase string, salt, filename, filetype string, passphrase ...string) (*uplink.Project, *bucket.Bucket_Error) {

	if reflect.DeepEqual(bo.NodeCredentials, &uplink.Access{}) {
		return &uplink.Project{}, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	iso_pub_key, err := uplink.DeriveEncryptionKey(passphrase[0], []byte(salt))
	if err != nil && !reflect.DeepEqual(iso_pub_key, &uplink.EncryptionKey{}) {
		log.Fatalln(err.Error())
		return &uplink.Project{}, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	// if err := bo.NodeCredentials.OverrideEncryptionKey(bo.Bucket, user_phrase+"/", iso_pub_key); err != nil {
	// 	log.Fatalln(err.Error())
	// 	return &uplink.Project{}, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	// }

	blockApp, err := uplink.OpenProject(bo.Ctx, bo.NodeCredentials)
	if err != nil && reflect.DeepEqual(blockApp, &uplink.Project{}) {
		log.Fatalln(err.Error())
		return &uplink.Project{}, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}
	defer blockApp.Close()

	blockBucket, err := blockApp.EnsureBucket(bo.Ctx, bo.Bucket)
	if err != nil && reflect.DeepEqual(blockBucket, &uplink.Bucket{}) {
		log.Fatalln(err.Error())
		return &uplink.Project{}, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	uploadObject, err := blockApp.UploadObject(bo.Ctx, bo.Bucket, bo.Key, nil)
	if err != nil && reflect.DeepEqual(uploadObject, nil) {
		log.Fatalln(err.Error())
		return &uplink.Project{}, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	read, err := os.ReadFile("app_data/" + filename + filetype)
	if err != nil {
		log.Fatalln(err.Error())
		return &uplink.Project{}, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	if _, err := io.Copy(uploadObject, bytes.NewBuffer(read)); err != nil {
		uploadObject.Abort()
		log.Fatalln(err.Error())
		return &uplink.Project{}, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	if err := uploadObject.Commit(); err != nil {
		log.Fatalln(err.Error())
		return &uplink.Project{}, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Error}
	}

	log.Println("File upload completed: ...", uploadObject.Info())

	return blockApp, &bucket.Bucket_Error{Bucket: bucket.Bucket_Error_Category_Ok}
}

type Storj_Proteins_Bucket interface {
	StoreObject(user_phrase string, salt, filename, filetype string, passphrase ...string) (*uplink.Project, *bucket.Bucket_Error)
	ListObject(project *uplink.Project) (*uplink.Object, *bucket.Bucket_Error)
	DownloadObject(project *uplink.Project, filetype string, filename ...string) *bucket.Bucket_Error
	StoreJCredentials(user_phrase string, passphrase ...string) (*bucket.Bucket_Error, *uplink.Access)
}

func New_Bucket(ctx context.Context, key, bucket string) Storj_Proteins_Bucket {
	return &BucketObject{Ctx: ctx, Key: key, Bucket: bucket, NodeCredentials: &uplink.Access{}}
}
