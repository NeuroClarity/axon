# Infrastructure set up for Axon

## AWS Access
AWS gets its credentials from ~/.aws/credentials. This path can be overriden by setting the `AWS_SHARED_CREDENTIALS_FILE` environment variable. The credentials file has been included here for reference.
The credentials can also be set with environment variables. More details here: https://docs.aws.amazon.com/cli/latest/topic/config-vars.html#credentials

## RabbitMQ Access
For now, the username and password will be harcoded until we implement a Kubernetes based solution. 
