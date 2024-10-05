#!/bin/bash

# Definir o diretório base como o diretório onde o script está localizado
BASE_DIR=$(dirname "$(realpath "$0")")

# Verificar se o arquivo .env existe
if [ ! -f "$BASE_DIR/dev.env" ]; then
  echo "Arquivo .env não encontrado em $BASE_DIR"
  exit 1
fi

# Carregar variáveis de ambiente do arquivo .env
set -a
source "$BASE_DIR/dev.env"
set +a

# Verificar se as variáveis de ambiente foram carregadas
if [ -z "$POSTGRES_USER" ] || [ -z "$POSTGRES_PASSWORD" ] || [ -z "$POSTGRES_DB" ] || [ -z "$DATA_DIR" ] || [ -z "$PGADMIN_DEFAULT_EMAIL" ] || [ -z "$PGADMIN_DEFAULT_PASSWORD" ]; then
  echo "Uma ou mais variáveis de ambiente não foram carregadas corretamente do arquivo .env"
  exit 1
fi

# Fazer uma cópia temporária do arquivo docker-compose-dev.yml
cp "$BASE_DIR/docker-compose-dev.yml" "$BASE_DIR/docker-compose-dev.yml.bak"

# Substituir os valores no arquivo docker-compose-dev.yml
sed -i "s|\${POSTGRES_USER}|$POSTGRES_USER|g" "$BASE_DIR/docker-compose-dev.yml"
sed -i "s|\${POSTGRES_PASSWORD}|$POSTGRES_PASSWORD|g" "$BASE_DIR/docker-compose-dev.yml"
sed -i "s|\${POSTGRES_DB}|$POSTGRES_DB|g" "$BASE_DIR/docker-compose-dev.yml"
sed -i "s|\${DATA_DIR}|$DATA_DIR|g" "$BASE_DIR/docker-compose-dev.yml"
sed -i "s|\${PGADMIN_DEFAULT_EMAIL}|$PGADMIN_DEFAULT_EMAIL|g" "$BASE_DIR/docker-compose-dev.yml"
sed -i "s|\${PGADMIN_DEFAULT_PASSWORD}|$PGADMIN_DEFAULT_PASSWORD|g" "$BASE_DIR/docker-compose-dev.yml"

# Criar o diretório de dados se não existir
mkdir -p $DATA_DIR/postgres-data

# Executar o Docker Compose usando o caminho relativo ao diretório base
sudo docker-compose -f "$BASE_DIR/docker-compose-dev.yml" up -d
sudo docker-compose -f "$BASE_DIR/docker-compose-dev.yml" up

# Restaurar o arquivo docker-compose-dev.yml original
sudo gnome-terminal -- bash -c "sleep 5; mv $BASE_DIR/docker-compose-dev.yml.bak $BASE_DIR/docker-compose-dev.yml; exec bash"


