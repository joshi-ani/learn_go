import boto3

# The bucket we created in docker-compose
BUCKET_NAME = 'my-bucket'
# Here we need to use the name of our service
# from the docker-compose file as domain name
ENDPOINT = 'http://localhost:9000'
# Credentials from the user we created in the
# setup (located in minio.env)
AWS_ACCESS_KEY_ID = 'VPP0fkoCyBZx8YU0QTjH'
AWS_SECRET_KEY_ID = 'iFq6k8RLJw5B0faz0cKCXeQk0w9Q8UdtaFzHuw4J'

if __name__ == "__main__":
    data = "Start uploading"

    s3_client = boto3.client(
        's3',
        aws_access_key_id=AWS_ACCESS_KEY_ID,
        aws_secret_access_key=AWS_SECRET_KEY_ID,
        endpoint_url=ENDPOINT,
    )

    s3_client.upload_file('s3-minio.py', BUCKET_NAME, 's3-minio-py-in-minio')

    print("Done, file is uploaded")