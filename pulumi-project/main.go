package main

import (
    "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Crea un bucket en Santiago con Uniform Bucket-Level Access habilitado
        bucket, err := storage.NewBucket(ctx, "santiago-bucket", &storage.BucketArgs{
            Location:                   pulumi.String("southamerica-west1"),
            UniformBucketLevelAccess:   pulumi.Bool(true), // Activar UBLA
        })
        if err != nil {
            return err
        }

        // Exporta el nombre del bucket como salida
        ctx.Export("bucketName", bucket.Name)
        return nil
    })
}
