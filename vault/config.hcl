storage "raft"{
    path = "./vault/data"
    node_id = "1"
}


listener "tcp"{
    address = "127.0.0.1:82000"
    tls_disable = "true"
}


api_addr = "127.0.0.1:82000"
cluster_addr = "127.0.0.1:82001"
ui = true