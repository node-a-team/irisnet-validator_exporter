# irisnet-validator_exporter :satellite:
![CreatePlan](https://img.shields.io/badge/relase-v0.2.0-red)
![CreatePlan](https://img.shields.io/badge/go-1.13.1%2B-blue)
![CreatePlan](https://img.shields.io/badge/license-Apache--2.0-green)

Prometheus exporter for IRISnet Validators


## Introduction
This exporter is for monitoring information which is not provided from Tendermint’s basic Prometheus exporter (localhost:26660), and other specific information monitoring purposes

## Install
```bash
mkdir exporter && cd exporter

wget https://github.com/node-a-team/irisnet-validator_exporter/releases/download/v0.2.0/irisnet-validator_exporter.tar.gz  && sha256sum irisnet-validator_exporter.tar.gz | fgrep 756eee5d9af2214ca10378198de690d1d6d1ba84ab3c7a514ec80d4d3e8986ce && tar -zxvf irisnet-validator_exporter.tar.gz ||  echo "Bad Binary!"
```

## Config
1. Modify to the appropriate RPC and REST server address
2. Modify the value of ```operatorAddr``` to the operator address of the validator you want to monitor.
3. You can change the service port(default: 26661)
```bash
vi config.toml
```
```bash
# TOML Document for IRISnet-Validator Exporter(Pometheus & Grafana)

title = "TOML Document"

[Servers]
        [Servers.addr]
        rpc = "localhost:26657"
        rest = "localhost:1317"

[Validator]
operatorAddr = "iva1w2dakpuvh9mglcs54wayta5dyv8vj853y6jz9e"

[Options]
listenPort = "26661"

```

## Start
  
```bash
./irisnet-validator_exporter {path to config.toml}

// ex)
./irisnet-validator_exporter /data/iris/exporter
```

## Use systemd service
  
```sh
# Make log directory & file
sudo mkdir /var/log/userLog  
sudo touch /var/log/userLog/irisnet-validator_exporter.log  
# user: iris
sudo chown iris:iris /var/log/userLog/irisnet-validator_exporter.log

# $HOME: /data/iris
# Path to config.toml: /data/iris/exporter
sudo tee /etc/systemd/system/irisnet-validator_exporter.service > /dev/null <<EOF
[Unit]
Description=IRISnet Validator Exporter
After=network-online.target

[Service]
User=cosmos
WorkingDirectory=/data/iris
ExecStart=/data/iris/exporter/irisnet-validator_exporter \
        /data/iris/exporter
StandardOutput=file:/var/log/userLog/irisnet-validator_exporter.log
StandardError=file:/var/log/userLog/irisnet-validator_exporter.log
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl enable irisnet-validator_exporter.service
sudo systemctl restart irisnet-validator_exporter.service


## log
tail -f /var/log/userLog/irisnet-validator_exporter.log
```
