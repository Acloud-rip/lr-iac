## Requisitos previos

1.  Instalar [Go](https://go.dev/dl/).
    Instalar Go versión 1.21.5:

    ````bash
    curl -OL https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin

        ```

    ````

2.  Instalar el [Google Cloud SDK](https://cloud.google.com/sdk/docs/install).

    ````bash
    curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-412.0.0-linux-x86_64.tar.gz
    tar -xvzf google-cloud-sdk-412.0.0-linux-x86_64.tar.gz
    ./google-cloud-sdk/install.sh
    export PATH=$PATH:$HOME/google-cloud-sdk/bin
        ```

    ````

3.  Instalar Pulumi:
    ```bash
    curl -fsSL https://get.pulumi.com | sh
    export PATH=$PATH:$HOME/.pulumi/bin
    ```
